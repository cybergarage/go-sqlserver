# go-sqlserver

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-sqlserver)
[![test](https://github.com/cybergarage/go-sqlserver/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/go-sqlserver/actions/workflows/make.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-sqlserver.svg)](https://pkg.go.dev/github.com/cybergarage/go-sqlserver) [![codecov](https://codecov.io/gh/cybergarage/go-sqlserver/graph/badge.svg?token=2RYOJPQRDM)](https://codecov.io/gh/cybergarage/go-sqlserver)

The `go-sqlserver` is a database framework for implementing a [MySQL](https://www.mysql.com/)-compatible and a [PostgreSQL](https://www.postgresql.org/)-compatible server using Go easily.
This framework is ideal for building mock servers for testing, custom database solutions, or learning purposes.

## What is the go-sqlserver?

The `go-sqlserver` provides a set of tools and libraries to create custom SQL servers that can handle MySQL and PostgreSQL protocols. It abstracts the complexities of protocol handling, query parsing, and execution, allowing developers to focus on implementing the business logic and data storage mechanisms.

![](doc/img/framework.png)

The `go-sqlserver` provides a unified implementation framework of authentication and query handlers for both MySQL and PostgreSQL, allowing developers to build custom SQL servers that support both protocols. The framework is designed to be extensible, allowing developers to add custom query handlers, authentication mechanisms, and data storage backends.
By using the framework , developers can easily build an SQL server that supports both MySQL and PostgreSQL protocols.

The `go-sqlserver` is currently implemented as an in-memory database using SQLite as the default data store. This allows for quick and efficient data operations without the need for an external database server.

## Get Started

See the following guide to learn about how to get started.

- [Quick Start](doc/quick-start.md) [![Docker Image Version](https://img.shields.io/docker/v/cybergarage/go-sqlserver)](https://hub.docker.com/repository/docker/cybergarage/go-sqlserver/)
  - [Configuring go-sqlserver](doc/configuring.md)
- Secifications
  - [Data Model](doc/data_model.md)
  - [Query Model](doc/query_model.md)

# Related Projects

The `go-sqlserver` is being developed in collaboration with the following Cybergarage projects:

- [go-postgresql](https://github.com/cybergarage/go-postgresql) ![go postgresql](https://img.shields.io/github/v/tag/cybergarage/go-postgresql)
- [go-mysql](https://github.com/cybergarage/go-mysql) ![go mysql](https://img.shields.io/github/v/tag/cybergarage/go-mysql)
- [go-sqlparser](https://github.com/cybergarage/go-sqlparser) ![go sqlparser](https://img.shields.io/github/v/tag/cybergarage/go-sqlparser)
- [go-logger](https://github.com/cybergarage/go-logger) ![go logger](https://img.shields.io/github/v/tag/cybergarage/go-logger)
- [go-safecast](https://github.com/cybergarage/go-safecast) ![go safecast](https://img.shields.io/github/v/tag/cybergarage/go-safecast)
- [go-tracing](https://github.com/cybergarage/go-tracing) ![go tracing](https://img.shields.io/github/v/tag/cybergarage/go-tracing)
- [go-authenticator](https://github.com/cybergarage/go-authenticator) ![go authenticator](https://img.shields.io/github/v/tag/cybergarage/go-authenticator)
- [go-sasl](https://github.com/cybergarage/go-sasl) ![go sasl](https://img.shields.io/github/v/tag/cybergarage/go-sasl)
- [go-sqltest](https://github.com/cybergarage/go-sqltest) ![go sqltest](https://img.shields.io/github/v/tag/cybergarage/go-sqltest)
- [go-pict](https://github.com/cybergarage/go-pict) ![go pict](https://img.shields.io/github/v/tag/cybergarage/go-pict)
