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

package server

import (
	dbsql "database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-sqlparser/sql"
	"github.com/cybergarage/go-sqlparser/sql/errors"
	"github.com/cybergarage/go-sqlparser/sql/query"
)

// Store represents a data store.
type Store struct {
}

// Begin should handle a BEGIN statement.
func (store *Store) Begin(conn net.Conn, stmt query.Begin) error {
	return nil
}

// Commit should handle a COMMIT statement.
func (store *Store) Commit(conn net.Conn, stmt query.Commit) error {
	return nil
}

// Rollback should handle a ROLLBACK statement.
func (store *Store) Rollback(conn net.Conn, stmt query.Rollback) error {
	return nil
}

// Use should handle a USE statement.
func (store *Store) Use(conn net.Conn, stmt query.Use) error {
	return nil
}

// CreateDatabase should handle a CREATE database statement.
func (store *Store) CreateDatabase(conn net.Conn, stmt query.CreateDatabase) error {
	db, err := dbsql.Open("sqlite3", "./test.db")
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

// AlterDatabase should handle a ALTER database statement.
func (store *Store) AlterDatabase(conn net.Conn, stmt query.AlterDatabase) error {
	return nil
}

// DropDatabase should handle a DROP database statement.
func (store *Store) DropDatabase(conn net.Conn, stmt query.DropDatabase) error {
	return nil
}

// CreateTable should handle a CREATE table statement.
func (store *Store) CreateTable(conn net.Conn, stmt query.CreateTable) error {
	return nil
}

// AlterTable should handle a ALTER table statement.
func (store *Store) AlterTable(conn net.Conn, stmt query.AlterTable) error {
	return nil
}

// DropTable should handle a DROP table statement.
func (store *Store) DropTable(conn net.Conn, stmt query.DropTable) error {
	return nil
}

// Insert should handle a INSERT statement.
func (store *Store) Insert(conn net.Conn, stmt query.Insert) error {
	return nil
}

// Update should handle a UPDATE statement.
func (store *Store) Update(conn net.Conn, stmt query.Update) (sql.ResultSet, error) {
	return nil, nil
}

// Delete should handle a DELETE statement.
func (store *Store) Delete(conn net.Conn, stmt query.Delete) (sql.ResultSet, error) {
	return nil, nil
}

// Select should handle a SELECT statement.
func (store *Store) Select(conn net.Conn, stmt query.Select) (sql.ResultSet, error) {
	return nil, nil
}

// SystemSelect should handle a system SELECT statement.
func (store *Store) SystemSelect(conn net.Conn, stmt query.Select) (sql.ResultSet, error) {
	log.Debugf("%v", stmt)
	return nil, errors.NewErrNotImplemented("SystemSelect")
}
