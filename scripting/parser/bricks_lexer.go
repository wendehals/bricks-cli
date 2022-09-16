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
	modeNames []string
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
    "", "':='", "'save'", "'('", "','", "')'", "'export'", "'load'", "'allParts'", 
    "'lost'", "'set'", "'setList'", "'partList'", "'partLists'", "'sum'", 
    "'subtract'", "'max'", "'sort'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "INT", "BOOL", "STRING", "ID", "SET_NUM", "WS",
  }
  staticData.ruleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
    "T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
    "DIGIT", "INT", "BOOL", "STRING", "ID", "SET_NUM", "WS",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 23, 193, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 
	5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 
	7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 
	9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 
	10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 
	1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 
	13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 
	1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 
	17, 1, 17, 1, 18, 4, 18, 146, 8, 18, 11, 18, 12, 18, 147, 1, 19, 1, 19, 
	1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 159, 8, 19, 1, 
	20, 1, 20, 1, 20, 1, 20, 5, 20, 165, 8, 20, 10, 20, 12, 20, 168, 9, 20, 
	1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 174, 8, 21, 10, 21, 12, 21, 177, 9, 
	21, 1, 22, 4, 22, 180, 8, 22, 11, 22, 12, 22, 181, 1, 22, 1, 22, 1, 22, 
	1, 23, 4, 23, 188, 8, 23, 11, 23, 12, 23, 189, 1, 23, 1, 23, 1, 166, 0, 
	24, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 
	11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 0, 37, 18, 39, 
	19, 41, 20, 43, 21, 45, 22, 47, 23, 1, 0, 5, 3, 0, 48, 48, 57, 57, 8211, 
	8211, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 
	0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 198, 0, 1, 1, 0, 0, 0, 0, 3, 1, 
	0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 
	0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 
	1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 
	27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 
	0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 
	0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 1, 49, 1, 0, 0, 0, 3, 52, 1, 0, 
	0, 0, 5, 57, 1, 0, 0, 0, 7, 59, 1, 0, 0, 0, 9, 61, 1, 0, 0, 0, 11, 63, 
	1, 0, 0, 0, 13, 70, 1, 0, 0, 0, 15, 75, 1, 0, 0, 0, 17, 84, 1, 0, 0, 0, 
	19, 89, 1, 0, 0, 0, 21, 93, 1, 0, 0, 0, 23, 101, 1, 0, 0, 0, 25, 110, 1, 
	0, 0, 0, 27, 120, 1, 0, 0, 0, 29, 124, 1, 0, 0, 0, 31, 133, 1, 0, 0, 0, 
	33, 137, 1, 0, 0, 0, 35, 142, 1, 0, 0, 0, 37, 145, 1, 0, 0, 0, 39, 158, 
	1, 0, 0, 0, 41, 160, 1, 0, 0, 0, 43, 171, 1, 0, 0, 0, 45, 179, 1, 0, 0, 
	0, 47, 187, 1, 0, 0, 0, 49, 50, 5, 58, 0, 0, 50, 51, 5, 61, 0, 0, 51, 2, 
	1, 0, 0, 0, 52, 53, 5, 115, 0, 0, 53, 54, 5, 97, 0, 0, 54, 55, 5, 118, 
	0, 0, 55, 56, 5, 101, 0, 0, 56, 4, 1, 0, 0, 0, 57, 58, 5, 40, 0, 0, 58, 
	6, 1, 0, 0, 0, 59, 60, 5, 44, 0, 0, 60, 8, 1, 0, 0, 0, 61, 62, 5, 41, 0, 
	0, 62, 10, 1, 0, 0, 0, 63, 64, 5, 101, 0, 0, 64, 65, 5, 120, 0, 0, 65, 
	66, 5, 112, 0, 0, 66, 67, 5, 111, 0, 0, 67, 68, 5, 114, 0, 0, 68, 69, 5, 
	116, 0, 0, 69, 12, 1, 0, 0, 0, 70, 71, 5, 108, 0, 0, 71, 72, 5, 111, 0, 
	0, 72, 73, 5, 97, 0, 0, 73, 74, 5, 100, 0, 0, 74, 14, 1, 0, 0, 0, 75, 76, 
	5, 97, 0, 0, 76, 77, 5, 108, 0, 0, 77, 78, 5, 108, 0, 0, 78, 79, 5, 80, 
	0, 0, 79, 80, 5, 97, 0, 0, 80, 81, 5, 114, 0, 0, 81, 82, 5, 116, 0, 0, 
	82, 83, 5, 115, 0, 0, 83, 16, 1, 0, 0, 0, 84, 85, 5, 108, 0, 0, 85, 86, 
	5, 111, 0, 0, 86, 87, 5, 115, 0, 0, 87, 88, 5, 116, 0, 0, 88, 18, 1, 0, 
	0, 0, 89, 90, 5, 115, 0, 0, 90, 91, 5, 101, 0, 0, 91, 92, 5, 116, 0, 0, 
	92, 20, 1, 0, 0, 0, 93, 94, 5, 115, 0, 0, 94, 95, 5, 101, 0, 0, 95, 96, 
	5, 116, 0, 0, 96, 97, 5, 76, 0, 0, 97, 98, 5, 105, 0, 0, 98, 99, 5, 115, 
	0, 0, 99, 100, 5, 116, 0, 0, 100, 22, 1, 0, 0, 0, 101, 102, 5, 112, 0, 
	0, 102, 103, 5, 97, 0, 0, 103, 104, 5, 114, 0, 0, 104, 105, 5, 116, 0, 
	0, 105, 106, 5, 76, 0, 0, 106, 107, 5, 105, 0, 0, 107, 108, 5, 115, 0, 
	0, 108, 109, 5, 116, 0, 0, 109, 24, 1, 0, 0, 0, 110, 111, 5, 112, 0, 0, 
	111, 112, 5, 97, 0, 0, 112, 113, 5, 114, 0, 0, 113, 114, 5, 116, 0, 0, 
	114, 115, 5, 76, 0, 0, 115, 116, 5, 105, 0, 0, 116, 117, 5, 115, 0, 0, 
	117, 118, 5, 116, 0, 0, 118, 119, 5, 115, 0, 0, 119, 26, 1, 0, 0, 0, 120, 
	121, 5, 115, 0, 0, 121, 122, 5, 117, 0, 0, 122, 123, 5, 109, 0, 0, 123, 
	28, 1, 0, 0, 0, 124, 125, 5, 115, 0, 0, 125, 126, 5, 117, 0, 0, 126, 127, 
	5, 98, 0, 0, 127, 128, 5, 116, 0, 0, 128, 129, 5, 114, 0, 0, 129, 130, 
	5, 97, 0, 0, 130, 131, 5, 99, 0, 0, 131, 132, 5, 116, 0, 0, 132, 30, 1, 
	0, 0, 0, 133, 134, 5, 109, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 120, 
	0, 0, 136, 32, 1, 0, 0, 0, 137, 138, 5, 115, 0, 0, 138, 139, 5, 111, 0, 
	0, 139, 140, 5, 114, 0, 0, 140, 141, 5, 116, 0, 0, 141, 34, 1, 0, 0, 0, 
	142, 143, 7, 0, 0, 0, 143, 36, 1, 0, 0, 0, 144, 146, 3, 35, 17, 0, 145, 
	144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 
	1, 0, 0, 0, 148, 38, 1, 0, 0, 0, 149, 150, 5, 116, 0, 0, 150, 151, 5, 114, 
	0, 0, 151, 152, 5, 117, 0, 0, 152, 159, 5, 101, 0, 0, 153, 154, 5, 102, 
	0, 0, 154, 155, 5, 97, 0, 0, 155, 156, 5, 108, 0, 0, 156, 157, 5, 115, 
	0, 0, 157, 159, 5, 101, 0, 0, 158, 149, 1, 0, 0, 0, 158, 153, 1, 0, 0, 
	0, 159, 40, 1, 0, 0, 0, 160, 166, 5, 34, 0, 0, 161, 162, 5, 92, 0, 0, 162, 
	165, 5, 34, 0, 0, 163, 165, 9, 0, 0, 0, 164, 161, 1, 0, 0, 0, 164, 163, 
	1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 166, 164, 1, 0, 
	0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 170, 5, 34, 0, 0, 
	170, 42, 1, 0, 0, 0, 171, 175, 7, 1, 0, 0, 172, 174, 7, 2, 0, 0, 173, 172, 
	1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 
	0, 0, 176, 44, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178, 180, 7, 3, 0, 0, 
	179, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 
	182, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 5, 45, 0, 0, 184, 185, 
	7, 3, 0, 0, 185, 46, 1, 0, 0, 0, 186, 188, 7, 4, 0, 0, 187, 186, 1, 0, 
	0, 0, 188, 189, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 
	190, 191, 1, 0, 0, 0, 191, 192, 6, 23, 0, 0, 192, 48, 1, 0, 0, 0, 8, 0, 
	147, 158, 164, 166, 175, 181, 189, 1, 6, 0, 0,
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
	BricksLexerT__0 = 1
	BricksLexerT__1 = 2
	BricksLexerT__2 = 3
	BricksLexerT__3 = 4
	BricksLexerT__4 = 5
	BricksLexerT__5 = 6
	BricksLexerT__6 = 7
	BricksLexerT__7 = 8
	BricksLexerT__8 = 9
	BricksLexerT__9 = 10
	BricksLexerT__10 = 11
	BricksLexerT__11 = 12
	BricksLexerT__12 = 13
	BricksLexerT__13 = 14
	BricksLexerT__14 = 15
	BricksLexerT__15 = 16
	BricksLexerT__16 = 17
	BricksLexerINT = 18
	BricksLexerBOOL = 19
	BricksLexerSTRING = 20
	BricksLexerID = 21
	BricksLexerSET_NUM = 22
	BricksLexerWS = 23
)

