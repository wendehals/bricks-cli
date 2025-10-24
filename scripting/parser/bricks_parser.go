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
		"", "':='", "'save'", "'('", "','", "')'", "'export'", "'build'", "'pause'",
		"'load'", "'import'", "'allParts'", "'lost'", "'set'", "'userSet'",
		"'setList'", "'partList'", "'partLists'", "'sum'", "'subtract'", "'max'",
		"'sort'", "'quantity'", "'descending'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "INT", "BOOL", "STRING", "ID", "SET_NUM",
		"WS",
	}
	staticData.RuleNames = []string{
		"bricks", "command", "assignment", "save", "export", "build", "pause",
		"exp", "identifier", "load", "import_", "allParts", "lost", "set", "userSet",
		"setList", "partList", "partLists", "sum", "subtract", "max", "sort",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 208, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 4, 0, 46, 8, 0, 11, 0, 12, 0, 47, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 55, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 84, 8, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 107, 8, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 130, 8, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 139, 8, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 148, 8, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 3, 17, 162, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		4, 18, 171, 8, 18, 11, 18, 12, 18, 172, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 4, 20,
		189, 8, 20, 11, 20, 12, 20, 190, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 3, 21, 200, 8, 21, 1, 21, 1, 21, 3, 21, 204, 8, 21, 1, 21, 1,
		21, 1, 21, 0, 0, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 0, 0, 212, 0, 45, 1, 0, 0, 0, 2, 54, 1,
		0, 0, 0, 4, 56, 1, 0, 0, 0, 6, 60, 1, 0, 0, 0, 8, 67, 1, 0, 0, 0, 10, 74,
		1, 0, 0, 0, 12, 87, 1, 0, 0, 0, 14, 106, 1, 0, 0, 0, 16, 108, 1, 0, 0,
		0, 18, 110, 1, 0, 0, 0, 20, 115, 1, 0, 0, 0, 22, 120, 1, 0, 0, 0, 24, 122,
		1, 0, 0, 0, 26, 124, 1, 0, 0, 0, 28, 133, 1, 0, 0, 0, 30, 142, 1, 0, 0,
		0, 32, 151, 1, 0, 0, 0, 34, 156, 1, 0, 0, 0, 36, 165, 1, 0, 0, 0, 38, 176,
		1, 0, 0, 0, 40, 183, 1, 0, 0, 0, 42, 194, 1, 0, 0, 0, 44, 46, 3, 2, 1,
		0, 45, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48,
		1, 0, 0, 0, 48, 1, 1, 0, 0, 0, 49, 55, 3, 4, 2, 0, 50, 55, 3, 6, 3, 0,
		51, 55, 3, 8, 4, 0, 52, 55, 3, 10, 5, 0, 53, 55, 3, 12, 6, 0, 54, 49, 1,
		0, 0, 0, 54, 50, 1, 0, 0, 0, 54, 51, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54,
		53, 1, 0, 0, 0, 55, 3, 1, 0, 0, 0, 56, 57, 5, 27, 0, 0, 57, 58, 5, 1, 0,
		0, 58, 59, 3, 14, 7, 0, 59, 5, 1, 0, 0, 0, 60, 61, 5, 2, 0, 0, 61, 62,
		5, 3, 0, 0, 62, 63, 3, 14, 7, 0, 63, 64, 5, 4, 0, 0, 64, 65, 5, 26, 0,
		0, 65, 66, 5, 5, 0, 0, 66, 7, 1, 0, 0, 0, 67, 68, 5, 6, 0, 0, 68, 69, 5,
		3, 0, 0, 69, 70, 3, 14, 7, 0, 70, 71, 5, 4, 0, 0, 71, 72, 5, 26, 0, 0,
		72, 73, 5, 5, 0, 0, 73, 9, 1, 0, 0, 0, 74, 75, 5, 7, 0, 0, 75, 76, 5, 3,
		0, 0, 76, 77, 3, 14, 7, 0, 77, 78, 5, 4, 0, 0, 78, 79, 3, 14, 7, 0, 79,
		80, 5, 4, 0, 0, 80, 83, 5, 26, 0, 0, 81, 82, 5, 4, 0, 0, 82, 84, 5, 27,
		0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86,
		5, 5, 0, 0, 86, 11, 1, 0, 0, 0, 87, 88, 5, 8, 0, 0, 88, 89, 5, 3, 0, 0,
		89, 90, 5, 24, 0, 0, 90, 91, 5, 5, 0, 0, 91, 13, 1, 0, 0, 0, 92, 107, 3,
		16, 8, 0, 93, 107, 3, 18, 9, 0, 94, 107, 3, 20, 10, 0, 95, 107, 3, 22,
		11, 0, 96, 107, 3, 24, 12, 0, 97, 107, 3, 26, 13, 0, 98, 107, 3, 28, 14,
		0, 99, 107, 3, 30, 15, 0, 100, 107, 3, 32, 16, 0, 101, 107, 3, 34, 17,
		0, 102, 107, 3, 36, 18, 0, 103, 107, 3, 38, 19, 0, 104, 107, 3, 40, 20,
		0, 105, 107, 3, 42, 21, 0, 106, 92, 1, 0, 0, 0, 106, 93, 1, 0, 0, 0, 106,
		94, 1, 0, 0, 0, 106, 95, 1, 0, 0, 0, 106, 96, 1, 0, 0, 0, 106, 97, 1, 0,
		0, 0, 106, 98, 1, 0, 0, 0, 106, 99, 1, 0, 0, 0, 106, 100, 1, 0, 0, 0, 106,
		101, 1, 0, 0, 0, 106, 102, 1, 0, 0, 0, 106, 103, 1, 0, 0, 0, 106, 104,
		1, 0, 0, 0, 106, 105, 1, 0, 0, 0, 107, 15, 1, 0, 0, 0, 108, 109, 5, 27,
		0, 0, 109, 17, 1, 0, 0, 0, 110, 111, 5, 9, 0, 0, 111, 112, 5, 3, 0, 0,
		112, 113, 5, 26, 0, 0, 113, 114, 5, 5, 0, 0, 114, 19, 1, 0, 0, 0, 115,
		116, 5, 10, 0, 0, 116, 117, 5, 3, 0, 0, 117, 118, 5, 26, 0, 0, 118, 119,
		5, 5, 0, 0, 119, 21, 1, 0, 0, 0, 120, 121, 5, 11, 0, 0, 121, 23, 1, 0,
		0, 0, 122, 123, 5, 12, 0, 0, 123, 25, 1, 0, 0, 0, 124, 125, 5, 13, 0, 0,
		125, 126, 5, 3, 0, 0, 126, 129, 5, 28, 0, 0, 127, 128, 5, 4, 0, 0, 128,
		130, 5, 25, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131,
		1, 0, 0, 0, 131, 132, 5, 5, 0, 0, 132, 27, 1, 0, 0, 0, 133, 134, 5, 14,
		0, 0, 134, 135, 5, 3, 0, 0, 135, 138, 5, 28, 0, 0, 136, 137, 5, 4, 0, 0,
		137, 139, 5, 25, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		140, 1, 0, 0, 0, 140, 141, 5, 5, 0, 0, 141, 29, 1, 0, 0, 0, 142, 143, 5,
		15, 0, 0, 143, 144, 5, 3, 0, 0, 144, 147, 5, 24, 0, 0, 145, 146, 5, 4,
		0, 0, 146, 148, 5, 25, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0,
		148, 149, 1, 0, 0, 0, 149, 150, 5, 5, 0, 0, 150, 31, 1, 0, 0, 0, 151, 152,
		5, 16, 0, 0, 152, 153, 5, 3, 0, 0, 153, 154, 5, 24, 0, 0, 154, 155, 5,
		5, 0, 0, 155, 33, 1, 0, 0, 0, 156, 157, 5, 17, 0, 0, 157, 158, 5, 3, 0,
		0, 158, 161, 5, 26, 0, 0, 159, 160, 5, 4, 0, 0, 160, 162, 5, 25, 0, 0,
		161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163,
		164, 5, 5, 0, 0, 164, 35, 1, 0, 0, 0, 165, 166, 5, 18, 0, 0, 166, 167,
		5, 3, 0, 0, 167, 170, 3, 14, 7, 0, 168, 169, 5, 4, 0, 0, 169, 171, 3, 14,
		7, 0, 170, 168, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0,
		172, 173, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 5, 5, 0, 0, 175,
		37, 1, 0, 0, 0, 176, 177, 5, 19, 0, 0, 177, 178, 5, 3, 0, 0, 178, 179,
		3, 14, 7, 0, 179, 180, 5, 4, 0, 0, 180, 181, 3, 14, 7, 0, 181, 182, 5,
		5, 0, 0, 182, 39, 1, 0, 0, 0, 183, 184, 5, 20, 0, 0, 184, 185, 5, 3, 0,
		0, 185, 188, 3, 14, 7, 0, 186, 187, 5, 4, 0, 0, 187, 189, 3, 14, 7, 0,
		188, 186, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 5, 5, 0, 0, 193, 41, 1,
		0, 0, 0, 194, 195, 5, 21, 0, 0, 195, 196, 5, 3, 0, 0, 196, 199, 3, 14,
		7, 0, 197, 198, 5, 4, 0, 0, 198, 200, 5, 22, 0, 0, 199, 197, 1, 0, 0, 0,
		199, 200, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 202, 5, 4, 0, 0, 202,
		204, 5, 23, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 205,
		1, 0, 0, 0, 205, 206, 5, 5, 0, 0, 206, 43, 1, 0, 0, 0, 12, 47, 54, 83,
		106, 129, 138, 147, 161, 172, 190, 199, 203,
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
	BricksParserEOF     = antlr.TokenEOF
	BricksParserT__0    = 1
	BricksParserT__1    = 2
	BricksParserT__2    = 3
	BricksParserT__3    = 4
	BricksParserT__4    = 5
	BricksParserT__5    = 6
	BricksParserT__6    = 7
	BricksParserT__7    = 8
	BricksParserT__8    = 9
	BricksParserT__9    = 10
	BricksParserT__10   = 11
	BricksParserT__11   = 12
	BricksParserT__12   = 13
	BricksParserT__13   = 14
	BricksParserT__14   = 15
	BricksParserT__15   = 16
	BricksParserT__16   = 17
	BricksParserT__17   = 18
	BricksParserT__18   = 19
	BricksParserT__19   = 20
	BricksParserT__20   = 21
	BricksParserT__21   = 22
	BricksParserT__22   = 23
	BricksParserINT     = 24
	BricksParserBOOL    = 25
	BricksParserSTRING  = 26
	BricksParserID      = 27
	BricksParserSET_NUM = 28
	BricksParserWS      = 29
)

