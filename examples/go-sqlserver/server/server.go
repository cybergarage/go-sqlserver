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

package server

import (
	"github.com/cybergarage/go-sqlserver/examples/go-sqlserver/server/store"
	"github.com/cybergarage/go-sqlserver/sql"
)

// Server is a SQL server.
type Server struct {
	sql.Server
	*store.Store
}

// NewServer creates a new Server.
func NewServer() *Server {
	server := &Server{
		Server: sql.NewServer(),
		Store:  store.NewStore(),
	}

	// Set common SQL executor for MySQL and PostgreSQL
	server.SetSQLExecutor(server.Store)

	// PostgreSQL server settings
	server.PostgreSQLServer().SetBulkQueryExecutor(server)
	server.PostgreSQLServer().SetErrorHandler(server)

	return server
}
