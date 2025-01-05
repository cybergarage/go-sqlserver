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
	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-mysql/mysql/errors"
	"github.com/cybergarage/go-sqlparser/sql/net"
	"github.com/cybergarage/go-sqlparser/sql/query"
	sql "github.com/cybergarage/go-sqlparser/sql/query/response/resultset"
)

// Begin should handle a BEGIN statement.
func (server *server) Begin(conn net.Conn, stmt query.Begin) error {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(conn.Database())
	if err != nil {
		return err
	}
	return db.Begin()
}

// Commit should handle a COMMIT statement.
func (server *server) Commit(conn net.Conn, stmt query.Commit) error {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(conn.Database())
	if err != nil {
		return err
	}
	return db.Commit()
}

// Rollback should handle a ROLLBACK statement.
func (server *server) Rollback(conn net.Conn, stmt query.Rollback) error {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(conn.Database())
	if err != nil {
		return err
	}
	return db.Rollback()
}

// Use should handle a USE statement.
func (server *server) Use(conn net.Conn, stmt query.Use) error {
	log.Debugf("%v", stmt)
	conn.SetDatabase(stmt.DatabaseName())
	return nil
}

// CreateDatabase should handle a CREATE database statement.
func (server *server) CreateDatabase(conn net.Conn, stmt query.CreateDatabase) error {
	log.Debugf("%v", stmt)

	dbName := stmt.DatabaseName()

	_, err := server.LookupDatabase(dbName)
	if err == nil {
		if stmt.IfNotExists() {
			return nil
		}
		return newErrDatabaseExist(dbName)
	}

	db, err := NewDatabaseWithName(dbName)
	if err != nil {
		return err
	}
	err = server.Databases.AddDatabase(db)
	if err != nil {
		return err
	}
	conn.SetDatabase(dbName)
	return nil
}

// AlterDatabase should handle a ALTER database statement.
func (server *server) AlterDatabase(conn net.Conn, stmt query.AlterDatabase) error {
	log.Debugf("%v", stmt)
	return errors.ErrNotImplemented
}

// DropDatabase should handle a DROP database statement.
func (server *server) DropDatabase(conn net.Conn, stmt query.DropDatabase) error {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(stmt.DatabaseName())
	if err != nil {
		if stmt.IfExists() {
			return nil
		}
		return err
	}
	return server.Databases.DropDatabase(db)
}

// CreateTable should handle a CREATE table statement.
func (server *server) CreateTable(conn net.Conn, stmt query.CreateTable) error {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(conn.Database())
	if err != nil {
		return err
	}
	_, err = db.Exec(stmt.String())
	return err
}

// AlterTable should handle a ALTER table statement.
func (server *server) AlterTable(conn net.Conn, stmt query.AlterTable) error {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(conn.Database())
	if err != nil {
		return err
	}
	_, err = db.Exec(stmt.String())
	return err
}

// DropTable should handle a DROP table statement.
func (server *server) DropTable(conn net.Conn, stmt query.DropTable) error {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(conn.Database())
	if err != nil {
		return err
	}
	_, err = db.Exec(stmt.String())
	return err
}

func (server *server) Insert(conn net.Conn, stmt query.Insert) error {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(conn.Database())
	if err != nil {
		return err
	}
	_, err = db.Exec(stmt.String())
	return err
}

// Update should handle a UPDATE statement.
func (server *server) Update(conn net.Conn, stmt query.Update) (sql.ResultSet, error) {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(conn.Database())
	if err != nil {
		return nil, err
	}
	result, err := db.Exec(stmt.String())
	if err != nil {
		return nil, err
	}
	return NewResultSet(
		WithResultSetResult(result),
	)
}

// Delete should handle a DELETE statement.
func (server *server) Delete(conn net.Conn, stmt query.Delete) (sql.ResultSet, error) {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(conn.Database())
	if err != nil {
		return nil, err
	}
	result, err := db.Exec(stmt.String())
	if err != nil {
		return nil, err
	}
	return NewResultSet(
		WithResultSetResult(result),
	)
}

// Select should handle a SELECT statement.
func (server *server) Select(conn net.Conn, stmt query.Select) (sql.ResultSet, error) {
	log.Debugf("%v", stmt)
	db, err := server.LookupDatabase(conn.Database())
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(stmt.String())
	if err != nil {
		return nil, err
	}
	return NewResultSet(
		WithResultSetRows(rows),
	)
}

// SystemSelect should handle a system SELECT statement.
func (server *server) SystemSelect(conn net.Conn, stmt query.Select) (sql.ResultSet, error) {
	log.Debugf("%v", stmt)
	return nil, errors.NewErrNotImplemented("SystemSelect")
}
