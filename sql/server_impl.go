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
	"errors"

	"github.com/cybergarage/go-mysql/mysql"
	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/go-tracing/tracer"
)

// Server represents a SQL server.
type server struct {
	myServer mysql.Server
	pgServer postgresql.Server
}

// NewServer creates a new SQL server.
func NewServer() Server {
	return &server{
		myServer: mysql.NewServer(),
		pgServer: postgresql.NewServer(),
	}
}

// SetTracer sets a tracing tracer.
func (server *server) SetTracer(tracer tracer.Tracer) {
	server.myServer.SetTracer(tracer)
	server.pgServer.SetTracer(tracer)
}

// SetSQLExecutor sets a SQL executor.
func (server *server) SetSQLExecutor(ex SQLExecutor) {
	server.myServer.SetSQLExecutor(ex)
	server.pgServer.SetSQLExecutor(ex)
}

// Start starts the SQL server.
func (server *server) Start() error {
	type starter interface {
		Start() error
	}
	starters := []starter{
		server.myServer,
		server.pgServer,
	}
	for _, s := range starters {
		if err := s.Start(); err != nil {
			return server.Stop()
		}
	}
	return nil
}

// Stop stops the SQL server.
func (server *server) Stop() error {
	type stopper interface {
		Stop() error
	}
	stoppers := []stopper{
		server.myServer,
		server.pgServer,
	}
	var err error
	for _, s := range stoppers {
		if e := s.Stop(); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}

// Restart restarts the SQL server.
func (server *server) Restart() error {
	if err := server.Stop(); err != nil {
		return err
	}
	return server.Start()
}
