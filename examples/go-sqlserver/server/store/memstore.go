// Copyright (C) 2020 The go-sqlserver Authors. All rights reserved.
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

package store

import (
	"fmt"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-mysql/mysql/errors"
	"github.com/cybergarage/go-mysql/mysql/net"
	"github.com/cybergarage/go-mysql/mysql/query"
	"github.com/cybergarage/go-sqlparser/sql"
)

type MemStore struct {
	Databases
}

// NewMemStore returns an in-memory storeinstance.
func NewMemStore() *MemStore {
	store := &MemStore{
		Databases: NewDatabases(),
	}
	return store
}

func (store *MemStore) LookupDatabaseTable(conn net.Conn, dbName string, tblName string) (*Database, *Table, error) {
	db, ok := store.LookupDatabase(dbName)
	if !ok {
		return nil, nil, errors.NewErrDatabaseNotExist(dbName)
	}

	tbl, ok := db.LookupTable(tblName)
	if !ok {
		return nil, nil, errors.NewErrTableNotExist(tblName)
	}

	return db, tbl, nil
}

// Begin should handle a BEGIN statement.
func (store *MemStore) Begin(conn net.Conn, stmt query.Begin) error {
	log.Debugf("%v", stmt)
	return errors.ErrNotImplemented
}

// Commit should handle a COMMIT statement.
func (store *MemStore) Commit(conn net.Conn, stmt query.Commit) error {
	log.Debugf("%v", stmt)
	return errors.ErrNotImplemented
}

// Rollback should handle a ROLLBACK statement.
func (store *MemStore) Rollback(conn net.Conn, stmt query.Rollback) error {
	log.Debugf("%v", stmt)
	return errors.ErrNotImplemented
}

// Use should handle a USE statement.
func (store *MemStore) Use(conn net.Conn, stmt query.Use) error {
	log.Debugf("%v", stmt)
	conn.SetDatabase(stmt.DatabaseName())
	return nil
}

// CreateDatabase should handle a CREATE database statement.
func (store *MemStore) CreateDatabase(conn net.Conn, stmt query.CreateDatabase) error {
	log.Debugf("%v", stmt)

	dbName := stmt.DatabaseName()
	_, ok := store.LookupDatabase(dbName)
	if ok {
		if stmt.IfNotExists() {
			return nil
		}
		return errors.NewErrDatabaseExist(dbName)
	}

	return store.AddDatabase(NewDatabaseWithName(dbName))
}

// AlterDatabase should handle a ALTER database statement.
func (store *MemStore) AlterDatabase(conn net.Conn, stmt query.AlterDatabase) error {
	log.Debugf("%v", stmt)
	return errors.ErrNotImplemented
}

// DropDatabase should handle a DROP database statement.
func (store *MemStore) DropDatabase(conn net.Conn, stmt query.DropDatabase) error {
	log.Debugf("%v", stmt)

	dbName := stmt.DatabaseName()
	db, ok := store.LookupDatabase(dbName)
	if !ok {
		if stmt.IfExists() {
			return nil
		}
		return errors.NewErrDatabaseNotExist(dbName)
	}
	return store.Databases.DropDatabase(db)
}

// CreateTable should handle a CREATE table statement.
func (store *MemStore) CreateTable(conn net.Conn, stmt query.CreateTable) error {
	log.Debugf("%v", stmt)

	dbName := conn.Database()
	db, ok := store.LookupDatabase(dbName)
	if !ok {
		return errors.NewErrDatabaseNotExist(dbName)
	}
	tableName := stmt.TableName()
	_, ok = db.LookupTable(tableName)
	if !ok {
		table := NewTableWith(tableName, stmt.Schema())
		db.AddTable(table)
	} else {
		if !stmt.IfNotExists() {
			return errors.NewErrTableExist(tableName)
		}
	}
	return nil
}

// AlterTable should handle a ALTER table statement.
func (store *MemStore) AlterTable(conn net.Conn, stmt query.AlterTable) error {
	// log.Debugf("%v", stmt)
	return errors.ErrNotImplemented
}

// DropTable should handle a DROP table statement.
func (store *MemStore) DropTable(conn net.Conn, stmt query.DropTable) error {
	log.Debugf("%v", stmt)

	dbName := conn.Database()
	db, ok := store.LookupDatabase(dbName)
	if !ok {
		return errors.NewErrDatabaseNotExist(dbName)
	}
	for _, table := range stmt.Tables() {
		tableName := table.TableName()
		table, ok := db.LookupTable(tableName)
		if !ok {
			if stmt.IfExists() {
				continue
			}
			return errors.NewErrTableNotExist(tableName)
		}

		if !db.DropTable(table) {
			return fmt.Errorf("%s could not deleted", table.TableName())
		}
	}
	return nil
}

