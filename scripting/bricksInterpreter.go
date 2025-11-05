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

	heap  map[string]model.CollectionType
	stack stack[model.CollectionType]
}

func newBricksInterpreter(usersAPI *api.UsersAPI, bricksAPI *api.BricksAPI) *bricksInterpreter {
	return &bricksInterpreter{
		usersAPI:  usersAPI,
		bricksAPI: bricksAPI,
		heap:      make(map[string]model.CollectionType),
		stack:     *newStack[model.CollectionType](),
	}
}

func (b *bricksInterpreter) ExitAssignment(ctx *parser.AssignmentContext) {
	b.heap[ctx.ID().GetText()] = b.stack.pop()
}

func (b *bricksInterpreter) ExitSave(ctx *parser.SaveContext) {
	b.stack.pop().Save(strings.Trim(ctx.STRING().GetText(), "\""))
}

func (b *bricksInterpreter) ExitExport(ctx *parser.ExportContext) {
	value := b.stack.pop()
	exportDir := strings.Trim(ctx.STRING().GetText(), "\"")

	switch collection := value.(type) {
	case model.BuildCollection:
		fileName := fmt.Sprintf("build_%s", collection.Set.Name)
		services.ExportBuildCollectionToHTML(&collection, exportDir, fileName)
	case model.Collection:
		fileName := "collection"
		for i := 0; i < len(collection.Sets) && i < 8; i++ {
			fileName += "_" + collection.Sets[i].Number
		}
		services.ExportCollectionToHTML(&collection, exportDir, fileName)
	}
}

func (b *bricksInterpreter) ExitPrint(ctx *parser.PrintContext) {
	value := b.stack.pop()
	value.Print()
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

func (b *bricksInterpreter) ExitBuild(ctx *parser.BuildContext) {
	firstValue := b.stack.pop()
	secondValue := b.stack.pop()

	var mode uint8 = 0
	if ctx.GetBuild_mode() != nil {
		mode = services.ModeToUInt8(ctx.GetBuild_mode().GetText())
	}

	if neededCollection, ok := firstValue.(model.Collection); ok {
		if providedCollection, ok := secondValue.(model.Collection); ok {
			result := services.Build(&neededCollection, &providedCollection, mode)
			b.stack.push(result)
		} else {
			log.Fatal("The second operand of build must be a collection")
			return
		}
	} else {
		log.Fatal("The first operand of build must be a collection")
		return
	}
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
	b.stack.push(collection)
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

	for range ctx.AllCollectionExp() {
		value := b.stack.pop()
		if summand, ok := value.(model.Collection); ok {
			sum.Add(&summand)
		} else {
			log.Fatal("Only collections can be summed")
		}
	}

	b.stack.push(sum)
}

func (b *bricksInterpreter) ExitSubtract(ctx *parser.SubtractContext) {
	log.Printf("Calculating difference")

	firstValue := b.stack.pop()
	secondValue := b.stack.pop()
	if subtrahend, ok := firstValue.(model.Collection); ok {
		if minuend, ok := secondValue.(model.Collection); ok {
			minuend.Subtract(&subtrahend)
			b.stack.push(minuend)
		} else {
			log.Fatal("The second operand of subtract must be a collection")
		}
	} else {
		log.Fatal("The first operand of subtract must be a collection")
	}
}

func (b *bricksInterpreter) ExitMax(ctx *parser.MaxContext) {
	log.Printf("Calculating maximum")

	max := model.NewCollection()
	for range ctx.AllCollectionExp() {
		value := b.stack.pop()
		if collection, ok := value.(model.Collection); ok {
			max.Max(&collection)
		} else {
			log.Fatal("Only collections can be compared")
		}

		b.stack.push(max)
	}
}

func (b *bricksInterpreter) ExitSort(ctx *parser.SortContext) {
	log.Printf("Sorting collection")

	value := b.stack.pop()
	if collection, ok := value.(model.Collection); ok {
		if ctx.GetQuantity() == nil {
			collection.SortByColorAndName(ctx.GetDescending() != nil)
		} else {
			collection.SortByQuantityAndName(ctx.GetDescending() != nil)
		}
		b.stack.push(collection)
	} else {
		log.Fatal("Only collections can be sorted")
	}
}
