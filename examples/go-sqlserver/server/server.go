// Copyright (C) 2020 The go-mysql Authors. All rights reserved.
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

// Server represents a test server.
// This Server struct behave as ${hoge}CommandExecutor.
type Server struct {
	*sql.Server
	Store
}

// NewServerWithStore returns a test server instance with the specified store.
func NewServerWithStore(store Store) *Server {
	server := &Server{
		Server: sql.NewServer(),
		Store:  store,
	}
	// server.SetQueryExecutor(store)
	return server
}

// NewServer returns a test server instance.
func NewServer() *Server {
	// NOTE: MemStore is a sample implementation. So, change to use your implementation.
	return NewServerWithStore(store.NewMemStore())
}

// GetStore returns a store in the server.
func (server *Server) GetStore() Store {
	return server.Store
}
