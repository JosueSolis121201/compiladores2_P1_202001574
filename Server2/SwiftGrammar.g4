grammar SwiftGrammar; 
// import SwiftLexer; 

options {
  tokenVocab = SwiftLexer;
}

@header {
    import "Server2/interfaces"
    import "Server2/environment"
    import "Server2/expressions"
    import "Server2/instructions"
    import "strings"
}


s returns [[]interface{} code]
: block EOF
    {   
        $code = $block.blk
    }
;

block returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listInt []IInstructionContext
  }
: ins+=instruction+
    {
        listInt = localctx.(*BlockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;

instruction returns [interfaces.Instruction inst]
: printstmt { $inst = $printstmt.prnt}
| ifstmt { $inst = $ifstmt.ifinst }
| declarationstmt { $inst = $declarationstmt.dec }
| asignacionstmt {}
| switchstmt {}
| whilestmt {}
| forstmt {}
| guardstmt {}
| jumpstmt {}
| declaracion_vectorstmt {}
| appendstmt {}
| removelaststmt {}
| removestmt {}
| isempystmt {}
| countstmt {}
| asignacion_vectorstmt {}
| declaracion_matrissts {}
| structstmt {}
| fuctionstmt {}
;

fuctionstmt returns []
: FUNC ID PARIZQ parametros PARDER SUB MAYOR types LLAVEIZQ block LLAVEDER {}
;

parametros returns []
: parametros COMA ID ID D_PTS types  {}
| ID ID D_PTS types  {}
|   
;


structstmt returns []
: STRUCT ID LLAVEIZQ block LLAVEDER {}
;



declaracion_matrissts returns []
: types ID D_PTS  IG declaracion_matrissts_supp {}
;

declaracion_matrissts_supp returns []
: typematris PARIZQ REPEATING D_PTS typematris {}
;

typematris returns []
: corizmatris types 
;




corizmatris returns []
: corizmatris CORIZQ  {}
|CORIZQ
;

asignacion_vectorstmt returns []
: ID CORIZQ NUMBER CORDER IG ID  CORIZQ NUMBER CORDER  { }
;
countstmt returns []
: ID D_COUNT { }
;

isempystmt returns []
: ID D_ISEMPTY { }
;

removestmt returns []
: ID D_REMOVE PARIZQ AT_DD expr PARDER { }
;

removelaststmt returns []
: ID D_REMOVELAST PARIZQ PARDER { }
;

appendstmt returns []
: ID D_APPEND PARIZQ expr PARDER { }
;
guardstmt returns []
: GUARD expr ELSE LLAVEIZQ block  LLAVEDER { }
;

declaracion_vectorstmt returns []
: reserv ID D_PTS CORIZQ  types CORDER IG declaracion_vectorstmt_SUPP { }
;

declaracion_vectorstmt_SUPP returns []
:  CORIZQ  listaval CORDER {}
| ID
| types CORIZQ CORDER
| PARIZQ PARDER
;


listaval returns []
: listaval COMA expr  {}
| expr  {}
|
;

jumpstmt returns []
: CONTINUE {}
| RETURN {}
| BREAK  {}
;

whilestmt returns []
: WHILE expr LLAVEIZQ block  LLAVEDER { }
;

forstmt returns []
: FOR ID IN expr LLAVEIZQ block LLAVEDER  { }
;

printstmt returns [interfaces.Instruction prnt]
: PRINT PARIZQ expr PARDER { $prnt = instructions.NewPrint($PRINT.line,$PRINT.pos,$expr.e)}
;

ifstmt returns [interfaces.Instruction ifinst]
: IF expr LLAVEIZQ block LLAVEDER else_ifstmt elsestmt {}
;

elsestmt returns []
: ELSE LLAVEIZQ  block LLAVEDER  {}
| 
;

else_ifstmt returns []
:  else_ifstmt ELSE_IF expr LLAVEIZQ block LLAVEDER   {}
| 
;

declarationstmt returns [interfaces.Instruction dec]
: reserv ID D_PTS types IG expr  {}
| reserv ID IG expr  {}
;


switchstmt returns []
: SWITCH expr LLAVEIZQ case defaultsts LLAVEDER   {}
;

defaultsts returns []
:  DEFAULT D_PTS  block  {}
|  
; 


case returns []
: case CASE  expr D_PTS block   {}
| 
;


asignacionstmt returns []
: ID ig_type expr {}
;

ig_type returns []
: IG {}
| IG_IG {}
| MAS_IGUAL {}
| MENOS_IGUAL {}
;


reserv returns []
: VAR {}
| LET {}
;

types returns[environment.TipoExpresion ty]
: INT { $ty = environment.INTEGER }
| FLOAT { $ty = environment.FLOAT }
| STR { $ty = environment.STRING }
| BOOL { $ty = environment.BOOLEAN }
| CHAR {}
| CORIZQ CORDER { $ty = environment.ARRAY }
;

expr returns [interfaces.Expression e]
: left=expr op=(MUL|DIV) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(ADD|SUB) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MEN_IG|MENOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=AND right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=OR right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=MAS_IGUAL right=expr {}
| left=expr op=MENOS_IGUAL right=expr {}
| left=expr op=MODULO right=expr {}
| left=expr op=MODULO right=expr {}
| ID D_APPEND {}
| ID D_REMOVE {}
| ID D_REMOVELAST {}
| ID D_ISEMPTY {}
| ID D_COUNT {}
| ID CORIZQ NUMBER CORDER {}
| list=listArray { $e = $list.p}
| CORIZQ listParams CORDER { $e = expressions.NewArray($CORIZQ.line, $CORIZQ.pos, $listParams.l) }
| NUMBER                             
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| STRING
    {
        str := $STRING.text
        $e = expressions.NewPrimitive($STRING.line, $STRING.pos, str[1:len(str)-1],environment.STRING)
    }                        
| TRU { $e = expressions.NewPrimitive($TRU.line, $TRU.pos, true, environment.BOOLEAN) }
| FAL { $e = expressions.NewPrimitive($FAL.line, $FAL.pos, false, environment.BOOLEAN) }
| GUION_B
;

listParams returns[[]interface{} l]
: list=listParams COMA expr {
                                var arr []interface{}
                                arr = append($list.l, $expr.e)
                                $l = arr
                            }   
| expr {
            $l = []interface{}{}
            $l = append($l, $expr.e)
        }
;

listArray returns[interfaces.Expression p]
: list = listArray CORIZQ expr CORDER { $p = expressions.NewArrayAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $expr.e) }
| ID { $p = expressions.NewCallVar($ID.line, $ID.pos, $ID.text)}
;