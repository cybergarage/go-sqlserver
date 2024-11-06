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

// Server represents a SQL server.
type Server struct {
}

// NewServer creates a new SQL server.
func NewServer() *Server {
	return &Server{}
}

// Start starts the SQL server.
func (server *Server) Start() error {
	return nil
}

// Stop stops the SQL server.
func (server *Server) Stop() error {
	return nil
}

// Restart restarts the SQL server.
func (server *Server) Restart() error {
	if err := server.Stop(); err != nil {
		return err
	}
	return server.Start()
}