// Insert should handle a INSERT statement.
func (store *MemStore) Insert(conn net.Conn, stmt query.Insert) error {
	log.Debugf("%v", stmt)

	dbName := conn.Database()
	tableName := stmt.TableName()
	table, ok := store.LookupTableWithDatabase(dbName, tableName)
	if !ok {
		return errors.NewErrTableNotExist(tableName)
	}

	row := NewRowWith(stmt.Columns())
	table.Lock()
	table.Rows = append(table.Rows, row)
	defer table.Unlock()

	return nil
}

// Update should handle a UPDATE statement.
func (store *MemStore) Update(conn net.Conn, stmt query.Update) (sql.ResultSet, error) {
	log.Debugf("%v", stmt)

	_, tbl, err := store.LookupDatabaseTable(conn, conn.Database(), stmt.TableName())
	if err != nil {
		return nil, err
	}

	n, err := tbl.Update(stmt.Columns(), stmt.Where())
	if err != nil {
		return nil, err
	}

	return sql.NewResultSet(
		sql.WithResultSetRowsAffected(uint64(n)),
	), nil
}

// Delete should handle a DELETE statement.
func (store *MemStore) Delete(conn net.Conn, stmt query.Delete) (sql.ResultSet, error) {
	log.Debugf("%v", stmt)

	_, tbl, err := store.LookupDatabaseTable(conn, conn.Database(), stmt.TableName())
	if err != nil {
		return nil, err
	}

	n, err := tbl.Delete(stmt.Where())
	if err != nil {
		return nil, err
	}

	return sql.NewResultSet(
		sql.WithResultSetRowsAffected(uint64(n)),
	), nil
}

// Select should handle a SELECT statement.
func (store *MemStore) Select(conn net.Conn, stmt query.Select) (sql.ResultSet, error) {
	log.Debugf("%v", stmt)

	from := stmt.From()
	if len(from) != 1 {
		return nil, errors.NewErrMultipleTableNotSupported(from.String())
	}

	tblName := from[0].TableName()

	_, tbl, err := store.LookupDatabaseTable(conn, conn.Database(), tblName)
	if err != nil {
		return nil, err
	}

	rows, err := tbl.Select(stmt.Where())
	if err != nil {
		return nil, err
	}

	// Row description response

	selectors := stmt.Selectors()
	if selectors.IsSelectAll() {
		selectors = tbl.Selectors()
	}

	schema := tbl.Schema
	rsSchemaColums := []sql.ResultSetColumn{}
	for _, selector := range selectors {
		colName := selector.Name()
		shemaColumn, err := schema.LookupColumn(colName)
		if err != nil {
			return nil, err
		}
		rsCchemaColumn, err := sql.NewResultSetColumnFrom(shemaColumn)
		if err != nil {
			return nil, err
		}
		rsSchemaColums = append(rsSchemaColums, rsCchemaColumn)
	}

	rsSchema := sql.NewResultSetSchema(
		sql.WithResultSetSchemaDatabaseName(conn.Database()),
		sql.WithResultSetSchemaTableName(tblName),
		sql.WithResultSetSchemaResultSetColumns(rsSchemaColums),
	)

	// Data row response

	rowIdx := 0
	rsRows := []sql.ResultSetRow{}
	if !selectors.HasAggregateFunction() {
		offset := stmt.Limit().Offset()
		limit := stmt.Limit().Limit()
		for rowNo, row := range rows {
			if 0 < offset && rowNo < offset {
				continue
			}
			rowValues := []any{}
			for _, selector := range selectors {
				colName := selector.Name()
				value, err := row.ValueByName(colName)
				if err != nil {
					return nil, err
				}
				rowValues = append(rowValues, value)
			}
			rsRow := sql.NewResultSetRow(
				sql.WithResultSetRowValues(rowValues),
			)
			rsRows = append(rsRows, rsRow)
			rowIdx++
			if 0 < limit && limit <= rowIdx {
				break
			}
		}
	}
	// Return a result set

	rs := sql.NewResultSet(
		sql.WithResultSetSchema(rsSchema),
		sql.WithResultSetRowsAffected(uint64(rowIdx)),
		sql.WithResultSetRows(rsRows),
	)

	return rs, nil
}

// ParserError should handle a parser error.
func (store *MemStore) ParserError(conn net.Conn, q string, err error) error {
	log.Debugf("%v", err)
	return err
}
