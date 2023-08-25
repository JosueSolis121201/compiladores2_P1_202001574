// Generated from c:\Users\josue\Desktop\compiladores2_P1_202001574\Server2\SwiftGrammar.g4 by ANTLR 4.9.2

    import "Server2/interfaces"
    import "Server2/environment"
    import "Server2/expressions"
    import "Server2/instructions"
    import "strings"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, BOOL=3, STR=4, CHAR=5, TRU=6, FAL=7, PRINT=8, IF=9, ELSE=10, 
		ELSE_IF=11, WHILE=12, FOR=13, IN=14, GUARD=15, CONTINUE=16, RETURN=17, 
		SWITCH=18, BREAK=19, CASE=20, VAR=21, LET=22, D_APPEND=23, D_REMOVE=24, 
		REPEATING=25, DEFAULT=26, COUNT=27, STRUCT=28, MUTATING=29, FUNC=30, SELF_D=31, 
		PUNTO=32, INOUT=33, GUION_B=34, NUMBER=35, STRING=36, ID=37, DIF=38, IG_IG=39, 
		MAS_IGUAL=40, MENOS_IGUAL=41, NOT=42, OR=43, AND=44, SOLO_AND=45, IG=46, 
		MAY_IG=47, MEN_IG=48, MAYOR=49, MENOR=50, MUL=51, DIV=52, ADD=53, MODULO=54, 
		SUB=55, PARIZQ=56, PARDER=57, LLAVEIZQ=58, LLAVEDER=59, D_PTS=60, CORIZQ=61, 
		CORDER=62, COMA=63, PUNTO_COMA=64, WHITESPACE=65, COMMENT=66, LINE_COMMENT=67;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_guardstmt = 3, 
		RULE_declaracion_vectorstmt = 4, RULE_listaval = 5, RULE_jumpstmt = 6, 
		RULE_whilestmt = 7, RULE_forstmt = 8, RULE_printstmt = 9, RULE_ifstmt = 10, 
		RULE_elsestmt = 11, RULE_else_ifstmt = 12, RULE_declarationstmt = 13, 
		RULE_switchstmt = 14, RULE_defaultsts = 15, RULE_case = 16, RULE_asignacionstmt = 17, 
		RULE_ig_type = 18, RULE_reserv = 19, RULE_types = 20, RULE_expr = 21, 
		RULE_listParams = 22, RULE_listArray = 23;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "guardstmt", "declaracion_vectorstmt", "listaval", 
			"jumpstmt", "whilestmt", "forstmt", "printstmt", "ifstmt", "elsestmt", 
			"else_ifstmt", "declarationstmt", "switchstmt", "defaultsts", "case", 
			"asignacionstmt", "ig_type", "reserv", "types", "expr", "listParams", 
			"listArray"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'", 
			"'false'", "'print'", "'if'", "'else'", "'else if'", "'while'", "'for'", 
			"'in'", "'guard'", "'continue'", "'return'", "'switch'", "'break'", "'case'", 
			"'var'", "'let'", "'.append'", "'.remove'", "'repeating'", "'default'", 
			"'count'", "'STRUCT'", "'mutating'", "'func'", "'self.'", "'.'", "'inout'", 
			"'_'", null, null, null, "'!='", "'=='", "'+='", "'-='", "'!'", "'||'", 
			"'&&'", "'&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", 
			"'%'", "'-'", "'('", "')'", "'{'", "'}'", "':'", "'['", "']'", "','", 
			"';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "CHAR", "TRU", "FAL", "PRINT", "IF", 
			"ELSE", "ELSE_IF", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", 
			"SWITCH", "BREAK", "CASE", "VAR", "LET", "D_APPEND", "D_REMOVE", "REPEATING", 
			"DEFAULT", "COUNT", "STRUCT", "MUTATING", "FUNC", "SELF_D", "PUNTO", 
			"INOUT", "GUION_B", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "MAS_IGUAL", 
			"MENOS_IGUAL", "NOT", "OR", "AND", "SOLO_AND", "IG", "MAY_IG", "MEN_IG", 
			"MAYOR", "MENOR", "MUL", "DIV", "ADD", "MODULO", "SUB", "PARIZQ", "PARDER", 
			"LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", "CORDER", "COMA", "PUNTO_COMA", 
			"WHITESPACE", "COMMENT", "LINE_COMMENT"
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
	public String getGrammarFileName() { return "SwiftGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SwiftGrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class SContext extends ParserRuleContext {
		public []interface{} code;
		public BlockContext block;
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SwiftGrammarParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			((SContext)_localctx).block = block();
			setState(49);
			match(EOF);
			   
			        _localctx.code = ((SContext)_localctx).block.blk
			    
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

	public static class BlockContext extends ParserRuleContext {
		public []interface{} blk;
		public InstructionContext instruction;
		public List<InstructionContext> ins = new ArrayList<InstructionContext>();
		public List<InstructionContext> instruction() {
			return getRuleContexts(InstructionContext.class);
		}
		public InstructionContext instruction(int i) {
			return getRuleContext(InstructionContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);

		    _localctx.blk = []interface{}{}
		    var listInt []IInstructionContext
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(53); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(52);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(55); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );

			        listInt = localctx.(*BlockContext).GetIns()
			        for _, e := range listInt {
			            _localctx.blk = append(_localctx.blk, e.GetInst())
			        }
			    
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

	public static class InstructionContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public PrintstmtContext printstmt;
		public IfstmtContext ifstmt;
		public DeclarationstmtContext declarationstmt;
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public DeclarationstmtContext declarationstmt() {
			return getRuleContext(DeclarationstmtContext.class,0);
		}
		public AsignacionstmtContext asignacionstmt() {
			return getRuleContext(AsignacionstmtContext.class,0);
		}
		public SwitchstmtContext switchstmt() {
			return getRuleContext(SwitchstmtContext.class,0);
		}
		public WhilestmtContext whilestmt() {
			return getRuleContext(WhilestmtContext.class,0);
		}
		public ForstmtContext forstmt() {
			return getRuleContext(ForstmtContext.class,0);
		}
		public GuardstmtContext guardstmt() {
			return getRuleContext(GuardstmtContext.class,0);
		}
		public JumpstmtContext jumpstmt() {
			return getRuleContext(JumpstmtContext.class,0);
		}
		public Declaracion_vectorstmtContext declaracion_vectorstmt() {
			return getRuleContext(Declaracion_vectorstmtContext.class,0);
		}
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(89);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(59);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(62);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinst 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(65);
				((InstructionContext)_localctx).declarationstmt = declarationstmt();
				 _localctx.inst = ((InstructionContext)_localctx).declarationstmt.dec 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(68);
				asignacionstmt();

				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(71);
				switchstmt();

				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(74);
				whilestmt();

				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(77);
				forstmt();

				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(80);
				guardstmt();

				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(83);
				jumpstmt();

				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(86);
				declaracion_vectorstmt();

				}
				break;
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

	public static class GuardstmtContext extends ParserRuleContext {
		public TerminalNode GUARD() { return getToken(SwiftGrammarParser.GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public GuardstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guardstmt; }
	}

	public final GuardstmtContext guardstmt() throws RecognitionException {
		GuardstmtContext _localctx = new GuardstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_guardstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			match(GUARD);
			setState(92);
			expr(0);
			setState(93);
			match(ELSE);
			setState(94);
			match(LLAVEIZQ);
			setState(95);
			block();
			setState(96);
			match(LLAVEDER);
			 
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

	public static class Declaracion_vectorstmtContext extends ParserRuleContext {
		public ReservContext reserv() {
			return getRuleContext(ReservContext.class,0);
		}
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public List<TerminalNode> CORIZQ() { return getTokens(SwiftGrammarParser.CORIZQ); }
		public TerminalNode CORIZQ(int i) {
			return getToken(SwiftGrammarParser.CORIZQ, i);
		}
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public List<TerminalNode> CORDER() { return getTokens(SwiftGrammarParser.CORDER); }
		public TerminalNode CORDER(int i) {
			return getToken(SwiftGrammarParser.CORDER, i);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ListavalContext listaval() {
			return getRuleContext(ListavalContext.class,0);
		}
		public Declaracion_vectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion_vectorstmt; }
	}

	public final Declaracion_vectorstmtContext declaracion_vectorstmt() throws RecognitionException {
		Declaracion_vectorstmtContext _localctx = new Declaracion_vectorstmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_declaracion_vectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(99);
			reserv();
			setState(100);
			match(ID);
			setState(101);
			match(D_PTS);
			setState(102);
			match(CORIZQ);
			setState(103);
			types();
			setState(104);
			match(CORDER);
			setState(105);
			match(IG);
			setState(106);
			match(CORIZQ);
			setState(107);
			listaval(0);
			setState(108);
			match(CORDER);
			 
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

	public static class ListavalContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ListavalContext listaval() {
			return getRuleContext(ListavalContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListavalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaval; }
	}

	public final ListavalContext listaval() throws RecognitionException {
		return listaval(0);
	}

	private ListavalContext listaval(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListavalContext _localctx = new ListavalContext(_ctx, _parentState);
		ListavalContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_listaval, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(112);
				expr(0);

				}
				break;
			case 2:
				{
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(125);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListavalContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_listaval);
					setState(118);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(119);
					match(COMA);
					setState(120);
					expr(0);

					}
					} 
				}
				setState(127);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class JumpstmtContext extends ParserRuleContext {
		public TerminalNode CONTINUE() { return getToken(SwiftGrammarParser.CONTINUE, 0); }
		public TerminalNode RETURN() { return getToken(SwiftGrammarParser.RETURN, 0); }
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public JumpstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_jumpstmt; }
	}

	public final JumpstmtContext jumpstmt() throws RecognitionException {
		JumpstmtContext _localctx = new JumpstmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_jumpstmt);
		try {
			setState(134);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CONTINUE:
				enterOuterAlt(_localctx, 1);
				{
				setState(128);
				match(CONTINUE);

				}
				break;
			case RETURN:
				enterOuterAlt(_localctx, 2);
				{
				setState(130);
				match(RETURN);

				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 3);
				{
				setState(132);
				match(BREAK);

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

	public static class WhilestmtContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(SwiftGrammarParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public WhilestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whilestmt; }
	}

	public final WhilestmtContext whilestmt() throws RecognitionException {
		WhilestmtContext _localctx = new WhilestmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			match(WHILE);
			setState(137);
			expr(0);
			setState(138);
			match(LLAVEIZQ);
			setState(139);
			block();
			setState(140);
			match(LLAVEDER);
			 
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

	public static class ForstmtContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(SwiftGrammarParser.FOR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IN() { return getToken(SwiftGrammarParser.IN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_forstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			match(FOR);
			setState(144);
			match(ID);
			setState(145);
			match(IN);
			setState(146);
			expr(0);
			setState(147);
			match(LLAVEIZQ);
			setState(148);
			block();
			setState(149);
			match(LLAVEDER);
			 
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

	public static class PrintstmtContext extends ParserRuleContext {
		public interfaces.Instruction prnt;
		public Token PRINT;
		public ExprContext expr;
		public TerminalNode PRINT() { return getToken(SwiftGrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(153);
			match(PARIZQ);
			setState(154);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(155);
			match(PARDER);
			 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).expr.e)
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

	public static class IfstmtContext extends ParserRuleContext {
		public interfaces.Instruction ifinst;
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public Else_ifstmtContext else_ifstmt() {
			return getRuleContext(Else_ifstmtContext.class,0);
		}
		public ElsestmtContext elsestmt() {
			return getRuleContext(ElsestmtContext.class,0);
		}
		public IfstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstmt; }
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_ifstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(158);
			match(IF);
			setState(159);
			expr(0);
			setState(160);
			match(LLAVEIZQ);
			setState(161);
			block();
			setState(162);
			match(LLAVEDER);
			setState(163);
			else_ifstmt(0);
			setState(164);
			elsestmt();

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

	public static class ElsestmtContext extends ParserRuleContext {
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ElsestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elsestmt; }
	}

	public final ElsestmtContext elsestmt() throws RecognitionException {
		ElsestmtContext _localctx = new ElsestmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_elsestmt);
		try {
			setState(174);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(167);
				match(ELSE);
				setState(168);
				match(LLAVEIZQ);
				setState(169);
				block();
				setState(170);
				match(LLAVEDER);

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				}
				break;
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

	public static class Else_ifstmtContext extends ParserRuleContext {
		public Else_ifstmtContext else_ifstmt() {
			return getRuleContext(Else_ifstmtContext.class,0);
		}
		public TerminalNode ELSE_IF() { return getToken(SwiftGrammarParser.ELSE_IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public Else_ifstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else_ifstmt; }
	}

	public final Else_ifstmtContext else_ifstmt() throws RecognitionException {
		return else_ifstmt(0);
	}

	private Else_ifstmtContext else_ifstmt(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Else_ifstmtContext _localctx = new Else_ifstmtContext(_ctx, _parentState);
		Else_ifstmtContext _prevctx = _localctx;
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_else_ifstmt, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			}
			_ctx.stop = _input.LT(-1);
			setState(187);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Else_ifstmtContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_else_ifstmt);
					setState(177);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(178);
					match(ELSE_IF);
					setState(179);
					expr(0);
					setState(180);
					match(LLAVEIZQ);
					setState(181);
					block();
					setState(182);
					match(LLAVEDER);

					}
					} 
				}
				setState(189);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class DeclarationstmtContext extends ParserRuleContext {
		public interfaces.Instruction dec;
		public ReservContext reserv() {
			return getRuleContext(ReservContext.class,0);
		}
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclarationstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarationstmt; }
	}

	public final DeclarationstmtContext declarationstmt() throws RecognitionException {
		DeclarationstmtContext _localctx = new DeclarationstmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_declarationstmt);
		try {
			setState(204);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(190);
				reserv();
				setState(191);
				match(ID);
				setState(192);
				match(D_PTS);
				setState(193);
				types();
				setState(194);
				match(IG);
				setState(195);
				expr(0);

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(198);
				reserv();
				setState(199);
				match(ID);
				setState(200);
				match(IG);
				setState(201);
				expr(0);

				}
				break;
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

	public static class SwitchstmtContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(SwiftGrammarParser.SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public CaseContext case() {
			return getRuleContext(CaseContext.class,0);
		}
		public DefaultstsContext defaultsts() {
			return getRuleContext(DefaultstsContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public SwitchstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchstmt; }
	}

	public final SwitchstmtContext switchstmt() throws RecognitionException {
		SwitchstmtContext _localctx = new SwitchstmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_switchstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(206);
			match(SWITCH);
			setState(207);
			expr(0);
			setState(208);
			match(LLAVEIZQ);
			setState(209);
			case(0);
			setState(210);
			defaultsts();
			setState(211);
			match(LLAVEDER);

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

	public static class DefaultstsContext extends ParserRuleContext {
		public TerminalNode DEFAULT() { return getToken(SwiftGrammarParser.DEFAULT, 0); }
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public DefaultstsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defaultsts; }
	}

	public final DefaultstsContext defaultsts() throws RecognitionException {
		DefaultstsContext _localctx = new DefaultstsContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_defaultsts);
		try {
			setState(220);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DEFAULT:
				enterOuterAlt(_localctx, 1);
				{
				setState(214);
				match(DEFAULT);
				setState(215);
				match(D_PTS);
				setState(216);
				block();

				}
				break;
			case LLAVEDER:
				enterOuterAlt(_localctx, 2);
				{
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

	public static class CaseContext extends ParserRuleContext {
		public CaseContext case() {
			return getRuleContext(CaseContext.class,0);
		}
		public TerminalNode CASE() { return getToken(SwiftGrammarParser.CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public CaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_case; }
	}

	public final CaseContext case() throws RecognitionException {
		return case(0);
	}

	private CaseContext case(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		CaseContext _localctx = new CaseContext(_ctx, _parentState);
		CaseContext _prevctx = _localctx;
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_case, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			}
			_ctx.stop = _input.LT(-1);
			setState(232);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new CaseContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_case);
					setState(223);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(224);
					match(CASE);
					setState(225);
					expr(0);
					setState(226);
					match(D_PTS);
					setState(227);
					block();

					}
					} 
				}
				setState(234);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AsignacionstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public Ig_typeContext ig_type() {
			return getRuleContext(Ig_typeContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AsignacionstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacionstmt; }
	}

	public final AsignacionstmtContext asignacionstmt() throws RecognitionException {
		AsignacionstmtContext _localctx = new AsignacionstmtContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_asignacionstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			match(ID);
			setState(236);
			ig_type();
			setState(237);
			expr(0);

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

	public static class Ig_typeContext extends ParserRuleContext {
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public TerminalNode IG_IG() { return getToken(SwiftGrammarParser.IG_IG, 0); }
		public TerminalNode MAS_IGUAL() { return getToken(SwiftGrammarParser.MAS_IGUAL, 0); }
		public TerminalNode MENOS_IGUAL() { return getToken(SwiftGrammarParser.MENOS_IGUAL, 0); }
		public Ig_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ig_type; }
	}

	public final Ig_typeContext ig_type() throws RecognitionException {
		Ig_typeContext _localctx = new Ig_typeContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_ig_type);
		try {
			setState(248);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IG:
				enterOuterAlt(_localctx, 1);
				{
				setState(240);
				match(IG);

				}
				break;
			case IG_IG:
				enterOuterAlt(_localctx, 2);
				{
				setState(242);
				match(IG_IG);

				}
				break;
			case MAS_IGUAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(244);
				match(MAS_IGUAL);

				}
				break;
			case MENOS_IGUAL:
				enterOuterAlt(_localctx, 4);
				{
				setState(246);
				match(MENOS_IGUAL);

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

	public static class ReservContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public ReservContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_reserv; }
	}

	public final ReservContext reserv() throws RecognitionException {
		ReservContext _localctx = new ReservContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_reserv);
		try {
			setState(254);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(250);
				match(VAR);

				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(252);
				match(LET);

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

	public static class TypesContext extends ParserRuleContext {
		public environment.TipoExpresion ty;
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode STR() { return getToken(SwiftGrammarParser.STR, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CHAR() { return getToken(SwiftGrammarParser.CHAR, 0); }
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public TypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types; }
	}

	public final TypesContext types() throws RecognitionException {
		TypesContext _localctx = new TypesContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_types);
		try {
			setState(269);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(256);
				match(INT);
				 _localctx.ty = environment.INTEGER 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(258);
				match(FLOAT);
				 _localctx.ty = environment.FLOAT 
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(260);
				match(STR);
				 _localctx.ty = environment.STRING 
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(262);
				match(BOOL);
				 _localctx.ty = environment.BOOLEAN 
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 5);
				{
				setState(264);
				match(CHAR);

				}
				break;
			case CORIZQ:
				enterOuterAlt(_localctx, 6);
				{
				setState(266);
				match(CORIZQ);
				setState(267);
				match(CORDER);
				 _localctx.ty = environment.ARRAY 
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

	public static class ExprContext extends ParserRuleContext {
		public interfaces.Expression e;
		public ExprContext left;
		public ListArrayContext list;
		public Token CORIZQ;
		public ListParamsContext listParams;
		public Token NUMBER;
		public Token STRING;
		public Token TRU;
		public Token FAL;
		public Token op;
		public ExprContext right;
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public TerminalNode NUMBER() { return getToken(SwiftGrammarParser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public TerminalNode TRU() { return getToken(SwiftGrammarParser.TRU, 0); }
		public TerminalNode FAL() { return getToken(SwiftGrammarParser.FAL, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode MUL() { return getToken(SwiftGrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(SwiftGrammarParser.DIV, 0); }
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public TerminalNode MAY_IG() { return getToken(SwiftGrammarParser.MAY_IG, 0); }
		public TerminalNode MAYOR() { return getToken(SwiftGrammarParser.MAYOR, 0); }
		public TerminalNode MEN_IG() { return getToken(SwiftGrammarParser.MEN_IG, 0); }
		public TerminalNode MENOR() { return getToken(SwiftGrammarParser.MENOR, 0); }
		public TerminalNode IG_IG() { return getToken(SwiftGrammarParser.IG_IG, 0); }
		public TerminalNode DIF() { return getToken(SwiftGrammarParser.DIF, 0); }
		public TerminalNode AND() { return getToken(SwiftGrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(SwiftGrammarParser.OR, 0); }
		public TerminalNode MAS_IGUAL() { return getToken(SwiftGrammarParser.MAS_IGUAL, 0); }
		public TerminalNode MENOS_IGUAL() { return getToken(SwiftGrammarParser.MENOS_IGUAL, 0); }
		public TerminalNode MODULO() { return getToken(SwiftGrammarParser.MODULO, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 42;
		enterRecursionRule(_localctx, 42, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(288);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				{
				setState(272);
				((ExprContext)_localctx).list = listArray(0);
				 _localctx.e = ((ExprContext)_localctx).list.p
				}
				break;
			case CORIZQ:
				{
				setState(275);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(276);
				((ExprContext)_localctx).listParams = listParams(0);
				setState(277);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), ((ExprContext)_localctx).listParams.l) 
				}
				break;
			case NUMBER:
				{
				setState(280);
				((ExprContext)_localctx).NUMBER = match(NUMBER);

				        if (strings.Contains((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null),".")){
				            num,err := strconv.ParseFloat((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null), 64);
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.FLOAT)
				        }else{
				            num,err := strconv.Atoi((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null))
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.INTEGER)
				        }
				    
				}
				break;
			case STRING:
				{
				setState(282);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case TRU:
				{
				setState(284);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case FAL:
				{
				setState(286);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(347);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(345);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(290);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(291);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MUL || _la==DIV) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(292);
						((ExprContext)_localctx).right = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(295);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(296);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(297);
						((ExprContext)_localctx).right = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(300);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(301);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MAY_IG || _la==MAYOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(302);
						((ExprContext)_localctx).right = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(305);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(306);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MEN_IG || _la==MENOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(307);
						((ExprContext)_localctx).right = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(310);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(311);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIF || _la==IG_IG) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(312);
						((ExprContext)_localctx).right = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(315);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(316);
						((ExprContext)_localctx).op = match(AND);
						setState(317);
						((ExprContext)_localctx).right = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(320);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(321);
						((ExprContext)_localctx).op = match(OR);
						setState(322);
						((ExprContext)_localctx).right = expr(12);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 8:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(325);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(326);
						((ExprContext)_localctx).op = match(MAS_IGUAL);
						setState(327);
						((ExprContext)_localctx).right = expr(11);

						}
						break;
					case 9:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(330);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(331);
						((ExprContext)_localctx).op = match(MENOS_IGUAL);
						setState(332);
						((ExprContext)_localctx).right = expr(10);

						}
						break;
					case 10:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(335);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(336);
						((ExprContext)_localctx).op = match(MODULO);
						setState(337);
						((ExprContext)_localctx).right = expr(9);

						}
						break;
					case 11:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(340);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(341);
						((ExprContext)_localctx).op = match(MODULO);
						setState(342);
						((ExprContext)_localctx).right = expr(8);

						}
						break;
					}
					} 
				}
				setState(349);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ListParamsContext extends ParserRuleContext {
		public []interface{} l;
		public ListParamsContext list;
		public ExprContext expr;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public ListParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParams; }
	}

	public final ListParamsContext listParams() throws RecognitionException {
		return listParams(0);
	}

	private ListParamsContext listParams(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsContext _localctx = new ListParamsContext(_ctx, _parentState);
		ListParamsContext _prevctx = _localctx;
		int _startState = 44;
		enterRecursionRule(_localctx, 44, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(351);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(361);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParams);
					setState(354);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(355);
					match(COMA);
					setState(356);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(363);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ListArrayContext extends ParserRuleContext {
		public interfaces.Expression p;
		public ListArrayContext list;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public ListArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listArray; }
	}

	public final ListArrayContext listArray() throws RecognitionException {
		return listArray(0);
	}

	private ListArrayContext listArray(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListArrayContext _localctx = new ListArrayContext(_ctx, _parentState);
		ListArrayContext _prevctx = _localctx;
		int _startState = 46;
		enterRecursionRule(_localctx, 46, RULE_listArray, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(365);
			((ListArrayContext)_localctx).ID = match(ID);
			 _localctx.p = expressions.NewCallVar((((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getCharPositionInLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(376);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListArrayContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listArray);
					setState(368);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(369);
					match(CORIZQ);
					setState(370);
					((ListArrayContext)_localctx).expr = expr(0);
					setState(371);
					match(CORDER);
					 _localctx.p = expressions.NewArrayAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, ((ListArrayContext)_localctx).expr.e) 
					}
					} 
				}
				setState(378);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 5:
			return listaval_sempred((ListavalContext)_localctx, predIndex);
		case 12:
			return else_ifstmt_sempred((Else_ifstmtContext)_localctx, predIndex);
		case 16:
			return case_sempred((CaseContext)_localctx, predIndex);
		case 21:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 22:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 23:
			return listArray_sempred((ListArrayContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean listaval_sempred(ListavalContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean else_ifstmt_sempred(Else_ifstmtContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean case_sempred(CaseContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 17);
		case 4:
			return precpred(_ctx, 16);
		case 5:
			return precpred(_ctx, 15);
		case 6:
			return precpred(_ctx, 14);
		case 7:
			return precpred(_ctx, 13);
		case 8:
			return precpred(_ctx, 12);
		case 9:
			return precpred(_ctx, 11);
		case 10:
			return precpred(_ctx, 10);
		case 11:
			return precpred(_ctx, 9);
		case 12:
			return precpred(_ctx, 8);
		case 13:
			return precpred(_ctx, 7);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 14:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listArray_sempred(ListArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 15:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3E\u017e\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\3\2\3\2\3\2\3\2\3\3\6\38\n\3\r\3\16\39\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\\\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\5"+
		"\7w\n\7\3\7\3\7\3\7\3\7\3\7\7\7~\n\7\f\7\16\7\u0081\13\7\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\5\b\u0089\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00b1\n\r\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\7\16\u00bc\n\16\f\16\16\16\u00bf\13\16"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\5\17\u00cf\n\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\5\21\u00df\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\7\22\u00e9\n\22\f\22\16\22\u00ec\13\22\3\23\3\23\3\23\3\23\3\23\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u00fb\n\24\3\25\3\25\3\25\3\25"+
		"\5\25\u0101\n\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\5\26\u0110\n\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u0123\n\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\7\27\u015c\n\27\f\27"+
		"\16\27\u015f\13\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\7\30\u016a"+
		"\n\30\f\30\16\30\u016d\13\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\3\31\7\31\u0179\n\31\f\31\16\31\u017c\13\31\3\31\2\b\f\32\",.\60\32"+
		"\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\2\7\3\2\65\66\4\2"+
		"\67\6799\4\2\61\61\63\63\4\2\62\62\64\64\3\2()\2\u0193\2\62\3\2\2\2\4"+
		"\67\3\2\2\2\6[\3\2\2\2\b]\3\2\2\2\ne\3\2\2\2\fv\3\2\2\2\16\u0088\3\2\2"+
		"\2\20\u008a\3\2\2\2\22\u0091\3\2\2\2\24\u009a\3\2\2\2\26\u00a0\3\2\2\2"+
		"\30\u00b0\3\2\2\2\32\u00b2\3\2\2\2\34\u00ce\3\2\2\2\36\u00d0\3\2\2\2 "+
		"\u00de\3\2\2\2\"\u00e0\3\2\2\2$\u00ed\3\2\2\2&\u00fa\3\2\2\2(\u0100\3"+
		"\2\2\2*\u010f\3\2\2\2,\u0122\3\2\2\2.\u0160\3\2\2\2\60\u016e\3\2\2\2\62"+
		"\63\5\4\3\2\63\64\7\2\2\3\64\65\b\2\1\2\65\3\3\2\2\2\668\5\6\4\2\67\66"+
		"\3\2\2\289\3\2\2\29\67\3\2\2\29:\3\2\2\2:;\3\2\2\2;<\b\3\1\2<\5\3\2\2"+
		"\2=>\5\24\13\2>?\b\4\1\2?\\\3\2\2\2@A\5\26\f\2AB\b\4\1\2B\\\3\2\2\2CD"+
		"\5\34\17\2DE\b\4\1\2E\\\3\2\2\2FG\5$\23\2GH\b\4\1\2H\\\3\2\2\2IJ\5\36"+
		"\20\2JK\b\4\1\2K\\\3\2\2\2LM\5\20\t\2MN\b\4\1\2N\\\3\2\2\2OP\5\22\n\2"+
		"PQ\b\4\1\2Q\\\3\2\2\2RS\5\b\5\2ST\b\4\1\2T\\\3\2\2\2UV\5\16\b\2VW\b\4"+
		"\1\2W\\\3\2\2\2XY\5\n\6\2YZ\b\4\1\2Z\\\3\2\2\2[=\3\2\2\2[@\3\2\2\2[C\3"+
		"\2\2\2[F\3\2\2\2[I\3\2\2\2[L\3\2\2\2[O\3\2\2\2[R\3\2\2\2[U\3\2\2\2[X\3"+
		"\2\2\2\\\7\3\2\2\2]^\7\21\2\2^_\5,\27\2_`\7\f\2\2`a\7<\2\2ab\5\4\3\2b"+
		"c\7=\2\2cd\b\5\1\2d\t\3\2\2\2ef\5(\25\2fg\7\'\2\2gh\7>\2\2hi\7?\2\2ij"+
		"\5*\26\2jk\7@\2\2kl\7\60\2\2lm\7?\2\2mn\5\f\7\2no\7@\2\2op\b\6\1\2p\13"+
		"\3\2\2\2qr\b\7\1\2rs\5,\27\2st\b\7\1\2tw\3\2\2\2uw\3\2\2\2vq\3\2\2\2v"+
		"u\3\2\2\2w\177\3\2\2\2xy\f\5\2\2yz\7A\2\2z{\5,\27\2{|\b\7\1\2|~\3\2\2"+
		"\2}x\3\2\2\2~\u0081\3\2\2\2\177}\3\2\2\2\177\u0080\3\2\2\2\u0080\r\3\2"+
		"\2\2\u0081\177\3\2\2\2\u0082\u0083\7\22\2\2\u0083\u0089\b\b\1\2\u0084"+
		"\u0085\7\23\2\2\u0085\u0089\b\b\1\2\u0086\u0087\7\25\2\2\u0087\u0089\b"+
		"\b\1\2\u0088\u0082\3\2\2\2\u0088\u0084\3\2\2\2\u0088\u0086\3\2\2\2\u0089"+
		"\17\3\2\2\2\u008a\u008b\7\16\2\2\u008b\u008c\5,\27\2\u008c\u008d\7<\2"+
		"\2\u008d\u008e\5\4\3\2\u008e\u008f\7=\2\2\u008f\u0090\b\t\1\2\u0090\21"+
		"\3\2\2\2\u0091\u0092\7\17\2\2\u0092\u0093\7\'\2\2\u0093\u0094\7\20\2\2"+
		"\u0094\u0095\5,\27\2\u0095\u0096\7<\2\2\u0096\u0097\5\4\3\2\u0097\u0098"+
		"\7=\2\2\u0098\u0099\b\n\1\2\u0099\23\3\2\2\2\u009a\u009b\7\n\2\2\u009b"+
		"\u009c\7:\2\2\u009c\u009d\5,\27\2\u009d\u009e\7;\2\2\u009e\u009f\b\13"+
		"\1\2\u009f\25\3\2\2\2\u00a0\u00a1\7\13\2\2\u00a1\u00a2\5,\27\2\u00a2\u00a3"+
		"\7<\2\2\u00a3\u00a4\5\4\3\2\u00a4\u00a5\7=\2\2\u00a5\u00a6\5\32\16\2\u00a6"+
		"\u00a7\5\30\r\2\u00a7\u00a8\b\f\1\2\u00a8\27\3\2\2\2\u00a9\u00aa\7\f\2"+
		"\2\u00aa\u00ab\7<\2\2\u00ab\u00ac\5\4\3\2\u00ac\u00ad\7=\2\2\u00ad\u00ae"+
		"\b\r\1\2\u00ae\u00b1\3\2\2\2\u00af\u00b1\3\2\2\2\u00b0\u00a9\3\2\2\2\u00b0"+
		"\u00af\3\2\2\2\u00b1\31\3\2\2\2\u00b2\u00bd\b\16\1\2\u00b3\u00b4\f\4\2"+
		"\2\u00b4\u00b5\7\r\2\2\u00b5\u00b6\5,\27\2\u00b6\u00b7\7<\2\2\u00b7\u00b8"+
		"\5\4\3\2\u00b8\u00b9\7=\2\2\u00b9\u00ba\b\16\1\2\u00ba\u00bc\3\2\2\2\u00bb"+
		"\u00b3\3\2\2\2\u00bc\u00bf\3\2\2\2\u00bd\u00bb\3\2\2\2\u00bd\u00be\3\2"+
		"\2\2\u00be\33\3\2\2\2\u00bf\u00bd\3\2\2\2\u00c0\u00c1\5(\25\2\u00c1\u00c2"+
		"\7\'\2\2\u00c2\u00c3\7>\2\2\u00c3\u00c4\5*\26\2\u00c4\u00c5\7\60\2\2\u00c5"+
		"\u00c6\5,\27\2\u00c6\u00c7\b\17\1\2\u00c7\u00cf\3\2\2\2\u00c8\u00c9\5"+
		"(\25\2\u00c9\u00ca\7\'\2\2\u00ca\u00cb\7\60\2\2\u00cb\u00cc\5,\27\2\u00cc"+
		"\u00cd\b\17\1\2\u00cd\u00cf\3\2\2\2\u00ce\u00c0\3\2\2\2\u00ce\u00c8\3"+
		"\2\2\2\u00cf\35\3\2\2\2\u00d0\u00d1\7\24\2\2\u00d1\u00d2\5,\27\2\u00d2"+
		"\u00d3\7<\2\2\u00d3\u00d4\5\"\22\2\u00d4\u00d5\5 \21\2\u00d5\u00d6\7="+
		"\2\2\u00d6\u00d7\b\20\1\2\u00d7\37\3\2\2\2\u00d8\u00d9\7\34\2\2\u00d9"+
		"\u00da\7>\2\2\u00da\u00db\5\4\3\2\u00db\u00dc\b\21\1\2\u00dc\u00df\3\2"+
		"\2\2\u00dd\u00df\3\2\2\2\u00de\u00d8\3\2\2\2\u00de\u00dd\3\2\2\2\u00df"+
		"!\3\2\2\2\u00e0\u00ea\b\22\1\2\u00e1\u00e2\f\4\2\2\u00e2\u00e3\7\26\2"+
		"\2\u00e3\u00e4\5,\27\2\u00e4\u00e5\7>\2\2\u00e5\u00e6\5\4\3\2\u00e6\u00e7"+
		"\b\22\1\2\u00e7\u00e9\3\2\2\2\u00e8\u00e1\3\2\2\2\u00e9\u00ec\3\2\2\2"+
		"\u00ea\u00e8\3\2\2\2\u00ea\u00eb\3\2\2\2\u00eb#\3\2\2\2\u00ec\u00ea\3"+
		"\2\2\2\u00ed\u00ee\7\'\2\2\u00ee\u00ef\5&\24\2\u00ef\u00f0\5,\27\2\u00f0"+
		"\u00f1\b\23\1\2\u00f1%\3\2\2\2\u00f2\u00f3\7\60\2\2\u00f3\u00fb\b\24\1"+
		"\2\u00f4\u00f5\7)\2\2\u00f5\u00fb\b\24\1\2\u00f6\u00f7\7*\2\2\u00f7\u00fb"+
		"\b\24\1\2\u00f8\u00f9\7+\2\2\u00f9\u00fb\b\24\1\2\u00fa\u00f2\3\2\2\2"+
		"\u00fa\u00f4\3\2\2\2\u00fa\u00f6\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fb\'\3"+
		"\2\2\2\u00fc\u00fd\7\27\2\2\u00fd\u0101\b\25\1\2\u00fe\u00ff\7\30\2\2"+
		"\u00ff\u0101\b\25\1\2\u0100\u00fc\3\2\2\2\u0100\u00fe\3\2\2\2\u0101)\3"+
		"\2\2\2\u0102\u0103\7\3\2\2\u0103\u0110\b\26\1\2\u0104\u0105\7\4\2\2\u0105"+
		"\u0110\b\26\1\2\u0106\u0107\7\6\2\2\u0107\u0110\b\26\1\2\u0108\u0109\7"+
		"\5\2\2\u0109\u0110\b\26\1\2\u010a\u010b\7\7\2\2\u010b\u0110\b\26\1\2\u010c"+
		"\u010d\7?\2\2\u010d\u010e\7@\2\2\u010e\u0110\b\26\1\2\u010f\u0102\3\2"+
		"\2\2\u010f\u0104\3\2\2\2\u010f\u0106\3\2\2\2\u010f\u0108\3\2\2\2\u010f"+
		"\u010a\3\2\2\2\u010f\u010c\3\2\2\2\u0110+\3\2\2\2\u0111\u0112\b\27\1\2"+
		"\u0112\u0113\5\60\31\2\u0113\u0114\b\27\1\2\u0114\u0123\3\2\2\2\u0115"+
		"\u0116\7?\2\2\u0116\u0117\5.\30\2\u0117\u0118\7@\2\2\u0118\u0119\b\27"+
		"\1\2\u0119\u0123\3\2\2\2\u011a\u011b\7%\2\2\u011b\u0123\b\27\1\2\u011c"+
		"\u011d\7&\2\2\u011d\u0123\b\27\1\2\u011e\u011f\7\b\2\2\u011f\u0123\b\27"+
		"\1\2\u0120\u0121\7\t\2\2\u0121\u0123\b\27\1\2\u0122\u0111\3\2\2\2\u0122"+
		"\u0115\3\2\2\2\u0122\u011a\3\2\2\2\u0122\u011c\3\2\2\2\u0122\u011e\3\2"+
		"\2\2\u0122\u0120\3\2\2\2\u0123\u015d\3\2\2\2\u0124\u0125\f\23\2\2\u0125"+
		"\u0126\t\2\2\2\u0126\u0127\5,\27\24\u0127\u0128\b\27\1\2\u0128\u015c\3"+
		"\2\2\2\u0129\u012a\f\22\2\2\u012a\u012b\t\3\2\2\u012b\u012c\5,\27\23\u012c"+
		"\u012d\b\27\1\2\u012d\u015c\3\2\2\2\u012e\u012f\f\21\2\2\u012f\u0130\t"+
		"\4\2\2\u0130\u0131\5,\27\22\u0131\u0132\b\27\1\2\u0132\u015c\3\2\2\2\u0133"+
		"\u0134\f\20\2\2\u0134\u0135\t\5\2\2\u0135\u0136\5,\27\21\u0136\u0137\b"+
		"\27\1\2\u0137\u015c\3\2\2\2\u0138\u0139\f\17\2\2\u0139\u013a\t\6\2\2\u013a"+
		"\u013b\5,\27\20\u013b\u013c\b\27\1\2\u013c\u015c\3\2\2\2\u013d\u013e\f"+
		"\16\2\2\u013e\u013f\7.\2\2\u013f\u0140\5,\27\17\u0140\u0141\b\27\1\2\u0141"+
		"\u015c\3\2\2\2\u0142\u0143\f\r\2\2\u0143\u0144\7-\2\2\u0144\u0145\5,\27"+
		"\16\u0145\u0146\b\27\1\2\u0146\u015c\3\2\2\2\u0147\u0148\f\f\2\2\u0148"+
		"\u0149\7*\2\2\u0149\u014a\5,\27\r\u014a\u014b\b\27\1\2\u014b\u015c\3\2"+
		"\2\2\u014c\u014d\f\13\2\2\u014d\u014e\7+\2\2\u014e\u014f\5,\27\f\u014f"+
		"\u0150\b\27\1\2\u0150\u015c\3\2\2\2\u0151\u0152\f\n\2\2\u0152\u0153\7"+
		"8\2\2\u0153\u0154\5,\27\13\u0154\u0155\b\27\1\2\u0155\u015c\3\2\2\2\u0156"+
		"\u0157\f\t\2\2\u0157\u0158\78\2\2\u0158\u0159\5,\27\n\u0159\u015a\b\27"+
		"\1\2\u015a\u015c\3\2\2\2\u015b\u0124\3\2\2\2\u015b\u0129\3\2\2\2\u015b"+
		"\u012e\3\2\2\2\u015b\u0133\3\2\2\2\u015b\u0138\3\2\2\2\u015b\u013d\3\2"+
		"\2\2\u015b\u0142\3\2\2\2\u015b\u0147\3\2\2\2\u015b\u014c\3\2\2\2\u015b"+
		"\u0151\3\2\2\2\u015b\u0156\3\2\2\2\u015c\u015f\3\2\2\2\u015d\u015b\3\2"+
		"\2\2\u015d\u015e\3\2\2\2\u015e-\3\2\2\2\u015f\u015d\3\2\2\2\u0160\u0161"+
		"\b\30\1\2\u0161\u0162\5,\27\2\u0162\u0163\b\30\1\2\u0163\u016b\3\2\2\2"+
		"\u0164\u0165\f\4\2\2\u0165\u0166\7A\2\2\u0166\u0167\5,\27\2\u0167\u0168"+
		"\b\30\1\2\u0168\u016a\3\2\2\2\u0169\u0164\3\2\2\2\u016a\u016d\3\2\2\2"+
		"\u016b\u0169\3\2\2\2\u016b\u016c\3\2\2\2\u016c/\3\2\2\2\u016d\u016b\3"+
		"\2\2\2\u016e\u016f\b\31\1\2\u016f\u0170\7\'\2\2\u0170\u0171\b\31\1\2\u0171"+
		"\u017a\3\2\2\2\u0172\u0173\f\4\2\2\u0173\u0174\7?\2\2\u0174\u0175\5,\27"+
		"\2\u0175\u0176\7@\2\2\u0176\u0177\b\31\1\2\u0177\u0179\3\2\2\2\u0178\u0172"+
		"\3\2\2\2\u0179\u017c\3\2\2\2\u017a\u0178\3\2\2\2\u017a\u017b\3\2\2\2\u017b"+
		"\61\3\2\2\2\u017c\u017a\3\2\2\2\249[v\177\u0088\u00b0\u00bd\u00ce\u00de"+
		"\u00ea\u00fa\u0100\u010f\u0122\u015b\u015d\u016b\u017a";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}