# Commands

    bricks is a command line interface to the rebrickable.com API. It
    enables you to access the Rebrickable database to retrieve for example a list
    of all parts of a certain set or user specific part collections.

    bricks is also able to merge, sort, add, or subtract collections of parts    
    to new collections.

    Usage:
          bricks [command]

    Available Commands:
        api         Groups all commands for the Rebrickable API
        build       Calculates needed parts from a given collection to build a set
        collection  Groups all commands for working with part collections
        download    Groups all download commands
        export      Exports the parts input file as HTML
        help        Help about any command
        script      Executes a bricks script

    Flags:
        -h, --help      help for bricks
            --verbose   verbose output
        -v, --version   version for bricks

    Use "bricks [command] --help" for more information about a command.

## bricks api Subcommands

    The api command groups all sub commands for the Rebrickable API.

    Usage:
        bricks api [command]

    Available Commands:
        details     Get details about a certain set, set list, or part list
        lists       Get the user's set or part lists
        parts       Get a list of parts
        sets        Get details about sets owned by the user

    Flags:
        -c, --credentials string   a JSON file containing the Rebrickable credentials (default "credentials.json")
        -h, --help                 help for api
        -o, --output string        a name for the output file

    Global Flags:
        --verbose   verbose output

    Use "bricks api [command] --help" for more information about a command.

## bricks collection Subcommands

    The collection command groups all sub commands for working with part collections.     

    Usage:
        bricks collection [command]

    Available Commands:
        max         Calculates the maximum quantity of each part of at least two collections
        sort        Sorts the parts of a collection by their number
        subtract    Subtracts one collection of parts from another
        sum         Sums up the parts of multiple collections

    Flags:
        -h, --help            help for collection
        -o, --output string   a name for the output file

    Global Flags:
        --verbose   verbose output

    Use "bricks collection [command] --help" for more information about a command.

## bricks download Subcommands

    The download command groups all sub commands to download Rebrickable data for offline usage.

    Usage:
        bricks download [command]

    Available Commands:
        images        Downloads part images
        relationships Downloads part relationships

    Flags:
        -h, --help     help for download
        -u, --update   update (overwrite) already downloaded files

    Global Flags:
        --verbose   verbose output

    Use "bricks download [command] --help" for more information about a command.