// BricksParser rules.
const (
	BricksParserRULE_bricks     = 0
	BricksParserRULE_command    = 1
	BricksParserRULE_assignment = 2
	BricksParserRULE_save       = 3
	BricksParserRULE_export     = 4
	BricksParserRULE_build      = 5
	BricksParserRULE_pause      = 6
	BricksParserRULE_exp        = 7
	BricksParserRULE_identifier = 8
	BricksParserRULE_load       = 9
	BricksParserRULE_import_    = 10
	BricksParserRULE_allParts   = 11
	BricksParserRULE_lost       = 12
	BricksParserRULE_set        = 13
	BricksParserRULE_userSet    = 14
	BricksParserRULE_setList    = 15
	BricksParserRULE_partList   = 16
	BricksParserRULE_partLists  = 17
	BricksParserRULE_sum        = 18
	BricksParserRULE_subtract   = 19
	BricksParserRULE_max        = 20
	BricksParserRULE_sort       = 21
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&134218180) != 0) {
		{
			p.SetState(44)
			p.Command()
		}

		p.SetState(47)
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
	Build() IBuildContext
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

func (s *CommandContext) Build() IBuildContext {
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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BricksParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Assignment()
		}

	case BricksParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Save()
		}

	case BricksParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)
			p.Export()
		}

	case BricksParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(52)
			p.Build()
		}

	case BricksParserT__7:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(53)
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Exp() IExpContext

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

