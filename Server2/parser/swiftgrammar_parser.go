// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "Server2/interfaces"
import "Server2/environment"
import "Server2/expressions"
import "Server2/instructions"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftGrammarParser struct {
	*antlr.BaseParser
}

var SwiftGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'at:'",
		"'true'", "'false'", "'print'", "'if'", "'else'", "'else if'", "'while'",
		"'for'", "'in'", "'guard'", "'continue'", "'return'", "'switch'", "'break'",
		"'case'", "'var'", "'let'", "'.append'", "'.remove'", "'.removelast'",
		"'.IsEmpty'", "'.count'", "'repeating'", "'default'", "'count'", "'STRUCT'",
		"'mutating'", "'func'", "'self.'", "'.'", "'inout'", "'_'", "", "",
		"", "'!='", "'=='", "'+='", "'-='", "'!'", "'||'", "'&&'", "'&'", "'='",
		"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'%'", "'-'", "'('",
		"')'", "'{'", "'}'", "':'", "'['", "']'", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "CHAR", "AT_DD", "TRU", "FAL", "PRINT",
		"IF", "ELSE", "ELSE_IF", "WHILE", "FOR", "IN", "GUARD", "CONTINUE",
		"RETURN", "SWITCH", "BREAK", "CASE", "VAR", "LET", "D_APPEND", "D_REMOVE",
		"D_REMOVELAST", "D_ISEMPTY", "D_COUNT", "REPEATING", "DEFAULT", "COUNT",
		"STRUCT", "MUTATING", "FUNC", "SELF_D", "PUNTO", "INOUT", "GUION_B",
		"NUMBER", "STRING", "ID", "DIF", "IG_IG", "MAS_IGUAL", "MENOS_IGUAL",
		"NOT", "OR", "AND", "SOLO_AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR",
		"MUL", "DIV", "ADD", "MODULO", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ",
		"LLAVEDER", "D_PTS", "CORIZQ", "CORDER", "COMA", "PUNTO_COMA", "WHITESPACE",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "asignacion_vectorstmt", "countstmt", "isempystmt",
		"removestmt", "removelaststmt", "appendstmt", "guardstmt", "declaracion_vectorstmt",
		"declaracion_vectorstmt_SUPP", "listaval", "jumpstmt", "whilestmt",
		"forstmt", "printstmt", "ifstmt", "elsestmt", "else_ifstmt", "declarationstmt",
		"switchstmt", "defaultsts", "case", "asignacionstmt", "ig_type", "reserv",
		"types", "expr", "listParams", "listArray",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 484, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 4, 1, 68, 8, 1, 11, 1, 12, 1, 69, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		122, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 3, 11, 194, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 3, 12, 201, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 208, 8,
		12, 10, 12, 12, 12, 211, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		3, 13, 219, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18,
		259, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 5, 19, 270, 8, 19, 10, 19, 12, 19, 273, 9, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 289, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 305, 8, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 315, 8, 23,
		10, 23, 12, 23, 318, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 333, 8, 25, 1, 26,
		1, 26, 1, 26, 1, 26, 3, 26, 339, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 354,
		8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 393, 8, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 5, 28, 450, 8, 28, 10, 28, 12, 28, 453, 9, 28, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 464, 8, 29, 10,
		29, 12, 29, 467, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 5, 30, 479, 8, 30, 10, 30, 12, 30, 482, 9, 30, 1,
		30, 0, 6, 24, 38, 46, 56, 58, 60, 31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 0, 5, 1, 0, 55, 56, 2, 0, 57, 57, 59, 59, 2, 0, 51, 51, 53,
		53, 2, 0, 52, 52, 54, 54, 1, 0, 42, 43, 513, 0, 62, 1, 0, 0, 0, 2, 67,
		1, 0, 0, 0, 4, 121, 1, 0, 0, 0, 6, 123, 1, 0, 0, 0, 8, 134, 1, 0, 0, 0,
		10, 138, 1, 0, 0, 0, 12, 142, 1, 0, 0, 0, 14, 150, 1, 0, 0, 0, 16, 156,
		1, 0, 0, 0, 18, 163, 1, 0, 0, 0, 20, 171, 1, 0, 0, 0, 22, 193, 1, 0, 0,
		0, 24, 200, 1, 0, 0, 0, 26, 218, 1, 0, 0, 0, 28, 220, 1, 0, 0, 0, 30, 227,
		1, 0, 0, 0, 32, 236, 1, 0, 0, 0, 34, 242, 1, 0, 0, 0, 36, 258, 1, 0, 0,
		0, 38, 260, 1, 0, 0, 0, 40, 288, 1, 0, 0, 0, 42, 290, 1, 0, 0, 0, 44, 304,
		1, 0, 0, 0, 46, 306, 1, 0, 0, 0, 48, 319, 1, 0, 0, 0, 50, 332, 1, 0, 0,
		0, 52, 338, 1, 0, 0, 0, 54, 353, 1, 0, 0, 0, 56, 392, 1, 0, 0, 0, 58, 454,
		1, 0, 0, 0, 60, 468, 1, 0, 0, 0, 62, 63, 3, 2, 1, 0, 63, 64, 5, 0, 0, 1,
		64, 65, 6, 0, -1, 0, 65, 1, 1, 0, 0, 0, 66, 68, 3, 4, 2, 0, 67, 66, 1,
		0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70,
		71, 1, 0, 0, 0, 71, 72, 6, 1, -1, 0, 72, 3, 1, 0, 0, 0, 73, 74, 3, 32,
		16, 0, 74, 75, 6, 2, -1, 0, 75, 122, 1, 0, 0, 0, 76, 77, 3, 34, 17, 0,
		77, 78, 6, 2, -1, 0, 78, 122, 1, 0, 0, 0, 79, 80, 3, 40, 20, 0, 80, 81,
		6, 2, -1, 0, 81, 122, 1, 0, 0, 0, 82, 83, 3, 48, 24, 0, 83, 84, 6, 2, -1,
		0, 84, 122, 1, 0, 0, 0, 85, 86, 3, 42, 21, 0, 86, 87, 6, 2, -1, 0, 87,
		122, 1, 0, 0, 0, 88, 89, 3, 28, 14, 0, 89, 90, 6, 2, -1, 0, 90, 122, 1,
		0, 0, 0, 91, 92, 3, 30, 15, 0, 92, 93, 6, 2, -1, 0, 93, 122, 1, 0, 0, 0,
		94, 95, 3, 18, 9, 0, 95, 96, 6, 2, -1, 0, 96, 122, 1, 0, 0, 0, 97, 98,
		3, 26, 13, 0, 98, 99, 6, 2, -1, 0, 99, 122, 1, 0, 0, 0, 100, 101, 3, 20,
		10, 0, 101, 102, 6, 2, -1, 0, 102, 122, 1, 0, 0, 0, 103, 104, 3, 16, 8,
		0, 104, 105, 6, 2, -1, 0, 105, 122, 1, 0, 0, 0, 106, 107, 3, 14, 7, 0,
		107, 108, 6, 2, -1, 0, 108, 122, 1, 0, 0, 0, 109, 110, 3, 12, 6, 0, 110,
		111, 6, 2, -1, 0, 111, 122, 1, 0, 0, 0, 112, 113, 3, 10, 5, 0, 113, 114,
		6, 2, -1, 0, 114, 122, 1, 0, 0, 0, 115, 116, 3, 8, 4, 0, 116, 117, 6, 2,
		-1, 0, 117, 122, 1, 0, 0, 0, 118, 119, 3, 6, 3, 0, 119, 120, 6, 2, -1,
		0, 120, 122, 1, 0, 0, 0, 121, 73, 1, 0, 0, 0, 121, 76, 1, 0, 0, 0, 121,
		79, 1, 0, 0, 0, 121, 82, 1, 0, 0, 0, 121, 85, 1, 0, 0, 0, 121, 88, 1, 0,
		0, 0, 121, 91, 1, 0, 0, 0, 121, 94, 1, 0, 0, 0, 121, 97, 1, 0, 0, 0, 121,
		100, 1, 0, 0, 0, 121, 103, 1, 0, 0, 0, 121, 106, 1, 0, 0, 0, 121, 109,
		1, 0, 0, 0, 121, 112, 1, 0, 0, 0, 121, 115, 1, 0, 0, 0, 121, 118, 1, 0,
		0, 0, 122, 5, 1, 0, 0, 0, 123, 124, 5, 41, 0, 0, 124, 125, 5, 65, 0, 0,
		125, 126, 5, 39, 0, 0, 126, 127, 5, 66, 0, 0, 127, 128, 5, 50, 0, 0, 128,
		129, 5, 41, 0, 0, 129, 130, 5, 65, 0, 0, 130, 131, 5, 39, 0, 0, 131, 132,
		5, 66, 0, 0, 132, 133, 6, 3, -1, 0, 133, 7, 1, 0, 0, 0, 134, 135, 5, 41,
		0, 0, 135, 136, 5, 28, 0, 0, 136, 137, 6, 4, -1, 0, 137, 9, 1, 0, 0, 0,
		138, 139, 5, 41, 0, 0, 139, 140, 5, 27, 0, 0, 140, 141, 6, 5, -1, 0, 141,
		11, 1, 0, 0, 0, 142, 143, 5, 41, 0, 0, 143, 144, 5, 25, 0, 0, 144, 145,
		5, 60, 0, 0, 145, 146, 5, 6, 0, 0, 146, 147, 3, 56, 28, 0, 147, 148, 5,
		61, 0, 0, 148, 149, 6, 6, -1, 0, 149, 13, 1, 0, 0, 0, 150, 151, 5, 41,
		0, 0, 151, 152, 5, 26, 0, 0, 152, 153, 5, 60, 0, 0, 153, 154, 5, 61, 0,
		0, 154, 155, 6, 7, -1, 0, 155, 15, 1, 0, 0, 0, 156, 157, 5, 41, 0, 0, 157,
		158, 5, 24, 0, 0, 158, 159, 5, 60, 0, 0, 159, 160, 3, 56, 28, 0, 160, 161,
		5, 61, 0, 0, 161, 162, 6, 8, -1, 0, 162, 17, 1, 0, 0, 0, 163, 164, 5, 16,
		0, 0, 164, 165, 3, 56, 28, 0, 165, 166, 5, 11, 0, 0, 166, 167, 5, 62, 0,
		0, 167, 168, 3, 2, 1, 0, 168, 169, 5, 63, 0, 0, 169, 170, 6, 9, -1, 0,
		170, 19, 1, 0, 0, 0, 171, 172, 3, 52, 26, 0, 172, 173, 5, 41, 0, 0, 173,
		174, 5, 64, 0, 0, 174, 175, 5, 65, 0, 0, 175, 176, 3, 54, 27, 0, 176, 177,
		5, 66, 0, 0, 177, 178, 5, 50, 0, 0, 178, 179, 3, 22, 11, 0, 179, 180, 6,
		10, -1, 0, 180, 21, 1, 0, 0, 0, 181, 182, 5, 65, 0, 0, 182, 183, 3, 24,
		12, 0, 183, 184, 5, 66, 0, 0, 184, 185, 6, 11, -1, 0, 185, 194, 1, 0, 0,
		0, 186, 194, 5, 41, 0, 0, 187, 188, 3, 54, 27, 0, 188, 189, 5, 65, 0, 0,
		189, 190, 5, 66, 0, 0, 190, 194, 1, 0, 0, 0, 191, 192, 5, 60, 0, 0, 192,
		194, 5, 61, 0, 0, 193, 181, 1, 0, 0, 0, 193, 186, 1, 0, 0, 0, 193, 187,
		1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 23, 1, 0, 0, 0, 195, 196, 6, 12,
		-1, 0, 196, 197, 3, 56, 28, 0, 197, 198, 6, 12, -1, 0, 198, 201, 1, 0,
		0, 0, 199, 201, 1, 0, 0, 0, 200, 195, 1, 0, 0, 0, 200, 199, 1, 0, 0, 0,
		201, 209, 1, 0, 0, 0, 202, 203, 10, 3, 0, 0, 203, 204, 5, 67, 0, 0, 204,
		205, 3, 56, 28, 0, 205, 206, 6, 12, -1, 0, 206, 208, 1, 0, 0, 0, 207, 202,
		1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0,
		0, 0, 210, 25, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212, 213, 5, 17, 0, 0,
		213, 219, 6, 13, -1, 0, 214, 215, 5, 18, 0, 0, 215, 219, 6, 13, -1, 0,
		216, 217, 5, 20, 0, 0, 217, 219, 6, 13, -1, 0, 218, 212, 1, 0, 0, 0, 218,
		214, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 27, 1, 0, 0, 0, 220, 221, 5,
		13, 0, 0, 221, 222, 3, 56, 28, 0, 222, 223, 5, 62, 0, 0, 223, 224, 3, 2,
		1, 0, 224, 225, 5, 63, 0, 0, 225, 226, 6, 14, -1, 0, 226, 29, 1, 0, 0,
		0, 227, 228, 5, 14, 0, 0, 228, 229, 5, 41, 0, 0, 229, 230, 5, 15, 0, 0,
		230, 231, 3, 56, 28, 0, 231, 232, 5, 62, 0, 0, 232, 233, 3, 2, 1, 0, 233,
		234, 5, 63, 0, 0, 234, 235, 6, 15, -1, 0, 235, 31, 1, 0, 0, 0, 236, 237,
		5, 9, 0, 0, 237, 238, 5, 60, 0, 0, 238, 239, 3, 56, 28, 0, 239, 240, 5,
		61, 0, 0, 240, 241, 6, 16, -1, 0, 241, 33, 1, 0, 0, 0, 242, 243, 5, 10,
		0, 0, 243, 244, 3, 56, 28, 0, 244, 245, 5, 62, 0, 0, 245, 246, 3, 2, 1,
		0, 246, 247, 5, 63, 0, 0, 247, 248, 3, 38, 19, 0, 248, 249, 3, 36, 18,
		0, 249, 250, 6, 17, -1, 0, 250, 35, 1, 0, 0, 0, 251, 252, 5, 11, 0, 0,
		252, 253, 5, 62, 0, 0, 253, 254, 3, 2, 1, 0, 254, 255, 5, 63, 0, 0, 255,
		256, 6, 18, -1, 0, 256, 259, 1, 0, 0, 0, 257, 259, 1, 0, 0, 0, 258, 251,
		1, 0, 0, 0, 258, 257, 1, 0, 0, 0, 259, 37, 1, 0, 0, 0, 260, 271, 6, 19,
		-1, 0, 261, 262, 10, 2, 0, 0, 262, 263, 5, 12, 0, 0, 263, 264, 3, 56, 28,
		0, 264, 265, 5, 62, 0, 0, 265, 266, 3, 2, 1, 0, 266, 267, 5, 63, 0, 0,
		267, 268, 6, 19, -1, 0, 268, 270, 1, 0, 0, 0, 269, 261, 1, 0, 0, 0, 270,
		273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 39, 1,
		0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 275, 3, 52, 26, 0, 275, 276, 5, 41,
		0, 0, 276, 277, 5, 64, 0, 0, 277, 278, 3, 54, 27, 0, 278, 279, 5, 50, 0,
		0, 279, 280, 3, 56, 28, 0, 280, 281, 6, 20, -1, 0, 281, 289, 1, 0, 0, 0,
		282, 283, 3, 52, 26, 0, 283, 284, 5, 41, 0, 0, 284, 285, 5, 50, 0, 0, 285,
		286, 3, 56, 28, 0, 286, 287, 6, 20, -1, 0, 287, 289, 1, 0, 0, 0, 288, 274,
		1, 0, 0, 0, 288, 282, 1, 0, 0, 0, 289, 41, 1, 0, 0, 0, 290, 291, 5, 19,
		0, 0, 291, 292, 3, 56, 28, 0, 292, 293, 5, 62, 0, 0, 293, 294, 3, 46, 23,
		0, 294, 295, 3, 44, 22, 0, 295, 296, 5, 63, 0, 0, 296, 297, 6, 21, -1,
		0, 297, 43, 1, 0, 0, 0, 298, 299, 5, 30, 0, 0, 299, 300, 5, 64, 0, 0, 300,
		301, 3, 2, 1, 0, 301, 302, 6, 22, -1, 0, 302, 305, 1, 0, 0, 0, 303, 305,
		1, 0, 0, 0, 304, 298, 1, 0, 0, 0, 304, 303, 1, 0, 0, 0, 305, 45, 1, 0,
		0, 0, 306, 316, 6, 23, -1, 0, 307, 308, 10, 2, 0, 0, 308, 309, 5, 21, 0,
		0, 309, 310, 3, 56, 28, 0, 310, 311, 5, 64, 0, 0, 311, 312, 3, 2, 1, 0,
		312, 313, 6, 23, -1, 0, 313, 315, 1, 0, 0, 0, 314, 307, 1, 0, 0, 0, 315,
		318, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 47, 1,
		0, 0, 0, 318, 316, 1, 0, 0, 0, 319, 320, 5, 41, 0, 0, 320, 321, 3, 50,
		25, 0, 321, 322, 3, 56, 28, 0, 322, 323, 6, 24, -1, 0, 323, 49, 1, 0, 0,
		0, 324, 325, 5, 50, 0, 0, 325, 333, 6, 25, -1, 0, 326, 327, 5, 43, 0, 0,
		327, 333, 6, 25, -1, 0, 328, 329, 5, 44, 0, 0, 329, 333, 6, 25, -1, 0,
		330, 331, 5, 45, 0, 0, 331, 333, 6, 25, -1, 0, 332, 324, 1, 0, 0, 0, 332,
		326, 1, 0, 0, 0, 332, 328, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 333, 51, 1,
		0, 0, 0, 334, 335, 5, 22, 0, 0, 335, 339, 6, 26, -1, 0, 336, 337, 5, 23,
		0, 0, 337, 339, 6, 26, -1, 0, 338, 334, 1, 0, 0, 0, 338, 336, 1, 0, 0,
		0, 339, 53, 1, 0, 0, 0, 340, 341, 5, 1, 0, 0, 341, 354, 6, 27, -1, 0, 342,
		343, 5, 2, 0, 0, 343, 354, 6, 27, -1, 0, 344, 345, 5, 4, 0, 0, 345, 354,
		6, 27, -1, 0, 346, 347, 5, 3, 0, 0, 347, 354, 6, 27, -1, 0, 348, 349, 5,
		5, 0, 0, 349, 354, 6, 27, -1, 0, 350, 351, 5, 65, 0, 0, 351, 352, 5, 66,
		0, 0, 352, 354, 6, 27, -1, 0, 353, 340, 1, 0, 0, 0, 353, 342, 1, 0, 0,
		0, 353, 344, 1, 0, 0, 0, 353, 346, 1, 0, 0, 0, 353, 348, 1, 0, 0, 0, 353,
		350, 1, 0, 0, 0, 354, 55, 1, 0, 0, 0, 355, 356, 6, 28, -1, 0, 356, 357,
		5, 41, 0, 0, 357, 358, 5, 24, 0, 0, 358, 393, 6, 28, -1, 0, 359, 360, 5,
		41, 0, 0, 360, 361, 5, 25, 0, 0, 361, 393, 6, 28, -1, 0, 362, 363, 5, 41,
		0, 0, 363, 364, 5, 26, 0, 0, 364, 393, 6, 28, -1, 0, 365, 366, 5, 41, 0,
		0, 366, 367, 5, 27, 0, 0, 367, 393, 6, 28, -1, 0, 368, 369, 5, 41, 0, 0,
		369, 370, 5, 28, 0, 0, 370, 393, 6, 28, -1, 0, 371, 372, 5, 41, 0, 0, 372,
		373, 5, 65, 0, 0, 373, 374, 5, 39, 0, 0, 374, 375, 5, 66, 0, 0, 375, 393,
		6, 28, -1, 0, 376, 377, 3, 60, 30, 0, 377, 378, 6, 28, -1, 0, 378, 393,
		1, 0, 0, 0, 379, 380, 5, 65, 0, 0, 380, 381, 3, 58, 29, 0, 381, 382, 5,
		66, 0, 0, 382, 383, 6, 28, -1, 0, 383, 393, 1, 0, 0, 0, 384, 385, 5, 39,
		0, 0, 385, 393, 6, 28, -1, 0, 386, 387, 5, 40, 0, 0, 387, 393, 6, 28, -1,
		0, 388, 389, 5, 7, 0, 0, 389, 393, 6, 28, -1, 0, 390, 391, 5, 8, 0, 0,
		391, 393, 6, 28, -1, 0, 392, 355, 1, 0, 0, 0, 392, 359, 1, 0, 0, 0, 392,
		362, 1, 0, 0, 0, 392, 365, 1, 0, 0, 0, 392, 368, 1, 0, 0, 0, 392, 371,
		1, 0, 0, 0, 392, 376, 1, 0, 0, 0, 392, 379, 1, 0, 0, 0, 392, 384, 1, 0,
		0, 0, 392, 386, 1, 0, 0, 0, 392, 388, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0,
		393, 451, 1, 0, 0, 0, 394, 395, 10, 23, 0, 0, 395, 396, 7, 0, 0, 0, 396,
		397, 3, 56, 28, 24, 397, 398, 6, 28, -1, 0, 398, 450, 1, 0, 0, 0, 399,
		400, 10, 22, 0, 0, 400, 401, 7, 1, 0, 0, 401, 402, 3, 56, 28, 23, 402,
		403, 6, 28, -1, 0, 403, 450, 1, 0, 0, 0, 404, 405, 10, 21, 0, 0, 405, 406,
		7, 2, 0, 0, 406, 407, 3, 56, 28, 22, 407, 408, 6, 28, -1, 0, 408, 450,
		1, 0, 0, 0, 409, 410, 10, 20, 0, 0, 410, 411, 7, 3, 0, 0, 411, 412, 3,
		56, 28, 21, 412, 413, 6, 28, -1, 0, 413, 450, 1, 0, 0, 0, 414, 415, 10,
		19, 0, 0, 415, 416, 7, 4, 0, 0, 416, 417, 3, 56, 28, 20, 417, 418, 6, 28,
		-1, 0, 418, 450, 1, 0, 0, 0, 419, 420, 10, 18, 0, 0, 420, 421, 5, 48, 0,
		0, 421, 422, 3, 56, 28, 19, 422, 423, 6, 28, -1, 0, 423, 450, 1, 0, 0,
		0, 424, 425, 10, 17, 0, 0, 425, 426, 5, 47, 0, 0, 426, 427, 3, 56, 28,
		18, 427, 428, 6, 28, -1, 0, 428, 450, 1, 0, 0, 0, 429, 430, 10, 16, 0,
		0, 430, 431, 5, 44, 0, 0, 431, 432, 3, 56, 28, 17, 432, 433, 6, 28, -1,
		0, 433, 450, 1, 0, 0, 0, 434, 435, 10, 15, 0, 0, 435, 436, 5, 45, 0, 0,
		436, 437, 3, 56, 28, 16, 437, 438, 6, 28, -1, 0, 438, 450, 1, 0, 0, 0,
		439, 440, 10, 14, 0, 0, 440, 441, 5, 58, 0, 0, 441, 442, 3, 56, 28, 15,
		442, 443, 6, 28, -1, 0, 443, 450, 1, 0, 0, 0, 444, 445, 10, 13, 0, 0, 445,
		446, 5, 58, 0, 0, 446, 447, 3, 56, 28, 14, 447, 448, 6, 28, -1, 0, 448,
		450, 1, 0, 0, 0, 449, 394, 1, 0, 0, 0, 449, 399, 1, 0, 0, 0, 449, 404,
		1, 0, 0, 0, 449, 409, 1, 0, 0, 0, 449, 414, 1, 0, 0, 0, 449, 419, 1, 0,
		0, 0, 449, 424, 1, 0, 0, 0, 449, 429, 1, 0, 0, 0, 449, 434, 1, 0, 0, 0,
		449, 439, 1, 0, 0, 0, 449, 444, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451,
		449, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452, 57, 1, 0, 0, 0, 453, 451, 1,
		0, 0, 0, 454, 455, 6, 29, -1, 0, 455, 456, 3, 56, 28, 0, 456, 457, 6, 29,
		-1, 0, 457, 465, 1, 0, 0, 0, 458, 459, 10, 2, 0, 0, 459, 460, 5, 67, 0,
		0, 460, 461, 3, 56, 28, 0, 461, 462, 6, 29, -1, 0, 462, 464, 1, 0, 0, 0,
		463, 458, 1, 0, 0, 0, 464, 467, 1, 0, 0, 0, 465, 463, 1, 0, 0, 0, 465,
		466, 1, 0, 0, 0, 466, 59, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 468, 469, 6,
		30, -1, 0, 469, 470, 5, 41, 0, 0, 470, 471, 6, 30, -1, 0, 471, 480, 1,
		0, 0, 0, 472, 473, 10, 2, 0, 0, 473, 474, 5, 65, 0, 0, 474, 475, 3, 56,
		28, 0, 475, 476, 5, 66, 0, 0, 476, 477, 6, 30, -1, 0, 477, 479, 1, 0, 0,
		0, 478, 472, 1, 0, 0, 0, 479, 482, 1, 0, 0, 0, 480, 478, 1, 0, 0, 0, 480,
		481, 1, 0, 0, 0, 481, 61, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 19, 69, 121,
		193, 200, 209, 218, 258, 271, 288, 304, 316, 332, 338, 353, 392, 449, 451,
		465, 480,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SwiftGrammarParserInit initializes any static state used to implement SwiftGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftGrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.once.Do(swiftgrammarParserInit)
}

// NewSwiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftGrammarParser(input antlr.TokenStream) *SwiftGrammarParser {
	SwiftGrammarParserInit()
	this := new(SwiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SwiftGrammar.g4"

	return this
}

// SwiftGrammarParser tokens.
const (
	SwiftGrammarParserEOF          = antlr.TokenEOF
	SwiftGrammarParserINT          = 1
	SwiftGrammarParserFLOAT        = 2
	SwiftGrammarParserBOOL         = 3
	SwiftGrammarParserSTR          = 4
	SwiftGrammarParserCHAR         = 5
	SwiftGrammarParserAT_DD        = 6
	SwiftGrammarParserTRU          = 7
	SwiftGrammarParserFAL          = 8
	SwiftGrammarParserPRINT        = 9
	SwiftGrammarParserIF           = 10
	SwiftGrammarParserELSE         = 11
	SwiftGrammarParserELSE_IF      = 12
	SwiftGrammarParserWHILE        = 13
	SwiftGrammarParserFOR          = 14
	SwiftGrammarParserIN           = 15
	SwiftGrammarParserGUARD        = 16
	SwiftGrammarParserCONTINUE     = 17
	SwiftGrammarParserRETURN       = 18
	SwiftGrammarParserSWITCH       = 19
	SwiftGrammarParserBREAK        = 20
	SwiftGrammarParserCASE         = 21
	SwiftGrammarParserVAR          = 22
	SwiftGrammarParserLET          = 23
	SwiftGrammarParserD_APPEND     = 24
	SwiftGrammarParserD_REMOVE     = 25
	SwiftGrammarParserD_REMOVELAST = 26
	SwiftGrammarParserD_ISEMPTY    = 27
	SwiftGrammarParserD_COUNT      = 28
	SwiftGrammarParserREPEATING    = 29
	SwiftGrammarParserDEFAULT      = 30
	SwiftGrammarParserCOUNT        = 31
	SwiftGrammarParserSTRUCT       = 32
	SwiftGrammarParserMUTATING     = 33
	SwiftGrammarParserFUNC         = 34
	SwiftGrammarParserSELF_D       = 35
	SwiftGrammarParserPUNTO        = 36
	SwiftGrammarParserINOUT        = 37
	SwiftGrammarParserGUION_B      = 38
	SwiftGrammarParserNUMBER       = 39
	SwiftGrammarParserSTRING       = 40
	SwiftGrammarParserID           = 41
	SwiftGrammarParserDIF          = 42
	SwiftGrammarParserIG_IG        = 43
	SwiftGrammarParserMAS_IGUAL    = 44
	SwiftGrammarParserMENOS_IGUAL  = 45
	SwiftGrammarParserNOT          = 46
	SwiftGrammarParserOR           = 47
	SwiftGrammarParserAND          = 48
	SwiftGrammarParserSOLO_AND     = 49
	SwiftGrammarParserIG           = 50
	SwiftGrammarParserMAY_IG       = 51
	SwiftGrammarParserMEN_IG       = 52
	SwiftGrammarParserMAYOR        = 53
	SwiftGrammarParserMENOR        = 54
	SwiftGrammarParserMUL          = 55
	SwiftGrammarParserDIV          = 56
	SwiftGrammarParserADD          = 57
	SwiftGrammarParserMODULO       = 58
	SwiftGrammarParserSUB          = 59
	SwiftGrammarParserPARIZQ       = 60
	SwiftGrammarParserPARDER       = 61
	SwiftGrammarParserLLAVEIZQ     = 62
	SwiftGrammarParserLLAVEDER     = 63
	SwiftGrammarParserD_PTS        = 64
	SwiftGrammarParserCORIZQ       = 65
	SwiftGrammarParserCORDER       = 66
	SwiftGrammarParserCOMA         = 67
	SwiftGrammarParserPUNTO_COMA   = 68
	SwiftGrammarParserWHITESPACE   = 69
	SwiftGrammarParserCOMMENT      = 70
	SwiftGrammarParserLINE_COMMENT = 71
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s                           = 0
	SwiftGrammarParserRULE_block                       = 1
	SwiftGrammarParserRULE_instruction                 = 2
	SwiftGrammarParserRULE_asignacion_vectorstmt       = 3
	SwiftGrammarParserRULE_countstmt                   = 4
	SwiftGrammarParserRULE_isempystmt                  = 5
	SwiftGrammarParserRULE_removestmt                  = 6
	SwiftGrammarParserRULE_removelaststmt              = 7
	SwiftGrammarParserRULE_appendstmt                  = 8
	SwiftGrammarParserRULE_guardstmt                   = 9
	SwiftGrammarParserRULE_declaracion_vectorstmt      = 10
	SwiftGrammarParserRULE_declaracion_vectorstmt_SUPP = 11
	SwiftGrammarParserRULE_listaval                    = 12
	SwiftGrammarParserRULE_jumpstmt                    = 13
	SwiftGrammarParserRULE_whilestmt                   = 14
	SwiftGrammarParserRULE_forstmt                     = 15
	SwiftGrammarParserRULE_printstmt                   = 16
	SwiftGrammarParserRULE_ifstmt                      = 17
	SwiftGrammarParserRULE_elsestmt                    = 18
	SwiftGrammarParserRULE_else_ifstmt                 = 19
	SwiftGrammarParserRULE_declarationstmt             = 20
	SwiftGrammarParserRULE_switchstmt                  = 21
	SwiftGrammarParserRULE_defaultsts                  = 22
	SwiftGrammarParserRULE_case                        = 23
	SwiftGrammarParserRULE_asignacionstmt              = 24
	SwiftGrammarParserRULE_ig_type                     = 25
	SwiftGrammarParserRULE_reserv                      = 26
	SwiftGrammarParserRULE_types                       = 27
	SwiftGrammarParserRULE_expr                        = 28
	SwiftGrammarParserRULE_listParams                  = 29
	SwiftGrammarParserRULE_listArray                   = 30
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	code   []interface{}
	_block IBlockContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Get_block() IBlockContext { return s._block }

func (s *SContext) Set_block(v IBlockContext) { s._block = v }

func (s *SContext) GetCode() []interface{} { return s.code }

func (s *SContext) SetCode(v []interface{}) { s.code = v }

func (s *SContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftGrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(63)
		p.Match(SwiftGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*SContext).code = localctx.(*SContext).Get_block().GetBlk()

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IInstructionContext

	// SetIns sets the ins rule context list.
	SetIns([]IInstructionContext)

	// GetBlk returns the blk attribute.
	GetBlk() []interface{}

	// SetBlk sets the blk attribute.
	SetBlk([]interface{})

	// Getter signatures
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blk          []interface{}
	_instruction IInstructionContext
	ins          []IInstructionContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *BlockContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *BlockContext) GetIns() []IInstructionContext { return s.ins }

func (s *BlockContext) SetIns(v []IInstructionContext) { s.ins = v }

func (s *BlockContext) GetBlk() []interface{} { return s.blk }

func (s *BlockContext) SetBlk(v []interface{}) { s.blk = v }

func (s *BlockContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftGrammarParserRULE_block)

	localctx.(*BlockContext).blk = []interface{}{}
	var listInt []IInstructionContext

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(66)

				var _x = p.Instruction()

				localctx.(*BlockContext)._instruction = _x
			}
			localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	listInt = localctx.(*BlockContext).GetIns()
	for _, e := range listInt {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInst())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_printstmt returns the _printstmt rule contexts.
	Get_printstmt() IPrintstmtContext

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Get_declarationstmt returns the _declarationstmt rule contexts.
	Get_declarationstmt() IDeclarationstmtContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// Set_declarationstmt sets the _declarationstmt rule contexts.
	Set_declarationstmt(IDeclarationstmtContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	Ifstmt() IIfstmtContext
	Declarationstmt() IDeclarationstmtContext
	Asignacionstmt() IAsignacionstmtContext
	Switchstmt() ISwitchstmtContext
	Whilestmt() IWhilestmtContext
	Forstmt() IForstmtContext
	Guardstmt() IGuardstmtContext
	Jumpstmt() IJumpstmtContext
	Declaracion_vectorstmt() IDeclaracion_vectorstmtContext
	Appendstmt() IAppendstmtContext
	Removelaststmt() IRemovelaststmtContext
	Removestmt() IRemovestmtContext
	Isempystmt() IIsempystmtContext
	Countstmt() ICountstmtContext
	Asignacion_vectorstmt() IAsignacion_vectorstmtContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	inst             interfaces.Instruction
	_printstmt       IPrintstmtContext
	_ifstmt          IIfstmtContext
	_declarationstmt IDeclarationstmtContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_printstmt() IPrintstmtContext { return s._printstmt }

func (s *InstructionContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *InstructionContext) Get_declarationstmt() IDeclarationstmtContext { return s._declarationstmt }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_declarationstmt(v IDeclarationstmtContext) { s._declarationstmt = v }

func (s *InstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstructionContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *InstructionContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *InstructionContext) Declarationstmt() IDeclarationstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationstmtContext)
}

