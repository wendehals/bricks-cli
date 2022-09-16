// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Bricks

import (
	"fmt"
	"strconv"
  "sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type BricksParser struct {
	*antlr.BaseParser
}

var bricksParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func bricksParserInit() {
  staticData := &bricksParserStaticData
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
    "bricks", "command", "assignment", "save", "export", "exp", "identifier", 
    "load", "allParts", "lost", "set", "setList", "partList", "partLists", 
    "sum", "subtract", "max", "sort",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 23, 156, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 4, 0, 38, 8, 0, 11, 0, 12, 0, 39, 1, 
	1, 1, 1, 1, 1, 3, 1, 45, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 
	1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 
	1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 
	77, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 
	9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 95, 8, 10, 1, 10, 1, 10, 1, 
	11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 104, 8, 11, 1, 11, 1, 11, 1, 12, 
	1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 118, 
	8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 4, 14, 127, 8, 
	14, 11, 14, 12, 14, 128, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 
	1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 4, 16, 145, 8, 16, 11, 
	16, 12, 16, 146, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 
	0, 0, 18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 
	34, 0, 0, 156, 0, 37, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 
	6, 50, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 76, 1, 0, 0, 0, 12, 78, 1, 0, 
	0, 0, 14, 80, 1, 0, 0, 0, 16, 85, 1, 0, 0, 0, 18, 87, 1, 0, 0, 0, 20, 89, 
	1, 0, 0, 0, 22, 98, 1, 0, 0, 0, 24, 107, 1, 0, 0, 0, 26, 112, 1, 0, 0, 
	0, 28, 121, 1, 0, 0, 0, 30, 132, 1, 0, 0, 0, 32, 139, 1, 0, 0, 0, 34, 150, 
	1, 0, 0, 0, 36, 38, 3, 2, 1, 0, 37, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 
	39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 1, 1, 0, 0, 0, 41, 45, 3, 4, 
	2, 0, 42, 45, 3, 6, 3, 0, 43, 45, 3, 8, 4, 0, 44, 41, 1, 0, 0, 0, 44, 42, 
	1, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 3, 1, 0, 0, 0, 46, 47, 5, 21, 0, 0, 
	47, 48, 5, 1, 0, 0, 48, 49, 3, 10, 5, 0, 49, 5, 1, 0, 0, 0, 50, 51, 5, 
	2, 0, 0, 51, 52, 5, 3, 0, 0, 52, 53, 3, 10, 5, 0, 53, 54, 5, 4, 0, 0, 54, 
	55, 5, 20, 0, 0, 55, 56, 5, 5, 0, 0, 56, 7, 1, 0, 0, 0, 57, 58, 5, 6, 0, 
	0, 58, 59, 5, 3, 0, 0, 59, 60, 3, 10, 5, 0, 60, 61, 5, 4, 0, 0, 61, 62, 
	5, 20, 0, 0, 62, 63, 5, 5, 0, 0, 63, 9, 1, 0, 0, 0, 64, 77, 3, 12, 6, 0, 
	65, 77, 3, 14, 7, 0, 66, 77, 3, 16, 8, 0, 67, 77, 3, 18, 9, 0, 68, 77, 
	3, 20, 10, 0, 69, 77, 3, 22, 11, 0, 70, 77, 3, 24, 12, 0, 71, 77, 3, 26, 
	13, 0, 72, 77, 3, 28, 14, 0, 73, 77, 3, 30, 15, 0, 74, 77, 3, 32, 16, 0, 
	75, 77, 3, 34, 17, 0, 76, 64, 1, 0, 0, 0, 76, 65, 1, 0, 0, 0, 76, 66, 1, 
	0, 0, 0, 76, 67, 1, 0, 0, 0, 76, 68, 1, 0, 0, 0, 76, 69, 1, 0, 0, 0, 76, 
	70, 1, 0, 0, 0, 76, 71, 1, 0, 0, 0, 76, 72, 1, 0, 0, 0, 76, 73, 1, 0, 0, 
	0, 76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 11, 1, 0, 0, 0, 78, 79, 
	5, 21, 0, 0, 79, 13, 1, 0, 0, 0, 80, 81, 5, 7, 0, 0, 81, 82, 5, 3, 0, 0, 
	82, 83, 5, 20, 0, 0, 83, 84, 5, 5, 0, 0, 84, 15, 1, 0, 0, 0, 85, 86, 5, 
	8, 0, 0, 86, 17, 1, 0, 0, 0, 87, 88, 5, 9, 0, 0, 88, 19, 1, 0, 0, 0, 89, 
	90, 5, 10, 0, 0, 90, 91, 5, 3, 0, 0, 91, 94, 5, 22, 0, 0, 92, 93, 5, 4, 
	0, 0, 93, 95, 5, 19, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 
	96, 1, 0, 0, 0, 96, 97, 5, 5, 0, 0, 97, 21, 1, 0, 0, 0, 98, 99, 5, 11, 
	0, 0, 99, 100, 5, 3, 0, 0, 100, 103, 5, 18, 0, 0, 101, 102, 5, 4, 0, 0, 
	102, 104, 5, 19, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 
	105, 1, 0, 0, 0, 105, 106, 5, 5, 0, 0, 106, 23, 1, 0, 0, 0, 107, 108, 5, 
	12, 0, 0, 108, 109, 5, 3, 0, 0, 109, 110, 5, 18, 0, 0, 110, 111, 5, 5, 
	0, 0, 111, 25, 1, 0, 0, 0, 112, 113, 5, 13, 0, 0, 113, 114, 5, 3, 0, 0, 
	114, 117, 5, 20, 0, 0, 115, 116, 5, 4, 0, 0, 116, 118, 5, 19, 0, 0, 117, 
	115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 
	5, 5, 0, 0, 120, 27, 1, 0, 0, 0, 121, 122, 5, 14, 0, 0, 122, 123, 5, 3, 
	0, 0, 123, 126, 3, 10, 5, 0, 124, 125, 5, 4, 0, 0, 125, 127, 3, 10, 5, 
	0, 126, 124, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 
	129, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 5, 5, 0, 0, 131, 29, 1, 
	0, 0, 0, 132, 133, 5, 15, 0, 0, 133, 134, 5, 3, 0, 0, 134, 135, 3, 10, 
	5, 0, 135, 136, 5, 4, 0, 0, 136, 137, 3, 10, 5, 0, 137, 138, 5, 5, 0, 0, 
	138, 31, 1, 0, 0, 0, 139, 140, 5, 16, 0, 0, 140, 141, 5, 3, 0, 0, 141, 
	144, 3, 10, 5, 0, 142, 143, 5, 4, 0, 0, 143, 145, 3, 10, 5, 0, 144, 142, 
	1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 
	0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 5, 5, 0, 0, 149, 33, 1, 0, 0, 0, 
	150, 151, 5, 17, 0, 0, 151, 152, 5, 3, 0, 0, 152, 153, 3, 10, 5, 0, 153, 
	154, 5, 5, 0, 0, 154, 35, 1, 0, 0, 0, 8, 39, 44, 76, 94, 103, 117, 128, 
	146,
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
  staticData := &bricksParserStaticData
  staticData.once.Do(bricksParserInit)
}

// NewBricksParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBricksParser(input antlr.TokenStream) *BricksParser {
	BricksParserInit()
	this := new(BricksParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &bricksParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}


// BricksParser tokens.
const (
	BricksParserEOF = antlr.TokenEOF
	BricksParserT__0 = 1
	BricksParserT__1 = 2
	BricksParserT__2 = 3
	BricksParserT__3 = 4
	BricksParserT__4 = 5
	BricksParserT__5 = 6
	BricksParserT__6 = 7
	BricksParserT__7 = 8
	BricksParserT__8 = 9
	BricksParserT__9 = 10
	BricksParserT__10 = 11
	BricksParserT__11 = 12
	BricksParserT__12 = 13
	BricksParserT__13 = 14
	BricksParserT__14 = 15
	BricksParserT__15 = 16
	BricksParserT__16 = 17
	BricksParserINT = 18
	BricksParserBOOL = 19
	BricksParserSTRING = 20
	BricksParserID = 21
	BricksParserSET_NUM = 22
	BricksParserWS = 23
)

// BricksParser rules.
const (
	BricksParserRULE_bricks = 0
	BricksParserRULE_command = 1
	BricksParserRULE_assignment = 2
	BricksParserRULE_save = 3
	BricksParserRULE_export = 4
	BricksParserRULE_exp = 5
	BricksParserRULE_identifier = 6
	BricksParserRULE_load = 7
	BricksParserRULE_allParts = 8
	BricksParserRULE_lost = 9
	BricksParserRULE_set = 10
	BricksParserRULE_setList = 11
	BricksParserRULE_partList = 12
	BricksParserRULE_partLists = 13
	BricksParserRULE_sum = 14
	BricksParserRULE_subtract = 15
	BricksParserRULE_max = 16
	BricksParserRULE_sort = 17
)

