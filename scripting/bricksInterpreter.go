package scripting

import (
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
	model.Save(collection, filePath)
}

func (b *bricksInterpreter) ExitExport(ctx *parser.ExportContext) {
	collection := b.stack.pop()
	exportDir := strings.Trim(ctx.STRING().GetText(), "\"")
	collection.ExportToHTML(exportDir)
}

func (b *bricksInterpreter) ExitIdentifier(ctx *parser.IdentifierContext) {
	value, ok := b.heap[ctx.ID().GetText()]
	if !ok {
		log.Fatalf("The variable '%s' is not defined", ctx.ID().GetText())
	}
	b.stack.push(model.DeepClone(value, &model.Collection{}))
}

func (b *bricksInterpreter) ExitLoad(ctx *parser.LoadContext) {
	filePath := strings.Trim(ctx.STRING().GetText(), "\"")

	log.Printf("Loading collection from '%s'", filePath)
	b.stack.push(model.Load(&model.Collection{}, filePath))
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
	includeMiniFigs := ctx.BOOL() != nil && ctx.BOOL().GetText() == "true"
	setParts := api.RetrieveSetParts(b.bricksAPI, ctx.SET_NUM().GetText(), includeMiniFigs)
	b.stack.push(setParts)
}

func (b *bricksInterpreter) ExitSetList(ctx *parser.SetListContext) {
	setListId, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		log.Fatal(err)
	}
	includeMiniFigs := ctx.BOOL() != nil && ctx.BOOL().GetText() == "true"
	setListParts := api.RetrieveSetListParts(b.bricksAPI, b.usersAPI, uint(setListId), includeMiniFigs)
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
	includeNonBuildable := ctx.BOOL() != nil && ctx.BOOL().GetText() == "true"
	collections := api.RetrievePartListParts(b.usersAPI, filePath, includeNonBuildable)
	collection := model.MergeAllCollections(collections)
	b.stack.push(collection)
}

func (b *bricksInterpreter) ExitSum(ctx *parser.SumContext) {
	log.Printf("Calculating sum")

	sum := &model.Collection{}

	for range ctx.AllExp() {
		summand := b.stack.pop()
		sum.Add(summand)
	}

	b.stack.push(sum)
}

func (b *bricksInterpreter) ExitSubtract(ctx *parser.SubtractContext) {
	log.Printf("Calculating difference")

	subtrahend := b.stack.pop()
	minuend := b.stack.pop()
	difference := minuend.Subtract(subtrahend)
	b.stack.push(difference)
}

func (b *bricksInterpreter) ExitMax(ctx *parser.MaxContext) {
	log.Printf("Calculating maximum")

	max := &model.Collection{}

	for range ctx.AllExp() {
		collection := b.stack.pop()
		max.Max(collection)
	}

	b.stack.push(max)
}

func (b *bricksInterpreter) ExitSort(ctx *parser.SortContext) {
	log.Printf("Sorting collection")

	collection := b.stack.pop()
	collection.Sort()
	b.stack.push(collection)
}
