// Code generated from SwiftLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SwiftLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SwiftLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftlexerLexerInit() {
	staticData := &SwiftLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"INT", "FLOAT", "BOOL", "STR", "CHAR", "TRU", "FAL", "PRINT", "IF",
		"ELSE", "ELSE_IF", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN",
		"SWITCH", "BREAK", "CASE", "VAR", "LET", "D_APPEND", "D_REMOVE", "REPEATING",
		"DEFAULT", "COUNT", "STRUCT", "MUTATING", "FUNC", "SELF_D", "PUNTO",
		"INOUT", "GUION_B", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "MAS_IGUAL",
		"MENOS_IGUAL", "NOT", "OR", "AND", "SOLO_AND", "IG", "MAY_IG", "MEN_IG",
		"MAYOR", "MENOR", "MUL", "DIV", "ADD", "MODULO", "SUB", "PARIZQ", "PARDER",
		"LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", "CORDER", "COMA", "PUNTO_COMA",
		"WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 67, 466, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 4, 34, 342, 8,
		34, 11, 34, 12, 34, 343, 1, 34, 1, 34, 4, 34, 348, 8, 34, 11, 34, 12, 34,
		349, 3, 34, 352, 8, 34, 1, 35, 1, 35, 5, 35, 356, 8, 35, 10, 35, 12, 35,
		359, 9, 35, 1, 35, 1, 35, 1, 36, 1, 36, 5, 36, 365, 8, 36, 10, 36, 12,
		36, 368, 9, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39,
		1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1,
		43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47,
		1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1,
		52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57,
		1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1,
		63, 1, 63, 1, 64, 4, 64, 433, 8, 64, 11, 64, 12, 64, 434, 1, 64, 1, 64,
		1, 65, 1, 65, 1, 65, 1, 65, 5, 65, 443, 8, 65, 10, 65, 12, 65, 446, 9,
		65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 5, 66,
		457, 8, 66, 10, 66, 12, 66, 460, 9, 66, 1, 66, 1, 66, 1, 67, 1, 67, 1,
		67, 1, 444, 0, 68, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71,
		36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89,
		45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53,
		107, 54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61,
		123, 62, 125, 63, 127, 64, 129, 65, 131, 66, 133, 67, 135, 0, 1, 0, 7,
		1, 0, 48, 57, 1, 0, 34, 34, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 4, 0, 9, 10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13,
		13, 7, 0, 32, 33, 35, 35, 43, 43, 45, 46, 58, 58, 64, 64, 91, 93, 472,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1,
		0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85,
		1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0,
		93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0,
		0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1,
		0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0,
		115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0,
		0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129,
		1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 1, 137, 1, 0, 0, 0,
		3, 141, 1, 0, 0, 0, 5, 147, 1, 0, 0, 0, 7, 152, 1, 0, 0, 0, 9, 159, 1,
		0, 0, 0, 11, 169, 1, 0, 0, 0, 13, 174, 1, 0, 0, 0, 15, 180, 1, 0, 0, 0,
		17, 186, 1, 0, 0, 0, 19, 189, 1, 0, 0, 0, 21, 194, 1, 0, 0, 0, 23, 202,
		1, 0, 0, 0, 25, 208, 1, 0, 0, 0, 27, 212, 1, 0, 0, 0, 29, 215, 1, 0, 0,
		0, 31, 221, 1, 0, 0, 0, 33, 230, 1, 0, 0, 0, 35, 237, 1, 0, 0, 0, 37, 244,
		1, 0, 0, 0, 39, 250, 1, 0, 0, 0, 41, 255, 1, 0, 0, 0, 43, 259, 1, 0, 0,
		0, 45, 263, 1, 0, 0, 0, 47, 271, 1, 0, 0, 0, 49, 279, 1, 0, 0, 0, 51, 289,
		1, 0, 0, 0, 53, 297, 1, 0, 0, 0, 55, 303, 1, 0, 0, 0, 57, 310, 1, 0, 0,
		0, 59, 319, 1, 0, 0, 0, 61, 324, 1, 0, 0, 0, 63, 330, 1, 0, 0, 0, 65, 332,
		1, 0, 0, 0, 67, 338, 1, 0, 0, 0, 69, 341, 1, 0, 0, 0, 71, 353, 1, 0, 0,
		0, 73, 362, 1, 0, 0, 0, 75, 369, 1, 0, 0, 0, 77, 372, 1, 0, 0, 0, 79, 375,
		1, 0, 0, 0, 81, 378, 1, 0, 0, 0, 83, 381, 1, 0, 0, 0, 85, 383, 1, 0, 0,
		0, 87, 386, 1, 0, 0, 0, 89, 389, 1, 0, 0, 0, 91, 391, 1, 0, 0, 0, 93, 393,
		1, 0, 0, 0, 95, 396, 1, 0, 0, 0, 97, 399, 1, 0, 0, 0, 99, 401, 1, 0, 0,
		0, 101, 403, 1, 0, 0, 0, 103, 405, 1, 0, 0, 0, 105, 407, 1, 0, 0, 0, 107,
		409, 1, 0, 0, 0, 109, 411, 1, 0, 0, 0, 111, 413, 1, 0, 0, 0, 113, 415,
		1, 0, 0, 0, 115, 417, 1, 0, 0, 0, 117, 419, 1, 0, 0, 0, 119, 421, 1, 0,
		0, 0, 121, 423, 1, 0, 0, 0, 123, 425, 1, 0, 0, 0, 125, 427, 1, 0, 0, 0,
		127, 429, 1, 0, 0, 0, 129, 432, 1, 0, 0, 0, 131, 438, 1, 0, 0, 0, 133,
		452, 1, 0, 0, 0, 135, 463, 1, 0, 0, 0, 137, 138, 5, 73, 0, 0, 138, 139,
		5, 110, 0, 0, 139, 140, 5, 116, 0, 0, 140, 2, 1, 0, 0, 0, 141, 142, 5,
		70, 0, 0, 142, 143, 5, 108, 0, 0, 143, 144, 5, 111, 0, 0, 144, 145, 5,
		97, 0, 0, 145, 146, 5, 116, 0, 0, 146, 4, 1, 0, 0, 0, 147, 148, 5, 66,
		0, 0, 148, 149, 5, 111, 0, 0, 149, 150, 5, 111, 0, 0, 150, 151, 5, 108,
		0, 0, 151, 6, 1, 0, 0, 0, 152, 153, 5, 83, 0, 0, 153, 154, 5, 116, 0, 0,
		154, 155, 5, 114, 0, 0, 155, 156, 5, 105, 0, 0, 156, 157, 5, 110, 0, 0,
		157, 158, 5, 103, 0, 0, 158, 8, 1, 0, 0, 0, 159, 160, 5, 67, 0, 0, 160,
		161, 5, 104, 0, 0, 161, 162, 5, 97, 0, 0, 162, 163, 5, 114, 0, 0, 163,
		164, 5, 97, 0, 0, 164, 165, 5, 99, 0, 0, 165, 166, 5, 116, 0, 0, 166, 167,
		5, 101, 0, 0, 167, 168, 5, 114, 0, 0, 168, 10, 1, 0, 0, 0, 169, 170, 5,
		116, 0, 0, 170, 171, 5, 114, 0, 0, 171, 172, 5, 117, 0, 0, 172, 173, 5,
		101, 0, 0, 173, 12, 1, 0, 0, 0, 174, 175, 5, 102, 0, 0, 175, 176, 5, 97,
		0, 0, 176, 177, 5, 108, 0, 0, 177, 178, 5, 115, 0, 0, 178, 179, 5, 101,
		0, 0, 179, 14, 1, 0, 0, 0, 180, 181, 5, 112, 0, 0, 181, 182, 5, 114, 0,
		0, 182, 183, 5, 105, 0, 0, 183, 184, 5, 110, 0, 0, 184, 185, 5, 116, 0,
		0, 185, 16, 1, 0, 0, 0, 186, 187, 5, 105, 0, 0, 187, 188, 5, 102, 0, 0,
		188, 18, 1, 0, 0, 0, 189, 190, 5, 101, 0, 0, 190, 191, 5, 108, 0, 0, 191,
		192, 5, 115, 0, 0, 192, 193, 5, 101, 0, 0, 193, 20, 1, 0, 0, 0, 194, 195,
		5, 101, 0, 0, 195, 196, 5, 108, 0, 0, 196, 197, 5, 115, 0, 0, 197, 198,
		5, 101, 0, 0, 198, 199, 5, 32, 0, 0, 199, 200, 5, 105, 0, 0, 200, 201,
		5, 102, 0, 0, 201, 22, 1, 0, 0, 0, 202, 203, 5, 119, 0, 0, 203, 204, 5,
		104, 0, 0, 204, 205, 5, 105, 0, 0, 205, 206, 5, 108, 0, 0, 206, 207, 5,
		101, 0, 0, 207, 24, 1, 0, 0, 0, 208, 209, 5, 102, 0, 0, 209, 210, 5, 111,
		0, 0, 210, 211, 5, 114, 0, 0, 211, 26, 1, 0, 0, 0, 212, 213, 5, 105, 0,
		0, 213, 214, 5, 110, 0, 0, 214, 28, 1, 0, 0, 0, 215, 216, 5, 103, 0, 0,
		216, 217, 5, 117, 0, 0, 217, 218, 5, 97, 0, 0, 218, 219, 5, 114, 0, 0,
		219, 220, 5, 100, 0, 0, 220, 30, 1, 0, 0, 0, 221, 222, 5, 99, 0, 0, 222,
		223, 5, 111, 0, 0, 223, 224, 5, 110, 0, 0, 224, 225, 5, 116, 0, 0, 225,
		226, 5, 105, 0, 0, 226, 227, 5, 110, 0, 0, 227, 228, 5, 117, 0, 0, 228,
		229, 5, 101, 0, 0, 229, 32, 1, 0, 0, 0, 230, 231, 5, 114, 0, 0, 231, 232,
		5, 101, 0, 0, 232, 233, 5, 116, 0, 0, 233, 234, 5, 117, 0, 0, 234, 235,
		5, 114, 0, 0, 235, 236, 5, 110, 0, 0, 236, 34, 1, 0, 0, 0, 237, 238, 5,
		115, 0, 0, 238, 239, 5, 119, 0, 0, 239, 240, 5, 105, 0, 0, 240, 241, 5,
		116, 0, 0, 241, 242, 5, 99, 0, 0, 242, 243, 5, 104, 0, 0, 243, 36, 1, 0,
		0, 0, 244, 245, 5, 98, 0, 0, 245, 246, 5, 114, 0, 0, 246, 247, 5, 101,
		0, 0, 247, 248, 5, 97, 0, 0, 248, 249, 5, 107, 0, 0, 249, 38, 1, 0, 0,
		0, 250, 251, 5, 99, 0, 0, 251, 252, 5, 97, 0, 0, 252, 253, 5, 115, 0, 0,
		253, 254, 5, 101, 0, 0, 254, 40, 1, 0, 0, 0, 255, 256, 5, 118, 0, 0, 256,
		257, 5, 97, 0, 0, 257, 258, 5, 114, 0, 0, 258, 42, 1, 0, 0, 0, 259, 260,
		5, 108, 0, 0, 260, 261, 5, 101, 0, 0, 261, 262, 5, 116, 0, 0, 262, 44,
		1, 0, 0, 0, 263, 264, 5, 46, 0, 0, 264, 265, 5, 97, 0, 0, 265, 266, 5,
		112, 0, 0, 266, 267, 5, 112, 0, 0, 267, 268, 5, 101, 0, 0, 268, 269, 5,
		110, 0, 0, 269, 270, 5, 100, 0, 0, 270, 46, 1, 0, 0, 0, 271, 272, 5, 46,
		0, 0, 272, 273, 5, 114, 0, 0, 273, 274, 5, 101, 0, 0, 274, 275, 5, 109,
		0, 0, 275, 276, 5, 111, 0, 0, 276, 277, 5, 118, 0, 0, 277, 278, 5, 101,
		0, 0, 278, 48, 1, 0, 0, 0, 279, 280, 5, 114, 0, 0, 280, 281, 5, 101, 0,
		0, 281, 282, 5, 112, 0, 0, 282, 283, 5, 101, 0, 0, 283, 284, 5, 97, 0,
		0, 284, 285, 5, 116, 0, 0, 285, 286, 5, 105, 0, 0, 286, 287, 5, 110, 0,
		0, 287, 288, 5, 103, 0, 0, 288, 50, 1, 0, 0, 0, 289, 290, 5, 100, 0, 0,
		290, 291, 5, 101, 0, 0, 291, 292, 5, 102, 0, 0, 292, 293, 5, 97, 0, 0,
		293, 294, 5, 117, 0, 0, 294, 295, 5, 108, 0, 0, 295, 296, 5, 116, 0, 0,
		296, 52, 1, 0, 0, 0, 297, 298, 5, 99, 0, 0, 298, 299, 5, 111, 0, 0, 299,
		300, 5, 117, 0, 0, 300, 301, 5, 110, 0, 0, 301, 302, 5, 116, 0, 0, 302,
		54, 1, 0, 0, 0, 303, 304, 5, 83, 0, 0, 304, 305, 5, 84, 0, 0, 305, 306,
		5, 82, 0, 0, 306, 307, 5, 85, 0, 0, 307, 308, 5, 67, 0, 0, 308, 309, 5,
		84, 0, 0, 309, 56, 1, 0, 0, 0, 310, 311, 5, 109, 0, 0, 311, 312, 5, 117,
		0, 0, 312, 313, 5, 116, 0, 0, 313, 314, 5, 97, 0, 0, 314, 315, 5, 116,
		0, 0, 315, 316, 5, 105, 0, 0, 316, 317, 5, 110, 0, 0, 317, 318, 5, 103,
		0, 0, 318, 58, 1, 0, 0, 0, 319, 320, 5, 102, 0, 0, 320, 321, 5, 117, 0,
		0, 321, 322, 5, 110, 0, 0, 322, 323, 5, 99, 0, 0, 323, 60, 1, 0, 0, 0,
		324, 325, 5, 115, 0, 0, 325, 326, 5, 101, 0, 0, 326, 327, 5, 108, 0, 0,
		327, 328, 5, 102, 0, 0, 328, 329, 5, 46, 0, 0, 329, 62, 1, 0, 0, 0, 330,
		331, 5, 46, 0, 0, 331, 64, 1, 0, 0, 0, 332, 333, 5, 105, 0, 0, 333, 334,
		5, 110, 0, 0, 334, 335, 5, 111, 0, 0, 335, 336, 5, 117, 0, 0, 336, 337,
		5, 116, 0, 0, 337, 66, 1, 0, 0, 0, 338, 339, 5, 95, 0, 0, 339, 68, 1, 0,
		0, 0, 340, 342, 7, 0, 0, 0, 341, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0,
		343, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 351, 1, 0, 0, 0, 345,
		347, 5, 46, 0, 0, 346, 348, 7, 0, 0, 0, 347, 346, 1, 0, 0, 0, 348, 349,
		1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 352, 1, 0,
		0, 0, 351, 345, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 70, 1, 0, 0, 0,
		353, 357, 5, 34, 0, 0, 354, 356, 8, 1, 0, 0, 355, 354, 1, 0, 0, 0, 356,
		359, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 360,
		1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 360, 361, 5, 34, 0, 0, 361, 72, 1, 0,
		0, 0, 362, 366, 7, 2, 0, 0, 363, 365, 7, 3, 0, 0, 364, 363, 1, 0, 0, 0,
		365, 368, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367,
		74, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 369, 370, 5, 33, 0, 0, 370, 371,
		5, 61, 0, 0, 371, 76, 1, 0, 0, 0, 372, 373, 5, 61, 0, 0, 373, 374, 5, 61,
		0, 0, 374, 78, 1, 0, 0, 0, 375, 376, 5, 43, 0, 0, 376, 377, 5, 61, 0, 0,
		377, 80, 1, 0, 0, 0, 378, 379, 5, 45, 0, 0, 379, 380, 5, 61, 0, 0, 380,
		82, 1, 0, 0, 0, 381, 382, 5, 33, 0, 0, 382, 84, 1, 0, 0, 0, 383, 384, 5,
		124, 0, 0, 384, 385, 5, 124, 0, 0, 385, 86, 1, 0, 0, 0, 386, 387, 5, 38,
		0, 0, 387, 388, 5, 38, 0, 0, 388, 88, 1, 0, 0, 0, 389, 390, 5, 38, 0, 0,
		390, 90, 1, 0, 0, 0, 391, 392, 5, 61, 0, 0, 392, 92, 1, 0, 0, 0, 393, 394,
		5, 62, 0, 0, 394, 395, 5, 61, 0, 0, 395, 94, 1, 0, 0, 0, 396, 397, 5, 60,
		0, 0, 397, 398, 5, 61, 0, 0, 398, 96, 1, 0, 0, 0, 399, 400, 5, 62, 0, 0,
		400, 98, 1, 0, 0, 0, 401, 402, 5, 60, 0, 0, 402, 100, 1, 0, 0, 0, 403,
		404, 5, 42, 0, 0, 404, 102, 1, 0, 0, 0, 405, 406, 5, 47, 0, 0, 406, 104,
		1, 0, 0, 0, 407, 408, 5, 43, 0, 0, 408, 106, 1, 0, 0, 0, 409, 410, 5, 37,
		0, 0, 410, 108, 1, 0, 0, 0, 411, 412, 5, 45, 0, 0, 412, 110, 1, 0, 0, 0,
		413, 414, 5, 40, 0, 0, 414, 112, 1, 0, 0, 0, 415, 416, 5, 41, 0, 0, 416,
		114, 1, 0, 0, 0, 417, 418, 5, 123, 0, 0, 418, 116, 1, 0, 0, 0, 419, 420,
		5, 125, 0, 0, 420, 118, 1, 0, 0, 0, 421, 422, 5, 58, 0, 0, 422, 120, 1,
		0, 0, 0, 423, 424, 5, 91, 0, 0, 424, 122, 1, 0, 0, 0, 425, 426, 5, 93,
		0, 0, 426, 124, 1, 0, 0, 0, 427, 428, 5, 44, 0, 0, 428, 126, 1, 0, 0, 0,
		429, 430, 5, 59, 0, 0, 430, 128, 1, 0, 0, 0, 431, 433, 7, 4, 0, 0, 432,
		431, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 434, 435,
		1, 0, 0, 0, 435, 436, 1, 0, 0, 0, 436, 437, 6, 64, 0, 0, 437, 130, 1, 0,
		0, 0, 438, 439, 5, 47, 0, 0, 439, 440, 5, 42, 0, 0, 440, 444, 1, 0, 0,
		0, 441, 443, 9, 0, 0, 0, 442, 441, 1, 0, 0, 0, 443, 446, 1, 0, 0, 0, 444,
		445, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 445, 447, 1, 0, 0, 0, 446, 444,
		1, 0, 0, 0, 447, 448, 5, 42, 0, 0, 448, 449, 5, 47, 0, 0, 449, 450, 1,
		0, 0, 0, 450, 451, 6, 65, 0, 0, 451, 132, 1, 0, 0, 0, 452, 453, 5, 47,
		0, 0, 453, 454, 5, 47, 0, 0, 454, 458, 1, 0, 0, 0, 455, 457, 8, 5, 0, 0,
		456, 455, 1, 0, 0, 0, 457, 460, 1, 0, 0, 0, 458, 456, 1, 0, 0, 0, 458,
		459, 1, 0, 0, 0, 459, 461, 1, 0, 0, 0, 460, 458, 1, 0, 0, 0, 461, 462,
		6, 66, 0, 0, 462, 134, 1, 0, 0, 0, 463, 464, 5, 92, 0, 0, 464, 465, 7,
		6, 0, 0, 465, 136, 1, 0, 0, 0, 9, 0, 343, 349, 351, 357, 366, 434, 444,
		458, 1, 6, 0, 0,
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