func (s *InstructionContext) Asignacionstmt() IAsignacionstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionstmtContext)
}

func (s *InstructionContext) Switchstmt() ISwitchstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstmtContext)
}

func (s *InstructionContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *InstructionContext) Forstmt() IForstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForstmtContext)
}

func (s *InstructionContext) Guardstmt() IGuardstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardstmtContext)
}

func (s *InstructionContext) Jumpstmt() IJumpstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJumpstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJumpstmtContext)
}

func (s *InstructionContext) Declaracion_vectorstmt() IDeclaracion_vectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_vectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_vectorstmtContext)
}

func (s *InstructionContext) Appendstmt() IAppendstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendstmtContext)
}

func (s *InstructionContext) Removelaststmt() IRemovelaststmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemovelaststmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemovelaststmtContext)
}

func (s *InstructionContext) Removestmt() IRemovestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemovestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemovestmtContext)
}

func (s *InstructionContext) Isempystmt() IIsempystmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsempystmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsempystmtContext)
}

func (s *InstructionContext) Countstmt() ICountstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICountstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICountstmtContext)
}

func (s *InstructionContext) Asignacion_vectorstmt() IAsignacion_vectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacion_vectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacion_vectorstmtContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (s *InstructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitInstruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftGrammarParserRULE_instruction)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinst()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)

			var _x = p.Declarationstmt()

			localctx.(*InstructionContext)._declarationstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarationstmt().GetDec()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(82)
			p.Asignacionstmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(85)
			p.Switchstmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(88)
			p.Whilestmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(91)
			p.Forstmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(94)
			p.Guardstmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(97)
			p.Jumpstmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(100)
			p.Declaracion_vectorstmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(103)
			p.Appendstmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(106)
			p.Removelaststmt()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(109)
			p.Removestmt()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(112)
			p.Isempystmt()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(115)
			p.Countstmt()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(118)
			p.Asignacion_vectorstmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacion_vectorstmtContext is an interface to support dynamic dispatch.
