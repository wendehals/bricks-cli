package scripting

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/wendehals/bricks-cli/api"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-cli/scripting/parser"
	"github.com/wendehals/bricks-cli/services"
)

type bricksInterpreter struct {
	*parser.BaseBricksListener

	usersAPI  *api.UsersAPI
	bricksAPI *api.BricksAPI

	heap  map[string]*model.Collection
	stack stack
}

func newBricksInterpreter(usersAPI *api.UsersAPI, bricksAPI *api.BricksAPI) *bricksInterpreter {
	return &bricksInterpreter{
		usersAPI:  usersAPI,
		bricksAPI: bricksAPI,
		heap:      make(map[string]*model.Collection),
		stack:     *newStack(),
	}
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
	exportDir := strings.Trim(ctx.STRING().GetText(), "\"")
	services.ExportCollectionToHTML(collection, exportDir, exportDir)
}

func (b *bricksInterpreter) ExitBuild(ctx *parser.BuildContext) {
	providedCollection := b.stack.pop()
	neededCollection := b.stack.pop()
	exportDir := strings.Trim(ctx.STRING().GetText(), "\"")

	var mode uint8 = 0
	if ctx.GetBuild_mode() != nil {
		mode = services.ModeToUInt8(ctx.GetBuild_mode().GetText())
	}

	buildCollection := services.Build(neededCollection, providedCollection, mode)
	buildCollection.Save(fmt.Sprintf("%s/result.build", exportDir))
	services.ExportBuildCollectionToHTML(buildCollection, exportDir, "build")

	providedCollection.Save(fmt.Sprintf("%s/remaining.parts", exportDir))
	services.ExportCollectionToHTML(providedCollection, exportDir, "remaining")
}

// ExitPause is called when production pause is exited.
func (b *bricksInterpreter) ExitPause(ctx *parser.PauseContext) {
	seconds, err := strconv.Atoi(ctx.GetSeconds().GetText())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Pausing for %d seconds...", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
}

func (b *bricksInterpreter) ExitIdentifier(ctx *parser.IdentifierContext) {
	value, ok := b.heap[ctx.ID().GetText()]
	if !ok {
		log.Fatalf("The variable '%s' is not defined", ctx.ID().GetText())
	}
	b.stack.push(model.DeepClone(value))
}

func (b *bricksInterpreter) ExitLoad(ctx *parser.LoadContext) {
	filePath := strings.Trim(ctx.STRING().GetText(), "\"")

	log.Printf("Loading collection from '%s'", filePath)
	collection := model.Load[model.Collection](filePath)
	b.stack.push(&collection)
}

func (b *bricksInterpreter) ExitImport_(ctx *parser.Import_Context) {
	filePath := strings.Trim(ctx.STRING().GetText(), "\"")

	b.stack.push(services.ImportCSVCollection(filePath))
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

func (b *bricksInterpreter) ExitUserSet(ctx *parser.UserSetContext) {
	includeMiniFigs := ctx.BOOL() != nil && ctx.BOOL().GetText() == "true"
	userSetParts := api.RetrieveUserSetParts(b.bricksAPI, b.usersAPI, ctx.SET_NUM().GetText(), includeMiniFigs)
	b.stack.push(userSetParts)
}

func (b *bricksInterpreter) ExitSetList(ctx *parser.SetListContext) {
	setListId, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		log.Fatal(err)
	}

	includeMiniFigs := ctx.BOOL() != nil && ctx.BOOL().GetText() == "true"
	setListParts := api.RetrieveSetListParts(b.bricksAPI, b.usersAPI, uint(setListId), includeMiniFigs)
	collection := services.MergeAllCollections(setListParts)
	collection.SortByColorAndName(false)

	b.stack.push(collection)
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
	collections := api.RetrievePartListsParts(b.usersAPI, filePath, includeNonBuildable)
	collection := services.MergeAllCollections(collections)
	b.stack.push(collection)
}

func (b *bricksInterpreter) ExitSum(ctx *parser.SumContext) {
	log.Printf("Calculating sum")

	sum := model.NewCollection()

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

	max := model.NewCollection()

	for range ctx.AllExp() {
		collection := b.stack.pop()
		max.Max(collection)
	}

	b.stack.push(max)
}

func (b *bricksInterpreter) ExitSort(ctx *parser.SortContext) {
	log.Printf("Sorting collection")

	collection := b.stack.pop()

	if ctx.GetQuantity() == nil {
		collection.SortByColorAndName(ctx.GetDescending() != nil)
	} else {
		collection.SortByQuantityAndName(ctx.GetDescending() != nil)
	}

	b.stack.push(collection)
}
