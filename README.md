DP Recipe API
=============

### Getting started

This repo contains 2 executables - the recipe API and a recipe checker app.
The recipe API can be run by running `make debug` and the recipe checker can be
run by running `make checker` (but has some prerequisites, documented below).

### Recipe API

#### Configuration

| Environment variable | Default                                   | Description
| -------------------- | ----------------------------------------- | -----------
| BIND_ADDR            | :22300                                    | The host and port to bind to

### Recipe Checker

#### Configuration

| Command line flag    | Default                                   | Description
| -------------------- | ----------------------------------------- | -----------
| bind                 | :2222                                     | The host and port to bind to
| dev                  | ""                                        | The host for the develop environment
| beta                 | ""                                        | The host for the production environment

Using the [Makefile](Makefile) target `make checker` the environment flags will not be set by default.
The Makefile reads these flags from local environment variables to prevent internal hostnames
being committed. The environment variables you will need to set are:

`export $CMD_DEV_API_HOST=<cmd-dev host>`
`export $CMD_API_HOST=<cmd host>`

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2016-2017, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
