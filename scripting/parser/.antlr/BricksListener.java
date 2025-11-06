// Generated from p:/Entwicklung/bricks-cli/scripting/parser/Bricks.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link BricksParser}.
 */
public interface BricksListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link BricksParser#bricks}.
	 * @param ctx the parse tree
	 */
	void enterBricks(BricksParser.BricksContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#bricks}.
	 * @param ctx the parse tree
	 */
	void exitBricks(BricksParser.BricksContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#command}.
	 * @param ctx the parse tree
	 */
	void enterCommand(BricksParser.CommandContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#command}.
	 * @param ctx the parse tree
	 */
	void exitCommand(BricksParser.CommandContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(BricksParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(BricksParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#assignment}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(BricksParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#assignment}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(BricksParser.AssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#save}.
	 * @param ctx the parse tree
	 */
	void enterSave(BricksParser.SaveContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#save}.
	 * @param ctx the parse tree
	 */
	void exitSave(BricksParser.SaveContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#export}.
	 * @param ctx the parse tree
	 */
	void enterExport(BricksParser.ExportContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#export}.
	 * @param ctx the parse tree
	 */
	void exitExport(BricksParser.ExportContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#print}.
	 * @param ctx the parse tree
	 */
	void enterPrint(BricksParser.PrintContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#print}.
	 * @param ctx the parse tree
	 */
	void exitPrint(BricksParser.PrintContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#pause}.
	 * @param ctx the parse tree
	 */
	void enterPause(BricksParser.PauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#pause}.
	 * @param ctx the parse tree
	 */
	void exitPause(BricksParser.PauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#collectionOrId}.
	 * @param ctx the parse tree
	 */
	void enterCollectionOrId(BricksParser.CollectionOrIdContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#collectionOrId}.
	 * @param ctx the parse tree
	 */
	void exitCollectionOrId(BricksParser.CollectionOrIdContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#collection}.
	 * @param ctx the parse tree
	 */
	void enterCollection(BricksParser.CollectionContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#collection}.
	 * @param ctx the parse tree
	 */
	void exitCollection(BricksParser.CollectionContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#load}.
	 * @param ctx the parse tree
	 */
	void enterLoad(BricksParser.LoadContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#load}.
	 * @param ctx the parse tree
	 */
	void exitLoad(BricksParser.LoadContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#import_}.
	 * @param ctx the parse tree
	 */
	void enterImport_(BricksParser.Import_Context ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#import_}.
	 * @param ctx the parse tree
	 */
	void exitImport_(BricksParser.Import_Context ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#allParts}.
	 * @param ctx the parse tree
	 */
	void enterAllParts(BricksParser.AllPartsContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#allParts}.
	 * @param ctx the parse tree
	 */
	void exitAllParts(BricksParser.AllPartsContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#lost}.
	 * @param ctx the parse tree
	 */
	void enterLost(BricksParser.LostContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#lost}.
	 * @param ctx the parse tree
	 */
	void exitLost(BricksParser.LostContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#set}.
	 * @param ctx the parse tree
	 */
	void enterSet(BricksParser.SetContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#set}.
	 * @param ctx the parse tree
	 */
	void exitSet(BricksParser.SetContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#userSet}.
	 * @param ctx the parse tree
	 */
	void enterUserSet(BricksParser.UserSetContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#userSet}.
	 * @param ctx the parse tree
	 */
	void exitUserSet(BricksParser.UserSetContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#setList}.
	 * @param ctx the parse tree
	 */
	void enterSetList(BricksParser.SetListContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#setList}.
	 * @param ctx the parse tree
	 */
	void exitSetList(BricksParser.SetListContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#partList}.
	 * @param ctx the parse tree
	 */
	void enterPartList(BricksParser.PartListContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#partList}.
	 * @param ctx the parse tree
	 */
	void exitPartList(BricksParser.PartListContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#partLists}.
	 * @param ctx the parse tree
	 */
	void enterPartLists(BricksParser.PartListsContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#partLists}.
	 * @param ctx the parse tree
	 */
	void exitPartLists(BricksParser.PartListsContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#sum}.
	 * @param ctx the parse tree
	 */
	void enterSum(BricksParser.SumContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#sum}.
	 * @param ctx the parse tree
	 */
	void exitSum(BricksParser.SumContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#subtract}.
	 * @param ctx the parse tree
	 */
	void enterSubtract(BricksParser.SubtractContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#subtract}.
	 * @param ctx the parse tree
	 */
	void exitSubtract(BricksParser.SubtractContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#max}.
	 * @param ctx the parse tree
	 */
	void enterMax(BricksParser.MaxContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#max}.
	 * @param ctx the parse tree
	 */
	void exitMax(BricksParser.MaxContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#sort}.
	 * @param ctx the parse tree
	 */
	void enterSort(BricksParser.SortContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#sort}.
	 * @param ctx the parse tree
	 */
	void exitSort(BricksParser.SortContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#build}.
	 * @param ctx the parse tree
	 */
	void enterBuild(BricksParser.BuildContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#build}.
	 * @param ctx the parse tree
	 */
	void exitBuild(BricksParser.BuildContext ctx);
	/**
	 * Enter a parse tree produced by {@link BricksParser#identifier}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier(BricksParser.IdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link BricksParser#identifier}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier(BricksParser.IdentifierContext ctx);
}