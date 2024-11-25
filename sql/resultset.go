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

	sql "github.com/cybergarage/go-sqlparser/sql/query/response/resultset"
	_ "github.com/mattn/go-sqlite3"
)

type resultset struct {
	*dbsql.Rows
	cursor int
}

// NewResultSet creates a new result set.
func NewResultSetWith(rows *dbsql.Rows) sql.ResultSet {
	return &resultset{
		Rows:   rows,
		cursor: 0,
	}
}

// Schema returns the schema.
func (rs *resultset) Schema() sql.Schema {
	return nil
}

// Next returns the next row.
func (rs *resultset) Next() bool {
	return rs.Rows.Next()
}

// Row returns the current row.
func (rs *resultset) Row() sql.Row {
	return nil
}

// RowsAffected returns the number of rows affected.
func (rs *resultset) RowsAffected() uint64 {
	return 0
}

// Close closes the resultset.
func (rs *resultset) Close() error {
	return rs.Rows.Close()
}
