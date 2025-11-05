# Scripting

The bricks-cli supports executing scripts contained in *.bricks files. The scripts are made up of commands working with expressions describing part collections and build results. Have a look at the following example:

    a := allParts
    b := build(set(8880-1), a)
    save(a, "remainings_parts.json")
    save(b, "build_result.json")
    print(b)

The first command in the script above assigns the user's entire inventory (`allParts`) to the variable `a`. The second command attempts to build set `8880-1` using the parts from the `set(8880-1)` expression while consuming parts from the collection `a`; the build result is stored in variable `b`. Note that `build` has an explicit side effect and subtracts the used parts from the second collection, so `a` contains the remaining parts after the build. The third and fourth commands save the remaining parts (`a`) and the build result (`b`) to JSON files. The final command prints a short summary of the build result to the console.

## Commands

The following commands can be used in bricks-cli scripts:

* *Assignment*: `ID := COLLECTION_EXP` - assigns a collection of parts defined by the given collection expression COLLECTION_EXP to a variable with name ID. The variable can be used in any expression later on.
* *save*: `save(COLLECTION_EXP, FILE_NAME)` or `save(BUILD_EXP, FILE_NAME)` - saves the given collection or build result defined by the expression to a JSON file with the given FILE_NAME.
* *export*: `export(COLLECTION_EXP, DIR_NAME)` or `export(BUILD_EXP, DIR_NAME)` - exports the given collection or build result defined by the expression to an HTML file into the given DIR_NAME.
* *print*: `print(COLLECTION_EXP)` or `print(BUILD_EXP)` - prints information about the given collection or build result to the console.
* *pause*: `pause(INT)` - Pauses the script for the given time INT in seconds. Sometimes it's necessary to pause the script between API requests to avoid a "429 Too many requests" error message from the rebrickable.com server.

## Expressions

These expressions are available in the bricks-cli scripts:

### Collection Expressions

Collection expressions return part collections and can be used in assignments, save/export/print commands, and as inputs to other expressions:

* *Identifier* - Identifiers denote a variable defined by an assignment in advance. Identifiers start with a lower or upper case letter and can contain letters (lower and upper case), digits, and underscore `_`.
* `load(FILE_NAME)` - Loads a collection from a file with the given FILE_NAME.
* `import(FILE_NAME)` - Loads a collection from a Rebrickable CSV file with the given FILE_NAME.
* `allParts` - Denotes all parts owned by the user.
* `lost` - Denotes all parts lost by the user.
* `set(SET_NUM, BOOL?)` - Denotes all parts of set number SET_NUM (e.g. 8880-1). If the optional BOOL parameter is set to `true`, it includes all mini figures. If the BOOL parameter is not specified it defaults to `false`, which means that mini figures are not included.
* `userSet(SET_NUM, BOOL?)` - Denotes all parts of set number SET_NUM (e.g. 8880-1) owned by the user excluding lost parts of this set. If the optional BOOL parameter is set to `true`, it includes all mini figures. If the BOOL parameter is not specified it defaults to `false`, which means that mini figures are not included.
* `setList(LIST_ID, BOOL?)` - Denotes all parts contained in the set list LIST_ID (a number). If the optional BOOL parameter is set to `true`, it includes all mini figures. If the BOOL parameter is not specified it defaults to `false`, which means that mini figures are not included.
* `partList(LIST_ID)` - Denotes all parts contained in the part list LIST_ID (a number).
* `partLists(LIST_ID, BOOL?)` - Denotes all parts contained in the part list LIST_ID (a number). If the optional BOOL parameter is set to `true`, it also includes all part lists marked as non-buildable. If the BOOL parameter is not specified it defaults to `false`, which means that non-buildable part lists are not included.
* `sum(COLLECTION_EXP, COLLECTION_EXP, ...)` - Calculates the sum of all parts in the collection expressions' results.
* `subtract(COLLECTION_EXP, COLLECTION_EXP)` - Subtracts all parts in the second collection expression's result from the parts in the first collection expression's result.
* `max(COLLECTION_EXP, COLLECTION_EXP, ...)` - Calculates the maximum of all parts in the collection expressions' results.
* `sort(COLLECTION_EXP, quantity?, descending?)` - Sorts parts in the collection expression's result in ascending order by their color and name. If sorting should be done by quantity and name, add `quantity` to the argument list. For sorting in descending order, add `descending` as last argument.

### Build Expressions

Build expressions return build results and can be used with save/export/print commands:

* *Identifier* - Identifiers denoting a build variable defined by an assignment.
* `build(COLLECTION_EXP, COLLECTION_EXP, MODE?)` - Builds a set given by the first collection expression by using parts from the collection defined by the second collection expression. This operation has an explicit side effect: the parts that have been used for the build are subtracted from the second collection. So, if the second collection is given by an identifier, it can be reused, e.g. in further builds to check if another set can be built from the remaining parts. MODE is a combination of any of the following characters: C(olor), A(lternates), M(olds), and P(rints). If for example CAMP is provided, the build substitutes missing parts by parts of different color, alternates, molds, and prints. If no MODE is provided, no missing parts will be replaced.
