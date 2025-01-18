# Configuring go-sqlserver

**go-sqlserver** is configured using a YAML configuration file. You can override settings by defining environment variables.

## Configuration File (go-sqlserver.yaml)

The configuration file is divided into sections, each represented as a YAML map. If no configuration file is specified or `go-sqlserver.yaml` is missing from the local directory, **go-sqlserver** will activate the default configuration. Below is an example of the default configuration file:

    logger:
      enabled: true
      level: info
    tls:
      enabled: false
      key_file: key.pem
      cert_file: cert.pem
      ca_files: [ca.pem]
    auth:
      enabled: false
      plain:
        -
          username: admin
          password: password
    query:
      mysql:
        port: 3306
      postgresql:
        port: 5432
    store:
      sqlite:
        memory: true
    metrics:
      prometheus:
        enabled: true
        port: 9181

### store.sqlite.memory

By default, **go-sqlserver** uses an in-memory SQLite database. To switch to a file-based SQLite database, set the `store.sqlite.memory` option to `false`.

## Environment Variables

The location of the configuration file can be overridden by setting an environment variable. **go-sqlserver** expects environment variables to follow the format: `GO_SQLSERVER_` + the key name in uppercase.

For example, setting the environment variable `GO_SQLSERVER_LOGGER_ENABLED` will override the `logging:enabled` configuration in the file.
