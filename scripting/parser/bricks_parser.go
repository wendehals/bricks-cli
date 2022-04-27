// Code generated from .\scripting\parser\Bricks.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bricks

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
		"", "'allParts'", "'lost'", "'set'", "'('", "')'", "'setList'", "'partList'",
		"'sum'", "','", "'subtract'", "'max'", "'mergeByColor'", "'mergeByVariant'",
		"'sort'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "INT", "SET_NUM",
		"WS",
	}
	staticData.ruleNames = []string{
		"bricks", "exp", "allParts", "lost", "set", "setList", "partList", "sum",
		"subtract", "max", "mergeByColor", "mergeByVariant", "sort",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 105, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 40, 8, 1, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 66, 8, 7, 11,
		7, 12, 7, 67, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 84, 8, 9, 11, 9, 12, 9, 85, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 0, 0, 13, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 0, 0, 103, 0, 26, 1, 0, 0, 0, 2, 39, 1, 0, 0, 0,
		4, 41, 1, 0, 0, 0, 6, 43, 1, 0, 0, 0, 8, 45, 1, 0, 0, 0, 10, 50, 1, 0,
		0, 0, 12, 55, 1, 0, 0, 0, 14, 60, 1, 0, 0, 0, 16, 71, 1, 0, 0, 0, 18, 78,
		1, 0, 0, 0, 20, 89, 1, 0, 0, 0, 22, 94, 1, 0, 0, 0, 24, 99, 1, 0, 0, 0,
		26, 27, 3, 2, 1, 0, 27, 1, 1, 0, 0, 0, 28, 40, 3, 4, 2, 0, 29, 40, 3, 6,
		3, 0, 30, 40, 3, 8, 4, 0, 31, 40, 3, 10, 5, 0, 32, 40, 3, 12, 6, 0, 33,
		40, 3, 14, 7, 0, 34, 40, 3, 16, 8, 0, 35, 40, 3, 18, 9, 0, 36, 40, 3, 20,
		10, 0, 37, 40, 3, 22, 11, 0, 38, 40, 3, 24, 12, 0, 39, 28, 1, 0, 0, 0,
		39, 29, 1, 0, 0, 0, 39, 30, 1, 0, 0, 0, 39, 31, 1, 0, 0, 0, 39, 32, 1,
		0, 0, 0, 39, 33, 1, 0, 0, 0, 39, 34, 1, 0, 0, 0, 39, 35, 1, 0, 0, 0, 39,
		36, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 3, 1, 0, 0,
		0, 41, 42, 5, 1, 0, 0, 42, 5, 1, 0, 0, 0, 43, 44, 5, 2, 0, 0, 44, 7, 1,
		0, 0, 0, 45, 46, 5, 3, 0, 0, 46, 47, 5, 4, 0, 0, 47, 48, 5, 16, 0, 0, 48,
		49, 5, 5, 0, 0, 49, 9, 1, 0, 0, 0, 50, 51, 5, 6, 0, 0, 51, 52, 5, 4, 0,
		0, 52, 53, 5, 15, 0, 0, 53, 54, 5, 5, 0, 0, 54, 11, 1, 0, 0, 0, 55, 56,
		5, 7, 0, 0, 56, 57, 5, 4, 0, 0, 57, 58, 5, 15, 0, 0, 58, 59, 5, 5, 0, 0,
		59, 13, 1, 0, 0, 0, 60, 61, 5, 8, 0, 0, 61, 62, 5, 4, 0, 0, 62, 65, 3,
		2, 1, 0, 63, 64, 5, 9, 0, 0, 64, 66, 3, 2, 1, 0, 65, 63, 1, 0, 0, 0, 66,
		67, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 1, 0, 0,
		0, 69, 70, 5, 5, 0, 0, 70, 15, 1, 0, 0, 0, 71, 72, 5, 10, 0, 0, 72, 73,
		5, 4, 0, 0, 73, 74, 3, 2, 1, 0, 74, 75, 5, 9, 0, 0, 75, 76, 3, 2, 1, 0,
		76, 77, 5, 5, 0, 0, 77, 17, 1, 0, 0, 0, 78, 79, 5, 11, 0, 0, 79, 80, 5,
		4, 0, 0, 80, 83, 3, 2, 1, 0, 81, 82, 5, 9, 0, 0, 82, 84, 3, 2, 1, 0, 83,
		81, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0,
		0, 86, 87, 1, 0, 0, 0, 87, 88, 5, 5, 0, 0, 88, 19, 1, 0, 0, 0, 89, 90,
		5, 12, 0, 0, 90, 91, 5, 4, 0, 0, 91, 92, 3, 2, 1, 0, 92, 93, 5, 5, 0, 0,
		93, 21, 1, 0, 0, 0, 94, 95, 5, 13, 0, 0, 95, 96, 5, 4, 0, 0, 96, 97, 3,
		2, 1, 0, 97, 98, 5, 5, 0, 0, 98, 23, 1, 0, 0, 0, 99, 100, 5, 14, 0, 0,
		100, 101, 5, 4, 0, 0, 101, 102, 3, 2, 1, 0, 102, 103, 5, 5, 0, 0, 103,
		25, 1, 0, 0, 0, 3, 39, 67, 85,
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
	BricksParserINT     = 15
	BricksParserSET_NUM = 16
	BricksParserWS      = 17
)

