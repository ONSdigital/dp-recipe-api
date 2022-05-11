DP Recipe API
=============

### Getting started

The recipe API can be run by running `make debug`.

#### Import recipe data locally

To import recipe's into your local MongoDB instance you can use the [import-recipes script](import-recipes/README.md)

### Health check

The endpoint `/health` checks the connection to the database and returns
one of:

* success (200, JSON "status": "OK")
* failure (500, JSON "status": "error").

### Recipe API

#### Configuration

| Environment variable           | Default                     | Description                                                                                            |
|--------------------------------|-----------------------------|--------------------------------------------------------------------------------------------------------|
| BIND_ADDR                      | :22300                      | The host and port to bind to                                                                           |
| MONGODB_BIND_ADDR              | localhost:27017             | The MongoDB bind address                                                                               |
| MONGODB_USERNAME               |                             | The MongoDB Username                                                                                   |
| MONGODB_PASSWORD               |                             | The MongoDB Password                                                                                   |
| MONGODB_DATABASE               | recipes                     | The MongoDB database                                                                                   |
| MONGODB_COLLECTIONS            | RecipesCollection:recipes   | The MongoDB collections                                                                                |
| MONGODB_REPLICA_SET            |                             | The name of the MongoDB replica set                                                                    |
| MONGODB_ENABLE_READ_CONCERN    | false                       | Switch to use (or not) majority read concern                                                           |
| MONGODB_ENABLE_WRITE_CONCERN   | true                        | Switch to use (or not) majority write concern                                                          |
| MONGODB_CONNECT_TIMEOUT        | 5s                          | The timeout when connecting to MongoDB (`time.Duration` format)                                        |
| MONGODB_QUERY_TIMEOUT          | 15s                         | The timeout for querying MongoDB (`time.Duration` format)                                              |
| MONGODB_IS_SSL                 | false                       | Switch to use (or not) TLS when connecting to mongodb                                                  |
| GRACEFUL_SHUTDOWN_TIMEOUT      | 5s                          | The graceful shutdown timeout in seconds                                                               |
| HEALTHCHECK_INTERVAL           | 30s                         | The time between calling healthcheck endpoints for check subsystems                                    |
| HEALTHCHECK_CRITICAL_TIMEOUT   | 90s                         | The time taken for the health changes from warning state to critical due to subsystem check failures   |
| ZEBEDEE_URL                    | http://localhost:8082       | The URL to Zebedee (for authentication)                                                                |

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2016-2022, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
