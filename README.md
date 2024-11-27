# go-sqlserver

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-sqlserver)
[![test](https://github.com/cybergarage/go-sqlserver/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/go-sqlserver/actions/workflows/make.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-sqlserver.svg)](https://pkg.go.dev/github.com/cybergarage/go-sqlserver) [![codecov](https://codecov.io/gh/cybergarage/go-sqlserver/graph/badge.svg?token=2RYOJPQRDM)](https://codecov.io/gh/cybergarage/go-sqlserver)

The `go-sqlserver` is a database framework for implementing a [MySQL](https://www.mysql.com/)-compatible and a [PostgreSQL](https://www.postgresql.org/)-compatible server using Go easily.
This framework is ideal for building mock servers for testing, custom database solutions, or learning purposes.

# What is the go-sqlserver?

The `go-sqlserver` provides a set of tools and libraries to create custom SQL servers that can handle MySQL and PostgreSQL protocols. It abstracts the complexities of protocol handling, query parsing, and execution, allowing developers to focus on implementing the business logic and data storage mechanisms.

![](doc/img/framework.png)

The `go-sqlserver` provides a unified implementation framework of authentication and query handlers for both MySQL and PostgreSQL, allowing developers to build custom SQL servers that support both protocols. The framework is designed to be extensible, allowing developers to add custom query handlers, authentication mechanisms, and data storage backends.
By using the framework , developers can easily build an SQL server that supports both MySQL and PostgreSQL protocols.

The `go-sqlserver` is currently implemented as an in-memory database using SQLite as the default data store. This allows for quick and efficient data operations without the need for an external database server.
