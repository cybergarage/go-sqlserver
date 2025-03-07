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

package sqltest

import (
	"testing"

	"github.com/cybergarage/go-sqlserver/sql"
)

func TestDefaultConfig(t *testing.T) {
	cfg, err := sql.NewDefaultConfig()
	if err != nil {
		t.Error(err)
		return
	}

	type fn func() (any, error)
	fns := []fn{
		func() (any, error) {
			return cfg.MySQLPort()
		},
		func() (any, error) {
			return cfg.PostgresqlPort()
		},
	}

	for _, f := range fns {
		if _, err := f(); err != nil {
			t.Error(err)
		}
	}
}
