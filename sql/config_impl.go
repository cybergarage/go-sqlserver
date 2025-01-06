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
	"bufio"
	"bytes"
	"crypto/tls"
	_ "embed"
	"os"

	"github.com/cybergarage/go-sqlserver/sql/config"
	"github.com/spf13/viper"
)

type configImpl struct {
	config.Config
}

// NewConfig returns a new configuration.
func NewConfig() (Config, error) {
	conf := config.NewConfigWith(ProductName)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &configImpl{conf}, nil
}

// NewConfigWithPath returns a new configuration with the specified path.
func NewConfigWithPath(path string) (Config, error) {
	conf := config.NewConfigWith(ProductName)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &configImpl{conf}, nil
}

func NewConfigWithPaths(paths ...string) (Config, error) {
	conf := config.NewConfigWith(ProductName)
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &configImpl{conf}, nil
}

// NewConfigWithString returns a new configuration with the specified string.
func NewConfigWithString(conString string) (Config, error) {
	conf := config.NewConfigWith(ProductName)
	if err := viper.ReadConfig(bytes.NewBufferString(conString)); err != nil {
		return nil, err
	}
	return &configImpl{conf}, nil
}

// NewConfigWithFile returns a new configuration with the specified file.
func NewConfigWithFile(confFile string) (Config, error) {
	f, err := os.Open(confFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	conf := config.NewConfigWith(ProductName)
	if err := viper.ReadConfig(bufio.NewReader(f)); err != nil {
		return nil, err
	}
	return &configImpl{conf}, nil
}

// MySQLPort returns the MySQL port.
func (config *configImpl) MySQLPort() (int, error) {
	return config.LookupConfigInt(ConfigQuery, ConfigMySQL, ConfigPort)
}

// PostgresqlPort returns the Postgresql port.
func (config *configImpl) PostgresqlPort() (int, error) {
	return config.LookupConfigInt(ConfigQuery, ConfigPostgresql, ConfigPort)
}

// IsLoggerEnabled returns true if the logger is enabled.
func (config *configImpl) IsLoggerEnabled() (bool, error) {
	return config.LookupConfigBool(ConfigLogger, ConfigEnabled)
}

// LoggerLevel returns the logger level.
func (config *configImpl) LoggerLevel() (string, error) {
	return config.LookupConfigString(ConfigLogger, ConfigLevel)
}

// IsTLSEnabled returns true if the TLS is enabled.
func (config *configImpl) IsTLSEnabled() (bool, error) {
	return config.LookupConfigBool(ConfigTLS, ConfigEnabled)
}

// TLSCert returns the TLS certificate.
func (config *configImpl) TLSConfig() (*tls.Config, error) {
	type tlsParams struct {
		CertFile string   `mapstructure:"cert_file"`
		KeyFile  string   `mapstructure:"key_file"`
		CAFiles  []string `mapstructure:"ca_files"`
	}

	var tlsConfig tlsParams
	err := config.UnmarshallConfig([]string{ConfigTLS}, &tlsConfig)
	if err != nil {
		return nil, err
	}

	tlsConf := NewTLSConf()
	tlsConf.SetServerCertFile(tlsConfig.CertFile)
	tlsConf.SetServerKeyFile(tlsConfig.KeyFile)
	tlsConf.SetRootCertFiles(tlsConfig.CAFiles...)

	return tlsConf.TLSConfig()
}

// IsPrometheusEnabled returns true if the Prometheus is enabled.
func (config *configImpl) IsPrometheusEnabled() (bool, error) {
	return config.LookupConfigBool(ConfigMetrics, ConfigPrometheus, ConfigEnabled)
}

// PrometheusPort returns the Prometheus port.
func (config *configImpl) PrometheusPort() (int, error) {
	return config.LookupConfigInt(ConfigMetrics, ConfigPrometheus, ConfigPort)
}

// IsStoreMemory returns true if the store is memory.
func (config *configImpl) IsStoreMemory() (bool, error) {
	return config.LookupConfigBool(ConfigStore, ConfigSQLite, ConfigMemory)
}
