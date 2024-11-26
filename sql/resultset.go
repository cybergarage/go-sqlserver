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

	query "github.com/cybergarage/go-sqlparser/sql/query"
	sql "github.com/cybergarage/go-sqlparser/sql/query/response/resultset"
	_ "github.com/mattn/go-sqlite3"
)

// ResultSetOption is a result set option.
type ResultSetOption func(*resultset) error

type resultset struct {
	rows         *dbsql.Rows
	schema       sql.Schema
	rowsAffected uint64
}

// NewResultSetDataTypeFrom creates a new result set data type from a column type.
func NewResultSetDataTypeFrom(ct *dbsql.ColumnType) (sql.DataType, error) {
	s := ct.DatabaseTypeName()
	switch s {
	case "INTEGER":
		return query.IntegerData, nil
	case "REAL":
		return query.RealData, nil
	case "TEXT":
		return query.TextData, nil
	case "BLOB":
		return query.BlobData, nil
	case "NUMERIC":
		return query.RealData, nil
	}
	return 0, errors.New("unsupported data type")
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

// WithResultSetRows sets the result set rows.
func WithResultSetRows(rows *dbsql.Rows) ResultSetOption {
	return func(rs *resultset) error {
		rs.rows = rows
		if rs.rows == nil {
			return errors.New("rows is nil")
		}
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
}

// WithResultSetResult sets the result set result.
func WithResultSetResult(result dbsql.Result) ResultSetOption {
	return func(rs *resultset) error {
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		rs.rowsAffected = uint64(rowsAffected)
		return nil
	}
}

// NewResultSet creates a new result set.
func NewResultSet(opts ...ResultSetOption) (sql.ResultSet, error) {
	rs := &resultset{
		rows:         nil,
		schema:       nil,
		rowsAffected: 0,
	}
	for _, opt := range opts {
		err := opt(rs)
		if err != nil {
			return nil, err
		}
	}
	return rs, nil
}

// Schema returns the schema.
func (rs *resultset) Schema() sql.Schema {
	return rs.schema
}

// Next returns the next row.
func (rs *resultset) Next() bool {
	if rs.rows == nil {
		return false
	}
	if rs.rows.Err() != nil {
		return false
	}
	return rs.rows.Next()
}

// Row returns the current row.
func (rs *resultset) Row() (sql.Row, error) {
	if rs.schema == nil {
		return nil, errors.New("schema is nil")
	}
	if rs.rows == nil {
		return nil, errors.New("rows is nil")
	}
	dest := make([]any, len(rs.schema.Columns()))
	for n, column := range rs.schema.Columns() {
		switch column.DataType() {
		case query.IntegerData:
			dest[n] = new(int64)
		case query.RealData:
			dest[n] = new(float64)
		case query.TextData:
			dest[n] = new(string)
		case query.BlobData:
			dest[n] = new([]byte)
		default:
			dest[n] = new(any)
		}
	}
	err := rs.rows.Scan(dest...)
	if err != nil {
		return nil, err
	}
	obj := map[string]any{}
	for n, column := range rs.schema.Columns() {
		obj[column.Name()] = dest[n]
	}
	return sql.NewRow(
			sql.WithRowSchema(rs.schema),
			sql.WithRowObject(obj),
		),
		nil
}

// RowsAffected returns the number of rows affected.
func (rs *resultset) RowsAffected() uint64 {
	return rs.rowsAffected
}

// Close closes the resultset.
func (rs *resultset) Close() error {
	if rs.rows == nil {
		return nil
	}
	return rs.rows.Close()
}