// BricksParser rules.
const (
	BricksParserRULE_bricks         = 0
	BricksParserRULE_exp            = 1
	BricksParserRULE_allParts       = 2
	BricksParserRULE_lost           = 3
	BricksParserRULE_set            = 4
	BricksParserRULE_setList        = 5
	BricksParserRULE_partList       = 6
	BricksParserRULE_sum            = 7
	BricksParserRULE_subtract       = 8
	BricksParserRULE_max            = 9
	BricksParserRULE_mergeByColor   = 10
	BricksParserRULE_mergeByVariant = 11
	BricksParserRULE_sort           = 12
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

func (s *BricksContext) Exp() IExpContext {
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
		p.SetState(26)
		p.Exp()
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

func (s *ExpContext) MergeByColor() IMergeByColorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMergeByColorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMergeByColorContext)
}

func (s *ExpContext) MergeByVariant() IMergeByVariantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMergeByVariantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMergeByVariantContext)
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
	this := p
	_ = this

	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BricksParserRULE_exp)

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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BricksParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.AllParts()
		}

	case BricksParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.Lost()
		}

	case BricksParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
			p.Set()
		}

	case BricksParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(31)
			p.SetList()
		}

	case BricksParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(32)
			p.PartList()
		}

	case BricksParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(33)
			p.Sum()
		}

	case BricksParserT__9:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(34)
			p.Subtract()
		}

	case BricksParserT__10:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(35)
			p.Max()
		}

	case BricksParserT__11:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(36)
			p.MergeByColor()
		}

	case BricksParserT__12:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(37)
			p.MergeByVariant()
		}

	case BricksParserT__13:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(38)
			p.Sort()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 4, BricksParserRULE_allParts)

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
		p.SetState(41)
		p.Match(BricksParserT__0)
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
	p.EnterRule(localctx, 6, BricksParserRULE_lost)

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
		p.SetState(43)
		p.Match(BricksParserT__1)
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
	p.EnterRule(localctx, 8, BricksParserRULE_set)

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
		p.SetState(45)
		p.Match(BricksParserT__2)
	}
	{
		p.SetState(46)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(47)
		p.Match(BricksParserSET_NUM)
	}
	{
		p.SetState(48)
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
	p.EnterRule(localctx, 10, BricksParserRULE_setList)

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
		p.Match(BricksParserT__5)
	}
	{
		p.SetState(51)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(52)
		p.Match(BricksParserINT)
	}
	{
		p.SetState(53)
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
	p.EnterRule(localctx, 12, BricksParserRULE_partList)

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
		p.SetState(55)
		p.Match(BricksParserT__6)
	}
	{
		p.SetState(56)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(57)
		p.Match(BricksParserINT)
	}
	{
		p.SetState(58)
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
	this := p
	_ = this

	localctx = NewSumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BricksParserRULE_sum)
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
		p.SetState(60)
		p.Match(BricksParserT__7)
	}
	{
		p.SetState(61)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(62)
		p.Exp()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BricksParserT__8 {
		{
			p.SetState(63)
			p.Match(BricksParserT__8)
		}
		{
			p.SetState(64)
			p.Exp()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(69)
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
	this := p
	_ = this

	localctx = NewSubtractContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BricksParserRULE_subtract)

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
		p.SetState(71)
		p.Match(BricksParserT__9)
	}
	{
		p.SetState(72)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(73)
		p.Exp()
	}
	{
		p.SetState(74)
		p.Match(BricksParserT__8)
	}
	{
		p.SetState(75)
		p.Exp()
	}
	{
		p.SetState(76)
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
	this := p
	_ = this

	localctx = NewMaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BricksParserRULE_max)
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
		p.SetState(78)
		p.Match(BricksParserT__10)
	}
	{
		p.SetState(79)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(80)
		p.Exp()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BricksParserT__8 {
		{
			p.SetState(81)
			p.Match(BricksParserT__8)
		}
		{
			p.SetState(82)
			p.Exp()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(BricksParserT__4)
	}

	return localctx
}

// IMergeByColorContext is an interface to support dynamic dispatch.
type IMergeByColorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMergeByColorContext differentiates from other interfaces.
	IsMergeByColorContext()
}

type MergeByColorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMergeByColorContext() *MergeByColorContext {
	var p = new(MergeByColorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_mergeByColor
	return p
}

func (*MergeByColorContext) IsMergeByColorContext() {}

func NewMergeByColorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MergeByColorContext {
	var p = new(MergeByColorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_mergeByColor

	return p
}

func (s *MergeByColorContext) GetParser() antlr.Parser { return s.parser }

func (s *MergeByColorContext) Exp() IExpContext {
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

func (s *MergeByColorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MergeByColorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MergeByColorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterMergeByColor(s)
	}
}

func (s *MergeByColorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitMergeByColor(s)
	}
}

func (p *BricksParser) MergeByColor() (localctx IMergeByColorContext) {
	this := p
	_ = this

	localctx = NewMergeByColorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BricksParserRULE_mergeByColor)

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
		p.Match(BricksParserT__11)
	}
	{
		p.SetState(90)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(91)
		p.Exp()
	}
	{
		p.SetState(92)
		p.Match(BricksParserT__4)
	}

	return localctx
}

// IMergeByVariantContext is an interface to support dynamic dispatch.
type IMergeByVariantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMergeByVariantContext differentiates from other interfaces.
	IsMergeByVariantContext()
}

type MergeByVariantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMergeByVariantContext() *MergeByVariantContext {
	var p = new(MergeByVariantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BricksParserRULE_mergeByVariant
	return p
}

func (*MergeByVariantContext) IsMergeByVariantContext() {}

func NewMergeByVariantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MergeByVariantContext {
	var p = new(MergeByVariantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BricksParserRULE_mergeByVariant

	return p
}

func (s *MergeByVariantContext) GetParser() antlr.Parser { return s.parser }

func (s *MergeByVariantContext) Exp() IExpContext {
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

func (s *MergeByVariantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MergeByVariantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MergeByVariantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.EnterMergeByVariant(s)
	}
}

func (s *MergeByVariantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BricksListener); ok {
		listenerT.ExitMergeByVariant(s)
	}
}

func (p *BricksParser) MergeByVariant() (localctx IMergeByVariantContext) {
	this := p
	_ = this

	localctx = NewMergeByVariantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BricksParserRULE_mergeByVariant)

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
		p.SetState(94)
		p.Match(BricksParserT__12)
	}
	{
		p.SetState(95)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(96)
		p.Exp()
	}
	{
		p.SetState(97)
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
	this := p
	_ = this

	localctx = NewSortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BricksParserRULE_sort)

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
		p.SetState(99)
		p.Match(BricksParserT__13)
	}
	{
		p.SetState(100)
		p.Match(BricksParserT__3)
	}
	{
		p.SetState(101)
		p.Exp()
	}
	{
		p.SetState(102)
		p.Match(BricksParserT__4)
	}

	return localctx
}
