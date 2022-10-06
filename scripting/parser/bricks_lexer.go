// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BricksLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var brickslexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func brickslexerLexerInit() {
	staticData := &brickslexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "':='", "'save'", "'('", "','", "')'", "'export'", "'build'", "'load'",
		"'allParts'", "'lost'", "'set'", "'setList'", "'partList'", "'partLists'",
		"'sum'", "'subtract'", "'max'", "'sort'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "INT", "BOOL", "STRING", "ID", "SET_NUM", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "DIGIT", "INT", "BOOL", "STRING", "ID", "SET_NUM", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 24, 201, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19,
		4, 19, 154, 8, 19, 11, 19, 12, 19, 155, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 167, 8, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 5, 21, 173, 8, 21, 10, 21, 12, 21, 176, 9, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 5, 22, 182, 8, 22, 10, 22, 12, 22, 185, 9, 22, 1, 23, 4, 23,
		188, 8, 23, 11, 23, 12, 23, 189, 1, 23, 1, 23, 1, 23, 1, 24, 4, 24, 196,
		8, 24, 11, 24, 12, 24, 197, 1, 24, 1, 24, 1, 174, 0, 25, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 0, 39, 19, 41, 20, 43,
		21, 45, 22, 47, 23, 49, 24, 1, 0, 5, 5, 0, 48, 48, 57, 57, 226, 226, 8220,
		8220, 8364, 8364, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 206, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		1, 51, 1, 0, 0, 0, 3, 54, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 61, 1, 0, 0,
		0, 9, 63, 1, 0, 0, 0, 11, 65, 1, 0, 0, 0, 13, 72, 1, 0, 0, 0, 15, 78, 1,
		0, 0, 0, 17, 83, 1, 0, 0, 0, 19, 92, 1, 0, 0, 0, 21, 97, 1, 0, 0, 0, 23,
		101, 1, 0, 0, 0, 25, 109, 1, 0, 0, 0, 27, 118, 1, 0, 0, 0, 29, 128, 1,
		0, 0, 0, 31, 132, 1, 0, 0, 0, 33, 141, 1, 0, 0, 0, 35, 145, 1, 0, 0, 0,
		37, 150, 1, 0, 0, 0, 39, 153, 1, 0, 0, 0, 41, 166, 1, 0, 0, 0, 43, 168,
		1, 0, 0, 0, 45, 179, 1, 0, 0, 0, 47, 187, 1, 0, 0, 0, 49, 195, 1, 0, 0,
		0, 51, 52, 5, 58, 0, 0, 52, 53, 5, 61, 0, 0, 53, 2, 1, 0, 0, 0, 54, 55,
		5, 115, 0, 0, 55, 56, 5, 97, 0, 0, 56, 57, 5, 118, 0, 0, 57, 58, 5, 101,
		0, 0, 58, 4, 1, 0, 0, 0, 59, 60, 5, 40, 0, 0, 60, 6, 1, 0, 0, 0, 61, 62,
		5, 44, 0, 0, 62, 8, 1, 0, 0, 0, 63, 64, 5, 41, 0, 0, 64, 10, 1, 0, 0, 0,
		65, 66, 5, 101, 0, 0, 66, 67, 5, 120, 0, 0, 67, 68, 5, 112, 0, 0, 68, 69,
		5, 111, 0, 0, 69, 70, 5, 114, 0, 0, 70, 71, 5, 116, 0, 0, 71, 12, 1, 0,
		0, 0, 72, 73, 5, 98, 0, 0, 73, 74, 5, 117, 0, 0, 74, 75, 5, 105, 0, 0,
		75, 76, 5, 108, 0, 0, 76, 77, 5, 100, 0, 0, 77, 14, 1, 0, 0, 0, 78, 79,
		5, 108, 0, 0, 79, 80, 5, 111, 0, 0, 80, 81, 5, 97, 0, 0, 81, 82, 5, 100,
		0, 0, 82, 16, 1, 0, 0, 0, 83, 84, 5, 97, 0, 0, 84, 85, 5, 108, 0, 0, 85,
		86, 5, 108, 0, 0, 86, 87, 5, 80, 0, 0, 87, 88, 5, 97, 0, 0, 88, 89, 5,
		114, 0, 0, 89, 90, 5, 116, 0, 0, 90, 91, 5, 115, 0, 0, 91, 18, 1, 0, 0,
		0, 92, 93, 5, 108, 0, 0, 93, 94, 5, 111, 0, 0, 94, 95, 5, 115, 0, 0, 95,
		96, 5, 116, 0, 0, 96, 20, 1, 0, 0, 0, 97, 98, 5, 115, 0, 0, 98, 99, 5,
		101, 0, 0, 99, 100, 5, 116, 0, 0, 100, 22, 1, 0, 0, 0, 101, 102, 5, 115,
		0, 0, 102, 103, 5, 101, 0, 0, 103, 104, 5, 116, 0, 0, 104, 105, 5, 76,
		0, 0, 105, 106, 5, 105, 0, 0, 106, 107, 5, 115, 0, 0, 107, 108, 5, 116,
		0, 0, 108, 24, 1, 0, 0, 0, 109, 110, 5, 112, 0, 0, 110, 111, 5, 97, 0,
		0, 111, 112, 5, 114, 0, 0, 112, 113, 5, 116, 0, 0, 113, 114, 5, 76, 0,
		0, 114, 115, 5, 105, 0, 0, 115, 116, 5, 115, 0, 0, 116, 117, 5, 116, 0,
		0, 117, 26, 1, 0, 0, 0, 118, 119, 5, 112, 0, 0, 119, 120, 5, 97, 0, 0,
		120, 121, 5, 114, 0, 0, 121, 122, 5, 116, 0, 0, 122, 123, 5, 76, 0, 0,
		123, 124, 5, 105, 0, 0, 124, 125, 5, 115, 0, 0, 125, 126, 5, 116, 0, 0,
		126, 127, 5, 115, 0, 0, 127, 28, 1, 0, 0, 0, 128, 129, 5, 115, 0, 0, 129,
		130, 5, 117, 0, 0, 130, 131, 5, 109, 0, 0, 131, 30, 1, 0, 0, 0, 132, 133,
		5, 115, 0, 0, 133, 134, 5, 117, 0, 0, 134, 135, 5, 98, 0, 0, 135, 136,
		5, 116, 0, 0, 136, 137, 5, 114, 0, 0, 137, 138, 5, 97, 0, 0, 138, 139,
		5, 99, 0, 0, 139, 140, 5, 116, 0, 0, 140, 32, 1, 0, 0, 0, 141, 142, 5,
		109, 0, 0, 142, 143, 5, 97, 0, 0, 143, 144, 5, 120, 0, 0, 144, 34, 1, 0,
		0, 0, 145, 146, 5, 115, 0, 0, 146, 147, 5, 111, 0, 0, 147, 148, 5, 114,
		0, 0, 148, 149, 5, 116, 0, 0, 149, 36, 1, 0, 0, 0, 150, 151, 7, 0, 0, 0,
		151, 38, 1, 0, 0, 0, 152, 154, 3, 37, 18, 0, 153, 152, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 40, 1,
		0, 0, 0, 157, 158, 5, 116, 0, 0, 158, 159, 5, 114, 0, 0, 159, 160, 5, 117,
		0, 0, 160, 167, 5, 101, 0, 0, 161, 162, 5, 102, 0, 0, 162, 163, 5, 97,
		0, 0, 163, 164, 5, 108, 0, 0, 164, 165, 5, 115, 0, 0, 165, 167, 5, 101,
		0, 0, 166, 157, 1, 0, 0, 0, 166, 161, 1, 0, 0, 0, 167, 42, 1, 0, 0, 0,
		168, 174, 5, 34, 0, 0, 169, 170, 5, 92, 0, 0, 170, 173, 5, 34, 0, 0, 171,
		173, 9, 0, 0, 0, 172, 169, 1, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 176,
		1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 177, 1, 0,
		0, 0, 176, 174, 1, 0, 0, 0, 177, 178, 5, 34, 0, 0, 178, 44, 1, 0, 0, 0,
		179, 183, 7, 1, 0, 0, 180, 182, 7, 2, 0, 0, 181, 180, 1, 0, 0, 0, 182,
		185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 46, 1,
		0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 188, 7, 3, 0, 0, 187, 186, 1, 0, 0,
		0, 188, 189, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 192, 5, 45, 0, 0, 192, 193, 7, 3, 0, 0, 193, 48,
		1, 0, 0, 0, 194, 196, 7, 4, 0, 0, 195, 194, 1, 0, 0, 0, 196, 197, 1, 0,
		0, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0,
		199, 200, 6, 24, 0, 0, 200, 50, 1, 0, 0, 0, 8, 0, 155, 166, 172, 174, 183,
		189, 197, 1, 6, 0, 0,
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

