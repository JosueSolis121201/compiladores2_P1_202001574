// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterElsestmt is called when entering the elsestmt production.
	EnterElsestmt(c *ElsestmtContext)

	// EnterElse_ifstmt is called when entering the else_ifstmt production.
	EnterElse_ifstmt(c *Else_ifstmtContext)

	// EnterDeclarationstmt is called when entering the declarationstmt production.
	EnterDeclarationstmt(c *DeclarationstmtContext)

	// EnterSwitchstmt is called when entering the switchstmt production.
	EnterSwitchstmt(c *SwitchstmtContext)

	// EnterDefaultsts is called when entering the defaultsts production.
	EnterDefaultsts(c *DefaultstsContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterAsignacionstmt is called when entering the asignacionstmt production.
	EnterAsignacionstmt(c *AsignacionstmtContext)

	// EnterIg_type is called when entering the ig_type production.
	EnterIg_type(c *Ig_typeContext)

	// EnterReserv is called when entering the reserv production.
	EnterReserv(c *ReservContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterListParams is called when entering the listParams production.
	EnterListParams(c *ListParamsContext)

	// EnterListArray is called when entering the listArray production.
	EnterListArray(c *ListArrayContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitElsestmt is called when exiting the elsestmt production.
	ExitElsestmt(c *ElsestmtContext)

	// ExitElse_ifstmt is called when exiting the else_ifstmt production.
	ExitElse_ifstmt(c *Else_ifstmtContext)

	// ExitDeclarationstmt is called when exiting the declarationstmt production.
	ExitDeclarationstmt(c *DeclarationstmtContext)

	// ExitSwitchstmt is called when exiting the switchstmt production.
	ExitSwitchstmt(c *SwitchstmtContext)

	// ExitDefaultsts is called when exiting the defaultsts production.
	ExitDefaultsts(c *DefaultstsContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitAsignacionstmt is called when exiting the asignacionstmt production.
	ExitAsignacionstmt(c *AsignacionstmtContext)

	// ExitIg_type is called when exiting the ig_type production.
	ExitIg_type(c *Ig_typeContext)

	// ExitReserv is called when exiting the reserv production.
	ExitReserv(c *ReservContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitListParams is called when exiting the listParams production.
	ExitListParams(c *ListParamsContext)

	// ExitListArray is called when exiting the listArray production.
	ExitListArray(c *ListArrayContext)
}
