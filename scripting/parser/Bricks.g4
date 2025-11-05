grammar Bricks;

bricks: command+;

command: assignment | save | export | print | pause;

assignment: ID ':=' collectionExp;

save: 'save' '(' (collectionExp | buildExp) ',' STRING ')';

export: 'export' '(' (collectionExp | buildExp) ',' STRING ')';

print: 'print' '(' (collectionExp | buildExp) ')';

pause: 'pause' '(' seconds=INT ')';


collectionExp: identifier | load | import_ | allParts | lost | set | userSet | setList | partList | partLists | sum | subtract | max | sort;

load: 'load' '(' STRING ')';

import_: 'import' '(' STRING ')';

allParts: 'allParts';

lost: 'lost';

set: 'set' '(' SET_NUM (',' BOOL)? ')';

userSet: 'userSet' '(' SET_NUM (',' BOOL)? ')';

setList: 'setList' '(' INT (',' BOOL)? ')';

partList: 'partList' '(' INT ')';

partLists: 'partLists' '(' STRING (',' BOOL)? ')';

sum: 'sum' '(' collectionExp (',' collectionExp)+ ')';

subtract: 'subtract' '(' collectionExp ',' collectionExp ')';

max: 'max' '(' collectionExp (',' collectionExp)+ ')';

sort: 'sort' '(' collectionExp (',' quantity='quantity')? (',' descending='descending')? ')';


buildExp : identifier | build;

build: 'build' '(' collectionExp ',' collectionExp (',' build_mode=ID )? ')';


ID: [a-zA-Z][a-zA-Z0-9_]*;

identifier: ID;

fragment DIGIT: ('0'..'9');

INT: DIGIT+;

SET_NUM: DIGIT+'-'DIGIT;

BOOL: 'true'|'false';

STRING : '"' ( '\\"' | . )*? '"' ;

WS: [ \t\r\n]+ -> skip;