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
