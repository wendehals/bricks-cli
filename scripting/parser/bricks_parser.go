// Code generated from ./scripting/parser/Bricks.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Bricks

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BricksParser struct {
	*antlr.BaseParser
}

var BricksParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bricksParserInit() {
	staticData := &BricksParserStaticData
	staticData.LiteralNames = []string{
		"", "':='", "'save'", "'('", "','", "')'", "'export'", "'print'", "'pause'",
		"'load'", "'import'", "'allParts'", "'lost'", "'set'", "'userSet'",
		"'setList'", "'partList'", "'partLists'", "'sum'", "'subtract'", "'max'",
		"'sort'", "'quantity'", "'descending'", "'build'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "ID", "INT", "SET_NUM", "BOOL", "STRING",
		"LINE_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"bricks", "command", "expression", "assignment", "save", "export", "print",
		"pause", "collectionOrId", "collection", "load", "import_", "allParts",
		"lost", "set", "userSet", "setList", "partList", "partLists", "sum",
		"subtract", "max", "sort", "build", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 225, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 4, 0, 52, 8,
		0, 11, 0, 12, 0, 53, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 61, 8, 1, 1, 2,
		1, 2, 1, 2, 3, 2, 66, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 98,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 113, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 3, 14, 134, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 3, 15, 143, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 152, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 166, 8, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 4, 19, 175, 8, 19, 11, 19, 12, 19,
		176, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 4, 21, 193, 8, 21, 11, 21, 12, 21, 194,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 204, 8, 22, 1,
		22, 1, 22, 3, 22, 208, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 3, 23, 219, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		24, 0, 0, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 0, 228, 0, 51, 1, 0, 0, 0, 2, 60,
		1, 0, 0, 0, 4, 65, 1, 0, 0, 0, 6, 67, 1, 0, 0, 0, 8, 71, 1, 0, 0, 0, 10,
		78, 1, 0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 90, 1, 0, 0, 0, 16, 97, 1, 0, 0,
		0, 18, 112, 1, 0, 0, 0, 20, 114, 1, 0, 0, 0, 22, 119, 1, 0, 0, 0, 24, 124,
		1, 0, 0, 0, 26, 126, 1, 0, 0, 0, 28, 128, 1, 0, 0, 0, 30, 137, 1, 0, 0,
		0, 32, 146, 1, 0, 0, 0, 34, 155, 1, 0, 0, 0, 36, 160, 1, 0, 0, 0, 38, 169,
		1, 0, 0, 0, 40, 180, 1, 0, 0, 0, 42, 187, 1, 0, 0, 0, 44, 198, 1, 0, 0,
		0, 46, 211, 1, 0, 0, 0, 48, 222, 1, 0, 0, 0, 50, 52, 3, 2, 1, 0, 51, 50,
		1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0,
		54, 1, 1, 0, 0, 0, 55, 61, 3, 6, 3, 0, 56, 61, 3, 8, 4, 0, 57, 61, 3, 10,
		5, 0, 58, 61, 3, 12, 6, 0, 59, 61, 3, 14, 7, 0, 60, 55, 1, 0, 0, 0, 60,
		56, 1, 0, 0, 0, 60, 57, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 59, 1, 0, 0,
		0, 61, 3, 1, 0, 0, 0, 62, 66, 3, 18, 9, 0, 63, 66, 3, 46, 23, 0, 64, 66,
		3, 48, 24, 0, 65, 62, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 64, 1, 0, 0,
		0, 66, 5, 1, 0, 0, 0, 67, 68, 5, 25, 0, 0, 68, 69, 5, 1, 0, 0, 69, 70,
		3, 4, 2, 0, 70, 7, 1, 0, 0, 0, 71, 72, 5, 2, 0, 0, 72, 73, 5, 3, 0, 0,
		73, 74, 3, 4, 2, 0, 74, 75, 5, 4, 0, 0, 75, 76, 5, 29, 0, 0, 76, 77, 5,
		5, 0, 0, 77, 9, 1, 0, 0, 0, 78, 79, 5, 6, 0, 0, 79, 80, 5, 3, 0, 0, 80,
		81, 3, 4, 2, 0, 81, 82, 5, 4, 0, 0, 82, 83, 5, 29, 0, 0, 83, 84, 5, 5,
		0, 0, 84, 11, 1, 0, 0, 0, 85, 86, 5, 7, 0, 0, 86, 87, 5, 3, 0, 0, 87, 88,
		3, 4, 2, 0, 88, 89, 5, 5, 0, 0, 89, 13, 1, 0, 0, 0, 90, 91, 5, 8, 0, 0,
		91, 92, 5, 3, 0, 0, 92, 93, 5, 26, 0, 0, 93, 94, 5, 5, 0, 0, 94, 15, 1,
		0, 0, 0, 95, 98, 3, 18, 9, 0, 96, 98, 3, 48, 24, 0, 97, 95, 1, 0, 0, 0,
		97, 96, 1, 0, 0, 0, 98, 17, 1, 0, 0, 0, 99, 113, 3, 20, 10, 0, 100, 113,
		3, 22, 11, 0, 101, 113, 3, 24, 12, 0, 102, 113, 3, 26, 13, 0, 103, 113,
		3, 28, 14, 0, 104, 113, 3, 30, 15, 0, 105, 113, 3, 32, 16, 0, 106, 113,
		3, 34, 17, 0, 107, 113, 3, 36, 18, 0, 108, 113, 3, 38, 19, 0, 109, 113,
		3, 40, 20, 0, 110, 113, 3, 42, 21, 0, 111, 113, 3, 44, 22, 0, 112, 99,
		1, 0, 0, 0, 112, 100, 1, 0, 0, 0, 112, 101, 1, 0, 0, 0, 112, 102, 1, 0,
		0, 0, 112, 103, 1, 0, 0, 0, 112, 104, 1, 0, 0, 0, 112, 105, 1, 0, 0, 0,
		112, 106, 1, 0, 0, 0, 112, 107, 1, 0, 0, 0, 112, 108, 1, 0, 0, 0, 112,
		109, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113, 19, 1,
		0, 0, 0, 114, 115, 5, 9, 0, 0, 115, 116, 5, 3, 0, 0, 116, 117, 5, 29, 0,
		0, 117, 118, 5, 5, 0, 0, 118, 21, 1, 0, 0, 0, 119, 120, 5, 10, 0, 0, 120,
		121, 5, 3, 0, 0, 121, 122, 5, 29, 0, 0, 122, 123, 5, 5, 0, 0, 123, 23,
		1, 0, 0, 0, 124, 125, 5, 11, 0, 0, 125, 25, 1, 0, 0, 0, 126, 127, 5, 12,
		0, 0, 127, 27, 1, 0, 0, 0, 128, 129, 5, 13, 0, 0, 129, 130, 5, 3, 0, 0,
		130, 133, 5, 27, 0, 0, 131, 132, 5, 4, 0, 0, 132, 134, 5, 28, 0, 0, 133,
		131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136,
		5, 5, 0, 0, 136, 29, 1, 0, 0, 0, 137, 138, 5, 14, 0, 0, 138, 139, 5, 3,
		0, 0, 139, 142, 5, 27, 0, 0, 140, 141, 5, 4, 0, 0, 141, 143, 5, 28, 0,
		0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144,
		145, 5, 5, 0, 0, 145, 31, 1, 0, 0, 0, 146, 147, 5, 15, 0, 0, 147, 148,
		5, 3, 0, 0, 148, 151, 5, 26, 0, 0, 149, 150, 5, 4, 0, 0, 150, 152, 5, 28,
		0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0,
		153, 154, 5, 5, 0, 0, 154, 33, 1, 0, 0, 0, 155, 156, 5, 16, 0, 0, 156,
		157, 5, 3, 0, 0, 157, 158, 5, 26, 0, 0, 158, 159, 5, 5, 0, 0, 159, 35,
		1, 0, 0, 0, 160, 161, 5, 17, 0, 0, 161, 162, 5, 3, 0, 0, 162, 165, 5, 29,
		0, 0, 163, 164, 5, 4, 0, 0, 164, 166, 5, 28, 0, 0, 165, 163, 1, 0, 0, 0,
		165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 5, 5, 0, 0, 168,
		37, 1, 0, 0, 0, 169, 170, 5, 18, 0, 0, 170, 171, 5, 3, 0, 0, 171, 174,
		3, 16, 8, 0, 172, 173, 5, 4, 0, 0, 173, 175, 3, 16, 8, 0, 174, 172, 1,
		0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0,
		0, 177, 178, 1, 0, 0, 0, 178, 179, 5, 5, 0, 0, 179, 39, 1, 0, 0, 0, 180,
		181, 5, 19, 0, 0, 181, 182, 5, 3, 0, 0, 182, 183, 3, 16, 8, 0, 183, 184,
		5, 4, 0, 0, 184, 185, 3, 16, 8, 0, 185, 186, 5, 5, 0, 0, 186, 41, 1, 0,
		0, 0, 187, 188, 5, 20, 0, 0, 188, 189, 5, 3, 0, 0, 189, 192, 3, 16, 8,
		0, 190, 191, 5, 4, 0, 0, 191, 193, 3, 16, 8, 0, 192, 190, 1, 0, 0, 0, 193,
		194, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196,
		1, 0, 0, 0, 196, 197, 5, 5, 0, 0, 197, 43, 1, 0, 0, 0, 198, 199, 5, 21,
		0, 0, 199, 200, 5, 3, 0, 0, 200, 203, 3, 16, 8, 0, 201, 202, 5, 4, 0, 0,
		202, 204, 5, 22, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204,
		207, 1, 0, 0, 0, 205, 206, 5, 4, 0, 0, 206, 208, 5, 23, 0, 0, 207, 205,
		1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 5, 5,
		0, 0, 210, 45, 1, 0, 0, 0, 211, 212, 5, 24, 0, 0, 212, 213, 5, 3, 0, 0,
		213, 214, 3, 16, 8, 0, 214, 215, 5, 4, 0, 0, 215, 218, 3, 16, 8, 0, 216,
		217, 5, 4, 0, 0, 217, 219, 5, 25, 0, 0, 218, 216, 1, 0, 0, 0, 218, 219,
		1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221, 5, 5, 0, 0, 221, 47, 1, 0,
		0, 0, 222, 223, 5, 25, 0, 0, 223, 49, 1, 0, 0, 0, 14, 53, 60, 65, 97, 112,
		133, 142, 151, 165, 176, 194, 203, 207, 218,
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

// BricksParserInit initializes any static state used to implement BricksParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBricksParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BricksParserInit() {
	staticData := &BricksParserStaticData
	staticData.once.Do(bricksParserInit)
}

// NewBricksParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBricksParser(input antlr.TokenStream) *BricksParser {
	BricksParserInit()
	this := new(BricksParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BricksParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Bricks.g4"

	return this
}

// BricksParser tokens.
const (
	BricksParserEOF          = antlr.TokenEOF
	BricksParserT__0         = 1
	BricksParserT__1         = 2
	BricksParserT__2         = 3
	BricksParserT__3         = 4
	BricksParserT__4         = 5
	BricksParserT__5         = 6
	BricksParserT__6         = 7
	BricksParserT__7         = 8
	BricksParserT__8         = 9
	BricksParserT__9         = 10
	BricksParserT__10        = 11
	BricksParserT__11        = 12
	BricksParserT__12        = 13
	BricksParserT__13        = 14
	BricksParserT__14        = 15
	BricksParserT__15        = 16
	BricksParserT__16        = 17
	BricksParserT__17        = 18
	BricksParserT__18        = 19
	BricksParserT__19        = 20
	BricksParserT__20        = 21
	BricksParserT__21        = 22
	BricksParserT__22        = 23
	BricksParserT__23        = 24
	BricksParserID           = 25
	BricksParserINT          = 26
	BricksParserSET_NUM      = 27
	BricksParserBOOL         = 28
	BricksParserSTRING       = 29
	BricksParserLINE_COMMENT = 30
	BricksParserWS           = 31
)

// BricksParser rules.
const (
	BricksParserRULE_bricks         = 0
	BricksParserRULE_command        = 1
	BricksParserRULE_expression     = 2
	BricksParserRULE_assignment     = 3
	BricksParserRULE_save           = 4
	BricksParserRULE_export         = 5
	BricksParserRULE_print          = 6
	BricksParserRULE_pause          = 7
	BricksParserRULE_collectionOrId = 8
	BricksParserRULE_collection     = 9
	BricksParserRULE_load           = 10
	BricksParserRULE_import_        = 11
	BricksParserRULE_allParts       = 12
	BricksParserRULE_lost           = 13
	BricksParserRULE_set            = 14
	BricksParserRULE_userSet        = 15
	BricksParserRULE_setList        = 16
	BricksParserRULE_partList       = 17
	BricksParserRULE_partLists      = 18
	BricksParserRULE_sum            = 19
	BricksParserRULE_subtract       = 20
	BricksParserRULE_max            = 21
	BricksParserRULE_sort           = 22
	BricksParserRULE_build          = 23
	BricksParserRULE_identifier     = 24
)

// IBricksContext is an interface to support dynamic dispatch.
type IBricksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCommand() []ICommandContext
	Command(i int) ICommandContext

	// IsBricksContext differentiates from other interfaces.
	IsBricksContext()
}

type BricksContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBricksContext() *BricksContext {
	var p = new(BricksContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_bricks
	return p
}

func InitEmptyBricksContext(p *BricksContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_bricks
}

func (*BricksContext) IsBricksContext() {}

func NewBricksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BricksContext {
	var p = new(BricksContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_bricks

	return p
}

func (s *BricksContext) GetParser() antlr.Parser { return s.parser }

func (s *BricksContext) AllCommand() []ICommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommandContext); ok {
			len++
		}
	}

	tst := make([]ICommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommandContext); ok {
			tst[i] = t.(ICommandContext)
			i++
		}
	}

	return tst
}

func (s *BricksContext) Command(i int) ICommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
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

	return t.(ICommandContext)
}

func (s *BricksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BricksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BricksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterBricks(s)
	}
}

func (s *BricksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitBricks(s)
	}
}

func (p *BricksParser) Bricks() (localctx IBricksContext) {
	localctx = NewBricksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BricksParserRULE_bricks)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33554884) != 0) {
		{
			p.SetState(50)
			p.Command()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Save() ISaveContext
	Export() IExportContext
	Print_() IPrintContext
	Pause() IPauseContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *CommandContext) Save() ISaveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISaveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISaveContext)
}

func (s *CommandContext) Export() IExportContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExportContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExportContext)
}

