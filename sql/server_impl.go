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
	"github.com/cybergarage/go-sqlserver/sql/auth"
	"github.com/cybergarage/go-tracing/tracer"
)

// Server represents a SQL server.
type server struct {
	Config
	auth.Manager
	*Databases
	myServer   mysql.Server
	pgServer   postgresql.Server
	ptExporter *PrometheusExporter
}

// NewServer creates a new SQL server.
func NewServer() Server {
	conf, err := NewDefaultConfig()
	if err != nil {
		panic(err)
	}

	server := &server{
		Config:     conf,
		Manager:    auth.NewManager(),
		Databases:  NewDatabases(),
		myServer:   mysql.NewServer(),
		pgServer:   postgresql.NewServer(),
		ptExporter: NewPrometheusExporter(),
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

func (server *server) setupPortConfig() error {
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

func (server *server) setupLogger() error {
	ok, err := server.IsLoggerEnabled()
	if err != nil {
		return err
	}

	if !ok {
		return nil
	}

	levelStr, err := server.LoggerLevel()
	if err != nil {
		return err
	}

	level := log.GetLevelFromString(levelStr)
	log.SetSharedLogger(log.NewStdoutLogger(level))

	return nil
}

func (server *server) setupPrometheus() error {
	ok, err := server.IsPrometheusEnabled()
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	port, err := server.PrometheusPort()
	if err != nil {
		return err
	}
	server.ptExporter.SetPort(port)
	return nil
}

func (server *server) setupTLSConfig() error {
	server.myServer.SetTLSConfig(nil)
	server.pgServer.SetTLSConfig(nil)

	ok, err := server.IsTLSEnabled()
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	tlsConfig, err := server.TLSConfig()
	if err != nil {
		return err
	}
	server.myServer.SetTLSConfig(tlsConfig)
	server.pgServer.SetTLSConfig(tlsConfig)

	return nil
}

func (server *server) setupCredentialConfig() error {
	plainConfigs, err := server.PlainCredentials()
	if err != nil {
		return err
	}

	creds := []auth.Credential{}
	for _, plainConfig := range plainConfigs {
		if !plainConfig.Enabled {
			continue
		}
		cred := auth.NewCredential(
			auth.WithCredentialUsername(plainConfig.Username),
			auth.WithCredentialPassword(plainConfig.Password),
		)
		creds = append(creds, cred)
	}

	if 0 < len(creds) {
		server.SetCredentials(creds...)
		server.myServer.SetCredentialStore(server)
		server.pgServer.SetCredentialStore(server)
	} else {
		server.myServer.SetCredentialStore(nil)
		server.pgServer.SetCredentialStore(nil)
	}

	return err
}

// Start starts the SQL server.
func (server *server) Start() error {
	setupper := []func() error{
		server.setupLogger,
		server.setupPortConfig,
		server.setupTLSConfig,
		server.setupCredentialConfig,
		server.setupPrometheus,
	}

	for _, setup := range setupper {
		if err := setup(); err != nil {
			return errors.Join(err, server.Stop())
		}
	}

	type starter interface {
		Start() error
	}

	starters := []starter{
		server.myServer,
		server.pgServer,
	}

	ok, err := server.IsPrometheusEnabled()
	if err != nil {
		return err
	}
	if ok {
		starters = append(starters, server.ptExporter)
	}

	for _, s := range starters {
		if err := s.Start(); err != nil {
			return errors.Join(err, server.Stop())
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

	ok, err := server.IsPrometheusEnabled()
	if err != nil {
		return err
	}
	if ok {
		stoppers = append(stoppers, server.ptExporter)
	}

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