// BricksLexerInit initializes any static state used to implement BricksLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBricksLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BricksLexerInit() {
	staticData := &brickslexerLexerStaticData
	staticData.once.Do(brickslexerLexerInit)
}

// NewBricksLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBricksLexer(input antlr.CharStream) *BricksLexer {
	BricksLexerInit()
	l := new(BricksLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &brickslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Bricks.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BricksLexer tokens.
const (
	BricksLexerT__0    = 1
	BricksLexerT__1    = 2
	BricksLexerT__2    = 3
	BricksLexerT__3    = 4
	BricksLexerT__4    = 5
	BricksLexerT__5    = 6
	BricksLexerT__6    = 7
	BricksLexerT__7    = 8
	BricksLexerT__8    = 9
	BricksLexerT__9    = 10
	BricksLexerT__10   = 11
	BricksLexerT__11   = 12
	BricksLexerT__12   = 13
	BricksLexerT__13   = 14
	BricksLexerT__14   = 15
	BricksLexerT__15   = 16
	BricksLexerT__16   = 17
	BricksLexerT__17   = 18
	BricksLexerINT     = 19
	BricksLexerBOOL    = 20
	BricksLexerSTRING  = 21
	BricksLexerID      = 22
	BricksLexerSET_NUM = 23
	BricksLexerWS      = 24
)
