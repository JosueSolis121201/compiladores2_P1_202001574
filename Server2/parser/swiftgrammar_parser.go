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
		"", "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'",
		"'false'", "'print'", "'if'", "'else'", "'else if'", "'while'", "'for'",
		"'in'", "'guard'", "'continue'", "'return'", "'switch'", "'break'",
		"'case'", "'var'", "'let'", "'.append'", "'.remove'", "'repeating'",
		"'default'", "'count'", "'STRUCT'", "'mutating'", "'func'", "'self.'",
		"'.'", "'inout'", "'_'", "", "", "", "'!='", "'=='", "'+='", "'-='",
		"'!'", "'||'", "'&&'", "'&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'",
		"'/'", "'+'", "'%'", "'-'", "'('", "')'", "'{'", "'}'", "':'", "'['",
		"']'", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "CHAR", "TRU", "FAL", "PRINT", "IF",
		"ELSE", "ELSE_IF", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN",
		"SWITCH", "BREAK", "CASE", "VAR", "LET", "D_APPEND", "D_REMOVE", "REPEATING",
		"DEFAULT", "COUNT", "STRUCT", "MUTATING", "FUNC", "SELF_D", "PUNTO",
		"INOUT", "GUION_B", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "MAS_IGUAL",
		"MENOS_IGUAL", "NOT", "OR", "AND", "SOLO_AND", "IG", "MAY_IG", "MEN_IG",
		"MAYOR", "MENOR", "MUL", "DIV", "ADD", "MODULO", "SUB", "PARIZQ", "PARDER",
		"LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", "CORDER", "COMA", "PUNTO_COMA",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "guardstmt", "declaracion_vectorstmt",
		"listaval", "jumpstmt", "whilestmt", "forstmt", "printstmt", "ifstmt",
		"elsestmt", "else_ifstmt", "declarationstmt", "switchstmt", "defaultsts",
		"case", "asignacionstmt", "ig_type", "reserv", "types", "expr", "listParams",
		"listArray",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 380, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4,
		1, 54, 8, 1, 11, 1, 12, 1, 55, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 90, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 117, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 5, 5, 124, 8, 5, 10, 5, 12, 5, 127, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 135, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 175, 8,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12,
		186, 8, 12, 10, 12, 12, 12, 189, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13,
		205, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 221, 8, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 231, 8, 16, 10, 16, 12,
		16, 234, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 249, 8, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 3, 19, 255, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 270, 8, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 3, 21, 294, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 5, 21, 346, 8, 21, 10, 21, 12, 21, 349, 9, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 360, 8, 22, 10, 22,
		12, 22, 363, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 5, 23, 375, 8, 23, 10, 23, 12, 23, 378, 9, 23, 1, 23,
		0, 6, 10, 24, 32, 42, 44, 46, 24, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 0, 5, 1, 0, 51, 52,
		2, 0, 53, 53, 55, 55, 2, 0, 47, 47, 49, 49, 2, 0, 48, 48, 50, 50, 1, 0,
		38, 39, 401, 0, 48, 1, 0, 0, 0, 2, 53, 1, 0, 0, 0, 4, 89, 1, 0, 0, 0, 6,
		91, 1, 0, 0, 0, 8, 99, 1, 0, 0, 0, 10, 116, 1, 0, 0, 0, 12, 134, 1, 0,
		0, 0, 14, 136, 1, 0, 0, 0, 16, 143, 1, 0, 0, 0, 18, 152, 1, 0, 0, 0, 20,
		158, 1, 0, 0, 0, 22, 174, 1, 0, 0, 0, 24, 176, 1, 0, 0, 0, 26, 204, 1,
		0, 0, 0, 28, 206, 1, 0, 0, 0, 30, 220, 1, 0, 0, 0, 32, 222, 1, 0, 0, 0,
		34, 235, 1, 0, 0, 0, 36, 248, 1, 0, 0, 0, 38, 254, 1, 0, 0, 0, 40, 269,
		1, 0, 0, 0, 42, 293, 1, 0, 0, 0, 44, 350, 1, 0, 0, 0, 46, 364, 1, 0, 0,
		0, 48, 49, 3, 2, 1, 0, 49, 50, 5, 0, 0, 1, 50, 51, 6, 0, -1, 0, 51, 1,
		1, 0, 0, 0, 52, 54, 3, 4, 2, 0, 53, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0,
		55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 6,
		1, -1, 0, 58, 3, 1, 0, 0, 0, 59, 60, 3, 18, 9, 0, 60, 61, 6, 2, -1, 0,
		61, 90, 1, 0, 0, 0, 62, 63, 3, 20, 10, 0, 63, 64, 6, 2, -1, 0, 64, 90,
		1, 0, 0, 0, 65, 66, 3, 26, 13, 0, 66, 67, 6, 2, -1, 0, 67, 90, 1, 0, 0,
		0, 68, 69, 3, 34, 17, 0, 69, 70, 6, 2, -1, 0, 70, 90, 1, 0, 0, 0, 71, 72,
		3, 28, 14, 0, 72, 73, 6, 2, -1, 0, 73, 90, 1, 0, 0, 0, 74, 75, 3, 14, 7,
		0, 75, 76, 6, 2, -1, 0, 76, 90, 1, 0, 0, 0, 77, 78, 3, 16, 8, 0, 78, 79,
		6, 2, -1, 0, 79, 90, 1, 0, 0, 0, 80, 81, 3, 6, 3, 0, 81, 82, 6, 2, -1,
		0, 82, 90, 1, 0, 0, 0, 83, 84, 3, 12, 6, 0, 84, 85, 6, 2, -1, 0, 85, 90,
		1, 0, 0, 0, 86, 87, 3, 8, 4, 0, 87, 88, 6, 2, -1, 0, 88, 90, 1, 0, 0, 0,
		89, 59, 1, 0, 0, 0, 89, 62, 1, 0, 0, 0, 89, 65, 1, 0, 0, 0, 89, 68, 1,
		0, 0, 0, 89, 71, 1, 0, 0, 0, 89, 74, 1, 0, 0, 0, 89, 77, 1, 0, 0, 0, 89,
		80, 1, 0, 0, 0, 89, 83, 1, 0, 0, 0, 89, 86, 1, 0, 0, 0, 90, 5, 1, 0, 0,
		0, 91, 92, 5, 15, 0, 0, 92, 93, 3, 42, 21, 0, 93, 94, 5, 10, 0, 0, 94,
		95, 5, 58, 0, 0, 95, 96, 3, 2, 1, 0, 96, 97, 5, 59, 0, 0, 97, 98, 6, 3,
		-1, 0, 98, 7, 1, 0, 0, 0, 99, 100, 3, 38, 19, 0, 100, 101, 5, 37, 0, 0,
		101, 102, 5, 60, 0, 0, 102, 103, 5, 61, 0, 0, 103, 104, 3, 40, 20, 0, 104,
		105, 5, 62, 0, 0, 105, 106, 5, 46, 0, 0, 106, 107, 5, 61, 0, 0, 107, 108,
		3, 10, 5, 0, 108, 109, 5, 62, 0, 0, 109, 110, 6, 4, -1, 0, 110, 9, 1, 0,
		0, 0, 111, 112, 6, 5, -1, 0, 112, 113, 3, 42, 21, 0, 113, 114, 6, 5, -1,
		0, 114, 117, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116, 111, 1, 0, 0, 0, 116,
		115, 1, 0, 0, 0, 117, 125, 1, 0, 0, 0, 118, 119, 10, 3, 0, 0, 119, 120,
		5, 63, 0, 0, 120, 121, 3, 42, 21, 0, 121, 122, 6, 5, -1, 0, 122, 124, 1,
		0, 0, 0, 123, 118, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123, 1, 0, 0,
		0, 125, 126, 1, 0, 0, 0, 126, 11, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128,
		129, 5, 16, 0, 0, 129, 135, 6, 6, -1, 0, 130, 131, 5, 17, 0, 0, 131, 135,
		6, 6, -1, 0, 132, 133, 5, 19, 0, 0, 133, 135, 6, 6, -1, 0, 134, 128, 1,
		0, 0, 0, 134, 130, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 13, 1, 0, 0,
		0, 136, 137, 5, 12, 0, 0, 137, 138, 3, 42, 21, 0, 138, 139, 5, 58, 0, 0,
		139, 140, 3, 2, 1, 0, 140, 141, 5, 59, 0, 0, 141, 142, 6, 7, -1, 0, 142,
		15, 1, 0, 0, 0, 143, 144, 5, 13, 0, 0, 144, 145, 5, 37, 0, 0, 145, 146,
		5, 14, 0, 0, 146, 147, 3, 42, 21, 0, 147, 148, 5, 58, 0, 0, 148, 149, 3,
		2, 1, 0, 149, 150, 5, 59, 0, 0, 150, 151, 6, 8, -1, 0, 151, 17, 1, 0, 0,
		0, 152, 153, 5, 8, 0, 0, 153, 154, 5, 56, 0, 0, 154, 155, 3, 42, 21, 0,
		155, 156, 5, 57, 0, 0, 156, 157, 6, 9, -1, 0, 157, 19, 1, 0, 0, 0, 158,
		159, 5, 9, 0, 0, 159, 160, 3, 42, 21, 0, 160, 161, 5, 58, 0, 0, 161, 162,
		3, 2, 1, 0, 162, 163, 5, 59, 0, 0, 163, 164, 3, 24, 12, 0, 164, 165, 3,
		22, 11, 0, 165, 166, 6, 10, -1, 0, 166, 21, 1, 0, 0, 0, 167, 168, 5, 10,
		0, 0, 168, 169, 5, 58, 0, 0, 169, 170, 3, 2, 1, 0, 170, 171, 5, 59, 0,
		0, 171, 172, 6, 11, -1, 0, 172, 175, 1, 0, 0, 0, 173, 175, 1, 0, 0, 0,
		174, 167, 1, 0, 0, 0, 174, 173, 1, 0, 0, 0, 175, 23, 1, 0, 0, 0, 176, 187,
		6, 12, -1, 0, 177, 178, 10, 2, 0, 0, 178, 179, 5, 11, 0, 0, 179, 180, 3,
		42, 21, 0, 180, 181, 5, 58, 0, 0, 181, 182, 3, 2, 1, 0, 182, 183, 5, 59,
		0, 0, 183, 184, 6, 12, -1, 0, 184, 186, 1, 0, 0, 0, 185, 177, 1, 0, 0,
		0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188,
		25, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 191, 3, 38, 19, 0, 191, 192,
		5, 37, 0, 0, 192, 193, 5, 60, 0, 0, 193, 194, 3, 40, 20, 0, 194, 195, 5,
		46, 0, 0, 195, 196, 3, 42, 21, 0, 196, 197, 6, 13, -1, 0, 197, 205, 1,
		0, 0, 0, 198, 199, 3, 38, 19, 0, 199, 200, 5, 37, 0, 0, 200, 201, 5, 46,
		0, 0, 201, 202, 3, 42, 21, 0, 202, 203, 6, 13, -1, 0, 203, 205, 1, 0, 0,
		0, 204, 190, 1, 0, 0, 0, 204, 198, 1, 0, 0, 0, 205, 27, 1, 0, 0, 0, 206,
		207, 5, 18, 0, 0, 207, 208, 3, 42, 21, 0, 208, 209, 5, 58, 0, 0, 209, 210,
		3, 32, 16, 0, 210, 211, 3, 30, 15, 0, 211, 212, 5, 59, 0, 0, 212, 213,
		6, 14, -1, 0, 213, 29, 1, 0, 0, 0, 214, 215, 5, 26, 0, 0, 215, 216, 5,
		60, 0, 0, 216, 217, 3, 2, 1, 0, 217, 218, 6, 15, -1, 0, 218, 221, 1, 0,
		0, 0, 219, 221, 1, 0, 0, 0, 220, 214, 1, 0, 0, 0, 220, 219, 1, 0, 0, 0,
		221, 31, 1, 0, 0, 0, 222, 232, 6, 16, -1, 0, 223, 224, 10, 2, 0, 0, 224,
		225, 5, 20, 0, 0, 225, 226, 3, 42, 21, 0, 226, 227, 5, 60, 0, 0, 227, 228,
		3, 2, 1, 0, 228, 229, 6, 16, -1, 0, 229, 231, 1, 0, 0, 0, 230, 223, 1,
		0, 0, 0, 231, 234, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0,
		0, 233, 33, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 235, 236, 5, 37, 0, 0, 236,
		237, 3, 36, 18, 0, 237, 238, 3, 42, 21, 0, 238, 239, 6, 17, -1, 0, 239,
		35, 1, 0, 0, 0, 240, 241, 5, 46, 0, 0, 241, 249, 6, 18, -1, 0, 242, 243,
		5, 39, 0, 0, 243, 249, 6, 18, -1, 0, 244, 245, 5, 40, 0, 0, 245, 249, 6,
		18, -1, 0, 246, 247, 5, 41, 0, 0, 247, 249, 6, 18, -1, 0, 248, 240, 1,
		0, 0, 0, 248, 242, 1, 0, 0, 0, 248, 244, 1, 0, 0, 0, 248, 246, 1, 0, 0,
		0, 249, 37, 1, 0, 0, 0, 250, 251, 5, 21, 0, 0, 251, 255, 6, 19, -1, 0,
		252, 253, 5, 22, 0, 0, 253, 255, 6, 19, -1, 0, 254, 250, 1, 0, 0, 0, 254,
		252, 1, 0, 0, 0, 255, 39, 1, 0, 0, 0, 256, 257, 5, 1, 0, 0, 257, 270, 6,
		20, -1, 0, 258, 259, 5, 2, 0, 0, 259, 270, 6, 20, -1, 0, 260, 261, 5, 4,
		0, 0, 261, 270, 6, 20, -1, 0, 262, 263, 5, 3, 0, 0, 263, 270, 6, 20, -1,
		0, 264, 265, 5, 5, 0, 0, 265, 270, 6, 20, -1, 0, 266, 267, 5, 61, 0, 0,
		267, 268, 5, 62, 0, 0, 268, 270, 6, 20, -1, 0, 269, 256, 1, 0, 0, 0, 269,
		258, 1, 0, 0, 0, 269, 260, 1, 0, 0, 0, 269, 262, 1, 0, 0, 0, 269, 264,
		1, 0, 0, 0, 269, 266, 1, 0, 0, 0, 270, 41, 1, 0, 0, 0, 271, 272, 6, 21,
		-1, 0, 272, 273, 5, 56, 0, 0, 273, 274, 3, 42, 21, 0, 274, 275, 5, 57,
		0, 0, 275, 276, 6, 21, -1, 0, 276, 294, 1, 0, 0, 0, 277, 278, 3, 46, 23,
		0, 278, 279, 6, 21, -1, 0, 279, 294, 1, 0, 0, 0, 280, 281, 5, 61, 0, 0,
		281, 282, 3, 44, 22, 0, 282, 283, 5, 62, 0, 0, 283, 284, 6, 21, -1, 0,
		284, 294, 1, 0, 0, 0, 285, 286, 5, 35, 0, 0, 286, 294, 6, 21, -1, 0, 287,
		288, 5, 36, 0, 0, 288, 294, 6, 21, -1, 0, 289, 290, 5, 6, 0, 0, 290, 294,
		6, 21, -1, 0, 291, 292, 5, 7, 0, 0, 292, 294, 6, 21, -1, 0, 293, 271, 1,
		0, 0, 0, 293, 277, 1, 0, 0, 0, 293, 280, 1, 0, 0, 0, 293, 285, 1, 0, 0,
		0, 293, 287, 1, 0, 0, 0, 293, 289, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294,
		347, 1, 0, 0, 0, 295, 296, 10, 17, 0, 0, 296, 297, 7, 0, 0, 0, 297, 298,
		3, 42, 21, 18, 298, 299, 6, 21, -1, 0, 299, 346, 1, 0, 0, 0, 300, 301,
		10, 16, 0, 0, 301, 302, 7, 1, 0, 0, 302, 303, 3, 42, 21, 17, 303, 304,
		6, 21, -1, 0, 304, 346, 1, 0, 0, 0, 305, 306, 10, 15, 0, 0, 306, 307, 7,
		2, 0, 0, 307, 308, 3, 42, 21, 16, 308, 309, 6, 21, -1, 0, 309, 346, 1,
		0, 0, 0, 310, 311, 10, 14, 0, 0, 311, 312, 7, 3, 0, 0, 312, 313, 3, 42,
		21, 15, 313, 314, 6, 21, -1, 0, 314, 346, 1, 0, 0, 0, 315, 316, 10, 13,
		0, 0, 316, 317, 7, 4, 0, 0, 317, 318, 3, 42, 21, 14, 318, 319, 6, 21, -1,
		0, 319, 346, 1, 0, 0, 0, 320, 321, 10, 12, 0, 0, 321, 322, 5, 44, 0, 0,
		322, 323, 3, 42, 21, 13, 323, 324, 6, 21, -1, 0, 324, 346, 1, 0, 0, 0,
		325, 326, 10, 11, 0, 0, 326, 327, 5, 43, 0, 0, 327, 328, 3, 42, 21, 12,
		328, 329, 6, 21, -1, 0, 329, 346, 1, 0, 0, 0, 330, 331, 10, 10, 0, 0, 331,
		332, 5, 40, 0, 0, 332, 333, 3, 42, 21, 11, 333, 334, 6, 21, -1, 0, 334,
		346, 1, 0, 0, 0, 335, 336, 10, 9, 0, 0, 336, 337, 5, 41, 0, 0, 337, 338,
		3, 42, 21, 10, 338, 339, 6, 21, -1, 0, 339, 346, 1, 0, 0, 0, 340, 341,
		10, 8, 0, 0, 341, 342, 5, 54, 0, 0, 342, 343, 3, 42, 21, 9, 343, 344, 6,
		21, -1, 0, 344, 346, 1, 0, 0, 0, 345, 295, 1, 0, 0, 0, 345, 300, 1, 0,
		0, 0, 345, 305, 1, 0, 0, 0, 345, 310, 1, 0, 0, 0, 345, 315, 1, 0, 0, 0,
		345, 320, 1, 0, 0, 0, 345, 325, 1, 0, 0, 0, 345, 330, 1, 0, 0, 0, 345,
		335, 1, 0, 0, 0, 345, 340, 1, 0, 0, 0, 346, 349, 1, 0, 0, 0, 347, 345,
		1, 0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 43, 1, 0, 0, 0, 349, 347, 1, 0,
		0, 0, 350, 351, 6, 22, -1, 0, 351, 352, 3, 42, 21, 0, 352, 353, 6, 22,
		-1, 0, 353, 361, 1, 0, 0, 0, 354, 355, 10, 2, 0, 0, 355, 356, 5, 63, 0,
		0, 356, 357, 3, 42, 21, 0, 357, 358, 6, 22, -1, 0, 358, 360, 1, 0, 0, 0,
		359, 354, 1, 0, 0, 0, 360, 363, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 361,
		362, 1, 0, 0, 0, 362, 45, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 364, 365, 6,
		23, -1, 0, 365, 366, 5, 37, 0, 0, 366, 367, 6, 23, -1, 0, 367, 376, 1,
		0, 0, 0, 368, 369, 10, 2, 0, 0, 369, 370, 5, 61, 0, 0, 370, 371, 3, 42,
		21, 0, 371, 372, 5, 62, 0, 0, 372, 373, 6, 23, -1, 0, 373, 375, 1, 0, 0,
		0, 374, 368, 1, 0, 0, 0, 375, 378, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376,
		377, 1, 0, 0, 0, 377, 47, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 18, 55, 89,
		116, 125, 134, 174, 187, 204, 220, 232, 248, 254, 269, 293, 345, 347, 361,
		376,
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
	SwiftGrammarParserTRU          = 6
	SwiftGrammarParserFAL          = 7
	SwiftGrammarParserPRINT        = 8
	SwiftGrammarParserIF           = 9
	SwiftGrammarParserELSE         = 10
	SwiftGrammarParserELSE_IF      = 11
	SwiftGrammarParserWHILE        = 12
	SwiftGrammarParserFOR          = 13
	SwiftGrammarParserIN           = 14
	SwiftGrammarParserGUARD        = 15
	SwiftGrammarParserCONTINUE     = 16
	SwiftGrammarParserRETURN       = 17
	SwiftGrammarParserSWITCH       = 18
	SwiftGrammarParserBREAK        = 19
	SwiftGrammarParserCASE         = 20
	SwiftGrammarParserVAR          = 21
	SwiftGrammarParserLET          = 22
	SwiftGrammarParserD_APPEND     = 23
	SwiftGrammarParserD_REMOVE     = 24
	SwiftGrammarParserREPEATING    = 25
	SwiftGrammarParserDEFAULT      = 26
	SwiftGrammarParserCOUNT        = 27
	SwiftGrammarParserSTRUCT       = 28
	SwiftGrammarParserMUTATING     = 29
	SwiftGrammarParserFUNC         = 30
	SwiftGrammarParserSELF_D       = 31
	SwiftGrammarParserPUNTO        = 32
	SwiftGrammarParserINOUT        = 33
	SwiftGrammarParserGUION_B      = 34
	SwiftGrammarParserNUMBER       = 35
	SwiftGrammarParserSTRING       = 36
	SwiftGrammarParserID           = 37
	SwiftGrammarParserDIF          = 38
	SwiftGrammarParserIG_IG        = 39
	SwiftGrammarParserMAS_IGUAL    = 40
	SwiftGrammarParserMENOS_IGUAL  = 41
	SwiftGrammarParserNOT          = 42
	SwiftGrammarParserOR           = 43
	SwiftGrammarParserAND          = 44
	SwiftGrammarParserSOLO_AND     = 45
	SwiftGrammarParserIG           = 46
	SwiftGrammarParserMAY_IG       = 47
	SwiftGrammarParserMEN_IG       = 48
	SwiftGrammarParserMAYOR        = 49
	SwiftGrammarParserMENOR        = 50
	SwiftGrammarParserMUL          = 51
	SwiftGrammarParserDIV          = 52
	SwiftGrammarParserADD          = 53
	SwiftGrammarParserMODULO       = 54
	SwiftGrammarParserSUB          = 55
	SwiftGrammarParserPARIZQ       = 56
	SwiftGrammarParserPARDER       = 57
	SwiftGrammarParserLLAVEIZQ     = 58
	SwiftGrammarParserLLAVEDER     = 59
	SwiftGrammarParserD_PTS        = 60
	SwiftGrammarParserCORIZQ       = 61
	SwiftGrammarParserCORDER       = 62
	SwiftGrammarParserCOMA         = 63
	SwiftGrammarParserPUNTO_COMA   = 64
	SwiftGrammarParserWHITESPACE   = 65
	SwiftGrammarParserCOMMENT      = 66
	SwiftGrammarParserLINE_COMMENT = 67
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s                      = 0
	SwiftGrammarParserRULE_block                  = 1
	SwiftGrammarParserRULE_instruction            = 2
	SwiftGrammarParserRULE_guardstmt              = 3
	SwiftGrammarParserRULE_declaracion_vectorstmt = 4
	SwiftGrammarParserRULE_listaval               = 5
	SwiftGrammarParserRULE_jumpstmt               = 6
	SwiftGrammarParserRULE_whilestmt              = 7
	SwiftGrammarParserRULE_forstmt                = 8
	SwiftGrammarParserRULE_printstmt              = 9
	SwiftGrammarParserRULE_ifstmt                 = 10
	SwiftGrammarParserRULE_elsestmt               = 11
	SwiftGrammarParserRULE_else_ifstmt            = 12
	SwiftGrammarParserRULE_declarationstmt        = 13
	SwiftGrammarParserRULE_switchstmt             = 14
	SwiftGrammarParserRULE_defaultsts             = 15
	SwiftGrammarParserRULE_case                   = 16
	SwiftGrammarParserRULE_asignacionstmt         = 17
	SwiftGrammarParserRULE_ig_type                = 18
	SwiftGrammarParserRULE_reserv                 = 19
	SwiftGrammarParserRULE_types                  = 20
	SwiftGrammarParserRULE_expr                   = 21
	SwiftGrammarParserRULE_listParams             = 22
	SwiftGrammarParserRULE_listArray              = 23
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
		p.SetState(48)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(49)
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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(52)

				var _x = p.Instruction()

				localctx.(*BlockContext)._instruction = _x
			}
			localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(55)
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
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinst()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)

			var _x = p.Declarationstmt()

			localctx.(*InstructionContext)._declarationstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarationstmt().GetDec()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Asignacionstmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(71)
			p.Switchstmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(74)
			p.Whilestmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(77)
			p.Forstmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(80)
			p.Guardstmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(83)
			p.Jumpstmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(86)
			p.Declaracion_vectorstmt()
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
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_guardstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(SwiftGrammarParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.expr(0)
	}
	{
		p.SetState(93)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Block()
	}
	{
		p.SetState(96)
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
	AllCORIZQ() []antlr.TerminalNode
	CORIZQ(i int) antlr.TerminalNode
	Types() ITypesContext
	AllCORDER() []antlr.TerminalNode
	CORDER(i int) antlr.TerminalNode
	IG() antlr.TerminalNode
	Listaval() IListavalContext

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

func (s *Declaracion_vectorstmtContext) AllCORIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORIZQ)
}

