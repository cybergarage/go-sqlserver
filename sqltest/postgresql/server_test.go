// Copyright (C) 2025 The go-sqlserver Authors. All rights reserved.
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

package postgresql

import (
	"testing"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/go-sqlserver/sqltest/server"
)

const (
	clientKey  = "../certs/key.pem"
	clientCert = "../certs/cert.pem"
	rootCert   = "../certs/root_cert.pem"
)

func TestServer(t *testing.T) {
	log.SetStdoutDebugEnbled(true)

	const (
		username = "testuser"
		password = "testpassword"
	)

	settings := []struct {
		isTLSEnabled      bool
		isPasswordEnabled bool
	}{
		{
			isTLSEnabled:      false,
			isPasswordEnabled: false,
		},
		{
			isTLSEnabled:      false,
			isPasswordEnabled: true,
		},
		{
			isTLSEnabled:      true,
			isPasswordEnabled: false,
		},
		{
			isTLSEnabled:      true,
			isPasswordEnabled: true,
		},
	}

	for _, setting := range settings {
		t.Logf("TLS: %v, Password: %v", setting.isTLSEnabled, setting.isPasswordEnabled)

		if setting.isPasswordEnabled {
			t.Setenv("GO_SQLSERVER_AUTH_ENABLED", "true")
		}

		server := server.NewServer()

		err := server.Start()
		if err != nil {
			t.Error(err)
			return
		}

		client := postgresql.NewDefaultClient()
		if setting.isTLSEnabled {
			client.SetClientKeyFile(clientKey)
			client.SetClientCertFile(clientCert)
			client.SetRootCertFile(rootCert)
		}
		if setting.isPasswordEnabled {
			client.SetUser(username)
			client.SetPassword(password)
		}
		client.SetDatabase("ycsb")

		err = client.Open()
		defer client.Close()
		if err != nil {
			t.Error(err)
			return
		}

		err = client.Ping()
		if err != nil {
			t.Error(err)
			return
		}

		err = server.Stop()
		if err != nil {
			t.Error(err)
			return
		}
	}
}
