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

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-mysql/mysql"
	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/go-tracing/tracer"
)

// Server represents a SQL server.
type server struct {
	Config
	*Databases
	myServer mysql.Server
	pgServer postgresql.Server
}

// NewServer creates a new SQL server.
func NewServer() Server {
	conf, err := NewDefaultConfig()
	if err != nil {
		panic(err)
	}

	server := &server{
		Config:    conf,
		Databases: NewDatabases(),
		myServer:  mysql.NewServer(),
		pgServer:  postgresql.NewServer(),
	}

	// Set common SQL executor for MySQL and PostgreSQL
	server.SetSQLExecutor(server)

	// PostgreSQL server settings
	server.PostgreSQLServer().SetBulkQueryExecutor(server)
	server.PostgreSQLServer().SetErrorHandler(server)

	return server
}

// SetConfig sets a configuration.
func (server *server) SetConfig(conf Config) {
	server.Config = conf
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

// MySQLServer returns a MySQL server.
func (server *server) MySQLServer() mysql.Server {
	return server.myServer
}

// PostgreSQLServer returns a PostgreSQL server.
func (server *server) PostgreSQLServer() postgresql.Server {
	return server.pgServer
}

func (server *server) applyConfig() error {
	port, err := server.MySQLPort()
	if err != nil {
		return err
	}
	server.myServer.SetPort(port)

	port, err = server.PostgresqlPort()
	if err != nil {
		return err
	}
	server.pgServer.SetPort(port)

	return nil
}

// Start starts the SQL server.
func (server *server) Start() error {
	log.SetSharedLogger(nil)
	if ok, err := server.IsLoggerEnabled(); err == nil {
		if ok {
			levelStr, err := server.LoggerLevel()
			if err != nil {
				return err
			}
			level := log.GetLevelFromString(levelStr)
			log.SetSharedLogger(log.NewStdoutLogger(level))
		}
	} else {
		return err
	}

	log.Infof("%s %s started", ProductName, Version)

	if err := server.applyConfig(); err != nil {
		return err
	}

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

	log.Infof("%s %s terminated", ProductName, Version)

	return err
}

// Restart restarts the SQL server.
func (server *server) Restart() error {
	if err := server.Stop(); err != nil {
		return err
	}
	return server.Start()
}
