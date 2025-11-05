grammar Bricks;

bricks: command+;

command: assignment | save | export | print | pause;

expression: collection | build | identifier;

assignment: ID ':=' expression;

save: 'save' '(' expression ',' STRING ')';

export: 'export' '(' expression ',' STRING ')';

print: 'print' '(' expression ')';

pause: 'pause' '(' seconds=INT ')';


collectionOrId: collection | identifier;

collection: load | import_ | allParts | lost | set | userSet | setList | partList | partLists | sum | subtract | max | sort;

load: 'load' '(' STRING ')';

import_: 'import' '(' STRING ')';

allParts: 'allParts';

lost: 'lost';

set: 'set' '(' SET_NUM (',' BOOL)? ')';

userSet: 'userSet' '(' SET_NUM (',' BOOL)? ')';

setList: 'setList' '(' INT (',' BOOL)? ')';

partList: 'partList' '(' INT ')';

partLists: 'partLists' '(' STRING (',' BOOL)? ')';

sum: 'sum' '(' collectionOrId (',' collectionOrId)+ ')';

subtract: 'subtract' '(' collectionOrId ',' collectionOrId ')';

max: 'max' '(' collectionOrId (',' collectionOrId)+ ')';

sort: 'sort' '(' collectionOrId (',' quantity='quantity')? (',' descending='descending')? ')';


build: 'build' '(' collectionOrId ',' collectionOrId (',' build_mode=ID )? ')';


ID: [a-zA-Z][a-zA-Z0-9_]*;

identifier: ID;

fragment DIGIT: ('0'..'9');

INT: DIGIT+;

SET_NUM: DIGIT+'-'DIGIT;

BOOL: 'true'|'false';

STRING : '"' ( '\\"' | . )*? '"' ;

WS: [ \t\r\n]+ -> skip;