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
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Database represents a destination or source database of query.
type Database struct {
	name string
	db   *sql.DB
	tx   *sql.Tx
}

// NewDatabaseWithName returns a new database with the specified string.
func NewDatabaseWithName(name string) (*Database, error) {
	var err error
	db := &Database{
		name: name,
		db:   nil,
		tx:   nil,
	}
	db.db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// DB returns the database.
func (db *Database) DB() *sql.DB {
	return db.db
}

// Name returns the database name.
func (db *Database) Name() string {
	return db.name
}

// Begin starts a transaction.
func (db *Database) Begin() error {
	if db.tx != nil {
		err := db.tx.Rollback()
		if err != nil {
			return err
		}
	}
	var err error
	db.tx, err = db.db.Begin()
	if err != nil {
		return err
	}
	return err
}

// Commit commits a transaction.
func (db *Database) Commit() error {
	if db.tx == nil {
		return nil
	}
	err := db.tx.Commit()
	if err != nil {
		return err
	}
	db.tx = nil
	return nil
}

// Rollback rolls back a transaction.
func (db *Database) Rollback() error {
	if db.tx == nil {
		return nil
	}
	err := db.tx.Rollback()
	if err != nil {
		return err
	}
	db.tx = nil
	return nil
}

// Exec executes a query.
func (db *Database) Exec(query string, args ...any) (sql.Result, error) {
	if db.tx != nil {
		return db.tx.Exec(query, args...)
	}
	return db.db.Exec(query, args...)
}
