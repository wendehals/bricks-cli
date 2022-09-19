// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Bricks

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

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

// EnterCommand is called when production command is entered.
func (s *BaseBricksListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseBricksListener) ExitCommand(ctx *CommandContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseBricksListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseBricksListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterSave is called when production save is entered.
func (s *BaseBricksListener) EnterSave(ctx *SaveContext) {}

// ExitSave is called when production save is exited.
func (s *BaseBricksListener) ExitSave(ctx *SaveContext) {}

// EnterExport is called when production export is entered.
func (s *BaseBricksListener) EnterExport(ctx *ExportContext) {}

// ExitExport is called when production export is exited.
func (s *BaseBricksListener) ExitExport(ctx *ExportContext) {}

// EnterBuild is called when production build is entered.
func (s *BaseBricksListener) EnterBuild(ctx *BuildContext) {}

// ExitBuild is called when production build is exited.
func (s *BaseBricksListener) ExitBuild(ctx *BuildContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseBricksListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseBricksListener) ExitExp(ctx *ExpContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseBricksListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseBricksListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLoad is called when production load is entered.
func (s *BaseBricksListener) EnterLoad(ctx *LoadContext) {}

// ExitLoad is called when production load is exited.
func (s *BaseBricksListener) ExitLoad(ctx *LoadContext) {}

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

// EnterPartLists is called when production partLists is entered.
func (s *BaseBricksListener) EnterPartLists(ctx *PartListsContext) {}

// ExitPartLists is called when production partLists is exited.
func (s *BaseBricksListener) ExitPartLists(ctx *PartListsContext) {}

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

// EnterSort is called when production sort is entered.
func (s *BaseBricksListener) EnterSort(ctx *SortContext) {}

// ExitSort is called when production sort is exited.
func (s *BaseBricksListener) ExitSort(ctx *SortContext) {}