func (s *CommandContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *CommandContext) Pause() IPauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPauseContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *BricksParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BricksParserRULE_command)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BricksParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Assignment()
		}

	case BricksParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Save()
		}

	case BricksParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Export()
		}

	case BricksParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.Print_()
		}

	case BricksParserT__7:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)
			p.Pause()
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Collection() ICollectionContext
	Build() IBuildContext
	Identifier() IIdentifierContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Collection() ICollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *ExpressionContext) Build() IBuildContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuildContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuildContext)
}

func (s *ExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BricksParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BricksParserRULE_expression)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BricksParserT__8, BricksParserT__9, BricksParserT__10, BricksParserT__11, BricksParserT__12, BricksParserT__13, BricksParserT__14, BricksParserT__15, BricksParserT__16, BricksParserT__17, BricksParserT__18, BricksParserT__19, BricksParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Collection()
		}

	case BricksParserT__23:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Build()
		}

	case BricksParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Identifier()
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(BricksParserID, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *BricksParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BricksParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(BricksParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(BricksParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Expression()
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

// ISaveContext is an interface to support dynamic dispatch.
type ISaveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	STRING() antlr.TerminalNode

	// IsSaveContext differentiates from other interfaces.
	IsSaveContext()
}

type SaveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySaveContext() *SaveContext {
	var p = new(SaveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_save
	return p
}

func InitEmptySaveContext(p *SaveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_save
}

func (*SaveContext) IsSaveContext() {}

func NewSaveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SaveContext {
	var p = new(SaveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_save

	return p
}

func (s *SaveContext) GetParser() antlr.Parser { return s.parser }

func (s *SaveContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SaveContext) STRING() antlr.TerminalNode {
	return s.GetToken(BricksParserSTRING, 0)
}

func (s *SaveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SaveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SaveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterSave(s)
	}
}

func (s *SaveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitSave(s)
	}
}

func (p *BricksParser) Save() (localctx ISaveContext) {
	localctx = NewSaveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BricksParserRULE_save)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(BricksParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Expression()
	}
	{
		p.SetState(74)
		p.Match(BricksParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(BricksParserT__4)
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

// IExportContext is an interface to support dynamic dispatch.
type IExportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	STRING() antlr.TerminalNode

	// IsExportContext differentiates from other interfaces.
	IsExportContext()
}

type ExportContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportContext() *ExportContext {
	var p = new(ExportContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_export
	return p
}

func InitEmptyExportContext(p *ExportContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_export
}

func (*ExportContext) IsExportContext() {}

func NewExportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportContext {
	var p = new(ExportContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_export

	return p
}

func (s *ExportContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExportContext) STRING() antlr.TerminalNode {
	return s.GetToken(BricksParserSTRING, 0)
}

func (s *ExportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterExport(s)
	}
}

func (s *ExportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitExport(s)
	}
}

func (p *BricksParser) Export() (localctx IExportContext) {
	localctx = NewExportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BricksParserRULE_export)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(BricksParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Expression()
	}
	{
		p.SetState(81)
		p.Match(BricksParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(BricksParserT__4)
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

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (p *BricksParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BricksParserRULE_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(BricksParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Expression()
	}
	{
		p.SetState(88)
		p.Match(BricksParserT__4)
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

// IPauseContext is an interface to support dynamic dispatch.
type IPauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSeconds returns the seconds token.
	GetSeconds() antlr.Token

	// SetSeconds sets the seconds token.
	SetSeconds(antlr.Token)

	// Getter signatures
	INT() antlr.TerminalNode

	// IsPauseContext differentiates from other interfaces.
	IsPauseContext()
}

type PauseContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	seconds antlr.Token
}

func NewEmptyPauseContext() *PauseContext {
	var p = new(PauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_pause
	return p
}

func InitEmptyPauseContext(p *PauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_pause
}

func (*PauseContext) IsPauseContext() {}

func NewPauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PauseContext {
	var p = new(PauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_pause

	return p
}

func (s *PauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PauseContext) GetSeconds() antlr.Token { return s.seconds }

func (s *PauseContext) SetSeconds(v antlr.Token) { s.seconds = v }

func (s *PauseContext) INT() antlr.TerminalNode {
	return s.GetToken(BricksParserINT, 0)
}

func (s *PauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterPause(s)
	}
}

func (s *PauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitPause(s)
	}
}

func (p *BricksParser) Pause() (localctx IPauseContext) {
	localctx = NewPauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BricksParserRULE_pause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(BricksParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)

		var _m = p.Match(BricksParserINT)

		localctx.(*PauseContext).seconds = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(BricksParserT__4)
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

// ICollectionOrIdContext is an interface to support dynamic dispatch.
type ICollectionOrIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Collection() ICollectionContext
	Identifier() IIdentifierContext

	// IsCollectionOrIdContext differentiates from other interfaces.
	IsCollectionOrIdContext()
}

type CollectionOrIdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionOrIdContext() *CollectionOrIdContext {
	var p = new(CollectionOrIdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_collectionOrId
	return p
}

func InitEmptyCollectionOrIdContext(p *CollectionOrIdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_collectionOrId
}

func (*CollectionOrIdContext) IsCollectionOrIdContext() {}

func NewCollectionOrIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionOrIdContext {
	var p = new(CollectionOrIdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_collectionOrId

	return p
}

func (s *CollectionOrIdContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionOrIdContext) Collection() ICollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *CollectionOrIdContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *CollectionOrIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionOrIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionOrIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterCollectionOrId(s)
	}
}

func (s *CollectionOrIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitCollectionOrId(s)
	}
}

func (p *BricksParser) CollectionOrId() (localctx ICollectionOrIdContext) {
	localctx = NewCollectionOrIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BricksParserRULE_collectionOrId)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BricksParserT__8, BricksParserT__9, BricksParserT__10, BricksParserT__11, BricksParserT__12, BricksParserT__13, BricksParserT__14, BricksParserT__15, BricksParserT__16, BricksParserT__17, BricksParserT__18, BricksParserT__19, BricksParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Collection()
		}

	case BricksParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Identifier()
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

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Load() ILoadContext
	Import_() IImport_Context
	AllParts() IAllPartsContext
	Lost() ILostContext
	Set() ISetContext
	UserSet() IUserSetContext
	SetList() ISetListContext
	PartList() IPartListContext
	PartLists() IPartListsContext
	Sum() ISumContext
	Subtract() ISubtractContext
	Max() IMaxContext
	Sort() ISortContext

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_collection
	return p
}

func InitEmptyCollectionContext(p *CollectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_collection
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) Load() ILoadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoadContext)
}

