DP Recipe API
=============

### Getting started

This repo contains 2 executables - the recipe API and a recipe checker app.
The recipe API can be run by running `make debug` and the recipe checker can be
run by running `make checker` (but has some prerequisites, documented below).

### Healthcheck

The endpoint `/health` checks the connection to the database and returns
one of:

* success (200, JSON "status": "OK")
* failure (500, JSON "status": "error").

### Recipe API

#### Configuration

| Environment variable         | Default                                | Description
| ---------------------------- | ---------------------------------------| -----------
| BIND_ADDR                    | :22300                                 | The host and port to bind to
| MONGODB_BIND_ADDR            | localhost:27017                        | The MongoDB bind address
| MONGODB_DATABASE             | recipes                                | The MongoDB dataset database
| MONGODB_COLLECTION           | recipes                                | MongoDB collection
| GRACEFUL_SHUTDOWN_TIMEOUT    | 5s                                     | The graceful shutdown timeout in seconds
| HEALTHCHECK_INTERVAL         | 30s                                    | The time between calling healthcheck endpoints for check subsystems
| HEALTHCHECK_CRITICAL_TIMEOUT | 90s                                    | The time taken for the health changes from warning state to critical due to subsystem check failures

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2016-2020, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
