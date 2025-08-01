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

package mysql

import (
	"testing"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-sqlserver/sqltest/server"
	"github.com/cybergarage/go-sqltest/sqltest"
)

// TestSQLTestSuite runs already passed scenario test files.
func TestSQLTestSuite(t *testing.T) {
	log.EnableStdoutDebug(true)

	server := server.NewServer()
	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Stop()

	testNames := []string{
		"SmplCrud.*",
		"SmplIndex.*",
		"YcsbWorkload",
	}

	sqltest.RunEmbedSuites(t, sqltest.NewMySQLClient(), testNames...)
}
