// Code generated from .\scripting\parser\Bricks.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bricks

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BricksListener is a complete listener for a parse tree produced by BricksParser.
type BricksListener interface {
	antlr.ParseTreeListener

	// EnterBricks is called when entering the bricks production.
	EnterBricks(c *BricksContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

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

	// EnterSum is called when entering the sum production.
	EnterSum(c *SumContext)

	// EnterSubtract is called when entering the subtract production.
	EnterSubtract(c *SubtractContext)

	// EnterMax is called when entering the max production.
	EnterMax(c *MaxContext)

	// EnterMergeByColor is called when entering the mergeByColor production.
	EnterMergeByColor(c *MergeByColorContext)

	// EnterMergeByVariant is called when entering the mergeByVariant production.
	EnterMergeByVariant(c *MergeByVariantContext)

	// EnterSort is called when entering the sort production.
	EnterSort(c *SortContext)

	// ExitBricks is called when exiting the bricks production.
	ExitBricks(c *BricksContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

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

	// ExitSum is called when exiting the sum production.
	ExitSum(c *SumContext)

	// ExitSubtract is called when exiting the subtract production.
	ExitSubtract(c *SubtractContext)

	// ExitMax is called when exiting the max production.
	ExitMax(c *MaxContext)

	// ExitMergeByColor is called when exiting the mergeByColor production.
	ExitMergeByColor(c *MergeByColorContext)

	// ExitMergeByVariant is called when exiting the mergeByVariant production.
	ExitMergeByVariant(c *MergeByVariantContext)

	// ExitSort is called when exiting the sort production.
	ExitSort(c *SortContext)
}
