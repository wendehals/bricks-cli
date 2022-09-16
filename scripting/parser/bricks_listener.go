// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Bricks

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BricksListener is a complete listener for a parse tree produced by BricksParser.
type BricksListener interface {
	antlr.ParseTreeListener

	// EnterBricks is called when entering the bricks production.
	EnterBricks(c *BricksContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterSave is called when entering the save production.
	EnterSave(c *SaveContext)

	// EnterExport is called when entering the export production.
	EnterExport(c *ExportContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLoad is called when entering the load production.
	EnterLoad(c *LoadContext)

	// EnterAllParts is called when entering the allParts production.
	EnterAllParts(c *AllPartsContext)

	// EnterLost is called when entering the lost production.
	EnterLost(c *LostContext)

	// EnterSet is called when entering the set production.
	EnterSet(c *SetContext)

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

	// ExitBricks is called when exiting the bricks production.
	ExitBricks(c *BricksContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitSave is called when exiting the save production.
	ExitSave(c *SaveContext)

	// ExitExport is called when exiting the export production.
	ExitExport(c *ExportContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLoad is called when exiting the load production.
	ExitLoad(c *LoadContext)

	// ExitAllParts is called when exiting the allParts production.
	ExitAllParts(c *AllPartsContext)

	// ExitLost is called when exiting the lost production.
	ExitLost(c *LostContext)

	// ExitSet is called when exiting the set production.
	ExitSet(c *SetContext)

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
}