type IAsignacion_vectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCORIZQ() []antlr.TerminalNode
	CORIZQ(i int) antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	AllCORDER() []antlr.TerminalNode
	CORDER(i int) antlr.TerminalNode
	IG() antlr.TerminalNode

	// IsAsignacion_vectorstmtContext differentiates from other interfaces.
	IsAsignacion_vectorstmtContext()
}

type Asignacion_vectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacion_vectorstmtContext() *Asignacion_vectorstmtContext {
	var p = new(Asignacion_vectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignacion_vectorstmt
	return p
}

func InitEmptyAsignacion_vectorstmtContext(p *Asignacion_vectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignacion_vectorstmt
}

func (*Asignacion_vectorstmtContext) IsAsignacion_vectorstmtContext() {}

func NewAsignacion_vectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asignacion_vectorstmtContext {
	var p = new(Asignacion_vectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_asignacion_vectorstmt

	return p
}

func (s *Asignacion_vectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Asignacion_vectorstmtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *Asignacion_vectorstmtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *Asignacion_vectorstmtContext) AllCORIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORIZQ)
}

func (s *Asignacion_vectorstmtContext) CORIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, i)
}

func (s *Asignacion_vectorstmtContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserNUMBER)
}

func (s *Asignacion_vectorstmtContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, i)
}

func (s *Asignacion_vectorstmtContext) AllCORDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORDER)
}

func (s *Asignacion_vectorstmtContext) CORDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, i)
}

func (s *Asignacion_vectorstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *Asignacion_vectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_vectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Asignacion_vectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAsignacion_vectorstmt(s)
	}
}

func (s *Asignacion_vectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAsignacion_vectorstmt(s)
	}
}

