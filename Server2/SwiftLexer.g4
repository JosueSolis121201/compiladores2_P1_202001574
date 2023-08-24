lexer grammar SwiftLexer;

// --------------- Tokens
// types;
INT:    'Int';
FLOAT:  'Float';
BOOL:   'Bool';
STR:    'String';
CHAR:   'Character';

// reserved words
TRU:    'true';
FAL:    'false';
PRINT:  'print';
IF:     'if';
ELSE:   'else';
ELSE_IF:   'else if';
WHILE:  'while';
FOR:    'for';
IN:    'in';
GUARD:    'guard';
CONTINUE:    'continue';
RETURN:    'return';
SWITCH: 'switch';
BREAK: 'break';
CASE:   'case';
VAR:    'var';
LET:    'let';
D_APPEND:    '.append';
D_REMOVE:    '.remove';
REPEATING:    'repeating';
DEFAULT:    'default';
COUNT:    'count';
STRUCT:    'STRUCT';
MUTATING:    'mutating';
FUNC:    'func';
SELF_D:    'self.';
PUNTO:    '.';
INOUT:    'inout';
GUION_B:    '_';

// primitives
NUMBER : [0-9]+ ('.'[0-9]+)?;
STRING: '"'~["]*'"';
ID: ([a-zA-Z])[a-zA-Z0-9_]*;

// symbols

DIF:            '!=';
IG_IG:          '==';
MAS_IGUAL:      '+=';
MENOS_IGUAL:    '-=';
NOT:            '!';
OR:             '||';
AND:            '&&';
SOLO_AND:        '&';
IG:             '=';
MAY_IG:         '>=';
MEN_IG:         '<=';
MAYOR:          '>';
MENOR:          '<';
MUL:            '*';
DIV:            '/';
ADD:            '+';
MODULO:         '%';
SUB:            '-';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
D_PTS:          ':';
CORIZQ:         '[';
CORDER:         ']';
COMA:           ',';
PUNTO_COMA:      ';';

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;

