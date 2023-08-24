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
		REPEATING=25, COUNT=26, STRUCT=27, MUTATING=28, FUNC=29, SELF_D=30, PUNTO=31, 
		INOUT=32, GUION_B=33, NUMBER=34, STRING=35, ID=36, DIF=37, IG_IG=38, MAS_IGUAL=39, 
		MENOS_IGUAL=40, NOT=41, OR=42, AND=43, SOLO_AND=44, IG=45, MAY_IG=46, 
		MEN_IG=47, MAYOR=48, MENOR=49, MUL=50, DIV=51, ADD=52, MODULO=53, SUB=54, 
		PARIZQ=55, PARDER=56, LLAVEIZQ=57, LLAVEDER=58, D_PTS=59, CORIZQ=60, CORDER=61, 
		COMA=62, PUNTO_COMA=63, WHITESPACE=64, COMMENT=65, LINE_COMMENT=66, DEFAULT=67;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_ifstmt = 4, RULE_elsestmt = 5, RULE_else_ifstmt = 6, RULE_declarationstmt = 7, 
		RULE_switchstmt = 8, RULE_defaultsts = 9, RULE_case = 10, RULE_asignacionstmt = 11, 
		RULE_ig_type = 12, RULE_reserv = 13, RULE_types = 14, RULE_expr = 15, 
		RULE_listParams = 16, RULE_listArray = 17;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "ifstmt", "elsestmt", "else_ifstmt", 
			"declarationstmt", "switchstmt", "defaultsts", "case", "asignacionstmt", 
			"ig_type", "reserv", "types", "expr", "listParams", "listArray"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'", 
			"'false'", "'print'", "'if'", "'else'", "'else if'", "'while'", "'for'", 
			"'in'", "'guard'", "'continue'", "'return'", "'switch'", "'break'", "'case'", 
			"'var'", "'let'", "'.append'", "'.remove'", "'repeating'", "'count'", 
			"'STRUCT'", "'mutating'", "'func'", "'self.'", "'.'", "'inout'", "'_'", 
			null, null, null, "'!='", "'=='", "'+='", "'-='", "'!'", "'||'", "'&&'", 
			"'&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'%'", 
			"'-'", "'('", "')'", "'{'", "'}'", "':'", "'['", "']'", "','", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "CHAR", "TRU", "FAL", "PRINT", "IF", 
			"ELSE", "ELSE_IF", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", 
			"SWITCH", "BREAK", "CASE", "VAR", "LET", "D_APPEND", "D_REMOVE", "REPEATING", 
			"COUNT", "STRUCT", "MUTATING", "FUNC", "SELF_D", "PUNTO", "INOUT", "GUION_B", 
			"NUMBER", "STRING", "ID", "DIF", "IG_IG", "MAS_IGUAL", "MENOS_IGUAL", 
			"NOT", "OR", "AND", "SOLO_AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", 
			"MUL", "DIV", "ADD", "MODULO", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", 
			"LLAVEDER", "D_PTS", "CORIZQ", "CORDER", "COMA", "PUNTO_COMA", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT", "DEFAULT"
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
			setState(36);
			((SContext)_localctx).block = block();
			setState(37);
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
			setState(41); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(40);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(43); 
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
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(62);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				enterOuterAlt(_localctx, 1);
				{
				setState(47);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 2);
				{
				setState(50);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinst 
				}
				break;
			case VAR:
			case LET:
				enterOuterAlt(_localctx, 3);
				{
				setState(53);
				((InstructionContext)_localctx).declarationstmt = declarationstmt();
				 _localctx.inst = ((InstructionContext)_localctx).declarationstmt.dec 
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 4);
				{
				setState(56);
				asignacionstmt();

				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 5);
				{
				setState(59);
				switchstmt();

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
		enterRule(_localctx, 6, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(65);
			match(PARIZQ);
			setState(66);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(67);
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
		enterRule(_localctx, 8, RULE_ifstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			match(IF);
			setState(71);
			expr(0);
			setState(72);
			match(LLAVEIZQ);
			setState(73);
			block();
			setState(74);
			match(LLAVEDER);
			setState(75);
			else_ifstmt(0);
			setState(76);
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
		enterRule(_localctx, 10, RULE_elsestmt);
		try {
			setState(86);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(79);
				match(ELSE);
				setState(80);
				match(LLAVEIZQ);
				setState(81);
				block();
				setState(82);
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
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_else_ifstmt, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			}
			_ctx.stop = _input.LT(-1);
			setState(99);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Else_ifstmtContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_else_ifstmt);
					setState(89);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(90);
					match(ELSE_IF);
					setState(91);
					expr(0);
					setState(92);
					match(LLAVEIZQ);
					setState(93);
					block();
					setState(94);
					match(LLAVEDER);

					}
					} 
				}
				setState(101);
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
		enterRule(_localctx, 14, RULE_declarationstmt);
		try {
			setState(116);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				reserv();
				setState(103);
				match(ID);
				setState(104);
				match(D_PTS);
				setState(105);
				types();
				setState(106);
				match(IG);
				setState(107);
				expr(0);

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(110);
				reserv();
				setState(111);
				match(ID);
				setState(112);
				match(IG);
				setState(113);
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
		enterRule(_localctx, 16, RULE_switchstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(118);
			match(SWITCH);
			setState(119);
			expr(0);
			setState(120);
			match(LLAVEIZQ);
			setState(121);
			case(0);
			setState(122);
			defaultsts();
			setState(123);
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
		enterRule(_localctx, 18, RULE_defaultsts);
		try {
			setState(132);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DEFAULT:
				enterOuterAlt(_localctx, 1);
				{
				setState(126);
				match(DEFAULT);
				setState(127);
				match(D_PTS);
				setState(128);
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
		int _startState = 20;
		enterRecursionRule(_localctx, 20, RULE_case, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			}
			_ctx.stop = _input.LT(-1);
			setState(144);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new CaseContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_case);
					setState(135);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(136);
					match(CASE);
					setState(137);
					expr(0);
					setState(138);
					match(D_PTS);
					setState(139);
					block();

					}
					} 
				}
				setState(146);
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
		enterRule(_localctx, 22, RULE_asignacionstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			match(ID);
			setState(148);
			ig_type();
			setState(149);
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
		enterRule(_localctx, 24, RULE_ig_type);
		try {
			setState(160);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IG:
				enterOuterAlt(_localctx, 1);
				{
				setState(152);
				match(IG);

				}
				break;
			case IG_IG:
				enterOuterAlt(_localctx, 2);
				{
				setState(154);
				match(IG_IG);

				}
				break;
			case MAS_IGUAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(156);
				match(MAS_IGUAL);

				}
				break;
			case MENOS_IGUAL:
				enterOuterAlt(_localctx, 4);
				{
				setState(158);
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
		enterRule(_localctx, 26, RULE_reserv);
		try {
			setState(166);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(162);
				match(VAR);

				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(164);
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
		enterRule(_localctx, 28, RULE_types);
		try {
			setState(181);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(168);
				match(INT);
				 _localctx.ty = environment.INTEGER 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(170);
				match(FLOAT);
				 _localctx.ty = environment.FLOAT 
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(172);
				match(STR);
				 _localctx.ty = environment.STRING 
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(174);
				match(BOOL);
				 _localctx.ty = environment.BOOLEAN 
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 5);
				{
				setState(176);
				match(CHAR);

				}
				break;
			case CORIZQ:
				enterOuterAlt(_localctx, 6);
				{
				setState(178);
				match(CORIZQ);
				setState(179);
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
		public ExprContext expr;
		public ListArrayContext list;
		public Token CORIZQ;
		public ListParamsContext listParams;
		public Token NUMBER;
		public Token STRING;
		public Token TRU;
		public Token FAL;
		public Token op;
		public ExprContext right;
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
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
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PARIZQ:
				{
				setState(184);
				match(PARIZQ);
				setState(185);
				((ExprContext)_localctx).expr = expr(0);
				setState(186);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case ID:
				{
				setState(189);
				((ExprContext)_localctx).list = listArray(0);
				 _localctx.e = ((ExprContext)_localctx).list.p
				}
				break;
			case CORIZQ:
				{
				setState(192);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(193);
				((ExprContext)_localctx).listParams = listParams(0);
				setState(194);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), ((ExprContext)_localctx).listParams.l) 
				}
				break;
			case NUMBER:
				{
				setState(197);
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
				setState(199);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case TRU:
				{
				setState(201);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case FAL:
				{
				setState(203);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(244);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(242);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(207);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(208);
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
						setState(209);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(212);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(213);
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
						setState(214);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(217);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(218);
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
						setState(219);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(222);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(223);
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
						setState(224);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(12);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(227);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(228);
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
						setState(229);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(11);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(232);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(233);
						((ExprContext)_localctx).op = match(AND);
						setState(234);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(10);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(237);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(238);
						((ExprContext)_localctx).op = match(OR);
						setState(239);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(9);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(246);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
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
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(248);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(258);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
					setState(251);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(252);
					match(COMA);
					setState(253);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(260);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
		int _startState = 34;
		enterRecursionRule(_localctx, 34, RULE_listArray, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(262);
			((ListArrayContext)_localctx).ID = match(ID);
			 _localctx.p = expressions.NewCallVar((((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getCharPositionInLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(273);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
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
					setState(265);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(266);
					match(CORIZQ);
					setState(267);
					((ListArrayContext)_localctx).expr = expr(0);
					setState(268);
					match(CORDER);
					 _localctx.p = expressions.NewArrayAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, ((ListArrayContext)_localctx).expr.e) 
					}
					} 
				}
				setState(275);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
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
		case 6:
			return else_ifstmt_sempred((Else_ifstmtContext)_localctx, predIndex);
		case 10:
			return case_sempred((CaseContext)_localctx, predIndex);
		case 15:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 16:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 17:
			return listArray_sempred((ListArrayContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean else_ifstmt_sempred(Else_ifstmtContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean case_sempred(CaseContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 14);
		case 3:
			return precpred(_ctx, 13);
		case 4:
			return precpred(_ctx, 12);
		case 5:
			return precpred(_ctx, 11);
		case 6:
			return precpred(_ctx, 10);
		case 7:
			return precpred(_ctx, 9);
		case 8:
			return precpred(_ctx, 8);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 9:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listArray_sempred(ListArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 10:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3E\u0117\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\3\2\3\2\3\2\3\2\3\3\6\3,\n\3\r\3\16\3-\3\3\3\3\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4A\n\4\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\5\7Y\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\7\bd\n\b\f\b\16\b"+
		"g\13\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\tw\n"+
		"\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\5\13"+
		"\u0087\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\7\f\u0091\n\f\f\f\16\f\u0094"+
		"\13\f\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16"+
		"\u00a3\n\16\3\17\3\17\3\17\3\17\5\17\u00a9\n\17\3\20\3\20\3\20\3\20\3"+
		"\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u00b8\n\20\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u00d0\n\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\7\21\u00f5\n\21\f\21\16\21\u00f8\13\21\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\7\22\u0103\n\22\f\22\16\22\u0106\13\22\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\7\23\u0112\n\23\f\23\16"+
		"\23\u0115\13\23\3\23\2\7\16\26 \"$\24\2\4\6\b\n\f\16\20\22\24\26\30\32"+
		"\34\36 \"$\2\7\3\2\64\65\4\2\66\6688\4\2\60\60\62\62\4\2\61\61\63\63\3"+
		"\2\'(\2\u0126\2&\3\2\2\2\4+\3\2\2\2\6@\3\2\2\2\bB\3\2\2\2\nH\3\2\2\2\f"+
		"X\3\2\2\2\16Z\3\2\2\2\20v\3\2\2\2\22x\3\2\2\2\24\u0086\3\2\2\2\26\u0088"+
		"\3\2\2\2\30\u0095\3\2\2\2\32\u00a2\3\2\2\2\34\u00a8\3\2\2\2\36\u00b7\3"+
		"\2\2\2 \u00cf\3\2\2\2\"\u00f9\3\2\2\2$\u0107\3\2\2\2&\'\5\4\3\2\'(\7\2"+
		"\2\3()\b\2\1\2)\3\3\2\2\2*,\5\6\4\2+*\3\2\2\2,-\3\2\2\2-+\3\2\2\2-.\3"+
		"\2\2\2./\3\2\2\2/\60\b\3\1\2\60\5\3\2\2\2\61\62\5\b\5\2\62\63\b\4\1\2"+
		"\63A\3\2\2\2\64\65\5\n\6\2\65\66\b\4\1\2\66A\3\2\2\2\678\5\20\t\289\b"+
		"\4\1\29A\3\2\2\2:;\5\30\r\2;<\b\4\1\2<A\3\2\2\2=>\5\22\n\2>?\b\4\1\2?"+
		"A\3\2\2\2@\61\3\2\2\2@\64\3\2\2\2@\67\3\2\2\2@:\3\2\2\2@=\3\2\2\2A\7\3"+
		"\2\2\2BC\7\n\2\2CD\79\2\2DE\5 \21\2EF\7:\2\2FG\b\5\1\2G\t\3\2\2\2HI\7"+
		"\13\2\2IJ\5 \21\2JK\7;\2\2KL\5\4\3\2LM\7<\2\2MN\5\16\b\2NO\5\f\7\2OP\b"+
		"\6\1\2P\13\3\2\2\2QR\7\f\2\2RS\7;\2\2ST\5\4\3\2TU\7<\2\2UV\b\7\1\2VY\3"+
		"\2\2\2WY\3\2\2\2XQ\3\2\2\2XW\3\2\2\2Y\r\3\2\2\2Ze\b\b\1\2[\\\f\4\2\2\\"+
		"]\7\r\2\2]^\5 \21\2^_\7;\2\2_`\5\4\3\2`a\7<\2\2ab\b\b\1\2bd\3\2\2\2c["+
		"\3\2\2\2dg\3\2\2\2ec\3\2\2\2ef\3\2\2\2f\17\3\2\2\2ge\3\2\2\2hi\5\34\17"+
		"\2ij\7&\2\2jk\7=\2\2kl\5\36\20\2lm\7/\2\2mn\5 \21\2no\b\t\1\2ow\3\2\2"+
		"\2pq\5\34\17\2qr\7&\2\2rs\7/\2\2st\5 \21\2tu\b\t\1\2uw\3\2\2\2vh\3\2\2"+
		"\2vp\3\2\2\2w\21\3\2\2\2xy\7\24\2\2yz\5 \21\2z{\7;\2\2{|\5\26\f\2|}\5"+
		"\24\13\2}~\7<\2\2~\177\b\n\1\2\177\23\3\2\2\2\u0080\u0081\7E\2\2\u0081"+
		"\u0082\7=\2\2\u0082\u0083\5\4\3\2\u0083\u0084\b\13\1\2\u0084\u0087\3\2"+
		"\2\2\u0085\u0087\3\2\2\2\u0086\u0080\3\2\2\2\u0086\u0085\3\2\2\2\u0087"+
		"\25\3\2\2\2\u0088\u0092\b\f\1\2\u0089\u008a\f\4\2\2\u008a\u008b\7\26\2"+
		"\2\u008b\u008c\5 \21\2\u008c\u008d\7=\2\2\u008d\u008e\5\4\3\2\u008e\u008f"+
		"\b\f\1\2\u008f\u0091\3\2\2\2\u0090\u0089\3\2\2\2\u0091\u0094\3\2\2\2\u0092"+
		"\u0090\3\2\2\2\u0092\u0093\3\2\2\2\u0093\27\3\2\2\2\u0094\u0092\3\2\2"+
		"\2\u0095\u0096\7&\2\2\u0096\u0097\5\32\16\2\u0097\u0098\5 \21\2\u0098"+
		"\u0099\b\r\1\2\u0099\31\3\2\2\2\u009a\u009b\7/\2\2\u009b\u00a3\b\16\1"+
		"\2\u009c\u009d\7(\2\2\u009d\u00a3\b\16\1\2\u009e\u009f\7)\2\2\u009f\u00a3"+
		"\b\16\1\2\u00a0\u00a1\7*\2\2\u00a1\u00a3\b\16\1\2\u00a2\u009a\3\2\2\2"+
		"\u00a2\u009c\3\2\2\2\u00a2\u009e\3\2\2\2\u00a2\u00a0\3\2\2\2\u00a3\33"+
		"\3\2\2\2\u00a4\u00a5\7\27\2\2\u00a5\u00a9\b\17\1\2\u00a6\u00a7\7\30\2"+
		"\2\u00a7\u00a9\b\17\1\2\u00a8\u00a4\3\2\2\2\u00a8\u00a6\3\2\2\2\u00a9"+
		"\35\3\2\2\2\u00aa\u00ab\7\3\2\2\u00ab\u00b8\b\20\1\2\u00ac\u00ad\7\4\2"+
		"\2\u00ad\u00b8\b\20\1\2\u00ae\u00af\7\6\2\2\u00af\u00b8\b\20\1\2\u00b0"+
		"\u00b1\7\5\2\2\u00b1\u00b8\b\20\1\2\u00b2\u00b3\7\7\2\2\u00b3\u00b8\b"+
		"\20\1\2\u00b4\u00b5\7>\2\2\u00b5\u00b6\7?\2\2\u00b6\u00b8\b\20\1\2\u00b7"+
		"\u00aa\3\2\2\2\u00b7\u00ac\3\2\2\2\u00b7\u00ae\3\2\2\2\u00b7\u00b0\3\2"+
		"\2\2\u00b7\u00b2\3\2\2\2\u00b7\u00b4\3\2\2\2\u00b8\37\3\2\2\2\u00b9\u00ba"+
		"\b\21\1\2\u00ba\u00bb\79\2\2\u00bb\u00bc\5 \21\2\u00bc\u00bd\7:\2\2\u00bd"+
		"\u00be\b\21\1\2\u00be\u00d0\3\2\2\2\u00bf\u00c0\5$\23\2\u00c0\u00c1\b"+
		"\21\1\2\u00c1\u00d0\3\2\2\2\u00c2\u00c3\7>\2\2\u00c3\u00c4\5\"\22\2\u00c4"+
		"\u00c5\7?\2\2\u00c5\u00c6\b\21\1\2\u00c6\u00d0\3\2\2\2\u00c7\u00c8\7$"+
		"\2\2\u00c8\u00d0\b\21\1\2\u00c9\u00ca\7%\2\2\u00ca\u00d0\b\21\1\2\u00cb"+
		"\u00cc\7\b\2\2\u00cc\u00d0\b\21\1\2\u00cd\u00ce\7\t\2\2\u00ce\u00d0\b"+
		"\21\1\2\u00cf\u00b9\3\2\2\2\u00cf\u00bf\3\2\2\2\u00cf\u00c2\3\2\2\2\u00cf"+
		"\u00c7\3\2\2\2\u00cf\u00c9\3\2\2\2\u00cf\u00cb\3\2\2\2\u00cf\u00cd\3\2"+
		"\2\2\u00d0\u00f6\3\2\2\2\u00d1\u00d2\f\20\2\2\u00d2\u00d3\t\2\2\2\u00d3"+
		"\u00d4\5 \21\21\u00d4\u00d5\b\21\1\2\u00d5\u00f5\3\2\2\2\u00d6\u00d7\f"+
		"\17\2\2\u00d7\u00d8\t\3\2\2\u00d8\u00d9\5 \21\20\u00d9\u00da\b\21\1\2"+
		"\u00da\u00f5\3\2\2\2\u00db\u00dc\f\16\2\2\u00dc\u00dd\t\4\2\2\u00dd\u00de"+
		"\5 \21\17\u00de\u00df\b\21\1\2\u00df\u00f5\3\2\2\2\u00e0\u00e1\f\r\2\2"+
		"\u00e1\u00e2\t\5\2\2\u00e2\u00e3\5 \21\16\u00e3\u00e4\b\21\1\2\u00e4\u00f5"+
		"\3\2\2\2\u00e5\u00e6\f\f\2\2\u00e6\u00e7\t\6\2\2\u00e7\u00e8\5 \21\r\u00e8"+
		"\u00e9\b\21\1\2\u00e9\u00f5\3\2\2\2\u00ea\u00eb\f\13\2\2\u00eb\u00ec\7"+
		"-\2\2\u00ec\u00ed\5 \21\f\u00ed\u00ee\b\21\1\2\u00ee\u00f5\3\2\2\2\u00ef"+
		"\u00f0\f\n\2\2\u00f0\u00f1\7,\2\2\u00f1\u00f2\5 \21\13\u00f2\u00f3\b\21"+
		"\1\2\u00f3\u00f5\3\2\2\2\u00f4\u00d1\3\2\2\2\u00f4\u00d6\3\2\2\2\u00f4"+
		"\u00db\3\2\2\2\u00f4\u00e0\3\2\2\2\u00f4\u00e5\3\2\2\2\u00f4\u00ea\3\2"+
		"\2\2\u00f4\u00ef\3\2\2\2\u00f5\u00f8\3\2\2\2\u00f6\u00f4\3\2\2\2\u00f6"+
		"\u00f7\3\2\2\2\u00f7!\3\2\2\2\u00f8\u00f6\3\2\2\2\u00f9\u00fa\b\22\1\2"+
		"\u00fa\u00fb\5 \21\2\u00fb\u00fc\b\22\1\2\u00fc\u0104\3\2\2\2\u00fd\u00fe"+
		"\f\4\2\2\u00fe\u00ff\7@\2\2\u00ff\u0100\5 \21\2\u0100\u0101\b\22\1\2\u0101"+
		"\u0103\3\2\2\2\u0102\u00fd\3\2\2\2\u0103\u0106\3\2\2\2\u0104\u0102\3\2"+
		"\2\2\u0104\u0105\3\2\2\2\u0105#\3\2\2\2\u0106\u0104\3\2\2\2\u0107\u0108"+
		"\b\23\1\2\u0108\u0109\7&\2\2\u0109\u010a\b\23\1\2\u010a\u0113\3\2\2\2"+
		"\u010b\u010c\f\4\2\2\u010c\u010d\7>\2\2\u010d\u010e\5 \21\2\u010e\u010f"+
		"\7?\2\2\u010f\u0110\b\23\1\2\u0110\u0112\3\2\2\2\u0111\u010b\3\2\2\2\u0112"+
		"\u0115\3\2\2\2\u0113\u0111\3\2\2\2\u0113\u0114\3\2\2\2\u0114%\3\2\2\2"+
		"\u0115\u0113\3\2\2\2\21-@Xev\u0086\u0092\u00a2\u00a8\u00b7\u00cf\u00f4"+
		"\u00f6\u0104\u0113";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}