func (s *Asignacion_vectorstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAsignacion_vectorstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Asignacion_vectorstmt() (localctx IAsignacion_vectorstmtContext) {
	localctx = NewAsignacion_vectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_asignacion_vectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(SwiftGrammarParserCORIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(SwiftGrammarParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Match(SwiftGrammarParserCORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(SwiftGrammarParserCORIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(SwiftGrammarParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(SwiftGrammarParserCORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICountstmtContext is an interface to support dynamic dispatch.
type ICountstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	D_COUNT() antlr.TerminalNode

	// IsCountstmtContext differentiates from other interfaces.
	IsCountstmtContext()
}

type CountstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountstmtContext() *CountstmtContext {
	var p = new(CountstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_countstmt
	return p
}

func InitEmptyCountstmtContext(p *CountstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_countstmt
}

func (*CountstmtContext) IsCountstmtContext() {}

func NewCountstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountstmtContext {
	var p = new(CountstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_countstmt

	return p
}

func (s *CountstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CountstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CountstmtContext) D_COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_COUNT, 0)
}

func (s *CountstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCountstmt(s)
	}
}

func (s *CountstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCountstmt(s)
	}
}

func (s *CountstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitCountstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Countstmt() (localctx ICountstmtContext) {
	localctx = NewCountstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_countstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Match(SwiftGrammarParserD_COUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIsempystmtContext is an interface to support dynamic dispatch.
type IIsempystmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	D_ISEMPTY() antlr.TerminalNode

	// IsIsempystmtContext differentiates from other interfaces.
	IsIsempystmtContext()
}

type IsempystmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsempystmtContext() *IsempystmtContext {
	var p = new(IsempystmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_isempystmt
	return p
}

func InitEmptyIsempystmtContext(p *IsempystmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_isempystmt
}

func (*IsempystmtContext) IsIsempystmtContext() {}

func NewIsempystmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsempystmtContext {
	var p = new(IsempystmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_isempystmt

	return p
}

func (s *IsempystmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IsempystmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *IsempystmtContext) D_ISEMPTY() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_ISEMPTY, 0)
}

func (s *IsempystmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsempystmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsempystmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIsempystmt(s)
	}
}

func (s *IsempystmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIsempystmt(s)
	}
}

func (s *IsempystmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitIsempystmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Isempystmt() (localctx IIsempystmtContext) {
	localctx = NewIsempystmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_isempystmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(SwiftGrammarParserD_ISEMPTY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRemovestmtContext is an interface to support dynamic dispatch.
type IRemovestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	D_REMOVE() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	AT_DD() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsRemovestmtContext differentiates from other interfaces.
	IsRemovestmtContext()
}

type RemovestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemovestmtContext() *RemovestmtContext {
	var p = new(RemovestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removestmt
	return p
}

func InitEmptyRemovestmtContext(p *RemovestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removestmt
}

func (*RemovestmtContext) IsRemovestmtContext() {}

func NewRemovestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemovestmtContext {
	var p = new(RemovestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_removestmt

	return p
}

func (s *RemovestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RemovestmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *RemovestmtContext) D_REMOVE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_REMOVE, 0)
}

func (s *RemovestmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *RemovestmtContext) AT_DD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAT_DD, 0)
}

func (s *RemovestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RemovestmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *RemovestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemovestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemovestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRemovestmt(s)
	}
}

func (s *RemovestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRemovestmt(s)
	}
}

func (s *RemovestmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitRemovestmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Removestmt() (localctx IRemovestmtContext) {
	localctx = NewRemovestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_removestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(SwiftGrammarParserD_REMOVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(SwiftGrammarParserAT_DD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.expr(0)
	}
	{
		p.SetState(147)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRemovelaststmtContext is an interface to support dynamic dispatch.
type IRemovelaststmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	D_REMOVELAST() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode

	// IsRemovelaststmtContext differentiates from other interfaces.
	IsRemovelaststmtContext()
}

type RemovelaststmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemovelaststmtContext() *RemovelaststmtContext {
	var p = new(RemovelaststmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removelaststmt
	return p
}

func InitEmptyRemovelaststmtContext(p *RemovelaststmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removelaststmt
}

func (*RemovelaststmtContext) IsRemovelaststmtContext() {}

func NewRemovelaststmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemovelaststmtContext {
	var p = new(RemovelaststmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_removelaststmt

	return p
}

func (s *RemovelaststmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RemovelaststmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *RemovelaststmtContext) D_REMOVELAST() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_REMOVELAST, 0)
}

func (s *RemovelaststmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *RemovelaststmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *RemovelaststmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemovelaststmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemovelaststmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRemovelaststmt(s)
	}
}

func (s *RemovelaststmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRemovelaststmt(s)
	}
}

func (s *RemovelaststmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitRemovelaststmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Removelaststmt() (localctx IRemovelaststmtContext) {
	localctx = NewRemovelaststmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_removelaststmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(SwiftGrammarParserD_REMOVELAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAppendstmtContext is an interface to support dynamic dispatch.
type IAppendstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	D_APPEND() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsAppendstmtContext differentiates from other interfaces.
	IsAppendstmtContext()
}

type AppendstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppendstmtContext() *AppendstmtContext {
	var p = new(AppendstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_appendstmt
	return p
}

func InitEmptyAppendstmtContext(p *AppendstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_appendstmt
}

func (*AppendstmtContext) IsAppendstmtContext() {}

func NewAppendstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendstmtContext {
	var p = new(AppendstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_appendstmt

	return p
}

func (s *AppendstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AppendstmtContext) D_APPEND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_APPEND, 0)
}

func (s *AppendstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *AppendstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AppendstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *AppendstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppendstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAppendstmt(s)
	}
}

func (s *AppendstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAppendstmt(s)
	}
}

func (s *AppendstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAppendstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Appendstmt() (localctx IAppendstmtContext) {
	localctx = NewAppendstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_appendstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(SwiftGrammarParserD_APPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.expr(0)
	}
	{
		p.SetState(160)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuardstmtContext is an interface to support dynamic dispatch.
type IGuardstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsGuardstmtContext differentiates from other interfaces.
	IsGuardstmtContext()
}

type GuardstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardstmtContext() *GuardstmtContext {
	var p = new(GuardstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
	return p
}

func InitEmptyGuardstmtContext(p *GuardstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
}

func (*GuardstmtContext) IsGuardstmtContext() {}

func NewGuardstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardstmtContext {
	var p = new(GuardstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt

	return p
}

func (s *GuardstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardstmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGUARD, 0)
}

func (s *GuardstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *GuardstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *GuardstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GuardstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *GuardstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterGuardstmt(s)
	}
}

func (s *GuardstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitGuardstmt(s)
	}
}

func (s *GuardstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitGuardstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Guardstmt() (localctx IGuardstmtContext) {
	localctx = NewGuardstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_guardstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(SwiftGrammarParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.expr(0)
	}
	{
		p.SetState(165)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Block()
	}
	{
		p.SetState(168)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracion_vectorstmtContext is an interface to support dynamic dispatch.
type IDeclaracion_vectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Reserv() IReservContext
	ID() antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	Types() ITypesContext
	CORDER() antlr.TerminalNode
	IG() antlr.TerminalNode
	Declaracion_vectorstmt_SUPP() IDeclaracion_vectorstmt_SUPPContext

	// IsDeclaracion_vectorstmtContext differentiates from other interfaces.
	IsDeclaracion_vectorstmtContext()
}

type Declaracion_vectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracion_vectorstmtContext() *Declaracion_vectorstmtContext {
	var p = new(Declaracion_vectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declaracion_vectorstmt
	return p
}

func InitEmptyDeclaracion_vectorstmtContext(p *Declaracion_vectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declaracion_vectorstmt
}

func (*Declaracion_vectorstmtContext) IsDeclaracion_vectorstmtContext() {}

func NewDeclaracion_vectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_vectorstmtContext {
	var p = new(Declaracion_vectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_declaracion_vectorstmt

	return p
}

func (s *Declaracion_vectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_vectorstmtContext) Reserv() IReservContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReservContext)
}

func (s *Declaracion_vectorstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *Declaracion_vectorstmtContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *Declaracion_vectorstmtContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *Declaracion_vectorstmtContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *Declaracion_vectorstmtContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *Declaracion_vectorstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *Declaracion_vectorstmtContext) Declaracion_vectorstmt_SUPP() IDeclaracion_vectorstmt_SUPPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_vectorstmt_SUPPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_vectorstmt_SUPPContext)
}

func (s *Declaracion_vectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_vectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracion_vectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclaracion_vectorstmt(s)
	}
}

func (s *Declaracion_vectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclaracion_vectorstmt(s)
	}
}

func (s *Declaracion_vectorstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDeclaracion_vectorstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Declaracion_vectorstmt() (localctx IDeclaracion_vectorstmtContext) {
	localctx = NewDeclaracion_vectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_declaracion_vectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Reserv()
	}
	{
		p.SetState(172)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(SwiftGrammarParserD_PTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(SwiftGrammarParserCORIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Types()
	}
	{
		p.SetState(176)
		p.Match(SwiftGrammarParserCORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Declaracion_vectorstmt_SUPP()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracion_vectorstmt_SUPPContext is an interface to support dynamic dispatch.
type IDeclaracion_vectorstmt_SUPPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CORIZQ() antlr.TerminalNode
	Listaval() IListavalContext
	CORDER() antlr.TerminalNode
	ID() antlr.TerminalNode
	Types() ITypesContext
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode

	// IsDeclaracion_vectorstmt_SUPPContext differentiates from other interfaces.
	IsDeclaracion_vectorstmt_SUPPContext()
}

type Declaracion_vectorstmt_SUPPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracion_vectorstmt_SUPPContext() *Declaracion_vectorstmt_SUPPContext {
	var p = new(Declaracion_vectorstmt_SUPPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declaracion_vectorstmt_SUPP
	return p
}

func InitEmptyDeclaracion_vectorstmt_SUPPContext(p *Declaracion_vectorstmt_SUPPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declaracion_vectorstmt_SUPP
}

func (*Declaracion_vectorstmt_SUPPContext) IsDeclaracion_vectorstmt_SUPPContext() {}

func NewDeclaracion_vectorstmt_SUPPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_vectorstmt_SUPPContext {
	var p = new(Declaracion_vectorstmt_SUPPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_declaracion_vectorstmt_SUPP

	return p
}

func (s *Declaracion_vectorstmt_SUPPContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_vectorstmt_SUPPContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *Declaracion_vectorstmt_SUPPContext) Listaval() IListavalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListavalContext)
}

func (s *Declaracion_vectorstmt_SUPPContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *Declaracion_vectorstmt_SUPPContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *Declaracion_vectorstmt_SUPPContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *Declaracion_vectorstmt_SUPPContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *Declaracion_vectorstmt_SUPPContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *Declaracion_vectorstmt_SUPPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_vectorstmt_SUPPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracion_vectorstmt_SUPPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclaracion_vectorstmt_SUPP(s)
	}
}

func (s *Declaracion_vectorstmt_SUPPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclaracion_vectorstmt_SUPP(s)
	}
}

func (s *Declaracion_vectorstmt_SUPPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDeclaracion_vectorstmt_SUPP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Declaracion_vectorstmt_SUPP() (localctx IDeclaracion_vectorstmt_SUPPContext) {
	localctx = NewDeclaracion_vectorstmt_SUPPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_declaracion_vectorstmt_SUPP)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.listaval(0)
		}
		{
			p.SetState(183)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.Types()
		}
		{
			p.SetState(188)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(191)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListavalContext is an interface to support dynamic dispatch.
type IListavalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Listaval() IListavalContext
	COMA() antlr.TerminalNode

	// IsListavalContext differentiates from other interfaces.
	IsListavalContext()
}

type ListavalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListavalContext() *ListavalContext {
	var p = new(ListavalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listaval
	return p
}

func InitEmptyListavalContext(p *ListavalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listaval
}

func (*ListavalContext) IsListavalContext() {}

func NewListavalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListavalContext {
	var p = new(ListavalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listaval

	return p
}

func (s *ListavalContext) GetParser() antlr.Parser { return s.parser }

func (s *ListavalContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListavalContext) Listaval() IListavalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListavalContext)
}

func (s *ListavalContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListavalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListavalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListavalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListaval(s)
	}
}

func (s *ListavalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListaval(s)
	}
}

func (s *ListavalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitListaval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Listaval() (localctx IListavalContext) {
	return p.listaval(0)
}

func (p *SwiftGrammarParser) listaval(_p int) (localctx IListavalContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListavalContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListavalContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, SwiftGrammarParserRULE_listaval, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(196)
			p.expr(0)
		}

	case 2:

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListavalContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listaval)
			p.SetState(202)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(203)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(204)
				p.expr(0)
			}

		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJumpstmtContext is an interface to support dynamic dispatch.
type IJumpstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	BREAK() antlr.TerminalNode

	// IsJumpstmtContext differentiates from other interfaces.
	IsJumpstmtContext()
}

type JumpstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJumpstmtContext() *JumpstmtContext {
	var p = new(JumpstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_jumpstmt
	return p
}

func InitEmptyJumpstmtContext(p *JumpstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_jumpstmt
}

func (*JumpstmtContext) IsJumpstmtContext() {}

func NewJumpstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JumpstmtContext {
	var p = new(JumpstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_jumpstmt

	return p
}

func (s *JumpstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *JumpstmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCONTINUE, 0)
}

func (s *JumpstmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRETURN, 0)
}

func (s *JumpstmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBREAK, 0)
}

