// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type BaseSwiftGrammarListener struct{}

var _ SwiftGrammarListener = &BaseSwiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseSwiftGrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseSwiftGrammarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseSwiftGrammarListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseSwiftGrammarListener) ExitInstruction(ctx *InstructionContext) {}

// EnterAsignacion_vectorstmt is called when production asignacion_vectorstmt is entered.
func (s *BaseSwiftGrammarListener) EnterAsignacion_vectorstmt(ctx *Asignacion_vectorstmtContext) {}

// ExitAsignacion_vectorstmt is called when production asignacion_vectorstmt is exited.
func (s *BaseSwiftGrammarListener) ExitAsignacion_vectorstmt(ctx *Asignacion_vectorstmtContext) {}

// EnterCountstmt is called when production countstmt is entered.
func (s *BaseSwiftGrammarListener) EnterCountstmt(ctx *CountstmtContext) {}

// ExitCountstmt is called when production countstmt is exited.
func (s *BaseSwiftGrammarListener) ExitCountstmt(ctx *CountstmtContext) {}

// EnterIsempystmt is called when production isempystmt is entered.
func (s *BaseSwiftGrammarListener) EnterIsempystmt(ctx *IsempystmtContext) {}

// ExitIsempystmt is called when production isempystmt is exited.
func (s *BaseSwiftGrammarListener) ExitIsempystmt(ctx *IsempystmtContext) {}

// EnterRemovestmt is called when production removestmt is entered.
func (s *BaseSwiftGrammarListener) EnterRemovestmt(ctx *RemovestmtContext) {}

// ExitRemovestmt is called when production removestmt is exited.
func (s *BaseSwiftGrammarListener) ExitRemovestmt(ctx *RemovestmtContext) {}

// EnterRemovelaststmt is called when production removelaststmt is entered.
func (s *BaseSwiftGrammarListener) EnterRemovelaststmt(ctx *RemovelaststmtContext) {}

// ExitRemovelaststmt is called when production removelaststmt is exited.
func (s *BaseSwiftGrammarListener) ExitRemovelaststmt(ctx *RemovelaststmtContext) {}

// EnterAppendstmt is called when production appendstmt is entered.
func (s *BaseSwiftGrammarListener) EnterAppendstmt(ctx *AppendstmtContext) {}

// ExitAppendstmt is called when production appendstmt is exited.
func (s *BaseSwiftGrammarListener) ExitAppendstmt(ctx *AppendstmtContext) {}

// EnterGuardstmt is called when production guardstmt is entered.
func (s *BaseSwiftGrammarListener) EnterGuardstmt(ctx *GuardstmtContext) {}

// ExitGuardstmt is called when production guardstmt is exited.
func (s *BaseSwiftGrammarListener) ExitGuardstmt(ctx *GuardstmtContext) {}

// EnterDeclaracion_vectorstmt is called when production declaracion_vectorstmt is entered.
func (s *BaseSwiftGrammarListener) EnterDeclaracion_vectorstmt(ctx *Declaracion_vectorstmtContext) {}

// ExitDeclaracion_vectorstmt is called when production declaracion_vectorstmt is exited.
func (s *BaseSwiftGrammarListener) ExitDeclaracion_vectorstmt(ctx *Declaracion_vectorstmtContext) {}

// EnterDeclaracion_vectorstmt_SUPP is called when production declaracion_vectorstmt_SUPP is entered.
func (s *BaseSwiftGrammarListener) EnterDeclaracion_vectorstmt_SUPP(ctx *Declaracion_vectorstmt_SUPPContext) {
}

// ExitDeclaracion_vectorstmt_SUPP is called when production declaracion_vectorstmt_SUPP is exited.
func (s *BaseSwiftGrammarListener) ExitDeclaracion_vectorstmt_SUPP(ctx *Declaracion_vectorstmt_SUPPContext) {
}

// EnterListaval is called when production listaval is entered.
func (s *BaseSwiftGrammarListener) EnterListaval(ctx *ListavalContext) {}

// ExitListaval is called when production listaval is exited.
func (s *BaseSwiftGrammarListener) ExitListaval(ctx *ListavalContext) {}

// EnterJumpstmt is called when production jumpstmt is entered.
func (s *BaseSwiftGrammarListener) EnterJumpstmt(ctx *JumpstmtContext) {}

