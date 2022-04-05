[![CI](https://github.com/wendehals/bricks-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/wendehals/bricks-cli/actions/workflows/ci.yml)

# bricks-cli

## bricks-cli Commands

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
    export      Exports the JSON input as HTML.
    help        Help about any command

    -h, --help      help for bricks
    -v, --version   version for bricks

    Use "bricks [command] --help" for more information about a command.


## bricks-cli api Subcommands

    The api command groups all commands for the Rebrickable API.

    bricks api [command]

    Available Commands:
    allParts      Get all parts owned by the user
    partListParts Get all parts of a user defined part list.
    partLists     Get a list of all the user's part lists
    setLists      Get a list of all the user's set lists
    setParts      Get all parts used in the given set(s)
    sets          Get a list of all the user's sets

    -c, --credentials string   A JSON file containing the Rebrickable credentials (default "credentials.json")
    -h, --help                 help for api
    -o, --json string          A name for the JSON output file

    Use "bricks api [command] --help" for more information about a command.


## bricks-cli collection Subcommands

    The collection command groups all commands for working with bricks collections.

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