// SwiftLexerInit initializes any static state used to implement SwiftLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSwiftLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftLexerInit() {
	staticData := &SwiftLexerLexerStaticData
	staticData.once.Do(swiftlexerLexerInit)
}

// NewSwiftLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSwiftLexer(input antlr.CharStream) *SwiftLexer {
	SwiftLexerInit()
	l := new(SwiftLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SwiftLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SwiftLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SwiftLexer tokens.
const (
	SwiftLexerINT          = 1
	SwiftLexerFLOAT        = 2
	SwiftLexerBOOL         = 3
	SwiftLexerSTR          = 4
	SwiftLexerCHAR         = 5
	SwiftLexerTRU          = 6
	SwiftLexerFAL          = 7
	SwiftLexerPRINT        = 8
	SwiftLexerIF           = 9
	SwiftLexerELSE         = 10
	SwiftLexerELSE_IF      = 11
	SwiftLexerWHILE        = 12
	SwiftLexerFOR          = 13
	SwiftLexerIN           = 14
	SwiftLexerGUARD        = 15
	SwiftLexerCONTINUE     = 16
	SwiftLexerRETURN       = 17
	SwiftLexerSWITCH       = 18
	SwiftLexerBREAK        = 19
	SwiftLexerCASE         = 20
	SwiftLexerVAR          = 21
	SwiftLexerLET          = 22
	SwiftLexerD_APPEND     = 23
	SwiftLexerD_REMOVE     = 24
	SwiftLexerREPEATING    = 25
	SwiftLexerDEFAULT      = 26
	SwiftLexerCOUNT        = 27
	SwiftLexerSTRUCT       = 28
	SwiftLexerMUTATING     = 29
	SwiftLexerFUNC         = 30
	SwiftLexerSELF_D       = 31
	SwiftLexerPUNTO        = 32
	SwiftLexerINOUT        = 33
	SwiftLexerGUION_B      = 34
	SwiftLexerNUMBER       = 35
	SwiftLexerSTRING       = 36
	SwiftLexerID           = 37
	SwiftLexerDIF          = 38
	SwiftLexerIG_IG        = 39
	SwiftLexerMAS_IGUAL    = 40
	SwiftLexerMENOS_IGUAL  = 41
	SwiftLexerNOT          = 42
	SwiftLexerOR           = 43
	SwiftLexerAND          = 44
	SwiftLexerSOLO_AND     = 45
	SwiftLexerIG           = 46
	SwiftLexerMAY_IG       = 47
	SwiftLexerMEN_IG       = 48
	SwiftLexerMAYOR        = 49
	SwiftLexerMENOR        = 50
	SwiftLexerMUL          = 51
	SwiftLexerDIV          = 52
	SwiftLexerADD          = 53
	SwiftLexerMODULO       = 54
	SwiftLexerSUB          = 55
	SwiftLexerPARIZQ       = 56
	SwiftLexerPARDER       = 57
	SwiftLexerLLAVEIZQ     = 58
	SwiftLexerLLAVEDER     = 59
	SwiftLexerD_PTS        = 60
	SwiftLexerCORIZQ       = 61
	SwiftLexerCORDER       = 62
	SwiftLexerCOMA         = 63
	SwiftLexerPUNTO_COMA   = 64
	SwiftLexerWHITESPACE   = 65
	SwiftLexerCOMMENT      = 66
	SwiftLexerLINE_COMMENT = 67
)