func (s *CollectionContext) Import_() IImport_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImport_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImport_Context)
}

func (s *CollectionContext) AllParts() IAllPartsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllPartsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllPartsContext)
}

func (s *CollectionContext) Lost() ILostContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILostContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILostContext)
}

func (s *CollectionContext) Set() ISetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetContext)
}

func (s *CollectionContext) UserSet() IUserSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserSetContext)
}

func (s *CollectionContext) SetList() ISetListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetListContext)
}

func (s *CollectionContext) PartList() IPartListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPartListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPartListContext)
}

func (s *CollectionContext) PartLists() IPartListsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPartListsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPartListsContext)
}

func (s *CollectionContext) Sum() ISumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISumContext)
}

func (s *CollectionContext) Subtract() ISubtractContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubtractContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubtractContext)
}

func (s *CollectionContext) Max() IMaxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMaxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMaxContext)
}

func (s *CollectionContext) Sort() ISortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortContext)
}

func (s *CollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterCollection(s)
	}
}

func (s *CollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitCollection(s)
	}
}

func (p *BricksParser) Collection() (localctx ICollectionContext) {
	localctx = NewCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BricksParserRULE_collection)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BricksParserT__8:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Load()
		}

	case BricksParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.Import_()
		}

	case BricksParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.AllParts()
		}

	case BricksParserT__11:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.Lost()
		}

	case BricksParserT__12:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(103)
			p.Set()
		}

	case BricksParserT__13:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(104)
			p.UserSet()
		}

	case BricksParserT__14:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(105)
			p.SetList()
		}

	case BricksParserT__15:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(106)
			p.PartList()
		}

	case BricksParserT__16:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(107)
			p.PartLists()
		}

	case BricksParserT__17:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(108)
			p.Sum()
		}

	case BricksParserT__18:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(109)
			p.Subtract()
		}

	case BricksParserT__19:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(110)
			p.Max()
		}

	case BricksParserT__20:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(111)
			p.Sort()
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

