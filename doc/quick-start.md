# Quick Start

This chapter shows you how to get started with **go-sqlserver** quickly: you can start a standalone **go-sqlserver** server with Docker and MySQL or PostgreSQL CLI commands to insert and read sample data.

## Starting **go-sqlserver** Server

### Using Docker image

**go-sqlserver** [Docker image](https://hub.docker.com/r/cybergarage/go-sqlserver) is the easiest way; if you do not have Docker installed, go there and install it first. To start the standalone server, run the following command:

```
docker run -it --rm \
 -p 3306:3306 \
 -p 5432:5432 \
 -p 9181:9181 \
 cybergarage/go-sqlserver:latest
```
### Building from Source

To start the latest **go-sqlserver**, refer to [Go Stared](https://go.dev/learn/) to set up your Go development environment and run the following command:

```
git clone https://github.com/cybergarage/go-sqlserver.git
cd go-sqlserver
make run
```

## Using database clients

The started **go-sqlserver** listens on the standard ports of the supported PostgreSQL and MySQL database protocols, and you can connect with **go-sqlserver** using the standard CLI commands.

## MySQL

To operate **go-sqlserver** with the MySQL protocol, use the standard MySQL shell [mysql](https://dev.mysql.com/doc/refman/8.0/en/mysql.html) as follows:

```
% mysql -h 127.0.0.1
mysql> CREATE DATABASE test;
mysql> USE test;
mysql> CREATE TABLE test (k VARCHAR(255) PRIMARY KEY, v int);
mysql> INSERT INTO test (k, v) VALUES ('foo', 0);
mysql> SELECT * FROM test WHERE k = 'foo';
+------+------+
| k  | v  |
+------+------+
| foo |  0 |
+------+------+
1 row in set (0.00 sec)
```

**go-sqlserver** currently supports the MySQL queries in stages. See [MySQL](doc/mysql.md) for current support status.

## PostgreSQL

To operate **go-sqlserver** with the PostgreSQL protocol, use the standard PostgreSQL shell [psql](https://www.postgresql.org/docs/current/app-psql.html) as follows:

```
% psql --host=localhost
> CREATE DATABASE test;
> \q
% psql --host=localhost test
> CREATE TABLE test (k VARCHAR(255) PRIMARY KEY, v int);
> INSERT INTO test (k, v) VALUES ('foo', 0);
> SELECT * FROM test WHERE k = 'foo';
  k  | v 
-----+---
 foo | 0
(1 row)
```

**go-sqlserver** currently supports the PostgreSQL queries in stages. See [PostgreSQL](doc/postgresql.md) for current support status.
