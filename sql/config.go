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
	"crypto/tls"

	"github.com/cybergarage/go-sqlserver/sql/auth"
)

const (
	ConfigLogger     = "logger"
	ConfigTLS        = "tls"
	ConfigAuth       = "auth"
	ConfigQuery      = "query"
	ConfigTracer     = "tracer"
	ConfigMetrics    = "metrics"
	ConfigMySQL      = "mysql"
	ConfigPostgresql = "postgresql"
	ConfigPort       = "port"
	ConfigEnabled    = "enabled"
	ConfigLevel      = "level"
	ConfigPrometheus = "prometheus"
	ConfigStore      = "store"
	ConfigSQLite     = "sqlite"
	ConfigMemory     = "memory"
	ConfigPlain      = "plain"
)

// Config represents a configuration interface for PuzzleDB.
type Config interface {
	// MySQLPort returns the MySQL port.
	MySQLPort() (int, error)
	// PostgresqlPort returns the Postgresql port.
	PostgresqlPort() (int, error)
	// IsLoggerEnabled returns true if the logger is enabled.
	IsLoggerEnabled() (bool, error)
	// LoggerLevel returns the logger level.
	LoggerLevel() (string, error)
	// IsTLSEnabled returns true if the TLS is enabled.
	IsTLSEnabled() (bool, error)
	// TLSCert returns the TLS certificate.
	TLSConfig() (*tls.Config, error)
	// IsPrometheusEnabled returns true if the Prometheus is enabled.
	IsPrometheusEnabled() (bool, error)
	// PrometheusPort returns the Prometheus port.
	PrometheusPort() (int, error)
	// IsMemoryStoreEnabled returns true if the store is memory.
	IsMemoryStoreEnabled() (bool, error)
	// PlainCredentials returns plain configurations.
	PlainCredentials() ([]auth.PlainConfig, error)
}