// ILoadContext is an interface to support dynamic dispatch.
type ILoadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsLoadContext differentiates from other interfaces.
	IsLoadContext()
}

type LoadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoadContext() *LoadContext {
	var p = new(LoadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_load
	return p
}

func InitEmptyLoadContext(p *LoadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_load
}

func (*LoadContext) IsLoadContext() {}

func NewLoadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoadContext {
	var p = new(LoadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_load

	return p
}

func (s *LoadContext) GetParser() antlr.Parser { return s.parser }

func (s *LoadContext) STRING() antlr.TerminalNode {
	return s.GetToken(BricksParserSTRING, 0)
}

func (s *LoadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterLoad(s)
	}
}

func (s *LoadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitLoad(s)
	}
}

func (p *BricksParser) Load() (localctx ILoadContext) {
	localctx = NewLoadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BricksParserRULE_load)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(BricksParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(BricksParserT__4)
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

// IImport_Context is an interface to support dynamic dispatch.
type IImport_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsImport_Context differentiates from other interfaces.
	IsImport_Context()
}

type Import_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_Context() *Import_Context {
	var p = new(Import_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_import_
	return p
}

func InitEmptyImport_Context(p *Import_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_import_
}

func (*Import_Context) IsImport_Context() {}

func NewImport_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_Context {
	var p = new(Import_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_import_

	return p
}

func (s *Import_Context) GetParser() antlr.Parser { return s.parser }

func (s *Import_Context) STRING() antlr.TerminalNode {
	return s.GetToken(BricksParserSTRING, 0)
}

func (s *Import_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterImport_(s)
	}
}

func (s *Import_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitImport_(s)
	}
}

func (p *BricksParser) Import_() (localctx IImport_Context) {
	localctx = NewImport_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BricksParserRULE_import_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(BricksParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(BricksParserT__4)
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

// IAllPartsContext is an interface to support dynamic dispatch.
type IAllPartsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAllPartsContext differentiates from other interfaces.
	IsAllPartsContext()
}

type AllPartsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllPartsContext() *AllPartsContext {
	var p = new(AllPartsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_allParts
	return p
}

func InitEmptyAllPartsContext(p *AllPartsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_allParts
}

func (*AllPartsContext) IsAllPartsContext() {}

func NewAllPartsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllPartsContext {
	var p = new(AllPartsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_allParts

	return p
}

func (s *AllPartsContext) GetParser() antlr.Parser { return s.parser }
func (s *AllPartsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllPartsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllPartsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterAllParts(s)
	}
}

func (s *AllPartsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitAllParts(s)
	}
}

func (p *BricksParser) AllParts() (localctx IAllPartsContext) {
	localctx = NewAllPartsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BricksParserRULE_allParts)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(BricksParserT__10)
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

// ILostContext is an interface to support dynamic dispatch.
type ILostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLostContext differentiates from other interfaces.
	IsLostContext()
}

type LostContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLostContext() *LostContext {
	var p = new(LostContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_lost
	return p
}

func InitEmptyLostContext(p *LostContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_lost
}

func (*LostContext) IsLostContext() {}

func NewLostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LostContext {
	var p = new(LostContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_lost

	return p
}

func (s *LostContext) GetParser() antlr.Parser { return s.parser }
func (s *LostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterLost(s)
	}
}

func (s *LostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitLost(s)
	}
}

func (p *BricksParser) Lost() (localctx ILostContext) {
	localctx = NewLostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BricksParserRULE_lost)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(BricksParserT__11)
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

// ISetContext is an interface to support dynamic dispatch.
type ISetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SET_NUM() antlr.TerminalNode
	BOOL() antlr.TerminalNode

	// IsSetContext differentiates from other interfaces.
	IsSetContext()
}

type SetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetContext() *SetContext {
	var p = new(SetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_set
	return p
}

func InitEmptySetContext(p *SetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_set
}

func (*SetContext) IsSetContext() {}

func NewSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetContext {
	var p = new(SetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_set

	return p
}

func (s *SetContext) GetParser() antlr.Parser { return s.parser }

func (s *SetContext) SET_NUM() antlr.TerminalNode {
	return s.GetToken(BricksParserSET_NUM, 0)
}

func (s *SetContext) BOOL() antlr.TerminalNode {
	return s.GetToken(BricksParserBOOL, 0)
}

func (s *SetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterSet(s)
	}
}

func (s *SetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitSet(s)
	}
}

