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

// PostgreSQL: Documentation: 16: 55.2. Message Flow
// https://www.postgresql.org/docs/16/protocol-flow.html
// PostgreSQL: Documentation: 16: 55.7. Message Formats
// https://www.postgresql.org/docs/16/protocol-message-formats.html

import (
	"fmt"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/go-postgresql/postgresql/protocol"
	"github.com/cybergarage/go-postgresql/postgresql/query"
)

// Copy handles a COPY query.
func (server *server) Copy(conn postgresql.Conn, q query.Copy) (protocol.Responses, error) {
	/*
		_, tbl, err := server.LookupDatabaseTable(conn, conn.Database(), q.TableName())
		if err != nil {
			return nil, err
		}

		return postgresql.NewCopyInResponsesFrom(q, tbl.Schema)
	*/
	return nil, nil
}

// Copy handles a COPY DATA protocol.
func (server *server) CopyData(conn postgresql.Conn, q query.Copy, stream *postgresql.CopyStream) (protocol.Responses, error) {
	/*
	   _, tbl, err := server.LookupDatabaseTable(conn, conn.Database(), q.TableName())

	   	if err != nil {
	   		log.Error(err)
	   		return nil, err
	   	}

	   return postgresql.NewCopyCompleteResponsesFrom(q, stream, conn, tbl.Schema, server.PostgreSQLServer().QueryExecutor())
	*/
	return nil, nil
}

// ParserError handles a parser error.
func (server *server) ParserError(conn postgresql.Conn, q string, err error) (protocol.Responses, error) {
	switch {
	case postgresql.IsPgbenchGetPartitionQuery(q):
		return postgresql.NewGetPartitionResponseForPgbench()
	}

	resErr := fmt.Errorf("parser error : %w", err)
	log.Warnf(err.Error())
	res, err := protocol.NewErrorResponseWith(resErr)
	if err != nil {
		return nil, err
	}
	return protocol.NewResponsesWith(res), nil
}