func (s *JumpstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JumpstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterJumpstmt(s)
	}
}

func (s *JumpstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitJumpstmt(s)
	}
}

func (s *JumpstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitJumpstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Jumpstmt() (localctx IJumpstmtContext) {
	localctx = NewJumpstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_jumpstmt)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserCONTINUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Match(SwiftGrammarParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserRETURN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Match(SwiftGrammarParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserBREAK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Match(SwiftGrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserWHILE, 0)
}

func (s *WhilestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhilestmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *WhilestmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhilestmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (s *WhilestmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitWhilestmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(SwiftGrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.expr(0)
	}
	{
		p.SetState(222)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Block()
	}
	{
		p.SetState(224)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForstmtContext is an interface to support dynamic dispatch.
type IForstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IN() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForstmtContext() *ForstmtContext {
	var p = new(ForstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
	return p
}

func InitEmptyForstmtContext(p *ForstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
}

func (*ForstmtContext) IsForstmtContext() {}

func NewForstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstmtContext {
	var p = new(ForstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_forstmt

	return p
}

func (s *ForstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFOR, 0)
}

func (s *ForstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ForstmtContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIN, 0)
}

func (s *ForstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ForstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ForstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterForstmt(s)
	}
}

func (s *ForstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitForstmt(s)
	}
}

func (s *ForstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitForstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Forstmt() (localctx IForstmtContext) {
	localctx = NewForstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_forstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(SwiftGrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(SwiftGrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.expr(0)
	}
	{
		p.SetState(231)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Block()
	}
	{
		p.SetState(233)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	prnt   interfaces.Instruction
	_PRINT antlr.Token
	_expr  IExprContext
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *PrintstmtContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *PrintstmtContext) Get_expr() IExprContext { return s._expr }

func (s *PrintstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (s *PrintstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitPrintstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)

		var _m = p.Match(SwiftGrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)

		var _x = p.expr(0)

		localctx.(*PrintstmtContext)._expr = _x
	}
	{
		p.SetState(239)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
		if localctx.(*PrintstmtContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
		}
	}()), (func() int {
		if localctx.(*PrintstmtContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
		}
	}()), localctx.(*PrintstmtContext).Get_expr().GetE())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIfinst returns the ifinst attribute.
	GetIfinst() interfaces.Instruction

	// SetIfinst sets the ifinst attribute.
	SetIfinst(interfaces.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	Else_ifstmt() IElse_ifstmtContext
	Elsestmt() IElsestmtContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ifinst interfaces.Instruction
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) GetIfinst() interfaces.Instruction { return s.ifinst }

func (s *IfstmtContext) SetIfinst(v interfaces.Instruction) { s.ifinst = v }

func (s *IfstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *IfstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *IfstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *IfstmtContext) Else_ifstmt() IElse_ifstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_ifstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_ifstmtContext)
}

func (s *IfstmtContext) Elsestmt() IElsestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElsestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElsestmtContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (s *IfstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitIfstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_ifstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(SwiftGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.expr(0)
	}
	{
		p.SetState(244)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Block()
	}
	{
		p.SetState(246)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.else_ifstmt(0)
	}
	{
		p.SetState(248)
		p.Elsestmt()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElsestmtContext is an interface to support dynamic dispatch.
type IElsestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsElsestmtContext differentiates from other interfaces.
	IsElsestmtContext()
}

type ElsestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElsestmtContext() *ElsestmtContext {
	var p = new(ElsestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elsestmt
	return p
}

func InitEmptyElsestmtContext(p *ElsestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elsestmt
}

func (*ElsestmtContext) IsElsestmtContext() {}

func NewElsestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsestmtContext {
	var p = new(ElsestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_elsestmt

	return p
}

func (s *ElsestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsestmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *ElsestmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ElsestmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElsestmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ElsestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElsestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElsestmt(s)
	}
}

func (s *ElsestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElsestmt(s)
	}
}

func (s *ElsestmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitElsestmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Elsestmt() (localctx IElsestmtContext) {
	localctx = NewElsestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_elsestmt)
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.Block()
		}
		{
			p.SetState(254)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElse_ifstmtContext is an interface to support dynamic dispatch.
type IElse_ifstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Else_ifstmt() IElse_ifstmtContext
	ELSE_IF() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsElse_ifstmtContext differentiates from other interfaces.
	IsElse_ifstmtContext()
}

type Else_ifstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_ifstmtContext() *Else_ifstmtContext {
	var p = new(Else_ifstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_else_ifstmt
	return p
}

func InitEmptyElse_ifstmtContext(p *Else_ifstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_else_ifstmt
}

func (*Else_ifstmtContext) IsElse_ifstmtContext() {}

func NewElse_ifstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_ifstmtContext {
	var p = new(Else_ifstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_else_ifstmt

	return p
}

func (s *Else_ifstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_ifstmtContext) Else_ifstmt() IElse_ifstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_ifstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_ifstmtContext)
}

func (s *Else_ifstmtContext) ELSE_IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE_IF, 0)
}

func (s *Else_ifstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Else_ifstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *Else_ifstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Else_ifstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *Else_ifstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_ifstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_ifstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElse_ifstmt(s)
	}
}

func (s *Else_ifstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElse_ifstmt(s)
	}
}

func (s *Else_ifstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitElse_ifstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Else_ifstmt() (localctx IElse_ifstmtContext) {
	return p.else_ifstmt(0)
}

func (p *SwiftGrammarParser) else_ifstmt(_p int) (localctx IElse_ifstmtContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewElse_ifstmtContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IElse_ifstmtContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, SwiftGrammarParserRULE_else_ifstmt, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewElse_ifstmtContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_else_ifstmt)
			p.SetState(261)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(262)
				p.Match(SwiftGrammarParserELSE_IF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(263)
				p.expr(0)
			}
			{
				p.SetState(264)
				p.Match(SwiftGrammarParserLLAVEIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(265)
				p.Block()
			}
			{
				p.SetState(266)
				p.Match(SwiftGrammarParserLLAVEDER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationstmtContext is an interface to support dynamic dispatch.
type IDeclarationstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// Getter signatures
	Reserv() IReservContext
	ID() antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Types() ITypesContext
	IG() antlr.TerminalNode
	Expr() IExprContext

	// IsDeclarationstmtContext differentiates from other interfaces.
	IsDeclarationstmtContext()
}

type DeclarationstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	dec    interfaces.Instruction
}

func NewEmptyDeclarationstmtContext() *DeclarationstmtContext {
	var p = new(DeclarationstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt
	return p
}

func InitEmptyDeclarationstmtContext(p *DeclarationstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt
}

func (*DeclarationstmtContext) IsDeclarationstmtContext() {}

func NewDeclarationstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationstmtContext {
	var p = new(DeclarationstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt

	return p
}

func (s *DeclarationstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationstmtContext) GetDec() interfaces.Instruction { return s.dec }

func (s *DeclarationstmtContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *DeclarationstmtContext) Reserv() IReservContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReservContext)
}

func (s *DeclarationstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *DeclarationstmtContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *DeclarationstmtContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *DeclarationstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *DeclarationstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclarationstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclarationstmt(s)
	}
}

func (s *DeclarationstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclarationstmt(s)
	}
}

func (s *DeclarationstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDeclarationstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Declarationstmt() (localctx IDeclarationstmtContext) {
	localctx = NewDeclarationstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_declarationstmt)
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.Reserv()
		}
		{
			p.SetState(275)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Types()
		}
		{
			p.SetState(278)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.Reserv()
		}
		{
			p.SetState(283)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchstmtContext is an interface to support dynamic dispatch.
type ISwitchstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Case_() ICaseContext
	Defaultsts() IDefaultstsContext
	LLAVEDER() antlr.TerminalNode

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchstmtContext() *SwitchstmtContext {
	var p = new(SwitchstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt
	return p
}

func InitEmptySwitchstmtContext(p *SwitchstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt
}

func (*SwitchstmtContext) IsSwitchstmtContext() {}

func NewSwitchstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstmtContext {
	var p = new(SwitchstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt

	return p
}

func (s *SwitchstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSWITCH, 0)
}

func (s *SwitchstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *SwitchstmtContext) Case_() ICaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseContext)
}