// IBricksContext is an interface to support dynamic dispatch.
type IBricksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBricksContext differentiates from other interfaces.
	IsBricksContext()
}

type BricksContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBricksContext() *BricksContext {
	var p = new(BricksContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_bricks
	return p
}

func (*BricksContext) IsBricksContext() {}

func NewBricksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BricksContext {
	var p = new(BricksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewBricksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BricksParserRULE_bricks)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2097220) != 0 {
		{
			p.SetState(36)
			p.Command()
		}


		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *CommandContext) Save() ISaveContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISaveContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISaveContext)
}

func (s *CommandContext) Export() IExportContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExportContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExportContext)
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
	this := p
	_ = this

	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BricksParserRULE_command)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BricksParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Assignment()
		}


	case BricksParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Save()
		}


	case BricksParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.Export()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(BricksParserID, 0)
}

func (s *AssignmentContext) Exp() IExpContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BricksParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(BricksParserID)
	}
	{
		p.SetState(47)
		p.Match(BricksParserT__0)
	}
	{
		p.SetState(48)
		p.Exp()
	}



	return localctx
}


// ISaveContext is an interface to support dynamic dispatch.
type ISaveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSaveContext differentiates from other interfaces.
	IsSaveContext()
}

type SaveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySaveContext() *SaveContext {
	var p = new(SaveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_save
	return p
}

func (*SaveContext) IsSaveContext() {}

func NewSaveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SaveContext {
	var p = new(SaveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_save

	return p
}

func (s *SaveContext) GetParser() antlr.Parser { return s.parser }

func (s *SaveContext) Exp() IExpContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewSaveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BricksParserRULE_save)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(BricksParserT__1)
	}
	{
		p.SetState(51)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(52)
		p.Exp()
	}
	{
		p.SetState(53)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(54)
		p.Match(BricksParserSTRING)
	}
	{
		p.SetState(55)
		p.Match(BricksParserT__4)
	}



	return localctx
}


// IExportContext is an interface to support dynamic dispatch.
type IExportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportContext differentiates from other interfaces.
	IsExportContext()
}

type ExportContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportContext() *ExportContext {
	var p = new(ExportContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_export
	return p
}

func (*ExportContext) IsExportContext() {}

func NewExportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportContext {
	var p = new(ExportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_export

	return p
}

func (s *ExportContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportContext) Exp() IExpContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewExportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BricksParserRULE_export)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(BricksParserT__5)
	}
	{
		p.SetState(58)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(59)
		p.Exp()
	}
	{
		p.SetState(60)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(61)
		p.Match(BricksParserSTRING)
	}
	{
		p.SetState(62)
		p.Match(BricksParserT__4)
	}



	return localctx
}


// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExpContext) Load() ILoadContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoadContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoadContext)
}

func (s *ExpContext) AllParts() IAllPartsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllPartsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllPartsContext)
}

func (s *ExpContext) Lost() ILostContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILostContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILostContext)
}

func (s *ExpContext) Set() ISetContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetContext)
}

func (s *ExpContext) SetList() ISetListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetListContext)
}

func (s *ExpContext) PartList() IPartListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPartListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPartListContext)
}

func (s *ExpContext) PartLists() IPartListsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPartListsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPartListsContext)
}

func (s *ExpContext) Sum() ISumContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISumContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISumContext)
}

