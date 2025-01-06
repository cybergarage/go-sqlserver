// Copyright (C) 2019 The go-sqlserver Authors. All rights reserved.
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
	"github.com/cybergarage/go-mysql/mysql"
	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/go-sqlparser/sql"
	"github.com/cybergarage/go-tracing/tracer"
)

// SQLExecutor represents a SQL executor.
type SQLExecutor = sql.Executor

// Server represents a PostgreSQL protocol server.
type Server interface {
	// SetConfig sets a configuration.
	SetConfig(Config)
	// SetTracer sets a tracing tracer.
	SetTracer(tracer.Tracer)
	// SetSQLExecutor sets a SQL executor.
	SetSQLExecutor(SQLExecutor)
	// MySQLServer returns a MySQL server.
	MySQLServer() mysql.Server
	// PostgreSQLServer returns a PostgreSQL server.
	PostgreSQLServer() postgresql.Server
	// Start starts the server.
	Start() error
	// Stop stops the server.
	Stop() error
	// Restart restarts the server.
	Restart() error
}
