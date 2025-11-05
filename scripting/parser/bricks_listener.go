// Code generated from ./scripting/parser/Bricks.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Bricks

import "github.com/antlr4-go/antlr/v4"

// BricksListener is a complete listener for a parse tree produced by BricksParser.
type BricksListener interface {
	antlr.ParseTreeListener

	// EnterBricks is called when entering the bricks production.
	EnterBricks(c *BricksContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterSave is called when entering the save production.
	EnterSave(c *SaveContext)

	// EnterExport is called when entering the export production.
	EnterExport(c *ExportContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterPause is called when entering the pause production.
	EnterPause(c *PauseContext)

	// EnterCollectionOrId is called when entering the collectionOrId production.
	EnterCollectionOrId(c *CollectionOrIdContext)

	// EnterCollection is called when entering the collection production.
	EnterCollection(c *CollectionContext)

	// EnterLoad is called when entering the load production.
	EnterLoad(c *LoadContext)

	// EnterImport_ is called when entering the import_ production.
	EnterImport_(c *Import_Context)

	// EnterAllParts is called when entering the allParts production.
	EnterAllParts(c *AllPartsContext)

	// EnterLost is called when entering the lost production.
	EnterLost(c *LostContext)

	// EnterSet is called when entering the set production.
	EnterSet(c *SetContext)

	// EnterUserSet is called when entering the userSet production.
	EnterUserSet(c *UserSetContext)

	// EnterSetList is called when entering the setList production.
	EnterSetList(c *SetListContext)

	// EnterPartList is called when entering the partList production.
	EnterPartList(c *PartListContext)

	// EnterPartLists is called when entering the partLists production.
	EnterPartLists(c *PartListsContext)

	// EnterSum is called when entering the sum production.
	EnterSum(c *SumContext)

	// EnterSubtract is called when entering the subtract production.
	EnterSubtract(c *SubtractContext)

	// EnterMax is called when entering the max production.
	EnterMax(c *MaxContext)

	// EnterSort is called when entering the sort production.
	EnterSort(c *SortContext)

	// EnterBuild is called when entering the build production.
	EnterBuild(c *BuildContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitBricks is called when exiting the bricks production.
	ExitBricks(c *BricksContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitSave is called when exiting the save production.
	ExitSave(c *SaveContext)

	// ExitExport is called when exiting the export production.
	ExitExport(c *ExportContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitPause is called when exiting the pause production.
	ExitPause(c *PauseContext)

	// ExitCollectionOrId is called when exiting the collectionOrId production.
	ExitCollectionOrId(c *CollectionOrIdContext)

	// ExitCollection is called when exiting the collection production.
	ExitCollection(c *CollectionContext)

	// ExitLoad is called when exiting the load production.
	ExitLoad(c *LoadContext)

	// ExitImport_ is called when exiting the import_ production.
	ExitImport_(c *Import_Context)

	// ExitAllParts is called when exiting the allParts production.
	ExitAllParts(c *AllPartsContext)

	// ExitLost is called when exiting the lost production.
	ExitLost(c *LostContext)

	// ExitSet is called when exiting the set production.
	ExitSet(c *SetContext)

	// ExitUserSet is called when exiting the userSet production.
	ExitUserSet(c *UserSetContext)

	// ExitSetList is called when exiting the setList production.
	ExitSetList(c *SetListContext)

	// ExitPartList is called when exiting the partList production.
	ExitPartList(c *PartListContext)

	// ExitPartLists is called when exiting the partLists production.
	ExitPartLists(c *PartListsContext)

	// ExitSum is called when exiting the sum production.
	ExitSum(c *SumContext)

	// ExitSubtract is called when exiting the subtract production.
	ExitSubtract(c *SubtractContext)

	// ExitMax is called when exiting the max production.
	ExitMax(c *MaxContext)

	// ExitSort is called when exiting the sort production.
	ExitSort(c *SortContext)

	// ExitBuild is called when exiting the build production.
	ExitBuild(c *BuildContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
