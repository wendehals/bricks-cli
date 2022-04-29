grammar Bricks;

bricks: command+;

command: assignment | save | export;

assignment: ID ':=' exp;

save: 'save' '(' exp ',' STRING ')';

export: 'export' '(' exp ',' STRING ')';


exp: identifier | load | allParts | lost | set | setList | partList | sum | subtract | max | mergeByColor | mergeByVariant | sort;

identifier: ID;

load: 'load' '(' STRING ')';

allParts: 'allParts';

lost: 'lost';

set: 'set' '(' SET_NUM ')';

setList: 'setList' '(' INT ')';

partList: 'partList' '(' INT ')';

partLists: 'partLists' '(' STRING ')';

sum: 'sum' '(' exp (',' exp)+ ')';

subtract: 'subtract' '(' exp ',' exp ')';

max: 'max' '(' exp (',' exp)+ ')';

mergeByColor: 'mergeByColor' '(' exp ')';

mergeByVariant: 'mergeByVariant' '(' exp ')';

sort: 'sort' '(' exp ')';


fragment DIGIT: [0–9];

INT: DIGIT+;

STRING : '"' ( '\\"' | . )*? '"' ;

ID: [a-zA-Z][a-zA-Z0-9_]*;

SET_NUM: [0-9]+'-'[0-9];

WS: [ \t\r\n]+ -> skip;