// Code generated from .\scripting\parser\Bricks.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
		"", "'allParts'", "'lost'", "'set'", "'('", "')'", "'setList'", "'partList'",
		"'sum'", "','", "'subtract'", "'max'", "'mergeByColor'", "'mergeByVariant'",
		"'sort'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "INT", "SET_NUM",
		"WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "DIGIT", "INT", "SET_NUM",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 150, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 4, 15, 132,
		8, 15, 11, 15, 12, 15, 133, 1, 16, 4, 16, 137, 8, 16, 11, 16, 12, 16, 138,
		1, 16, 1, 16, 1, 16, 1, 17, 4, 17, 145, 8, 17, 11, 17, 12, 17, 146, 1,
		17, 1, 17, 0, 0, 18, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 0, 31, 15, 33, 16, 35,
		17, 1, 0, 3, 5, 0, 48, 48, 57, 57, 226, 226, 8220, 8220, 8364, 8364, 1,
		0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 151, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		1, 37, 1, 0, 0, 0, 3, 46, 1, 0, 0, 0, 5, 51, 1, 0, 0, 0, 7, 55, 1, 0, 0,
		0, 9, 57, 1, 0, 0, 0, 11, 59, 1, 0, 0, 0, 13, 67, 1, 0, 0, 0, 15, 76, 1,
		0, 0, 0, 17, 80, 1, 0, 0, 0, 19, 82, 1, 0, 0, 0, 21, 91, 1, 0, 0, 0, 23,
		95, 1, 0, 0, 0, 25, 108, 1, 0, 0, 0, 27, 123, 1, 0, 0, 0, 29, 128, 1, 0,
		0, 0, 31, 131, 1, 0, 0, 0, 33, 136, 1, 0, 0, 0, 35, 144, 1, 0, 0, 0, 37,
		38, 5, 97, 0, 0, 38, 39, 5, 108, 0, 0, 39, 40, 5, 108, 0, 0, 40, 41, 5,
		80, 0, 0, 41, 42, 5, 97, 0, 0, 42, 43, 5, 114, 0, 0, 43, 44, 5, 116, 0,
		0, 44, 45, 5, 115, 0, 0, 45, 2, 1, 0, 0, 0, 46, 47, 5, 108, 0, 0, 47, 48,
		5, 111, 0, 0, 48, 49, 5, 115, 0, 0, 49, 50, 5, 116, 0, 0, 50, 4, 1, 0,
		0, 0, 51, 52, 5, 115, 0, 0, 52, 53, 5, 101, 0, 0, 53, 54, 5, 116, 0, 0,
		54, 6, 1, 0, 0, 0, 55, 56, 5, 40, 0, 0, 56, 8, 1, 0, 0, 0, 57, 58, 5, 41,
		0, 0, 58, 10, 1, 0, 0, 0, 59, 60, 5, 115, 0, 0, 60, 61, 5, 101, 0, 0, 61,
		62, 5, 116, 0, 0, 62, 63, 5, 76, 0, 0, 63, 64, 5, 105, 0, 0, 64, 65, 5,
		115, 0, 0, 65, 66, 5, 116, 0, 0, 66, 12, 1, 0, 0, 0, 67, 68, 5, 112, 0,
		0, 68, 69, 5, 97, 0, 0, 69, 70, 5, 114, 0, 0, 70, 71, 5, 116, 0, 0, 71,
		72, 5, 76, 0, 0, 72, 73, 5, 105, 0, 0, 73, 74, 5, 115, 0, 0, 74, 75, 5,
		116, 0, 0, 75, 14, 1, 0, 0, 0, 76, 77, 5, 115, 0, 0, 77, 78, 5, 117, 0,
		0, 78, 79, 5, 109, 0, 0, 79, 16, 1, 0, 0, 0, 80, 81, 5, 44, 0, 0, 81, 18,
		1, 0, 0, 0, 82, 83, 5, 115, 0, 0, 83, 84, 5, 117, 0, 0, 84, 85, 5, 98,
		0, 0, 85, 86, 5, 116, 0, 0, 86, 87, 5, 114, 0, 0, 87, 88, 5, 97, 0, 0,
		88, 89, 5, 99, 0, 0, 89, 90, 5, 116, 0, 0, 90, 20, 1, 0, 0, 0, 91, 92,
		5, 109, 0, 0, 92, 93, 5, 97, 0, 0, 93, 94, 5, 120, 0, 0, 94, 22, 1, 0,
		0, 0, 95, 96, 5, 109, 0, 0, 96, 97, 5, 101, 0, 0, 97, 98, 5, 114, 0, 0,
		98, 99, 5, 103, 0, 0, 99, 100, 5, 101, 0, 0, 100, 101, 5, 66, 0, 0, 101,
		102, 5, 121, 0, 0, 102, 103, 5, 67, 0, 0, 103, 104, 5, 111, 0, 0, 104,
		105, 5, 108, 0, 0, 105, 106, 5, 111, 0, 0, 106, 107, 5, 114, 0, 0, 107,
		24, 1, 0, 0, 0, 108, 109, 5, 109, 0, 0, 109, 110, 5, 101, 0, 0, 110, 111,
		5, 114, 0, 0, 111, 112, 5, 103, 0, 0, 112, 113, 5, 101, 0, 0, 113, 114,
		5, 66, 0, 0, 114, 115, 5, 121, 0, 0, 115, 116, 5, 86, 0, 0, 116, 117, 5,
		97, 0, 0, 117, 118, 5, 114, 0, 0, 118, 119, 5, 105, 0, 0, 119, 120, 5,
		97, 0, 0, 120, 121, 5, 110, 0, 0, 121, 122, 5, 116, 0, 0, 122, 26, 1, 0,
		0, 0, 123, 124, 5, 115, 0, 0, 124, 125, 5, 111, 0, 0, 125, 126, 5, 114,
		0, 0, 126, 127, 5, 116, 0, 0, 127, 28, 1, 0, 0, 0, 128, 129, 7, 0, 0, 0,
		129, 30, 1, 0, 0, 0, 130, 132, 3, 29, 14, 0, 131, 130, 1, 0, 0, 0, 132,
		133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 32, 1,
		0, 0, 0, 135, 137, 7, 1, 0, 0, 136, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0,
		0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140,
		141, 5, 45, 0, 0, 141, 142, 7, 1, 0, 0, 142, 34, 1, 0, 0, 0, 143, 145,
		7, 2, 0, 0, 144, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 144, 1, 0,
		0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 6, 17, 0, 0,
		149, 36, 1, 0, 0, 0, 4, 0, 133, 138, 146, 1, 6, 0, 0,
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
	BricksLexerINT     = 15
	BricksLexerSET_NUM = 16
	BricksLexerWS      = 17
)
