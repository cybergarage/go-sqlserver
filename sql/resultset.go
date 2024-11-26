// Copyright (C) 2024 The go-sqlserver Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

import (
	dbsql "database/sql"
	"errors"

	sql "github.com/cybergarage/go-sqlparser/sql/query/response/resultset"
	_ "github.com/mattn/go-sqlite3"
)

type resultset struct {
	rows   *dbsql.Rows
	schema sql.Schema
}

// NewResultSetDataTypeFrom creates a new result set data type from a column type.
func NewResultSetDataTypeFrom(ct *dbsql.ColumnType) (sql.DataType, error) {
	return nil, nil
}

// NewResultSetColumn creates a new result set column from a column name and type.
func NewResultSetColumnFrom(name string, ct *dbsql.ColumnType) (sql.Column, error) {
	dt, err := NewResultSetDataTypeFrom(ct)
	if err != nil {
		return nil, err
	}
	return sql.NewColumn(
		sql.WithColumnType(dt),
		sql.WithColumnName(name),
	), nil
}

// NewResultSet creates a new result set.
func NewResultSetWith(rows *dbsql.Rows) (sql.ResultSet, error) {
	rs := &resultset{
		rows: rows,
	}
	err := rs.updateSchema()
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (rs *resultset) updateSchema() error {
	rowColumnNames, err := rs.rows.Columns()
	if err != nil {
		return err
	}
	rowColumnTypes, err := rs.rows.ColumnTypes()
	if err != nil {
		return err
	}
	if len(rowColumnNames) != len(rowColumnTypes) {
		return errors.New("column name and type length mismatch")
	}
	rsColums := []sql.Column{}
	for i, name := range rowColumnNames {
		rsColumn, err := NewResultSetColumnFrom(name, rowColumnTypes[i])
		if err != nil {
			return err
		}
		rsColums = append(rsColums, rsColumn)
	}
	rs.schema = sql.NewSchema(
		sql.WithSchemaColumns(rsColums),
	)
	return nil
}

// Schema returns the schema.
func (rs *resultset) Schema() sql.Schema {
	return rs.schema
}

// Next returns the next row.
func (rs *resultset) Next() bool {
	if rs.rows.Err() != nil {
		return false
	}
	return rs.rows.Next()
}

// Row returns the current row.
func (rs *resultset) Row() (sql.Row, error) {
	dest := make([]any, len(rs.schema.Columns()))
	for i := range dest {
		dest[i] = new(any)
	}
	err := rs.rows.Scan(dest...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// RowsAffected returns the number of rows affected.
func (rs *resultset) RowsAffected() uint64 {
	return 0
}

// Close closes the resultset.
func (rs *resultset) Close() error {
	return rs.rows.Close()
}
