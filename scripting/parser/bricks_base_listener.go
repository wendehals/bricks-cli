// Code generated from .\scripting\parser\Bricks.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bricks

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBricksListener is a complete listener for a parse tree produced by BricksParser.
type BaseBricksListener struct{}

var _ BricksListener = &BaseBricksListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBricksListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBricksListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBricksListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBricksListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBricks is called when production bricks is entered.
func (s *BaseBricksListener) EnterBricks(ctx *BricksContext) {}

// ExitBricks is called when production bricks is exited.
func (s *BaseBricksListener) ExitBricks(ctx *BricksContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseBricksListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseBricksListener) ExitExp(ctx *ExpContext) {}

// EnterAllParts is called when production allParts is entered.
func (s *BaseBricksListener) EnterAllParts(ctx *AllPartsContext) {}

// ExitAllParts is called when production allParts is exited.
func (s *BaseBricksListener) ExitAllParts(ctx *AllPartsContext) {}

// EnterLost is called when production lost is entered.
func (s *BaseBricksListener) EnterLost(ctx *LostContext) {}

// ExitLost is called when production lost is exited.
func (s *BaseBricksListener) ExitLost(ctx *LostContext) {}

// EnterSet is called when production set is entered.
func (s *BaseBricksListener) EnterSet(ctx *SetContext) {}

// ExitSet is called when production set is exited.
func (s *BaseBricksListener) ExitSet(ctx *SetContext) {}

// EnterSetList is called when production setList is entered.
func (s *BaseBricksListener) EnterSetList(ctx *SetListContext) {}

// ExitSetList is called when production setList is exited.
func (s *BaseBricksListener) ExitSetList(ctx *SetListContext) {}

// EnterPartList is called when production partList is entered.
func (s *BaseBricksListener) EnterPartList(ctx *PartListContext) {}

// ExitPartList is called when production partList is exited.
func (s *BaseBricksListener) ExitPartList(ctx *PartListContext) {}

// EnterSum is called when production sum is entered.
func (s *BaseBricksListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production sum is exited.
func (s *BaseBricksListener) ExitSum(ctx *SumContext) {}

// EnterSubtract is called when production subtract is entered.
func (s *BaseBricksListener) EnterSubtract(ctx *SubtractContext) {}

// ExitSubtract is called when production subtract is exited.
func (s *BaseBricksListener) ExitSubtract(ctx *SubtractContext) {}

// EnterMax is called when production max is entered.
func (s *BaseBricksListener) EnterMax(ctx *MaxContext) {}

// ExitMax is called when production max is exited.
func (s *BaseBricksListener) ExitMax(ctx *MaxContext) {}

// EnterMergeByColor is called when production mergeByColor is entered.
func (s *BaseBricksListener) EnterMergeByColor(ctx *MergeByColorContext) {}

// ExitMergeByColor is called when production mergeByColor is exited.
func (s *BaseBricksListener) ExitMergeByColor(ctx *MergeByColorContext) {}

// EnterMergeByVariant is called when production mergeByVariant is entered.
func (s *BaseBricksListener) EnterMergeByVariant(ctx *MergeByVariantContext) {}

// ExitMergeByVariant is called when production mergeByVariant is exited.
func (s *BaseBricksListener) ExitMergeByVariant(ctx *MergeByVariantContext) {}

// EnterSort is called when production sort is entered.
func (s *BaseBricksListener) EnterSort(ctx *SortContext) {}

// ExitSort is called when production sort is exited.
func (s *BaseBricksListener) ExitSort(ctx *SortContext) {}
