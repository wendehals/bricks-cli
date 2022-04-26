[![CI](https://github.com/wendehals/bricks-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/wendehals/bricks-cli/actions/workflows/ci.yml)

# bricks-cli

## bricks Commands

    bricks is a command line interface to the rebrickable.com API. It
    enables you to access the Rebrickable database to retrieve for example a list
    of all parts of a certain set or user specific part collections.

    bricks is also able to merge, sort, add, or subtract collections of parts
    to new collections.

    Usage:
    bricks [command]

    Available Commands:
    api         Groups all commands for the Rebrickable API.
    collection  Groups all commands for working with part collections.
    export      Exports the JSON input as HTML.
    help        Help about any command

    Flags:
    -h, --help      help for bricks
    -v, --version   version for bricks

    Use "bricks [command] --help" for more information about a command.

## bricks api Subcommands

    The api command groups all commands for the Rebrickable API.

    Usage:
    bricks api [command]

    Available Commands:
    details     Get details about a certain set, set list, or part list
    lists       Get the user's set or part lists
    lostParts   Get all lost parts of the user
    parts       Get a list of parts
    sets        Get details about sets owned by the user

    Flags:
    -c, --credentials string   A JSON file containing the Rebrickable credentials (default "credentials.json")
    -h, --help                 help for api
    -o, --json string          A name for the JSON output file

    Use "bricks api [command] --help" for more information about a command.

## bricks collection Subcommands

    The collection command groups all commands for working with part collections.

    Usage:
    bricks collection [command]

    Available Commands:
    max         Calculates the maximum quantity of each part of at least two collections.
    merge       Merges the parts of a collection by their color or by their variant.
    sort        Sorts the parts of a collection by their number.
    subtract    Subtracts one collection of parts from another.
    sum         Sums up the parts of multiple collections.

    Flags:
    -h, --help          help for collection
    -o, --json string   A name for the JSON output file

    Use "bricks collection [command] --help" for more information about a command.

## credentials.json

The content of the `credentials.json` file looks like:

    {
        "username": "<rebrickable username>",
        "password": "<rebrickable password>",
        "api_key": "<rebrickable api key>"
    }