func (s *SwitchstmtContext) Defaultsts() IDefaultstsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultstsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultstsContext)
}

func (s *SwitchstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *SwitchstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitSwitchstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Switchstmt() (localctx ISwitchstmtContext) {
	localctx = NewSwitchstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SwiftGrammarParserRULE_switchstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(SwiftGrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.expr(0)
	}
	{
		p.SetState(292)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.case_(0)
	}
	{
		p.SetState(294)
		p.Defaultsts()
	}
	{
		p.SetState(295)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultstsContext is an interface to support dynamic dispatch.
type IDefaultstsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Block() IBlockContext

	// IsDefaultstsContext differentiates from other interfaces.
	IsDefaultstsContext()
}

type DefaultstsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultstsContext() *DefaultstsContext {
	var p = new(DefaultstsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defaultsts
	return p
}

func InitEmptyDefaultstsContext(p *DefaultstsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defaultsts
}

func (*DefaultstsContext) IsDefaultstsContext() {}

func NewDefaultstsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultstsContext {
	var p = new(DefaultstsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_defaultsts

	return p
}

func (s *DefaultstsContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultstsContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDEFAULT, 0)
}

func (s *DefaultstsContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *DefaultstsContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefaultstsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultstsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultstsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDefaultsts(s)
	}
}

func (s *DefaultstsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDefaultsts(s)
	}
}

func (s *DefaultstsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDefaultsts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Defaultsts() (localctx IDefaultstsContext) {
	localctx = NewDefaultstsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftGrammarParserRULE_defaultsts)
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserDEFAULT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.Match(SwiftGrammarParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Block()
		}

	case SwiftGrammarParserLLAVEDER:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseContext is an interface to support dynamic dispatch.
type ICaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Case_() ICaseContext
	CASE() antlr.TerminalNode
	Expr() IExprContext
	D_PTS() antlr.TerminalNode
	Block() IBlockContext

	// IsCaseContext differentiates from other interfaces.
	IsCaseContext()
}

type CaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseContext() *CaseContext {
	var p = new(CaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_case
	return p
}

func InitEmptyCaseContext(p *CaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_case
}

func (*CaseContext) IsCaseContext() {}

func NewCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseContext {
	var p = new(CaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_case

	return p
}

func (s *CaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseContext) Case_() ICaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseContext)
}

func (s *CaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCASE, 0)
}

func (s *CaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CaseContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *CaseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCase(s)
	}
}

func (s *CaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Case_() (localctx ICaseContext) {
	return p.case_(0)
}

func (p *SwiftGrammarParser) case_(_p int) (localctx ICaseContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCaseContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICaseContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, SwiftGrammarParserRULE_case, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCaseContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_case)
			p.SetState(307)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(308)
				p.Match(SwiftGrammarParserCASE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(309)
				p.expr(0)
			}
			{
				p.SetState(310)
				p.Match(SwiftGrammarParserD_PTS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(311)
				p.Block()
			}

		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionstmtContext is an interface to support dynamic dispatch.
type IAsignacionstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Ig_type() IIg_typeContext
	Expr() IExprContext

	// IsAsignacionstmtContext differentiates from other interfaces.
	IsAsignacionstmtContext()
}

type AsignacionstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionstmtContext() *AsignacionstmtContext {
	var p = new(AsignacionstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignacionstmt
	return p
}

func InitEmptyAsignacionstmtContext(p *AsignacionstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignacionstmt
}

func (*AsignacionstmtContext) IsAsignacionstmtContext() {}

func NewAsignacionstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionstmtContext {
	var p = new(AsignacionstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_asignacionstmt

	return p
}

func (s *AsignacionstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AsignacionstmtContext) Ig_type() IIg_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIg_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIg_typeContext)
}

func (s *AsignacionstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAsignacionstmt(s)
	}
}

func (s *AsignacionstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAsignacionstmt(s)
	}
}

func (s *AsignacionstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAsignacionstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Asignacionstmt() (localctx IAsignacionstmtContext) {
	localctx = NewAsignacionstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftGrammarParserRULE_asignacionstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)
		p.Ig_type()
	}
	{
		p.SetState(321)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIg_typeContext is an interface to support dynamic dispatch.
type IIg_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IG() antlr.TerminalNode
	IG_IG() antlr.TerminalNode
	MAS_IGUAL() antlr.TerminalNode
	MENOS_IGUAL() antlr.TerminalNode

	// IsIg_typeContext differentiates from other interfaces.
	IsIg_typeContext()
}

type Ig_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIg_typeContext() *Ig_typeContext {
	var p = new(Ig_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ig_type
	return p
}

func InitEmptyIg_typeContext(p *Ig_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ig_type
}

func (*Ig_typeContext) IsIg_typeContext() {}

func NewIg_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ig_typeContext {
	var p = new(Ig_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_ig_type

	return p
}

func (s *Ig_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Ig_typeContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *Ig_typeContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG_IG, 0)
}

func (s *Ig_typeContext) MAS_IGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAS_IGUAL, 0)
}

func (s *Ig_typeContext) MENOS_IGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOS_IGUAL, 0)
}

func (s *Ig_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ig_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ig_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIg_type(s)
	}
}

func (s *Ig_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIg_type(s)
	}
}

func (s *Ig_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitIg_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Ig_type() (localctx IIg_typeContext) {
	localctx = NewIg_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SwiftGrammarParserRULE_ig_type)
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserIG:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(324)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserIG_IG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)
			p.Match(SwiftGrammarParserIG_IG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserMAS_IGUAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(328)
			p.Match(SwiftGrammarParserMAS_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserMENOS_IGUAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(330)
			p.Match(SwiftGrammarParserMENOS_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReservContext is an interface to support dynamic dispatch.
type IReservContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	LET() antlr.TerminalNode

	// IsReservContext differentiates from other interfaces.
	IsReservContext()
}

type ReservContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservContext() *ReservContext {
	var p = new(ReservContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_reserv
	return p
}

func InitEmptyReservContext(p *ReservContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_reserv
}

func (*ReservContext) IsReservContext() {}

func NewReservContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservContext {
	var p = new(ReservContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_reserv

	return p
}

func (s *ReservContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *ReservContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
}

func (s *ReservContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterReserv(s)
	}
}

func (s *ReservContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitReserv(s)
	}
}

func (s *ReservContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitReserv(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Reserv() (localctx IReservContext) {
	localctx = NewReservContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SwiftGrammarParserRULE_reserv)
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(334)
			p.Match(SwiftGrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.Match(SwiftGrammarParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty attribute.
	GetTy() environment.TipoExpresion

	// SetTy sets the ty attribute.
	SetTy(environment.TipoExpresion)

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STR() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	CORDER() antlr.TerminalNode

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     environment.TipoExpresion
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_types
	return p
}

func InitEmptyTypesContext(p *TypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_types
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) GetTy() environment.TipoExpresion { return s.ty }

func (s *TypesContext) SetTy(v environment.TipoExpresion) { s.ty = v }

func (s *TypesContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *TypesContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *TypesContext) STR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTR, 0)
}

func (s *TypesContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBOOL, 0)
}

func (s *TypesContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCHAR, 0)
}

func (s *TypesContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *TypesContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (s *TypesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitTypes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SwiftGrammarParserRULE_types)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(340)
			p.Match(SwiftGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.INTEGER

	case SwiftGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.Match(SwiftGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.FLOAT

	case SwiftGrammarParserSTR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(344)
			p.Match(SwiftGrammarParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.STRING

	case SwiftGrammarParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(346)
			p.Match(SwiftGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.BOOLEAN

	case SwiftGrammarParserCHAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(348)
			p.Match(SwiftGrammarParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserCORIZQ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(350)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.ARRAY

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CORIZQ returns the _CORIZQ token.
	Get_CORIZQ() antlr.Token

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_TRU returns the _TRU token.
	Get_TRU() antlr.Token

	// Get_FAL returns the _FAL token.
	Get_FAL() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_CORIZQ sets the _CORIZQ token.
	Set_CORIZQ(antlr.Token)

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_TRU sets the _TRU token.
	Set_TRU(antlr.Token)

	// Set_FAL sets the _FAL token.
	Set_FAL(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	D_APPEND() antlr.TerminalNode
	D_REMOVE() antlr.TerminalNode
	D_REMOVELAST() antlr.TerminalNode
	D_ISEMPTY() antlr.TerminalNode
	D_COUNT() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	CORDER() antlr.TerminalNode
	ListArray() IListArrayContext
	ListParams() IListParamsContext
	STRING() antlr.TerminalNode
	TRU() antlr.TerminalNode
	FAL() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MAY_IG() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	MEN_IG() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	IG_IG() antlr.TerminalNode
	DIF() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	MAS_IGUAL() antlr.TerminalNode
	MENOS_IGUAL() antlr.TerminalNode
	MODULO() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	e           interfaces.Expression
	left        IExprContext
	_CORIZQ     antlr.Token
	_NUMBER     antlr.Token
	list        IListArrayContext
	_listParams IListParamsContext
	_STRING     antlr.Token
	_TRU        antlr.Token
	_FAL        antlr.Token
	op          antlr.Token
	right       IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Get_CORIZQ() antlr.Token { return s._CORIZQ }

func (s *ExprContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExprContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ExprContext) Get_TRU() antlr.Token { return s._TRU }

func (s *ExprContext) Get_FAL() antlr.Token { return s._FAL }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) Set_CORIZQ(v antlr.Token) { s._CORIZQ = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExprContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *ExprContext) Set_FAL(v antlr.Token) { s._FAL = v }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) GetList() IListArrayContext { return s.list }

func (s *ExprContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) SetList(v IListArrayContext) { s.list = v }

func (s *ExprContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ExprContext) D_APPEND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_APPEND, 0)
}

func (s *ExprContext) D_REMOVE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_REMOVE, 0)
}

func (s *ExprContext) D_REMOVELAST() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_REMOVELAST, 0)
}

