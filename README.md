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
    script      Executes a bricks script

    Flags:
    -h, --help      help for bricks
        --verbose   Verbose output
    -v, --version   version for bricks

    Use "bricks [command] --help" for more information about a command.

## bricks api Subcommands

    The api command groups all commands for the Rebrickable API.

    Usage:
    bricks api [command]

    Available Commands:
    details     Get details about a certain set, set list, or part list
    lists       Get the user's set or part lists
    parts       Get a list of parts
    sets        Get details about sets owned by the user

    Flags:
    -c, --credentials string   A JSON file containing the Rebrickable credentials (default "credentials.json")
    -h, --help                 help for api
    -o, --output string        A name for the JSON output file

    Global Flags:
        --verbose   Verbose output

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
    -h, --help            help for collection
    -o, --output string   A name for the JSON output file

    Global Flags:
        --verbose   Verbose output

    Use "bricks collection [command] --help" for more information about a command.

## credentials.json

The content of the `credentials.json` file looks like:

    {
        "username": "<rebrickable username>",
        "password": "<rebrickable password>",
        "api_key": "<rebrickable api key>"
    }

## Scripting

The bricks-cli supports executing scripts contained in *.bricks files. The scripts are made up of commands working with expressions describing part collections. Have a look at the following example:

    save(sort(subtract(allParts, set(8880-1))))

The script above first retrieves all parts owned by the user and subtracts those parts contained in set 8880-1. After that it sorts the parts. The overall result of the expression is saved into a JSON file.

### Commands

The following commands can be used in bricks scripts:

* *Assignment*: `ID := exp` - assign a collection of parts defined by the given expression to a variable with name ID. The variable can be used in any expression later on.
* *Save*: `save(exp, FILE_PATH)` - save the given collection defined by the expression to a JSON file with path FILE_PATH.
* *export*: `export(exp, FILE_PATH)` - export the given collection defined by the expression to an HTML file with path FILE_PATH.

### Expressions

These expressions are available in the bricks scripts:

* *Identifier* - Identifiers denote a variable defined by an assignment in advance.
* `load(FILE_PATH)` - Loads a collection from the given file with path FILE_PATH.
* `allParts` - All parts owned by the user.
* `lost` - All parts lost by the user.
* `set(SET_NUM, BOOL?)` - All parts of set number SET_NUM (e.g. 8880-1). If the optional BOOL paramater is set to true, it includes all mini figures.
* `setList(LIST_ID, BOOL?)` - All parts contained in the set list LIST_ID (a number). If the optional BOOL paramater is set to true, it includes all mini figures.
* `partList(LIST_ID)` - All parts contained in the part list LIST_ID (a number).
* `partLists(LIST_ID, BOOL?)` - All parts contained in the part list LIST_ID (a number). If the optional BOOL paramater is set to true, it also includes all part lists marked as non-buildable.
* `sum(EXP, EXP, ...)` - Calculate the sum of all parts in the expression's results.
* `subtract(EXP, EXP)` - Subtract all parts in the second expression's result from the parts in the first expression's result.
* `max(EXP, EXP, ...)` - Calculate the maximum of all parts in the expression's results.
* `mergeByColor(EXP)` - Merge all parts in the expression's result by ignoring their color.
* `mergeByVariant(EXP)` - Merge all parts in the expression's result by replacing parts by their variants.
* `sort(EXP)` - Sort parts in the expression's result.
