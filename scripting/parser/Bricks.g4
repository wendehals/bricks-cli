grammar Bricks;

bricks: exp;

exp: allParts | lost | set | setList | partList | sum | subtract | max | mergeByColor | mergeByVariant | sort;

allParts: 'allParts';

lost: 'lost';

set: 'set' '(' SET_NUM ')';

setList: 'setList' '(' INT ')';

partList: 'partList' '(' INT ')';

sum: 'sum' '(' exp (',' exp)+ ')';

subtract: 'subtract' '(' exp ',' exp ')';

max: 'max' '(' exp (',' exp)+ ')';

mergeByColor: 'mergeByColor' '(' exp ')';

mergeByVariant: 'mergeByVariant' '(' exp ')';

sort: 'sort' '(' exp ')';

fragment DIGIT: [0–9];
INT: DIGIT+;
SET_NUM: [0-9]+'-'[0-9];
WS: [ \t\r\n]+ -> skip;