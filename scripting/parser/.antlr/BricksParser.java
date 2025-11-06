// Generated from p:/Entwicklung/bricks-cli/scripting/parser/Bricks.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class BricksParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		ID=25, INT=26, SET_NUM=27, BOOL=28, STRING=29, LINE_COMMENT=30, WS=31;
	public static final int
		RULE_bricks = 0, RULE_command = 1, RULE_expression = 2, RULE_assignment = 3, 
		RULE_save = 4, RULE_export = 5, RULE_print = 6, RULE_pause = 7, RULE_collectionOrId = 8, 
		RULE_collection = 9, RULE_load = 10, RULE_import_ = 11, RULE_allParts = 12, 
		RULE_lost = 13, RULE_set = 14, RULE_userSet = 15, RULE_setList = 16, RULE_partList = 17, 
		RULE_partLists = 18, RULE_sum = 19, RULE_subtract = 20, RULE_max = 21, 
		RULE_sort = 22, RULE_build = 23, RULE_identifier = 24;
	private static String[] makeRuleNames() {
		return new String[] {
			"bricks", "command", "expression", "assignment", "save", "export", "print", 
			"pause", "collectionOrId", "collection", "load", "import_", "allParts", 
			"lost", "set", "userSet", "setList", "partList", "partLists", "sum", 
			"subtract", "max", "sort", "build", "identifier"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':='", "'save'", "'('", "','", "')'", "'export'", "'print'", "'pause'", 
			"'load'", "'import'", "'allParts'", "'lost'", "'set'", "'userSet'", "'setList'", 
			"'partList'", "'partLists'", "'sum'", "'subtract'", "'max'", "'sort'", 
			"'quantity'", "'descending'", "'build'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, "ID", "INT", "SET_NUM", "BOOL", "STRING", "LINE_COMMENT", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Bricks.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public BricksParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BricksContext extends ParserRuleContext {
		public List<CommandContext> command() {
			return getRuleContexts(CommandContext.class);
		}
		public CommandContext command(int i) {
			return getRuleContext(CommandContext.class,i);
		}
		public BricksContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bricks; }
	}

	public final BricksContext bricks() throws RecognitionException {
		BricksContext _localctx = new BricksContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_bricks);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(51); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(50);
				command();
				}
				}
				setState(53); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 33554884L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CommandContext extends ParserRuleContext {
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public SaveContext save() {
			return getRuleContext(SaveContext.class,0);
		}
		public ExportContext export() {
			return getRuleContext(ExportContext.class,0);
		}
		public PrintContext print() {
			return getRuleContext(PrintContext.class,0);
		}
		public PauseContext pause() {
			return getRuleContext(PauseContext.class,0);
		}
		public CommandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_command; }
	}

	public final CommandContext command() throws RecognitionException {
		CommandContext _localctx = new CommandContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_command);
		try {
			setState(60);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				setState(55);
				assignment();
				}
				break;
			case T__1:
				enterOuterAlt(_localctx, 2);
				{
				setState(56);
				save();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 3);
				{
				setState(57);
				export();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 4);
				{
				setState(58);
				print();
				}
				break;
			case T__7:
				enterOuterAlt(_localctx, 5);
				{
				setState(59);
				pause();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public CollectionContext collection() {
			return getRuleContext(CollectionContext.class,0);
		}
		public BuildContext build() {
			return getRuleContext(BuildContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_expression);
		try {
			setState(65);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__8:
			case T__9:
			case T__10:
			case T__11:
			case T__12:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
			case T__17:
			case T__18:
			case T__19:
			case T__20:
				enterOuterAlt(_localctx, 1);
				{
				setState(62);
				collection();
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 2);
				{
				setState(63);
				build();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(64);
				identifier();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(BricksParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(67);
			match(ID);
			setState(68);
			match(T__0);
			setState(69);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SaveContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode STRING() { return getToken(BricksParser.STRING, 0); }
		public SaveContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_save; }
	}

	public final SaveContext save() throws RecognitionException {
		SaveContext _localctx = new SaveContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_save);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			match(T__1);
			setState(72);
			match(T__2);
			setState(73);
			expression();
			setState(74);
			match(T__3);
			setState(75);
			match(STRING);
			setState(76);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExportContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode STRING() { return getToken(BricksParser.STRING, 0); }
		public ExportContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_export; }
	}

	public final ExportContext export() throws RecognitionException {
		ExportContext _localctx = new ExportContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_export);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			match(T__5);
			setState(79);
			match(T__2);
			setState(80);
			expression();
			setState(81);
			match(T__3);
			setState(82);
			match(STRING);
			setState(83);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PrintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print; }
	}

	public final PrintContext print() throws RecognitionException {
		PrintContext _localctx = new PrintContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_print);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(85);
			match(T__6);
			setState(86);
			match(T__2);
			setState(87);
			expression();
			setState(88);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PauseContext extends ParserRuleContext {
		public Token seconds;
		public TerminalNode INT() { return getToken(BricksParser.INT, 0); }
		public PauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pause; }
	}

	public final PauseContext pause() throws RecognitionException {
		PauseContext _localctx = new PauseContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_pause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			match(T__7);
			setState(91);
			match(T__2);
			setState(92);
			((PauseContext)_localctx).seconds = match(INT);
			setState(93);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CollectionOrIdContext extends ParserRuleContext {
		public CollectionContext collection() {
			return getRuleContext(CollectionContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public CollectionOrIdContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_collectionOrId; }
	}

	public final CollectionOrIdContext collectionOrId() throws RecognitionException {
		CollectionOrIdContext _localctx = new CollectionOrIdContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_collectionOrId);
		try {
			setState(97);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__8:
			case T__9:
			case T__10:
			case T__11:
			case T__12:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
			case T__17:
			case T__18:
			case T__19:
			case T__20:
				enterOuterAlt(_localctx, 1);
				{
				setState(95);
				collection();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				setState(96);
				identifier();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CollectionContext extends ParserRuleContext {
		public LoadContext load() {
			return getRuleContext(LoadContext.class,0);
		}
		public Import_Context import_() {
			return getRuleContext(Import_Context.class,0);
		}
		public AllPartsContext allParts() {
			return getRuleContext(AllPartsContext.class,0);
		}
		public LostContext lost() {
			return getRuleContext(LostContext.class,0);
		}
		public SetContext set() {
			return getRuleContext(SetContext.class,0);
		}
		public UserSetContext userSet() {
			return getRuleContext(UserSetContext.class,0);
		}
		public SetListContext setList() {
			return getRuleContext(SetListContext.class,0);
		}
		public PartListContext partList() {
			return getRuleContext(PartListContext.class,0);
		}
		public PartListsContext partLists() {
			return getRuleContext(PartListsContext.class,0);
		}
		public SumContext sum() {
			return getRuleContext(SumContext.class,0);
		}
		public SubtractContext subtract() {
			return getRuleContext(SubtractContext.class,0);
		}
		public MaxContext max() {
			return getRuleContext(MaxContext.class,0);
		}
		public SortContext sort() {
			return getRuleContext(SortContext.class,0);
		}
		public CollectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_collection; }
	}

	public final CollectionContext collection() throws RecognitionException {
		CollectionContext _localctx = new CollectionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_collection);
		try {
			setState(112);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__8:
				enterOuterAlt(_localctx, 1);
				{
				setState(99);
				load();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 2);
				{
				setState(100);
				import_();
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 3);
				{
				setState(101);
				allParts();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 4);
				{
				setState(102);
				lost();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 5);
				{
				setState(103);
				set();
				}
				break;
			case T__13:
				enterOuterAlt(_localctx, 6);
				{
				setState(104);
				userSet();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 7);
				{
				setState(105);
				setList();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 8);
				{
				setState(106);
				partList();
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 9);
				{
				setState(107);
				partLists();
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 10);
				{
				setState(108);
				sum();
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 11);
				{
				setState(109);
				subtract();
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 12);
				{
				setState(110);
				max();
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 13);
				{
				setState(111);
				sort();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LoadContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(BricksParser.STRING, 0); }
		public LoadContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_load; }
	}

	public final LoadContext load() throws RecognitionException {
		LoadContext _localctx = new LoadContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_load);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			match(T__8);
			setState(115);
			match(T__2);
			setState(116);
			match(STRING);
			setState(117);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Import_Context extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(BricksParser.STRING, 0); }
		public Import_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_; }
	}

	public final Import_Context import_() throws RecognitionException {
		Import_Context _localctx = new Import_Context(_ctx, getState());
		enterRule(_localctx, 22, RULE_import_);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			match(T__9);
			setState(120);
			match(T__2);
			setState(121);
			match(STRING);
			setState(122);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AllPartsContext extends ParserRuleContext {
		public AllPartsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_allParts; }
	}

	public final AllPartsContext allParts() throws RecognitionException {
		AllPartsContext _localctx = new AllPartsContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_allParts);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			match(T__10);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LostContext extends ParserRuleContext {
		public LostContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lost; }
	}

	public final LostContext lost() throws RecognitionException {
		LostContext _localctx = new LostContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_lost);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(126);
			match(T__11);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SetContext extends ParserRuleContext {
		public TerminalNode SET_NUM() { return getToken(BricksParser.SET_NUM, 0); }
		public TerminalNode BOOL() { return getToken(BricksParser.BOOL, 0); }
		public SetContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_set; }
	}

	public final SetContext set() throws RecognitionException {
		SetContext _localctx = new SetContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_set);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			match(T__12);
			setState(129);
			match(T__2);
			setState(130);
			match(SET_NUM);
			setState(133);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(131);
				match(T__3);
				setState(132);
				match(BOOL);
				}
			}

			setState(135);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UserSetContext extends ParserRuleContext {
		public TerminalNode SET_NUM() { return getToken(BricksParser.SET_NUM, 0); }
		public TerminalNode BOOL() { return getToken(BricksParser.BOOL, 0); }
		public UserSetContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_userSet; }
	}

	public final UserSetContext userSet() throws RecognitionException {
		UserSetContext _localctx = new UserSetContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_userSet);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137);
			match(T__13);
			setState(138);
			match(T__2);
			setState(139);
			match(SET_NUM);
			setState(142);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(140);
				match(T__3);
				setState(141);
				match(BOOL);
				}
			}

			setState(144);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SetListContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(BricksParser.INT, 0); }
		public TerminalNode BOOL() { return getToken(BricksParser.BOOL, 0); }
		public SetListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_setList; }
	}

	public final SetListContext setList() throws RecognitionException {
		SetListContext _localctx = new SetListContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_setList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(146);
			match(T__14);
			setState(147);
			match(T__2);
			setState(148);
			match(INT);
			setState(151);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(149);
				match(T__3);
				setState(150);
				match(BOOL);
				}
			}

			setState(153);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PartListContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(BricksParser.INT, 0); }
		public PartListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_partList; }
	}

	public final PartListContext partList() throws RecognitionException {
		PartListContext _localctx = new PartListContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_partList);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(155);
			match(T__15);
			setState(156);
			match(T__2);
			setState(157);
			match(INT);
			setState(158);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PartListsContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(BricksParser.STRING, 0); }
		public TerminalNode BOOL() { return getToken(BricksParser.BOOL, 0); }
		public PartListsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_partLists; }
	}

	public final PartListsContext partLists() throws RecognitionException {
		PartListsContext _localctx = new PartListsContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_partLists);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			match(T__16);
			setState(161);
			match(T__2);
			setState(162);
			match(STRING);
			setState(165);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(163);
				match(T__3);
				setState(164);
				match(BOOL);
				}
			}

			setState(167);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SumContext extends ParserRuleContext {
		public List<CollectionOrIdContext> collectionOrId() {
			return getRuleContexts(CollectionOrIdContext.class);
		}
		public CollectionOrIdContext collectionOrId(int i) {
			return getRuleContext(CollectionOrIdContext.class,i);
		}
		public SumContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sum; }
	}

	public final SumContext sum() throws RecognitionException {
		SumContext _localctx = new SumContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_sum);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			match(T__17);
			setState(170);
			match(T__2);
			setState(171);
			collectionOrId();
			setState(174); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(172);
				match(T__3);
				setState(173);
				collectionOrId();
				}
				}
				setState(176); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__3 );
			setState(178);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SubtractContext extends ParserRuleContext {
		public List<CollectionOrIdContext> collectionOrId() {
			return getRuleContexts(CollectionOrIdContext.class);
		}
		public CollectionOrIdContext collectionOrId(int i) {
			return getRuleContext(CollectionOrIdContext.class,i);
		}
		public SubtractContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subtract; }
	}

	public final SubtractContext subtract() throws RecognitionException {
		SubtractContext _localctx = new SubtractContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_subtract);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(180);
			match(T__18);
			setState(181);
			match(T__2);
			setState(182);
			collectionOrId();
			setState(183);
			match(T__3);
			setState(184);
			collectionOrId();
			setState(185);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MaxContext extends ParserRuleContext {
		public List<CollectionOrIdContext> collectionOrId() {
			return getRuleContexts(CollectionOrIdContext.class);
		}
		public CollectionOrIdContext collectionOrId(int i) {
			return getRuleContext(CollectionOrIdContext.class,i);
		}
		public MaxContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_max; }
	}

	public final MaxContext max() throws RecognitionException {
		MaxContext _localctx = new MaxContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_max);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(187);
			match(T__19);
			setState(188);
			match(T__2);
			setState(189);
			collectionOrId();
			setState(192); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(190);
				match(T__3);
				setState(191);
				collectionOrId();
				}
				}
				setState(194); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__3 );
			setState(196);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SortContext extends ParserRuleContext {
		public Token quantity;
		public Token descending;
		public CollectionOrIdContext collectionOrId() {
			return getRuleContext(CollectionOrIdContext.class,0);
		}
		public SortContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sort; }
	}

	public final SortContext sort() throws RecognitionException {
		SortContext _localctx = new SortContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_sort);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(198);
			match(T__20);
			setState(199);
			match(T__2);
			setState(200);
			collectionOrId();
			setState(203);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(201);
				match(T__3);
				setState(202);
				((SortContext)_localctx).quantity = match(T__21);
				}
				break;
			}
			setState(207);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(205);
				match(T__3);
				setState(206);
				((SortContext)_localctx).descending = match(T__22);
				}
			}

			setState(209);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BuildContext extends ParserRuleContext {
		public Token build_mode;
		public List<CollectionOrIdContext> collectionOrId() {
			return getRuleContexts(CollectionOrIdContext.class);
		}
		public CollectionOrIdContext collectionOrId(int i) {
			return getRuleContext(CollectionOrIdContext.class,i);
		}
		public TerminalNode ID() { return getToken(BricksParser.ID, 0); }
		public BuildContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_build; }
	}

	public final BuildContext build() throws RecognitionException {
		BuildContext _localctx = new BuildContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_build);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(211);
			match(T__23);
			setState(212);
			match(T__2);
			setState(213);
			collectionOrId();
			setState(214);
			match(T__3);
			setState(215);
			collectionOrId();
			setState(218);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(216);
				match(T__3);
				setState(217);
				((BuildContext)_localctx).build_mode = match(ID);
				}
			}

			setState(220);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(BricksParser.ID, 0); }
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_identifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(222);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u001f\u00e1\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001"+
		"\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004"+
		"\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007"+
		"\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b"+
		"\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007"+
		"\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007"+
		"\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007"+
		"\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007"+
		"\u0018\u0001\u0000\u0004\u00004\b\u0000\u000b\u0000\f\u00005\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001=\b\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002B\b\u0002\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0003\bb\b\b\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0003\tq\b\t\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0003\u000e\u0086\b\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u008f"+
		"\b\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0003\u0010\u0098\b\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u00a6\b\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0004\u0013\u00af\b\u0013\u000b\u0013\f\u0013\u00b0\u0001\u0013"+
		"\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0004\u0015\u00c1\b\u0015\u000b\u0015\f\u0015\u00c2\u0001"+
		"\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0003\u0016\u00cc\b\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u00d0"+
		"\b\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u00db\b\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0000\u0000\u0019"+
		"\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a"+
		"\u001c\u001e \"$&(*,.0\u0000\u0000\u00e4\u00003\u0001\u0000\u0000\u0000"+
		"\u0002<\u0001\u0000\u0000\u0000\u0004A\u0001\u0000\u0000\u0000\u0006C"+
		"\u0001\u0000\u0000\u0000\bG\u0001\u0000\u0000\u0000\nN\u0001\u0000\u0000"+
		"\u0000\fU\u0001\u0000\u0000\u0000\u000eZ\u0001\u0000\u0000\u0000\u0010"+
		"a\u0001\u0000\u0000\u0000\u0012p\u0001\u0000\u0000\u0000\u0014r\u0001"+
		"\u0000\u0000\u0000\u0016w\u0001\u0000\u0000\u0000\u0018|\u0001\u0000\u0000"+
		"\u0000\u001a~\u0001\u0000\u0000\u0000\u001c\u0080\u0001\u0000\u0000\u0000"+
		"\u001e\u0089\u0001\u0000\u0000\u0000 \u0092\u0001\u0000\u0000\u0000\""+
		"\u009b\u0001\u0000\u0000\u0000$\u00a0\u0001\u0000\u0000\u0000&\u00a9\u0001"+
		"\u0000\u0000\u0000(\u00b4\u0001\u0000\u0000\u0000*\u00bb\u0001\u0000\u0000"+
		"\u0000,\u00c6\u0001\u0000\u0000\u0000.\u00d3\u0001\u0000\u0000\u00000"+
		"\u00de\u0001\u0000\u0000\u000024\u0003\u0002\u0001\u000032\u0001\u0000"+
		"\u0000\u000045\u0001\u0000\u0000\u000053\u0001\u0000\u0000\u000056\u0001"+
		"\u0000\u0000\u00006\u0001\u0001\u0000\u0000\u00007=\u0003\u0006\u0003"+
		"\u00008=\u0003\b\u0004\u00009=\u0003\n\u0005\u0000:=\u0003\f\u0006\u0000"+
		";=\u0003\u000e\u0007\u0000<7\u0001\u0000\u0000\u0000<8\u0001\u0000\u0000"+
		"\u0000<9\u0001\u0000\u0000\u0000<:\u0001\u0000\u0000\u0000<;\u0001\u0000"+
		"\u0000\u0000=\u0003\u0001\u0000\u0000\u0000>B\u0003\u0012\t\u0000?B\u0003"+
		".\u0017\u0000@B\u00030\u0018\u0000A>\u0001\u0000\u0000\u0000A?\u0001\u0000"+
		"\u0000\u0000A@\u0001\u0000\u0000\u0000B\u0005\u0001\u0000\u0000\u0000"+
		"CD\u0005\u0019\u0000\u0000DE\u0005\u0001\u0000\u0000EF\u0003\u0004\u0002"+
		"\u0000F\u0007\u0001\u0000\u0000\u0000GH\u0005\u0002\u0000\u0000HI\u0005"+
		"\u0003\u0000\u0000IJ\u0003\u0004\u0002\u0000JK\u0005\u0004\u0000\u0000"+
		"KL\u0005\u001d\u0000\u0000LM\u0005\u0005\u0000\u0000M\t\u0001\u0000\u0000"+
		"\u0000NO\u0005\u0006\u0000\u0000OP\u0005\u0003\u0000\u0000PQ\u0003\u0004"+
		"\u0002\u0000QR\u0005\u0004\u0000\u0000RS\u0005\u001d\u0000\u0000ST\u0005"+
		"\u0005\u0000\u0000T\u000b\u0001\u0000\u0000\u0000UV\u0005\u0007\u0000"+
		"\u0000VW\u0005\u0003\u0000\u0000WX\u0003\u0004\u0002\u0000XY\u0005\u0005"+
		"\u0000\u0000Y\r\u0001\u0000\u0000\u0000Z[\u0005\b\u0000\u0000[\\\u0005"+
		"\u0003\u0000\u0000\\]\u0005\u001a\u0000\u0000]^\u0005\u0005\u0000\u0000"+
		"^\u000f\u0001\u0000\u0000\u0000_b\u0003\u0012\t\u0000`b\u00030\u0018\u0000"+
		"a_\u0001\u0000\u0000\u0000a`\u0001\u0000\u0000\u0000b\u0011\u0001\u0000"+
		"\u0000\u0000cq\u0003\u0014\n\u0000dq\u0003\u0016\u000b\u0000eq\u0003\u0018"+
		"\f\u0000fq\u0003\u001a\r\u0000gq\u0003\u001c\u000e\u0000hq\u0003\u001e"+
		"\u000f\u0000iq\u0003 \u0010\u0000jq\u0003\"\u0011\u0000kq\u0003$\u0012"+
		"\u0000lq\u0003&\u0013\u0000mq\u0003(\u0014\u0000nq\u0003*\u0015\u0000"+
		"oq\u0003,\u0016\u0000pc\u0001\u0000\u0000\u0000pd\u0001\u0000\u0000\u0000"+
		"pe\u0001\u0000\u0000\u0000pf\u0001\u0000\u0000\u0000pg\u0001\u0000\u0000"+
		"\u0000ph\u0001\u0000\u0000\u0000pi\u0001\u0000\u0000\u0000pj\u0001\u0000"+
		"\u0000\u0000pk\u0001\u0000\u0000\u0000pl\u0001\u0000\u0000\u0000pm\u0001"+
		"\u0000\u0000\u0000pn\u0001\u0000\u0000\u0000po\u0001\u0000\u0000\u0000"+
		"q\u0013\u0001\u0000\u0000\u0000rs\u0005\t\u0000\u0000st\u0005\u0003\u0000"+
		"\u0000tu\u0005\u001d\u0000\u0000uv\u0005\u0005\u0000\u0000v\u0015\u0001"+
		"\u0000\u0000\u0000wx\u0005\n\u0000\u0000xy\u0005\u0003\u0000\u0000yz\u0005"+
		"\u001d\u0000\u0000z{\u0005\u0005\u0000\u0000{\u0017\u0001\u0000\u0000"+
		"\u0000|}\u0005\u000b\u0000\u0000}\u0019\u0001\u0000\u0000\u0000~\u007f"+
		"\u0005\f\u0000\u0000\u007f\u001b\u0001\u0000\u0000\u0000\u0080\u0081\u0005"+
		"\r\u0000\u0000\u0081\u0082\u0005\u0003\u0000\u0000\u0082\u0085\u0005\u001b"+
		"\u0000\u0000\u0083\u0084\u0005\u0004\u0000\u0000\u0084\u0086\u0005\u001c"+
		"\u0000\u0000\u0085\u0083\u0001\u0000\u0000\u0000\u0085\u0086\u0001\u0000"+
		"\u0000\u0000\u0086\u0087\u0001\u0000\u0000\u0000\u0087\u0088\u0005\u0005"+
		"\u0000\u0000\u0088\u001d\u0001\u0000\u0000\u0000\u0089\u008a\u0005\u000e"+
		"\u0000\u0000\u008a\u008b\u0005\u0003\u0000\u0000\u008b\u008e\u0005\u001b"+
		"\u0000\u0000\u008c\u008d\u0005\u0004\u0000\u0000\u008d\u008f\u0005\u001c"+
		"\u0000\u0000\u008e\u008c\u0001\u0000\u0000\u0000\u008e\u008f\u0001\u0000"+
		"\u0000\u0000\u008f\u0090\u0001\u0000\u0000\u0000\u0090\u0091\u0005\u0005"+
		"\u0000\u0000\u0091\u001f\u0001\u0000\u0000\u0000\u0092\u0093\u0005\u000f"+
		"\u0000\u0000\u0093\u0094\u0005\u0003\u0000\u0000\u0094\u0097\u0005\u001a"+
		"\u0000\u0000\u0095\u0096\u0005\u0004\u0000\u0000\u0096\u0098\u0005\u001c"+
		"\u0000\u0000\u0097\u0095\u0001\u0000\u0000\u0000\u0097\u0098\u0001\u0000"+
		"\u0000\u0000\u0098\u0099\u0001\u0000\u0000\u0000\u0099\u009a\u0005\u0005"+
		"\u0000\u0000\u009a!\u0001\u0000\u0000\u0000\u009b\u009c\u0005\u0010\u0000"+
		"\u0000\u009c\u009d\u0005\u0003\u0000\u0000\u009d\u009e\u0005\u001a\u0000"+
		"\u0000\u009e\u009f\u0005\u0005\u0000\u0000\u009f#\u0001\u0000\u0000\u0000"+
		"\u00a0\u00a1\u0005\u0011\u0000\u0000\u00a1\u00a2\u0005\u0003\u0000\u0000"+
		"\u00a2\u00a5\u0005\u001d\u0000\u0000\u00a3\u00a4\u0005\u0004\u0000\u0000"+
		"\u00a4\u00a6\u0005\u001c\u0000\u0000\u00a5\u00a3\u0001\u0000\u0000\u0000"+
		"\u00a5\u00a6\u0001\u0000\u0000\u0000\u00a6\u00a7\u0001\u0000\u0000\u0000"+
		"\u00a7\u00a8\u0005\u0005\u0000\u0000\u00a8%\u0001\u0000\u0000\u0000\u00a9"+
		"\u00aa\u0005\u0012\u0000\u0000\u00aa\u00ab\u0005\u0003\u0000\u0000\u00ab"+
		"\u00ae\u0003\u0010\b\u0000\u00ac\u00ad\u0005\u0004\u0000\u0000\u00ad\u00af"+
		"\u0003\u0010\b\u0000\u00ae\u00ac\u0001\u0000\u0000\u0000\u00af\u00b0\u0001"+
		"\u0000\u0000\u0000\u00b0\u00ae\u0001\u0000\u0000\u0000\u00b0\u00b1\u0001"+
		"\u0000\u0000\u0000\u00b1\u00b2\u0001\u0000\u0000\u0000\u00b2\u00b3\u0005"+
		"\u0005\u0000\u0000\u00b3\'\u0001\u0000\u0000\u0000\u00b4\u00b5\u0005\u0013"+
		"\u0000\u0000\u00b5\u00b6\u0005\u0003\u0000\u0000\u00b6\u00b7\u0003\u0010"+
		"\b\u0000\u00b7\u00b8\u0005\u0004\u0000\u0000\u00b8\u00b9\u0003\u0010\b"+
		"\u0000\u00b9\u00ba\u0005\u0005\u0000\u0000\u00ba)\u0001\u0000\u0000\u0000"+
		"\u00bb\u00bc\u0005\u0014\u0000\u0000\u00bc\u00bd\u0005\u0003\u0000\u0000"+
		"\u00bd\u00c0\u0003\u0010\b\u0000\u00be\u00bf\u0005\u0004\u0000\u0000\u00bf"+
		"\u00c1\u0003\u0010\b\u0000\u00c0\u00be\u0001\u0000\u0000\u0000\u00c1\u00c2"+
		"\u0001\u0000\u0000\u0000\u00c2\u00c0\u0001\u0000\u0000\u0000\u00c2\u00c3"+
		"\u0001\u0000\u0000\u0000\u00c3\u00c4\u0001\u0000\u0000\u0000\u00c4\u00c5"+
		"\u0005\u0005\u0000\u0000\u00c5+\u0001\u0000\u0000\u0000\u00c6\u00c7\u0005"+
		"\u0015\u0000\u0000\u00c7\u00c8\u0005\u0003\u0000\u0000\u00c8\u00cb\u0003"+
		"\u0010\b\u0000\u00c9\u00ca\u0005\u0004\u0000\u0000\u00ca\u00cc\u0005\u0016"+
		"\u0000\u0000\u00cb\u00c9\u0001\u0000\u0000\u0000\u00cb\u00cc\u0001\u0000"+
		"\u0000\u0000\u00cc\u00cf\u0001\u0000\u0000\u0000\u00cd\u00ce\u0005\u0004"+
		"\u0000\u0000\u00ce\u00d0\u0005\u0017\u0000\u0000\u00cf\u00cd\u0001\u0000"+
		"\u0000\u0000\u00cf\u00d0\u0001\u0000\u0000\u0000\u00d0\u00d1\u0001\u0000"+
		"\u0000\u0000\u00d1\u00d2\u0005\u0005\u0000\u0000\u00d2-\u0001\u0000\u0000"+
		"\u0000\u00d3\u00d4\u0005\u0018\u0000\u0000\u00d4\u00d5\u0005\u0003\u0000"+
		"\u0000\u00d5\u00d6\u0003\u0010\b\u0000\u00d6\u00d7\u0005\u0004\u0000\u0000"+
		"\u00d7\u00da\u0003\u0010\b\u0000\u00d8\u00d9\u0005\u0004\u0000\u0000\u00d9"+
		"\u00db\u0005\u0019\u0000\u0000\u00da\u00d8\u0001\u0000\u0000\u0000\u00da"+
		"\u00db\u0001\u0000\u0000\u0000\u00db\u00dc\u0001\u0000\u0000\u0000\u00dc"+
		"\u00dd\u0005\u0005\u0000\u0000\u00dd/\u0001\u0000\u0000\u0000\u00de\u00df"+
		"\u0005\u0019\u0000\u0000\u00df1\u0001\u0000\u0000\u0000\u000e5<Aap\u0085"+
		"\u008e\u0097\u00a5\u00b0\u00c2\u00cb\u00cf\u00da";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}