func (s *ExprContext) D_ISEMPTY() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_ISEMPTY, 0)
}

func (s *ExprContext) D_COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_COUNT, 0)
}

func (s *ExprContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
}

func (s *ExprContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ExprContext) ListArray() IListArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ExprContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *ExprContext) TRU() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTRU, 0)
}

func (s *ExprContext) FAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFAL, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIV, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *ExprContext) MAY_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAY_IG, 0)
}

func (s *ExprContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYOR, 0)
}

func (s *ExprContext) MEN_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMEN_IG, 0)
}

func (s *ExprContext) MENOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOR, 0)
}

func (s *ExprContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG_IG, 0)
}

func (s *ExprContext) DIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIF, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOR, 0)
}

func (s *ExprContext) MAS_IGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAS_IGUAL, 0)
}

func (s *ExprContext) MENOS_IGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOS_IGUAL, 0)
}

func (s *ExprContext) MODULO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMODULO, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(356)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Match(SwiftGrammarParserD_APPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(359)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.Match(SwiftGrammarParserD_REMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(362)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.Match(SwiftGrammarParserD_REMOVELAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(365)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Match(SwiftGrammarParserD_ISEMPTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(368)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(SwiftGrammarParserD_COUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		{
			p.SetState(371)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)

			var _m = p.Match(SwiftGrammarParserCORIZQ)

			localctx.(*ExprContext)._CORIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*ExprContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		{
			p.SetState(376)

			var _x = p.listArray(0)

			localctx.(*ExprContext).list = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).GetList().GetP()

	case 8:
		{
			p.SetState(379)

			var _m = p.Match(SwiftGrammarParserCORIZQ)

			localctx.(*ExprContext)._CORIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)

			var _x = p.listParams(0)

			localctx.(*ExprContext)._listParams = _x
		}
		{
			p.SetState(381)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewArray((func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetColumn()
			}
		}()), localctx.(*ExprContext).Get_listParams().GetL())

	case 9:
		{
			p.SetState(384)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*ExprContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if strings.Contains((func() string {
			if localctx.(*ExprContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.INTEGER)
		}

	case 10:
		{
			p.SetState(386)

			var _m = p.Match(SwiftGrammarParserSTRING)

			localctx.(*ExprContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_STRING().GetText()
			}
		}())
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case 11:
		{
			p.SetState(388)

			var _m = p.Match(SwiftGrammarParserTRU)

			localctx.(*ExprContext)._TRU = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetColumn()
			}
		}()), true, environment.BOOLEAN)

	case 12:
		{
			p.SetState(390)

			var _m = p.Match(SwiftGrammarParserFAL)

			localctx.(*ExprContext)._FAL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetColumn()
			}
		}()), false, environment.BOOLEAN)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(449)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(394)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(395)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMUL || _la == SwiftGrammarParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(396)

					var _x = p.expr(24)

					localctx.(*ExprContext).right = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(399)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(400)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserADD || _la == SwiftGrammarParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(401)

					var _x = p.expr(23)

					localctx.(*ExprContext).right = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(404)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(405)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMAY_IG || _la == SwiftGrammarParserMAYOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(406)

					var _x = p.expr(22)

					localctx.(*ExprContext).right = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(409)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(410)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMEN_IG || _la == SwiftGrammarParserMENOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(411)

					var _x = p.expr(21)

					localctx.(*ExprContext).right = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(414)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(415)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserDIF || _la == SwiftGrammarParserIG_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(416)

					var _x = p.expr(20)

					localctx.(*ExprContext).right = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(419)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(420)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(421)

					var _x = p.expr(19)

					localctx.(*ExprContext).right = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(424)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(425)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(426)

					var _x = p.expr(18)

					localctx.(*ExprContext).right = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(429)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(430)

					var _m = p.Match(SwiftGrammarParserMAS_IGUAL)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(431)

					var _x = p.expr(17)

					localctx.(*ExprContext).right = _x
				}

			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(434)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(435)

					var _m = p.Match(SwiftGrammarParserMENOS_IGUAL)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(436)

					var _x = p.expr(16)

					localctx.(*ExprContext).right = _x
				}

			case 10:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(439)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(440)

					var _m = p.Match(SwiftGrammarParserMODULO)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(441)

					var _x = p.expr(15)

					localctx.(*ExprContext).right = _x
				}

			case 11:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(444)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(445)

					var _m = p.Match(SwiftGrammarParserMODULO)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(446)

					var _x = p.expr(14)

					localctx.(*ExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListParamsContext is an interface to support dynamic dispatch.
type IListParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListParamsContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListParams() IListParamsContext

	// IsListParamsContext differentiates from other interfaces.
	IsListParamsContext()
}

type ListParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListParamsContext
	_expr  IExprContext
}

func NewEmptyListParamsContext() *ListParamsContext {
	var p = new(ListParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParams
	return p
}

func InitEmptyListParamsContext(p *ListParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParams
}

func (*ListParamsContext) IsListParamsContext() {}

func NewListParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsContext {
	var p = new(ListParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParams

	return p
}

func (s *ListParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsContext) GetList() IListParamsContext { return s.list }

func (s *ListParamsContext) Get_expr() IExprContext { return s._expr }

func (s *ListParamsContext) SetList(v IListParamsContext) { s.list = v }

func (s *ListParamsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListParamsContext) GetL() []interface{} { return s.l }

func (s *ListParamsContext) SetL(v []interface{}) { s.l = v }

func (s *ListParamsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListParamsContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ListParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParams(s)
	}
}

func (s *ListParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParams(s)
	}
}

func (s *ListParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitListParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) ListParams() (localctx IListParamsContext) {
	return p.listParams(0)
}

func (p *SwiftGrammarParser) listParams(_p int) (localctx IListParamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, SwiftGrammarParserRULE_listParams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)

		var _x = p.expr(0)

		localctx.(*ListParamsContext)._expr = _x
	}

	localctx.(*ListParamsContext).l = []interface{}{}
	localctx.(*ListParamsContext).l = append(localctx.(*ListParamsContext).l, localctx.(*ListParamsContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParams)
			p.SetState(458)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(459)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(460)

				var _x = p.expr(0)

				localctx.(*ListParamsContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsContext).GetList().GetL(), localctx.(*ListParamsContext).Get_expr().GetE())
			localctx.(*ListParamsContext).l = arr

		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListArrayContext is an interface to support dynamic dispatch.
type IListArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	Expr() IExprContext
	CORDER() antlr.TerminalNode
	ListArray() IListArrayContext

	// IsListArrayContext differentiates from other interfaces.
	IsListArrayContext()
}

type ListArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	p      interfaces.Expression
	list   IListArrayContext
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyListArrayContext() *ListArrayContext {
	var p = new(ListArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listArray
	return p
}

func InitEmptyListArrayContext(p *ListArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listArray
}

func (*ListArrayContext) IsListArrayContext() {}

func NewListArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArrayContext {
	var p = new(ListArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listArray

	return p
}

func (s *ListArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArrayContext) Get_ID() antlr.Token { return s._ID }

func (s *ListArrayContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListArrayContext) GetList() IListArrayContext { return s.list }

func (s *ListArrayContext) Get_expr() IExprContext { return s._expr }

func (s *ListArrayContext) SetList(v IListArrayContext) { s.list = v }

func (s *ListArrayContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListArrayContext) GetP() interfaces.Expression { return s.p }

func (s *ListArrayContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ListArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListArrayContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *ListArrayContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListArrayContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ListArrayContext) ListArray() IListArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ListArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListArray(s)
	}
}

func (s *ListArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListArray(s)
	}
}

func (s *ListArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitListArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) ListArray() (localctx IListArrayContext) {
	return p.listArray(0)
}

func (p *SwiftGrammarParser) listArray(_p int) (localctx IListArrayContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, SwiftGrammarParserRULE_listArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*ListArrayContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ListArrayContext).p = expressions.NewCallVar((func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListArrayContext(p, _parentctx, _parentState)
			localctx.(*ListArrayContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
			p.SetState(472)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(473)
				p.Match(SwiftGrammarParserCORIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(474)

				var _x = p.expr(0)

				localctx.(*ListArrayContext)._expr = _x
			}
			{
				p.SetState(475)
				p.Match(SwiftGrammarParserCORDER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			localctx.(*ListArrayContext).p = expressions.NewArrayAccess((func() antlr.Token {
				if localctx.(*ListArrayContext).GetList() == nil {
					return nil
				} else {
					return localctx.(*ListArrayContext).GetList().GetStart()
				}
			}()).GetLine(), (func() antlr.Token {
				if localctx.(*ListArrayContext).GetList() == nil {
					return nil
				} else {
					return localctx.(*ListArrayContext).GetList().GetStart()
				}
			}()).GetColumn(), localctx.(*ListArrayContext).GetList().GetP(), localctx.(*ListArrayContext).Get_expr().GetE())

		}
		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ListavalContext = nil
		if localctx != nil {
			t = localctx.(*ListavalContext)
		}
		return p.Listaval_Sempred(t, predIndex)

	case 19:
		var t *Else_ifstmtContext = nil
		if localctx != nil {
			t = localctx.(*Else_ifstmtContext)
		}
		return p.Else_ifstmt_Sempred(t, predIndex)

	case 23:
		var t *CaseContext = nil
		if localctx != nil {
			t = localctx.(*CaseContext)
		}
		return p.Case__Sempred(t, predIndex)

	case 28:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 29:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 30:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) Listaval_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Else_ifstmt_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Case__Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 13)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
