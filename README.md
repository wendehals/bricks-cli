[![CI](https://github.com/wendehals/bricks-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/wendehals/bricks-cli/actions/workflows/ci.yml)

# bricks-cli

    bricks is a command line interface to the Rebrickable.com API. It
    enables you to access the Rebrickable database to retrieve for example a list
    of all bricks of a certain set or user specific brick collections.

    bricks is also able to merge, sort, add, or subtract collections of bricks
    to new collections.

    Usage:
    bricks [command]

    Available Commands:
    api         Groups all commands for the Rebrickable API.
    collection  Groups all commands for working with bricks collections.
    help        Help about any command

    Flags:
    -h, --help      help for bricks
    -v, --version   version for bricks

    Use "bricks [command] --help" for more information about a command.

## credentials.json

The content of the `credentials.json` file looks like:

    {
        "username": "<rebrickable username>",
        "password": "<rebrickable password>",
        "api_key": "<rebrickable api key>"
    }
