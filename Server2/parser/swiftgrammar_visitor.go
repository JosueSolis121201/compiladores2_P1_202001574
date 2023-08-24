// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"



// A complete Visitor for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SwiftGrammarParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#instruction.
	VisitInstruction(ctx *InstructionContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#printstmt.
	VisitPrintstmt(ctx *PrintstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#ifstmt.
	VisitIfstmt(ctx *IfstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#elsestmt.
	VisitElsestmt(ctx *ElsestmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#else_ifstmt.
	VisitElse_ifstmt(ctx *Else_ifstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#declarationstmt.
	VisitDeclarationstmt(ctx *DeclarationstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#switchstmt.
	VisitSwitchstmt(ctx *SwitchstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#defaultsts.
	VisitDefaultsts(ctx *DefaultstsContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#case.
	VisitCase(ctx *CaseContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#asignacionstmt.
	VisitAsignacionstmt(ctx *AsignacionstmtContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#ig_type.
	VisitIg_type(ctx *Ig_typeContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#reserv.
	VisitReserv(ctx *ReservContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#types.
	VisitTypes(ctx *TypesContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#listParams.
	VisitListParams(ctx *ListParamsContext) interface{}

	// Visit a parse tree produced by SwiftGrammarParser#listArray.
	VisitListArray(ctx *ListArrayContext) interface{}
}
