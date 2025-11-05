// Code generated from ./scripting/parser/Bricks.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Bricks

import "github.com/antlr4-go/antlr/v4"

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

// EnterExpression is called when production expression is entered.
func (s *BaseBricksListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBricksListener) ExitExpression(ctx *ExpressionContext) {}

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

// EnterPrint is called when production print is entered.
func (s *BaseBricksListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseBricksListener) ExitPrint(ctx *PrintContext) {}

// EnterPause is called when production pause is entered.
func (s *BaseBricksListener) EnterPause(ctx *PauseContext) {}

// ExitPause is called when production pause is exited.
func (s *BaseBricksListener) ExitPause(ctx *PauseContext) {}

// EnterCollectionOrId is called when production collectionOrId is entered.
func (s *BaseBricksListener) EnterCollectionOrId(ctx *CollectionOrIdContext) {}

// ExitCollectionOrId is called when production collectionOrId is exited.
func (s *BaseBricksListener) ExitCollectionOrId(ctx *CollectionOrIdContext) {}

// EnterCollection is called when production collection is entered.
func (s *BaseBricksListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BaseBricksListener) ExitCollection(ctx *CollectionContext) {}

// EnterLoad is called when production load is entered.
func (s *BaseBricksListener) EnterLoad(ctx *LoadContext) {}

// ExitLoad is called when production load is exited.
func (s *BaseBricksListener) ExitLoad(ctx *LoadContext) {}

// EnterImport_ is called when production import_ is entered.
func (s *BaseBricksListener) EnterImport_(ctx *Import_Context) {}

// ExitImport_ is called when production import_ is exited.
func (s *BaseBricksListener) ExitImport_(ctx *Import_Context) {}

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

// EnterUserSet is called when production userSet is entered.
func (s *BaseBricksListener) EnterUserSet(ctx *UserSetContext) {}

// ExitUserSet is called when production userSet is exited.
func (s *BaseBricksListener) ExitUserSet(ctx *UserSetContext) {}

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

// EnterBuild is called when production build is entered.
func (s *BaseBricksListener) EnterBuild(ctx *BuildContext) {}

// ExitBuild is called when production build is exited.
func (s *BaseBricksListener) ExitBuild(ctx *BuildContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseBricksListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseBricksListener) ExitIdentifier(ctx *IdentifierContext) {}
