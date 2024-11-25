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
	"fmt"
	"sync"

	"github.com/cybergarage/go-sqlparser/sql/errors"
)

// Databases represents a collection of databases.
type Databases struct {
	dbmap sync.Map
}

// NewDatabases returns a databases instance.
func NewDatabases() *Databases {
	return &Databases{
		dbmap: sync.Map{},
	}
}

// AddDatabase adds a specified database.
func (dbs *Databases) AddDatabase(db *Database) error {
	dbName := db.Name()
	if _, ok := dbs.dbmap.Load(dbName); ok {
		return fmt.Errorf("database %s already %w", dbName, errors.ErrExist)
	}
	dbs.dbmap.Store(dbName, db)
	return nil
}

// DropDatabase remove the specified database.
func (dbs *Databases) DropDatabase(db *Database) error {
	name := db.Name()
	dbs.dbmap.Delete(name)
	return nil
}

// LookupDatabase returns a database with the specified name.
func (dbs *Databases) LookupDatabase(name string) (*Database, error) {
	v, ok := dbs.dbmap.Load(name)
	if !ok {
		return nil, fmt.Errorf("database %s %w", name, errors.ErrNotExist)
	}
	db, ok := v.(*Database)
	if !ok {
		return nil, fmt.Errorf("database %s %w", name, errors.ErrNotExist)
	}
	return db, nil
}
