# Scripting

The bricks-cli supports executing scripts contained in *.bricks files. The scripts are made up of commands working with expressions describing part collections. Have a look at the following example:

    a := set(8880-1)
    save(sort(subtract(allParts, a)))

The first command in the script above retrieves all parts of set 8880-1 and stores the resulting collection in the variable `a`. The next command first retrieves all parts owned by the user, subtracts those parts from the collection saved in `a`, and sorts the parts. The overall result of the expression is saved into a JSON file.

## Commands

The following commands can be used in bricks scripts:

* *Assignment*: `ID := exp` - assigns a collection of parts defined by the given expression to a variable with name ID. The variable can be used in any expression later on.
* *save*: `save(exp, FILE_NAME)` - saves the given collection defined by the expression to a JSON file with the given FILE_NAME.
* *export*: `export(exp, EXPORT_DIR)` - exports the given collection defined by the expression to an HTML file into the given EXPORT_DIR.
* *build*: `build(EXP, EXP, DIR_NAME, MODE?)` - Builds a set given by the first expression by using parts from the collection defined by the second expression. MODE is a combination of any of the following characters: C(olor), A(lternates), M(olds), and P(rints). If for example CAMP is provided, the build substitutes missing parts by parts of different color, alternates, molds, and prints. If no MODE is provided, no missing parts will be replaced. Exports the result to the given directory.
* *pause*: `pause(INT)` - Pauses the script for the given time in seconds. Sometimes it's neccessary to pause the script between API requests to avoid a "429 Too many requests" error message from the rebrickable.com server.

## Expressions

These expressions are available in the bricks scripts:

* *Identifier* - Identifiers denote a variable defined by an assignment in advance. Identifiers start with a lower or upper case letter and can contain letters (lower and upper case), digits, and underscore `_`.
* `load(FILE_NAME)` - Loads a collection from a file with the given FILE_NAME.
* `import(FILE_NAME)` - Loads a collection from a Rebrickable CSV file with the given FILE_NAME.
* `allParts` - Denotes all parts owned by the user.
* `lost` - Denotes all parts lost by the user.
* `set(SET_NUM, BOOL?)` - Denotes all parts of set number SET_NUM (e.g. 8880-1). If the optional BOOL paramater is set to `true`, it includes all mini figures. If the BOOL parameter is not specified it defaults to `false`, which means that mini figures are not included.
* `userSet(SET_NUM, BOOL?)` - Denotes all parts of set number SET_NUM (e.g. 8880-1) owned by the user excluding lost parts of this set. If the optional BOOL paramater is set to `true`, it includes all mini figures. If the BOOL parameter is not specified it defaults to `false`, which means that mini figures are not included.
* `setList(LIST_ID, BOOL?)` - Denotes all parts contained in the set list LIST_ID (a number). If the optional BOOL paramater is set to `true`, it includes all mini figures. If the BOOL parameter is not specified it defaults to `false`, which means that mini figures are not included.
* `partList(LIST_ID)` - Denotes all parts contained in the part list LIST_ID (a number).
* `partLists(LIST_ID, BOOL?)` - Denotes all parts contained in the part list LIST_ID (a number). If the optional BOOL paramater is set to `true`, it also includes all part lists marked as non-buildable. If the BOOL parameter is not specified it defaults to `false`, which means that non-buildable part lists are not included.
* `sum(EXP, EXP, ...)` - Calculates the sum of all parts in the expression's results.
* `subtract(EXP, EXP)` - Subtracts all parts in the second expression's result from the parts in the first expression's result.
* `max(EXP, EXP, ...)` - Calculates the maximum of all parts in the expression's results.
* `sort(EXP, quantity?, descending?)` - Sorts parts in the expression's result in ascending order by their color and name. If sorting should be done by quantity and name, add `quantity` to the argument list. For sorting in descending order, add `descending` as last argument.