func (s *Declaracion_vectorstmtContext) CORIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, i)
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

func (s *Declaracion_vectorstmtContext) AllCORDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORDER)
}

func (s *Declaracion_vectorstmtContext) CORDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, i)
}

func (s *Declaracion_vectorstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *Declaracion_vectorstmtContext) Listaval() IListavalContext {
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
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_declaracion_vectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Reserv()
	}
	{
		p.SetState(100)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(SwiftGrammarParserD_PTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(SwiftGrammarParserCORIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Types()
	}
	{
		p.SetState(104)
		p.Match(SwiftGrammarParserCORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(SwiftGrammarParserCORIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.listaval(0)
	}
	{
		p.SetState(108)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, SwiftGrammarParserRULE_listaval, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(112)
			p.expr(0)
		}

	case 2:

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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
			p.SetState(118)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(119)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(120)
				p.expr(0)
			}

		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_jumpstmt)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserCONTINUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(SwiftGrammarParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserRETURN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Match(SwiftGrammarParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserBREAK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
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
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(SwiftGrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.expr(0)
	}
	{
		p.SetState(138)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Block()
	}
	{
		p.SetState(140)
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
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_forstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(SwiftGrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(SwiftGrammarParserIN)
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
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Block()
	}
	{
		p.SetState(149)
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
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)

		var _m = p.Match(SwiftGrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)

		var _x = p.expr(0)

		localctx.(*PrintstmtContext)._expr = _x
	}
	{
		p.SetState(155)
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
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_ifstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(SwiftGrammarParserIF)
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
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Block()
	}
	{
		p.SetState(162)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.else_ifstmt(0)
	}
	{
		p.SetState(164)
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
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_elsestmt)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Block()
		}
		{
			p.SetState(170)
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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, SwiftGrammarParserRULE_else_ifstmt, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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
			p.SetState(177)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(178)
				p.Match(SwiftGrammarParserELSE_IF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(179)
				p.expr(0)
			}
			{
				p.SetState(180)
				p.Match(SwiftGrammarParserLLAVEIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(181)
				p.Block()
			}
			{
				p.SetState(182)
				p.Match(SwiftGrammarParserLLAVEDER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_declarationstmt)
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Reserv()
		}
		{
			p.SetState(191)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Types()
		}
		{
			p.SetState(194)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Reserv()
		}
		{
			p.SetState(199)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
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
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_switchstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(SwiftGrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.expr(0)
	}
	{
		p.SetState(208)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.case_(0)
	}
	{
		p.SetState(210)
		p.Defaultsts()
	}
	{
		p.SetState(211)
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
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_defaultsts)
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserDEFAULT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Match(SwiftGrammarParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, SwiftGrammarParserRULE_case, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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
			p.SetState(223)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(224)
				p.Match(SwiftGrammarParserCASE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(225)
				p.expr(0)
			}
			{
				p.SetState(226)
				p.Match(SwiftGrammarParserD_PTS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(227)
				p.Block()
			}

		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_asignacionstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Ig_type()
	}
	{
		p.SetState(237)
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
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_ig_type)
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserIG:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserIG_IG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Match(SwiftGrammarParserIG_IG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserMAS_IGUAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(244)
			p.Match(SwiftGrammarParserMAS_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserMENOS_IGUAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(246)
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
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_reserv)
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Match(SwiftGrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)
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
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_types)
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
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
			p.SetState(258)
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
			p.SetState(260)
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
			p.SetState(262)
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
			p.SetState(264)
			p.Match(SwiftGrammarParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftGrammarParserCORIZQ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(266)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
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

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

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
	PARIZQ() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PARDER() antlr.TerminalNode
	ListArray() IListArrayContext
	CORIZQ() antlr.TerminalNode
	ListParams() IListParamsContext
	CORDER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRU() antlr.TerminalNode
	FAL() antlr.TerminalNode
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
	_expr       IExprContext
	list        IListArrayContext
	_CORIZQ     antlr.Token
	_listParams IListParamsContext
	_NUMBER     antlr.Token
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

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) GetList() IListArrayContext { return s.list }

func (s *ExprContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) SetList(v IListArrayContext) { s.list = v }

func (s *ExprContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
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

func (s *ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
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

func (s *ExprContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
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

func (s *ExprContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
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
	_startState := 42
	p.EnterRecursionRule(localctx, 42, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserPARIZQ:
		{
			p.SetState(272)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(274)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case SwiftGrammarParserID:
		{
			p.SetState(277)

			var _x = p.listArray(0)

			localctx.(*ExprContext).list = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).GetList().GetP()

	case SwiftGrammarParserCORIZQ:
		{
			p.SetState(280)

			var _m = p.Match(SwiftGrammarParserCORIZQ)

			localctx.(*ExprContext)._CORIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)

			var _x = p.listParams(0)

			localctx.(*ExprContext)._listParams = _x
		}
		{
			p.SetState(282)
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

	case SwiftGrammarParserNUMBER:
		{
			p.SetState(285)

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

	case SwiftGrammarParserSTRING:
		{
			p.SetState(287)

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

	case SwiftGrammarParserTRU:
		{
			p.SetState(289)

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

	case SwiftGrammarParserFAL:
		{
			p.SetState(291)

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

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(345)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(296)

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
					p.SetState(297)

					var _x = p.expr(18)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
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
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(301)

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
					p.SetState(302)

					var _x = p.expr(17)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
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
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(306)

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
					p.SetState(307)

					var _x = p.expr(16)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
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
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(311)

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
					p.SetState(312)

					var _x = p.expr(15)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
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
				p.SetState(315)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(316)

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
					p.SetState(317)

					var _x = p.expr(14)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
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
				p.SetState(320)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(321)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(322)

					var _x = p.expr(13)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
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
				p.SetState(325)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(326)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(327)

					var _x = p.expr(12)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
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
				p.SetState(330)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(331)

					var _m = p.Match(SwiftGrammarParserMAS_IGUAL)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(332)

					var _x = p.expr(11)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}

			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(335)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(336)

					var _m = p.Match(SwiftGrammarParserMENOS_IGUAL)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(337)

					var _x = p.expr(10)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}

			case 10:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(340)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(341)

					var _m = p.Match(SwiftGrammarParserMODULO)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(342)

					var _x = p.expr(9)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, SwiftGrammarParserRULE_listParams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)

		var _x = p.expr(0)

		localctx.(*ListParamsContext)._expr = _x
	}

	localctx.(*ListParamsContext).l = []interface{}{}
	localctx.(*ListParamsContext).l = append(localctx.(*ListParamsContext).l, localctx.(*ListParamsContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(361)
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
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParams)
			p.SetState(354)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(355)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(356)

				var _x = p.expr(0)

				localctx.(*ListParamsContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsContext).GetList().GetL(), localctx.(*ListParamsContext).Get_expr().GetE())
			localctx.(*ListParamsContext).l = arr

		}
		p.SetState(363)
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
	_startState := 46
	p.EnterRecursionRule(localctx, 46, SwiftGrammarParserRULE_listArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)

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
	p.SetState(376)
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
			localctx = NewListArrayContext(p, _parentctx, _parentState)
			localctx.(*ListArrayContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
			p.SetState(368)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(369)
				p.Match(SwiftGrammarParserCORIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(370)

				var _x = p.expr(0)

				localctx.(*ListArrayContext)._expr = _x
			}
			{
				p.SetState(371)
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
		p.SetState(378)
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

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ListavalContext = nil
		if localctx != nil {
			t = localctx.(*ListavalContext)
		}
		return p.Listaval_Sempred(t, predIndex)

	case 12:
		var t *Else_ifstmtContext = nil
		if localctx != nil {
			t = localctx.(*Else_ifstmtContext)
		}
		return p.Else_ifstmt_Sempred(t, predIndex)

	case 16:
		var t *CaseContext = nil
		if localctx != nil {
			t = localctx.(*CaseContext)
		}
		return p.Case__Sempred(t, predIndex)

	case 21:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 22:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 23:
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
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