// ExitJumpstmt is called when production jumpstmt is exited.
func (s *BaseSwiftGrammarListener) ExitJumpstmt(ctx *JumpstmtContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *BaseSwiftGrammarListener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *BaseSwiftGrammarListener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterForstmt is called when production forstmt is entered.
func (s *BaseSwiftGrammarListener) EnterForstmt(ctx *ForstmtContext) {}

// ExitForstmt is called when production forstmt is exited.
func (s *BaseSwiftGrammarListener) ExitForstmt(ctx *ForstmtContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseSwiftGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseSwiftGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterElsestmt is called when production elsestmt is entered.
func (s *BaseSwiftGrammarListener) EnterElsestmt(ctx *ElsestmtContext) {}

// ExitElsestmt is called when production elsestmt is exited.
func (s *BaseSwiftGrammarListener) ExitElsestmt(ctx *ElsestmtContext) {}

// EnterElse_ifstmt is called when production else_ifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterElse_ifstmt(ctx *Else_ifstmtContext) {}

// ExitElse_ifstmt is called when production else_ifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitElse_ifstmt(ctx *Else_ifstmtContext) {}

// EnterDeclarationstmt is called when production declarationstmt is entered.
func (s *BaseSwiftGrammarListener) EnterDeclarationstmt(ctx *DeclarationstmtContext) {}

// ExitDeclarationstmt is called when production declarationstmt is exited.
func (s *BaseSwiftGrammarListener) ExitDeclarationstmt(ctx *DeclarationstmtContext) {}

// EnterSwitchstmt is called when production switchstmt is entered.
func (s *BaseSwiftGrammarListener) EnterSwitchstmt(ctx *SwitchstmtContext) {}

// ExitSwitchstmt is called when production switchstmt is exited.
func (s *BaseSwiftGrammarListener) ExitSwitchstmt(ctx *SwitchstmtContext) {}

// EnterDefaultsts is called when production defaultsts is entered.
func (s *BaseSwiftGrammarListener) EnterDefaultsts(ctx *DefaultstsContext) {}

// ExitDefaultsts is called when production defaultsts is exited.
func (s *BaseSwiftGrammarListener) ExitDefaultsts(ctx *DefaultstsContext) {}

// EnterCase is called when production case is entered.
func (s *BaseSwiftGrammarListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production case is exited.
func (s *BaseSwiftGrammarListener) ExitCase(ctx *CaseContext) {}

// EnterAsignacionstmt is called when production asignacionstmt is entered.
func (s *BaseSwiftGrammarListener) EnterAsignacionstmt(ctx *AsignacionstmtContext) {}

// ExitAsignacionstmt is called when production asignacionstmt is exited.
func (s *BaseSwiftGrammarListener) ExitAsignacionstmt(ctx *AsignacionstmtContext) {}

// EnterIg_type is called when production ig_type is entered.
func (s *BaseSwiftGrammarListener) EnterIg_type(ctx *Ig_typeContext) {}

// ExitIg_type is called when production ig_type is exited.
func (s *BaseSwiftGrammarListener) ExitIg_type(ctx *Ig_typeContext) {}

// EnterReserv is called when production reserv is entered.
func (s *BaseSwiftGrammarListener) EnterReserv(ctx *ReservContext) {}

// ExitReserv is called when production reserv is exited.
func (s *BaseSwiftGrammarListener) ExitReserv(ctx *ReservContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseSwiftGrammarListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseSwiftGrammarListener) ExitTypes(ctx *TypesContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSwiftGrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSwiftGrammarListener) ExitExpr(ctx *ExprContext) {}

// EnterListParams is called when production listParams is entered.
func (s *BaseSwiftGrammarListener) EnterListParams(ctx *ListParamsContext) {}

// ExitListParams is called when production listParams is exited.
func (s *BaseSwiftGrammarListener) ExitListParams(ctx *ListParamsContext) {}

// EnterListArray is called when production listArray is entered.
func (s *BaseSwiftGrammarListener) EnterListArray(ctx *ListArrayContext) {}

// ExitListArray is called when production listArray is exited.
func (s *BaseSwiftGrammarListener) ExitListArray(ctx *ListArrayContext) {}
