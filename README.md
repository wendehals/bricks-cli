# bricks-cli

    bricks is command line interface to the Rebrickable.com API. It
    enables you to access the Rebrickable database to retrieve for example a list
    of all bricks of a certain set or user specific brick collections.

    bricks is also able to merge, sort, add, or subtract collections of bricks
    to new collections.

    Usage:
    bricks [command]

    Available Commands:
    help        Help about any command
    max         Calculates the maximum quantity of each part of multiple collections.
    merge       Merges the parts of a collection by their color or by their variant.
    partLists   Get a list of all the user's part lists
    setLists    Get a list of all the user's set lists
    setParts    Returns a list of parts used in the given set or sets
    sort        Sorts the parts of a collection by their number.
    subtract    Subtracts one collection of parts from another.
    sum         Sums up the parts of multiple collections.

    Flags:
    -c, --credentials string   A JSON file containing the Rebrickable credentials (default "credentials.json")
    -h, --help                 help for bricks
    -o, --html string          A name for the HTML output file (default "output.html")
    -j, --json string          A name for the JSON output file
    -v, --version              version for bricks

    Use "bricks [command] --help" for more information about a command.