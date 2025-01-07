# Configuring go-sqlserver

**go-sqlserver** is configured using a configuration file. The configuration file is a YAML file, and you can override the configuration by setting environment variables.

## Configuration File (go-sqlserver.yaml)

The configuration file is divided into sections. Each section is a YAML map. **go-sqlserver** will activate a default configuration if a configuration file is not specified or if there is no go-sqlserver.yaml in the local directory. The following is the default configuration file:

    logger:
      enabled: true
      level: info
    tls:
      enabled: false
      key_file: key.pem
      cert_file: cert.pem
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

## Environment Variables

You can override the configuration file location by setting the **go-sqlserver** environment variable. **go-sqlserver** assumes that the environment variable matches the following format: SQLSERVER + "\_" + the key name in ALL CAPS.

For example, if the environment variable `SQLSERVER_LOGGING_ENABLED` is set, then **go-sqlserver** will override the `logging:enabled` setting.
