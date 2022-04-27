package scripting

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wendehals/bricks/api"
	cmdapi "github.com/wendehals/bricks/cmd/api"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/scripting/parser"
)

type bricksTreeListener struct {
	*parser.BaseBricksListener

	stack     scriptStack
	usersAPI  *api.UsersAPI
	bricksAPI *api.BricksAPI
}

func newBricksTreeListener(usersAPI *api.UsersAPI, bricksAPI *api.BricksAPI) *bricksTreeListener {
	b := &bricksTreeListener{}
	b.stack = *newScriptStack()
	b.usersAPI = usersAPI
	b.bricksAPI = bricksAPI

	return b
}

func (s *bricksTreeListener) VisitTerminal(node antlr.TerminalNode) {
	if node.GetSymbol().GetTokenType() == parser.BricksLexerSET_NUM ||
		node.GetSymbol().GetTokenType() == parser.BricksLexerINT {
		s.stack.push(node.GetText())
	}
}

func (b *bricksTreeListener) ExitBricks(ctx *parser.BricksContext) {
	collection := b.stack.pop().(*model.Collection)
	model.ExportToJSON("script_result_parts.json", collection)
}

func (b *bricksTreeListener) ExitAllParts(ctx *parser.AllPartsContext) {
	allParts := b.usersAPI.GetAllParts()
	b.stack.push(allParts)
}

func (b *bricksTreeListener) ExitLost(ctx *parser.LostContext) {
	lostParts := b.usersAPI.GetLostParts()
	b.stack.push(lostParts)
}

func (b *bricksTreeListener) ExitSet(ctx *parser.SetContext) {
	setNum := b.stack.pop().(string)
	collection := cmdapi.RetrieveSetParts(b.bricksAPI, setNum)
	b.stack.push(collection)
}

func (b *bricksTreeListener) ExitSetList(ctx *parser.SetListContext) {
}

func (b *bricksTreeListener) ExitPartList(ctx *parser.PartListContext) {
}

func (b *bricksTreeListener) EnterSum(ctx *parser.SumContext) {
	b.stack.push("sum")
}

func (b *bricksTreeListener) ExitSum(ctx *parser.SumContext) {
	sum := b.stack.pop().(*model.Collection)

	end := false
	for !end {
		element := b.stack.pop()
		switch v := element.(type) {
		case string:
			if element.(string) == "sum" {
				end = true
			} else {
				log.Fatalf("unexpected element on stack: %v", v)
			}
		case *model.Collection:
			summand := element.(*model.Collection)
			sum.Add(summand)
		}
	}

	b.stack.push(sum)
}

func (b *bricksTreeListener) ExitSubtract(ctx *parser.SubtractContext) {
	subtrahend := b.stack.pop().(*model.Collection)
	minuend := b.stack.pop().(*model.Collection)
	result := minuend.Subtract(subtrahend)
	b.stack.push(result)
}

func (b *bricksTreeListener) EnterMax(ctx *parser.MaxContext) {
	b.stack.push("max")
}

func (b *bricksTreeListener) ExitMax(ctx *parser.MaxContext) {
	collection1 := b.stack.pop().(*model.Collection)

	end := false
	for !end {
		element := b.stack.pop()
		switch v := element.(type) {
		case string:
			if element.(string) == "max" {
				end = true
			} else {
				log.Fatalf("unexpected element on stack: %v", v)
			}
		case *model.Collection:
			collection2 := element.(*model.Collection)
			collection1.Max(collection2)
		}
	}

	b.stack.push(collection1)
}

func (b *bricksTreeListener) ExitMergeByColor(ctx *parser.MergeByColorContext) {
	collection := b.stack.pop().(*model.Collection)
	collection.MergeByColor()
	b.stack.push(collection)
}

func (b *bricksTreeListener) ExitMergeByVariant(ctx *parser.MergeByVariantContext) {
	collection := b.stack.pop().(*model.Collection)
	collection.MergeByVariant()
	b.stack.push(collection)
}

func (b *bricksTreeListener) ExitSort(ctx *parser.SortContext) {
	collection := b.stack.pop().(*model.Collection)
	collection.Sort()
	b.stack.push(collection)
}
