grammar Bricks;

bricks: command+;

command: assignment | save | export | build;

assignment: ID ':=' exp;

save: 'save' '(' exp ',' STRING ')';

export: 'export' '(' exp ',' STRING ')';

build: 'build' '(' exp ',' exp ',' STRING ')';


exp: identifier | load | allParts | lost | set | setList | partList | partLists | sum | subtract | max | sort;

identifier: ID;

load: 'load' '(' STRING ')';

allParts: 'allParts';

lost: 'lost';

set: 'set' '(' SET_NUM (',' BOOL)?')';

setList: 'setList' '(' INT (',' BOOL)?')';

partList: 'partList' '(' INT ')';

partLists: 'partLists' '(' STRING (',' BOOL)?')';

sum: 'sum' '(' exp (',' exp)+ ')';

subtract: 'subtract' '(' exp ',' exp ')';

max: 'max' '(' exp (',' exp)+ ')';

sort: 'sort' '(' exp (',' quantity='quantity')? (',' descending='descending')? ')';


fragment DIGIT: [0–9];

INT: DIGIT+;

BOOL: 'true'|'false';

STRING : '"' ( '\\"' | . )*? '"' ;

ID: [a-zA-Z][a-zA-Z0-9_]*;

SET_NUM: [0-9]+'-'[0-9];

WS: [ \t\r\n]+ -> skip;