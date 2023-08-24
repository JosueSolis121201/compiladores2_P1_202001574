// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseSwiftGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSwiftGrammarVisitor) VisitS(ctx *SContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitInstruction(ctx *InstructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitPrintstmt(ctx *PrintstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitIfstmt(ctx *IfstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitDeclarationstmt(ctx *DeclarationstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitAsignacionstmt(ctx *AsignacionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitTypes(ctx *TypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitListParams(ctx *ListParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftGrammarVisitor) VisitListArray(ctx *ListArrayContext) interface{} {
	return v.VisitChildren(ctx)
}