func (p *BricksParser) Set() (localctx ISetContext) {
	localctx = NewSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BricksParserRULE_set)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(BricksParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(BricksParserSET_NUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
		{
			p.SetState(131)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(BricksParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(135)
		p.Match(BricksParserT__4)
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

// IUserSetContext is an interface to support dynamic dispatch.
type IUserSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SET_NUM() antlr.TerminalNode
	BOOL() antlr.TerminalNode

	// IsUserSetContext differentiates from other interfaces.
	IsUserSetContext()
}

type UserSetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserSetContext() *UserSetContext {
	var p = new(UserSetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_userSet
	return p
}

func InitEmptyUserSetContext(p *UserSetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_userSet
}

func (*UserSetContext) IsUserSetContext() {}

func NewUserSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserSetContext {
	var p = new(UserSetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_userSet

	return p
}

func (s *UserSetContext) GetParser() antlr.Parser { return s.parser }

func (s *UserSetContext) SET_NUM() antlr.TerminalNode {
	return s.GetToken(BricksParserSET_NUM, 0)
}

func (s *UserSetContext) BOOL() antlr.TerminalNode {
	return s.GetToken(BricksParserBOOL, 0)
}

func (s *UserSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterUserSet(s)
	}
}

func (s *UserSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitUserSet(s)
	}
}

func (p *BricksParser) UserSet() (localctx IUserSetContext) {
	localctx = NewUserSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BricksParserRULE_userSet)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(BricksParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(BricksParserSET_NUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
		{
			p.SetState(140)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(BricksParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(144)
		p.Match(BricksParserT__4)
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

// ISetListContext is an interface to support dynamic dispatch.
type ISetListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	BOOL() antlr.TerminalNode

	// IsSetListContext differentiates from other interfaces.
	IsSetListContext()
}

type SetListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetListContext() *SetListContext {
	var p = new(SetListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_setList
	return p
}

func InitEmptySetListContext(p *SetListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_setList
}

func (*SetListContext) IsSetListContext() {}

func NewSetListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetListContext {
	var p = new(SetListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_setList

	return p
}

func (s *SetListContext) GetParser() antlr.Parser { return s.parser }

func (s *SetListContext) INT() antlr.TerminalNode {
	return s.GetToken(BricksParserINT, 0)
}

func (s *SetListContext) BOOL() antlr.TerminalNode {
	return s.GetToken(BricksParserBOOL, 0)
}

func (s *SetListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterSetList(s)
	}
}

func (s *SetListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitSetList(s)
	}
}

func (p *BricksParser) SetList() (localctx ISetListContext) {
	localctx = NewSetListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BricksParserRULE_setList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(BricksParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(BricksParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
		{
			p.SetState(149)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(BricksParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(153)
		p.Match(BricksParserT__4)
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

// IPartListContext is an interface to support dynamic dispatch.
type IPartListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsPartListContext differentiates from other interfaces.
	IsPartListContext()
}

type PartListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartListContext() *PartListContext {
	var p = new(PartListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_partList
	return p
}

func InitEmptyPartListContext(p *PartListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_partList
}

func (*PartListContext) IsPartListContext() {}

func NewPartListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartListContext {
	var p = new(PartListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_partList

	return p
}

func (s *PartListContext) GetParser() antlr.Parser { return s.parser }

func (s *PartListContext) INT() antlr.TerminalNode {
	return s.GetToken(BricksParserINT, 0)
}

func (s *PartListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PartListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterPartList(s)
	}
}

func (s *PartListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitPartList(s)
	}
}

func (p *BricksParser) PartList() (localctx IPartListContext) {
	localctx = NewPartListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BricksParserRULE_partList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(BricksParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(BricksParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(BricksParserT__4)
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

// IPartListsContext is an interface to support dynamic dispatch.
type IPartListsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode

	// IsPartListsContext differentiates from other interfaces.
	IsPartListsContext()
}

type PartListsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartListsContext() *PartListsContext {
	var p = new(PartListsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_partLists
	return p
}

func InitEmptyPartListsContext(p *PartListsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_partLists
}

func (*PartListsContext) IsPartListsContext() {}

func NewPartListsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartListsContext {
	var p = new(PartListsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_partLists

	return p
}

func (s *PartListsContext) GetParser() antlr.Parser { return s.parser }

func (s *PartListsContext) STRING() antlr.TerminalNode {
	return s.GetToken(BricksParserSTRING, 0)
}

func (s *PartListsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(BricksParserBOOL, 0)
}

func (s *PartListsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartListsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PartListsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterPartLists(s)
	}
}

func (s *PartListsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitPartLists(s)
	}
}

func (p *BricksParser) PartLists() (localctx IPartListsContext) {
	localctx = NewPartListsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BricksParserRULE_partLists)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(BricksParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
		{
			p.SetState(163)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Match(BricksParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(167)
		p.Match(BricksParserT__4)
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

// ISumContext is an interface to support dynamic dispatch.
type ISumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCollectionOrId() []ICollectionOrIdContext
	CollectionOrId(i int) ICollectionOrIdContext

	// IsSumContext differentiates from other interfaces.
	IsSumContext()
}

type SumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySumContext() *SumContext {
	var p = new(SumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_sum
	return p
}

func InitEmptySumContext(p *SumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_sum
}

func (*SumContext) IsSumContext() {}

func NewSumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SumContext {
	var p = new(SumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_sum

	return p
}

func (s *SumContext) GetParser() antlr.Parser { return s.parser }

func (s *SumContext) AllCollectionOrId() []ICollectionOrIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICollectionOrIdContext); ok {
			len++
		}
	}

	tst := make([]ICollectionOrIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICollectionOrIdContext); ok {
			tst[i] = t.(ICollectionOrIdContext)
			i++
		}
	}

	return tst
}

func (s *SumContext) CollectionOrId(i int) ICollectionOrIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionOrIdContext); ok {
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

	return t.(ICollectionOrIdContext)
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterSum(s)
	}
}

func (s *SumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitSum(s)
	}
}

func (p *BricksParser) Sum() (localctx ISumContext) {
	localctx = NewSumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BricksParserRULE_sum)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(BricksParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.CollectionOrId()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BricksParserT__3 {
		{
			p.SetState(172)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.CollectionOrId()
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(178)
		p.Match(BricksParserT__4)
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

// ISubtractContext is an interface to support dynamic dispatch.
type ISubtractContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCollectionOrId() []ICollectionOrIdContext
	CollectionOrId(i int) ICollectionOrIdContext

	// IsSubtractContext differentiates from other interfaces.
	IsSubtractContext()
}

type SubtractContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubtractContext() *SubtractContext {
	var p = new(SubtractContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_subtract
	return p
}

func InitEmptySubtractContext(p *SubtractContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_subtract
}

func (*SubtractContext) IsSubtractContext() {}

func NewSubtractContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubtractContext {
	var p = new(SubtractContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_subtract

	return p
}

func (s *SubtractContext) GetParser() antlr.Parser { return s.parser }

func (s *SubtractContext) AllCollectionOrId() []ICollectionOrIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICollectionOrIdContext); ok {
			len++
		}
	}

	tst := make([]ICollectionOrIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICollectionOrIdContext); ok {
			tst[i] = t.(ICollectionOrIdContext)
			i++
		}
	}

	return tst
}

func (s *SubtractContext) CollectionOrId(i int) ICollectionOrIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionOrIdContext); ok {
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

	return t.(ICollectionOrIdContext)
}

func (s *SubtractContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubtractContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterSubtract(s)
	}
}

func (s *SubtractContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitSubtract(s)
	}
}

func (p *BricksParser) Subtract() (localctx ISubtractContext) {
	localctx = NewSubtractContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BricksParserRULE_subtract)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(BricksParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.CollectionOrId()
	}
	{
		p.SetState(183)
		p.Match(BricksParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.CollectionOrId()
	}
	{
		p.SetState(185)
		p.Match(BricksParserT__4)
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

// IMaxContext is an interface to support dynamic dispatch.
type IMaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCollectionOrId() []ICollectionOrIdContext
	CollectionOrId(i int) ICollectionOrIdContext

	// IsMaxContext differentiates from other interfaces.
	IsMaxContext()
}

type MaxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaxContext() *MaxContext {
	var p = new(MaxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_max
	return p
}

func InitEmptyMaxContext(p *MaxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_max
}

func (*MaxContext) IsMaxContext() {}

func NewMaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaxContext {
	var p = new(MaxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_max

	return p
}

func (s *MaxContext) GetParser() antlr.Parser { return s.parser }

func (s *MaxContext) AllCollectionOrId() []ICollectionOrIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICollectionOrIdContext); ok {
			len++
		}
	}

	tst := make([]ICollectionOrIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICollectionOrIdContext); ok {
			tst[i] = t.(ICollectionOrIdContext)
			i++
		}
	}

	return tst
}

func (s *MaxContext) CollectionOrId(i int) ICollectionOrIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionOrIdContext); ok {
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

	return t.(ICollectionOrIdContext)
}

func (s *MaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterMax(s)
	}
}

func (s *MaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitMax(s)
	}
}

func (p *BricksParser) Max() (localctx IMaxContext) {
	localctx = NewMaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BricksParserRULE_max)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(BricksParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.CollectionOrId()
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BricksParserT__3 {
		{
			p.SetState(190)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.CollectionOrId()
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(196)
		p.Match(BricksParserT__4)
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

// ISortContext is an interface to support dynamic dispatch.
type ISortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetQuantity returns the quantity token.
	GetQuantity() antlr.Token

	// GetDescending returns the descending token.
	GetDescending() antlr.Token

	// SetQuantity sets the quantity token.
	SetQuantity(antlr.Token)

	// SetDescending sets the descending token.
	SetDescending(antlr.Token)

	// Getter signatures
	CollectionOrId() ICollectionOrIdContext

	// IsSortContext differentiates from other interfaces.
	IsSortContext()
}

type SortContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	quantity   antlr.Token
	descending antlr.Token
}

func NewEmptySortContext() *SortContext {
	var p = new(SortContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_sort
	return p
}

func InitEmptySortContext(p *SortContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_sort
}

func (*SortContext) IsSortContext() {}

func NewSortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortContext {
	var p = new(SortContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_sort

	return p
}

func (s *SortContext) GetParser() antlr.Parser { return s.parser }

func (s *SortContext) GetQuantity() antlr.Token { return s.quantity }

func (s *SortContext) GetDescending() antlr.Token { return s.descending }

func (s *SortContext) SetQuantity(v antlr.Token) { s.quantity = v }

func (s *SortContext) SetDescending(v antlr.Token) { s.descending = v }

func (s *SortContext) CollectionOrId() ICollectionOrIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionOrIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionOrIdContext)
}

func (s *SortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterSort(s)
	}
}

func (s *SortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitSort(s)
	}
}

func (p *BricksParser) Sort() (localctx ISortContext) {
	localctx = NewSortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BricksParserRULE_sort)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(BricksParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.CollectionOrId()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(201)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)

			var _m = p.Match(BricksParserT__21)

			localctx.(*SortContext).quantity = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
		{
			p.SetState(205)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)

			var _m = p.Match(BricksParserT__22)

			localctx.(*SortContext).descending = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(209)
		p.Match(BricksParserT__4)
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

// IBuildContext is an interface to support dynamic dispatch.
type IBuildContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBuild_mode returns the build_mode token.
	GetBuild_mode() antlr.Token

	// SetBuild_mode sets the build_mode token.
	SetBuild_mode(antlr.Token)

	// Getter signatures
	AllCollectionOrId() []ICollectionOrIdContext
	CollectionOrId(i int) ICollectionOrIdContext
	ID() antlr.TerminalNode

	// IsBuildContext differentiates from other interfaces.
	IsBuildContext()
}

type BuildContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	build_mode antlr.Token
}

func NewEmptyBuildContext() *BuildContext {
	var p = new(BuildContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_build
	return p
}

func InitEmptyBuildContext(p *BuildContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_build
}

func (*BuildContext) IsBuildContext() {}

func NewBuildContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildContext {
	var p = new(BuildContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_build

	return p
}

func (s *BuildContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildContext) GetBuild_mode() antlr.Token { return s.build_mode }

func (s *BuildContext) SetBuild_mode(v antlr.Token) { s.build_mode = v }

func (s *BuildContext) AllCollectionOrId() []ICollectionOrIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICollectionOrIdContext); ok {
			len++
		}
	}

	tst := make([]ICollectionOrIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICollectionOrIdContext); ok {
			tst[i] = t.(ICollectionOrIdContext)
			i++
		}
	}

	return tst
}

func (s *BuildContext) CollectionOrId(i int) ICollectionOrIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionOrIdContext); ok {
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

	return t.(ICollectionOrIdContext)
}

func (s *BuildContext) ID() antlr.TerminalNode {
	return s.GetToken(BricksParserID, 0)
}

func (s *BuildContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterBuild(s)
	}
}

func (s *BuildContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitBuild(s)
	}
}

func (p *BricksParser) Build() (localctx IBuildContext) {
	localctx = NewBuildContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BricksParserRULE_build)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(BricksParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.CollectionOrId()
	}
	{
		p.SetState(214)
		p.Match(BricksParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.CollectionOrId()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
		{
			p.SetState(216)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)

			var _m = p.Match(BricksParserID)

			localctx.(*BuildContext).build_mode = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(220)
		p.Match(BricksParserT__4)
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(BricksParserID, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *BricksParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BricksParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(BricksParserID)
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
