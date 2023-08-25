package main

import (
	"Server2/parser"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseSwiftGrammarListener
	Code []interface{}
}

type Resp struct {
	Output  string
	Flag    bool
	Message string
}

type Message struct {
	Content string `json:"content"`
}

func main() {
	//Entrada
	code := "print(vec2.count)"
	//Leyendo entrada
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	//creacion de parser
	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true
	tree := p.S()
	//listener
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	/*Code := listener.Code
	//create ast
	var Ast environment.AST
	//ejecución
	for _, inst := range Code {
		inst.(interfaces.Instruction).Ejecutar(&Ast, nil)
		fmt.Print("Hello, ")
	}
	fmt.Println(Ast.GetPrint())*/
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}