func (s *ExpContext) Subtract() ISubtractContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubtractContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubtractContext)
}

func (s *ExpContext) Max() IMaxContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMaxContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMaxContext)
}

func (s *ExpContext) Sort() ISortContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortContext); ok {
			t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BricksParserRULE_exp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BricksParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Identifier()
		}


	case BricksParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Load()
		}


	case BricksParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.AllParts()
		}


	case BricksParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.Lost()
		}


	case BricksParserT__9:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.Set()
		}


	case BricksParserT__10:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(69)
			p.SetList()
		}


	case BricksParserT__11:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(70)
			p.PartList()
		}


	case BricksParserT__12:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(71)
			p.PartLists()
		}


	case BricksParserT__13:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(72)
			p.Sum()
		}


	case BricksParserT__14:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(73)
			p.Subtract()
		}


	case BricksParserT__15:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(74)
			p.Max()
		}


	case BricksParserT__16:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(75)
			p.Sort()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BricksParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(BricksParserID)
	}



	return localctx
}


// ILoadContext is an interface to support dynamic dispatch.
type ILoadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoadContext differentiates from other interfaces.
	IsLoadContext()
}

type LoadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoadContext() *LoadContext {
	var p = new(LoadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_load
	return p
}

func (*LoadContext) IsLoadContext() {}

func NewLoadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoadContext {
	var p = new(LoadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewLoadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BricksParserRULE_load)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(BricksParserT__6)
	}
	{
		p.SetState(81)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(82)
		p.Match(BricksParserSTRING)
	}
	{
		p.SetState(83)
		p.Match(BricksParserT__4)
	}



	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllPartsContext() *AllPartsContext {
	var p = new(AllPartsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_allParts
	return p
}

func (*AllPartsContext) IsAllPartsContext() {}

func NewAllPartsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllPartsContext {
	var p = new(AllPartsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewAllPartsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BricksParserRULE_allParts)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(BricksParserT__7)
	}



	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLostContext() *LostContext {
	var p = new(LostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_lost
	return p
}

func (*LostContext) IsLostContext() {}

func NewLostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LostContext {
	var p = new(LostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewLostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BricksParserRULE_lost)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(BricksParserT__8)
	}



	return localctx
}


// ISetContext is an interface to support dynamic dispatch.
type ISetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetContext differentiates from other interfaces.
	IsSetContext()
}

type SetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetContext() *SetContext {
	var p = new(SetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_set
	return p
}

func (*SetContext) IsSetContext() {}

func NewSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetContext {
	var p = new(SetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BricksParserRULE_set)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(BricksParserT__9)
	}
	{
		p.SetState(90)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(91)
		p.Match(BricksParserSET_NUM)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == BricksParserT__3 {
		{
			p.SetState(92)
			p.Match(BricksParserT__3)
		}
		{
			p.SetState(93)
			p.Match(BricksParserBOOL)
		}

	}
	{
		p.SetState(96)
		p.Match(BricksParserT__4)
	}



	return localctx
}


// ISetListContext is an interface to support dynamic dispatch.
type ISetListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetListContext differentiates from other interfaces.
	IsSetListContext()
}

type SetListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetListContext() *SetListContext {
	var p = new(SetListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_setList
	return p
}

func (*SetListContext) IsSetListContext() {}

func NewSetListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetListContext {
	var p = new(SetListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewSetListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BricksParserRULE_setList)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(BricksParserT__10)
	}
	{
		p.SetState(99)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(100)
		p.Match(BricksParserINT)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == BricksParserT__3 {
		{
			p.SetState(101)
			p.Match(BricksParserT__3)
		}
		{
			p.SetState(102)
			p.Match(BricksParserBOOL)
		}

	}
	{
		p.SetState(105)
		p.Match(BricksParserT__4)
	}



	return localctx
}


// IPartListContext is an interface to support dynamic dispatch.
type IPartListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPartListContext differentiates from other interfaces.
	IsPartListContext()
}

type PartListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartListContext() *PartListContext {
	var p = new(PartListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_partList
	return p
}

func (*PartListContext) IsPartListContext() {}

func NewPartListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartListContext {
	var p = new(PartListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewPartListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BricksParserRULE_partList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(BricksParserT__11)
	}
	{
		p.SetState(108)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(109)
		p.Match(BricksParserINT)
	}
	{
		p.SetState(110)
		p.Match(BricksParserT__4)
	}



	return localctx
}


// IPartListsContext is an interface to support dynamic dispatch.
type IPartListsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPartListsContext differentiates from other interfaces.
	IsPartListsContext()
}

type PartListsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartListsContext() *PartListsContext {
	var p = new(PartListsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_partLists
	return p
}

func (*PartListsContext) IsPartListsContext() {}

func NewPartListsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartListsContext {
	var p = new(PartListsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewPartListsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BricksParserRULE_partLists)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(BricksParserT__12)
	}
	{
		p.SetState(113)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(114)
		p.Match(BricksParserSTRING)
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == BricksParserT__3 {
		{
			p.SetState(115)
			p.Match(BricksParserT__3)
		}
		{
			p.SetState(116)
			p.Match(BricksParserBOOL)
		}

	}
	{
		p.SetState(119)
		p.Match(BricksParserT__4)
	}



	return localctx
}


// ISumContext is an interface to support dynamic dispatch.
type ISumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSumContext differentiates from other interfaces.
	IsSumContext()
}

type SumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySumContext() *SumContext {
	var p = new(SumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_sum
	return p
}

func (*SumContext) IsSumContext() {}

func NewSumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SumContext {
	var p = new(SumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewSumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BricksParserRULE_sum)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(BricksParserT__13)
	}
	{
		p.SetState(122)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(123)
		p.Exp()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == BricksParserT__3 {
		{
			p.SetState(124)
			p.Match(BricksParserT__3)
		}
		{
			p.SetState(125)
			p.Exp()
		}


		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(130)
		p.Match(BricksParserT__4)
	}



	return localctx
}


// ISubtractContext is an interface to support dynamic dispatch.
type ISubtractContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubtractContext differentiates from other interfaces.
	IsSubtractContext()
}

type SubtractContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubtractContext() *SubtractContext {
	var p = new(SubtractContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_subtract
	return p
}

func (*SubtractContext) IsSubtractContext() {}

func NewSubtractContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubtractContext {
	var p = new(SubtractContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewSubtractContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BricksParserRULE_subtract)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(BricksParserT__14)
	}
	{
		p.SetState(133)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(134)
		p.Exp()
	}
	{
		p.SetState(135)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(136)
		p.Exp()
	}
	{
		p.SetState(137)
		p.Match(BricksParserT__4)
	}



	return localctx
}


// IMaxContext is an interface to support dynamic dispatch.
type IMaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaxContext differentiates from other interfaces.
	IsMaxContext()
}

type MaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaxContext() *MaxContext {
	var p = new(MaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_max
	return p
}

func (*MaxContext) IsMaxContext() {}

func NewMaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaxContext {
	var p = new(MaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewMaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BricksParserRULE_max)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(BricksParserT__15)
	}
	{
		p.SetState(140)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(141)
		p.Exp()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == BricksParserT__3 {
		{
			p.SetState(142)
			p.Match(BricksParserT__3)
		}
		{
			p.SetState(143)
			p.Exp()
		}


		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(148)
		p.Match(BricksParserT__4)
	}



	return localctx
}


// ISortContext is an interface to support dynamic dispatch.
type ISortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSortContext differentiates from other interfaces.
	IsSortContext()
}

type SortContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortContext() *SortContext {
	var p = new(SortContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_sort
	return p
}

func (*SortContext) IsSortContext() {}

func NewSortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortContext {
	var p = new(SortContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_sort

	return p
}

func (s *SortContext) GetParser() antlr.Parser { return s.parser }

func (s *SortContext) Exp() IExpContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewSortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BricksParserRULE_sort)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(BricksParserT__16)
	}
	{
		p.SetState(151)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(152)
		p.Exp()
	}
	{
		p.SetState(153)
		p.Match(BricksParserT__4)
	}



	return localctx
}