func (s *AssignmentContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
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
	p.EnterRule(localctx, 4, BricksParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(BricksParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(BricksParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Exp()
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
	Exp() IExpContext
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

func (s *SaveContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
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
	p.EnterRule(localctx, 6, BricksParserRULE_save)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(BricksParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Exp()
	}
	{
		p.SetState(63)
		p.Match(BricksParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
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
	Exp() IExpContext
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

func (s *ExportContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
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
	p.EnterRule(localctx, 8, BricksParserRULE_export)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(BricksParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Exp()
	}
	{
		p.SetState(70)
		p.Match(BricksParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
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
	AllExp() []IExpContext
	Exp(i int) IExpContext
	STRING() antlr.TerminalNode
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

func (s *BuildContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *BuildContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
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

	return t.(IExpContext)
}

func (s *BuildContext) STRING() antlr.TerminalNode {
	return s.GetToken(BricksParserSTRING, 0)
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
	p.EnterRule(localctx, 10, BricksParserRULE_build)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(BricksParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Exp()
	}
	{
		p.SetState(77)
		p.Match(BricksParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Exp()
	}
	{
		p.SetState(79)
		p.Match(BricksParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
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

			var _m = p.Match(BricksParserID)

			localctx.(*BuildContext).build_mode = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(85)
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
	p.EnterRule(localctx, 12, BricksParserRULE_pause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(BricksParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)

		var _m = p.Match(BricksParserINT)

		localctx.(*PauseContext).seconds = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
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

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
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

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BricksParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Identifier() IIdentifierContext {
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

func (s *ExpContext) Load() ILoadContext {
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

func (s *ExpContext) Import_() IImport_Context {
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

func (s *ExpContext) AllParts() IAllPartsContext {
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

func (s *ExpContext) Lost() ILostContext {
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

func (s *ExpContext) Set() ISetContext {
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

func (s *ExpContext) UserSet() IUserSetContext {
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

func (s *ExpContext) SetList() ISetListContext {
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

func (s *ExpContext) PartList() IPartListContext {
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

func (s *ExpContext) PartLists() IPartListsContext {
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

func (s *ExpContext) Sum() ISumContext {
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

func (s *ExpContext) Subtract() ISubtractContext {
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

func (s *ExpContext) Max() IMaxContext {
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

func (s *ExpContext) Sort() ISortContext {
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

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *BricksParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BricksParserRULE_exp)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BricksParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Identifier()
		}

	case BricksParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Load()
		}

	case BricksParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Import_()
		}

	case BricksParserT__10:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(95)
			p.AllParts()
		}

	case BricksParserT__11:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(96)
			p.Lost()
		}

	case BricksParserT__12:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(97)
			p.Set()
		}

	case BricksParserT__13:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(98)
			p.UserSet()
		}

	case BricksParserT__14:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(99)
			p.SetList()
		}

	case BricksParserT__15:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(100)
			p.PartList()
		}

	case BricksParserT__16:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(101)
			p.PartLists()
		}

	case BricksParserT__17:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(102)
			p.Sum()
		}

	case BricksParserT__18:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(103)
			p.Subtract()
		}

	case BricksParserT__19:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(104)
			p.Max()
		}

	case BricksParserT__20:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(105)
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
	p.EnterRule(localctx, 16, BricksParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
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
	p.EnterRule(localctx, 18, BricksParserRULE_load)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(BricksParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
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
	p.EnterRule(localctx, 20, BricksParserRULE_import_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(BricksParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
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
	p.EnterRule(localctx, 22, BricksParserRULE_allParts)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
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
	p.EnterRule(localctx, 24, BricksParserRULE_lost)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
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
	p.EnterRule(localctx, 26, BricksParserRULE_set)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(BricksParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Match(BricksParserSET_NUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
		{
			p.SetState(127)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(BricksParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(131)
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
	p.EnterRule(localctx, 28, BricksParserRULE_userSet)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(BricksParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Match(BricksParserSET_NUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
		{
			p.SetState(136)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(BricksParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(140)
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
	p.EnterRule(localctx, 30, BricksParserRULE_setList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(BricksParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(BricksParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
		{
			p.SetState(145)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(BricksParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(149)
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
	p.EnterRule(localctx, 32, BricksParserRULE_partList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(BricksParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(BricksParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
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
	p.EnterRule(localctx, 34, BricksParserRULE_partLists)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(BricksParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(BricksParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
		{
			p.SetState(159)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(BricksParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(163)
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
	AllExp() []IExpContext
	Exp(i int) IExpContext

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

func (s *SumContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *SumContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
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

	return t.(IExpContext)
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
	p.EnterRule(localctx, 36, BricksParserRULE_sum)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(BricksParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Exp()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BricksParserT__3 {
		{
			p.SetState(168)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Exp()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(174)
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
	AllExp() []IExpContext
	Exp(i int) IExpContext

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

func (s *SubtractContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *SubtractContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
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

	return t.(IExpContext)
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
	p.EnterRule(localctx, 38, BricksParserRULE_subtract)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(BricksParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Exp()
	}
	{
		p.SetState(179)
		p.Match(BricksParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Exp()
	}
	{
		p.SetState(181)
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
	AllExp() []IExpContext
	Exp(i int) IExpContext

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

func (s *MaxContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *MaxContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
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

	return t.(IExpContext)
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
	p.EnterRule(localctx, 40, BricksParserRULE_max)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(BricksParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Exp()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BricksParserT__3 {
		{
			p.SetState(186)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Exp()
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(192)
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
	Exp() IExpContext

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

func (s *SortContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
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
	p.EnterRule(localctx, 42, BricksParserRULE_sort)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(BricksParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Match(BricksParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Exp()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(197)
			p.Match(BricksParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)

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
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BricksParserT__3 {
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

			var _m = p.Match(BricksParserT__22)

			localctx.(*SortContext).descending = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(205)
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
