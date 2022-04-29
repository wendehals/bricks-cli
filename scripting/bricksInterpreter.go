package scripting

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/scripting/parser"
)

type bricksInterpreter struct {
	*parser.BaseBricksListener

	usersAPI  *api.UsersAPI
	bricksAPI *api.BricksAPI

	heap  map[string]*model.Collection
	stack stack
}

func newBricksInterpreter(usersAPI *api.UsersAPI, bricksAPI *api.BricksAPI) *bricksInterpreter {
	b := &bricksInterpreter{}
	b.usersAPI = usersAPI
	b.bricksAPI = bricksAPI
	b.heap = make(map[string]*model.Collection)
	b.stack = *newStack()

	return b
}

func (b *bricksInterpreter) ExitAssignment(ctx *parser.AssignmentContext) {
	b.heap[ctx.ID().GetText()] = b.stack.pop()
}

func (b *bricksInterpreter) ExitSave(ctx *parser.SaveContext) {
	collection := b.stack.pop()
	filePath := strings.Trim(ctx.STRING().GetText(), "\"")
	collection.Save(filePath)
}

func (b *bricksInterpreter) ExitExport(ctx *parser.ExportContext) {
	collection := b.stack.pop()
	filePath := strings.Trim(ctx.STRING().GetText(), "\"")
	collection.ExportToHTML(filePath)
}

func (b *bricksInterpreter) ExitIdentifier(ctx *parser.IdentifierContext) {
	value, ok := b.heap[ctx.ID().GetText()]
	if !ok {
		log.Fatalf("The variable '%s' is not defined", ctx.ID().GetText())
	}
	b.stack.push(value)
}

func (b *bricksInterpreter) ExitLoad(ctx *parser.LoadContext) {
	filePath := strings.Trim(ctx.STRING().GetText(), "\"")

	fmt.Printf("Loading collection from '%s'", filePath)
	b.stack.push(model.ImportCollection(filePath))
}

func (b *bricksInterpreter) ExitAllParts(ctx *parser.AllPartsContext) {
	allParts := b.usersAPI.GetAllParts()
	b.stack.push(allParts)
}

func (b *bricksInterpreter) ExitLost(ctx *parser.LostContext) {
	lostParts := b.usersAPI.GetLostParts()
	b.stack.push(lostParts)
}

func (b *bricksInterpreter) ExitSet(ctx *parser.SetContext) {
	setParts := api.RetrieveSetParts(b.bricksAPI, ctx.SET_NUM().GetText())
	b.stack.push(setParts)
}

func (b *bricksInterpreter) ExitSetList(ctx *parser.SetListContext) {
	setListId, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		log.Fatal(err)
	}
	setListParts := api.RetrieveSetListParts(b.bricksAPI, b.usersAPI, uint(setListId))
	b.stack.push(model.MergeAllCollections(setListParts))
}

func (b *bricksInterpreter) ExitPartList(ctx *parser.PartListContext) {
	partListId, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		log.Fatal(err)
	}
	partListParts := b.usersAPI.GetPartListParts(uint(partListId))
	b.stack.push(partListParts)
}

func (b *bricksInterpreter) ExitPartLists(ctx *parser.PartListsContext) {
	filePath := strings.Trim(ctx.STRING().GetText(), "\"")
	collections := api.RetrievePartListParts(b.usersAPI, filePath, true)
	collection := model.MergeAllCollections(collections)
	b.stack.push(collection)
}

func (b *bricksInterpreter) ExitSum(ctx *parser.SumContext) {
	sum := &model.Collection{}

	for range ctx.AllExp() {
		summand := b.stack.pop()
		sum.Add(summand)
	}

	b.stack.push(sum)
}

func (b *bricksInterpreter) ExitSubtract(ctx *parser.SubtractContext) {
	subtrahend := b.stack.pop()
	minuend := b.stack.pop()
	result := minuend.Subtract(subtrahend)
	b.stack.push(result)
}

func (b *bricksInterpreter) ExitMax(ctx *parser.MaxContext) {
	max := &model.Collection{}

	for range ctx.AllExp() {
		collection := b.stack.pop()
		max.Max(collection)
	}

	b.stack.push(max)
}

func (b *bricksInterpreter) ExitMergeByColor(ctx *parser.MergeByColorContext) {
	collection := b.stack.pop()
	collection.MergeByColor()
	b.stack.push(collection)
}

func (b *bricksInterpreter) ExitMergeByVariant(ctx *parser.MergeByVariantContext) {
	collection := b.stack.pop()
	collection.MergeByVariant()
	b.stack.push(collection)
}

func (b *bricksInterpreter) ExitSort(ctx *parser.SortContext) {
	collection := b.stack.pop()
	collection.Sort()
	b.stack.push(collection)
}
