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

import "errors"

// Datatypes In SQLite
// https://www.sqlite.org/datatype3.html

// DataType represents a data type.
type DataType = int

// Data types.
const (
	// Integer.
	Integer DataType = iota
	// Real.
	Real
	// Text.
	Text
	// Blob.
	Blob
	// Numeric.
	Numeric
)

// NewDataTypeFrom creates a new data type from a string.
func NewDataTypeFrom(s string) (DataType, error) {
	switch s {
	case "INTEGER":
		return Integer, nil
	case "REAL":
		return Real, nil
	case "TEXT":
		return Text, nil
	case "BLOB":
		return Blob, nil
	case "NUMERIC":
		return Numeric, nil
	}
	return 0, errors.New("unsupported data type")
}
