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
		INT=1, FLOAT=2, BOOL=3, STR=4, CHAR=5, TRU=6, FAL=7, PRINT=8, IF=9, ELSE_IF=10, 
		ELSE=11, WHILE=12, FOR=13, IN=14, GUARD=15, CONTINUE=16, RETURN=17, SWITCH=18, 
		BREAK=19, CASE=20, VAR=21, LET=22, D_APPEND=23, D_REMOVE=24, REPEATING=25, 
		COUNT=26, STRUCT=27, MUTATING=28, FUNC=29, SELF_D=30, PUNTO=31, INOUT=32, 
		GUION_B=33, NUMBER=34, STRING=35, ID=36, DIF=37, IG_IG=38, MAS_IGUAL=39, 
		MENOS_IGUAL=40, NOT=41, OR=42, AND=43, SOLO_AND=44, IG=45, MAY_IG=46, 
		MEN_IG=47, MAYOR=48, MENOR=49, MUL=50, DIV=51, ADD=52, MODULO=53, SUB=54, 
		PARIZQ=55, PARDER=56, LLAVEIZQ=57, LLAVEDER=58, D_PTS=59, CORIZQ=60, CORDER=61, 
		COMA=62, PUNTO_COMA=63, WHITESPACE=64, COMMENT=65, LINE_COMMENT=66;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_ifstmt = 4, RULE_elsestmt = 5, RULE_else_ifstmt = 6, RULE_declarationstmt = 7, 
		RULE_asignacionstmt = 8, RULE_ig_type = 9, RULE_reserv = 10, RULE_types = 11, 
		RULE_expr = 12, RULE_listParams = 13, RULE_listArray = 14;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "ifstmt", "elsestmt", "else_ifstmt", 
			"declarationstmt", "asignacionstmt", "ig_type", "reserv", "types", "expr", 
			"listParams", "listArray"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'", 
			"'false'", "'print'", "'if'", "'else if'", "'else'", "'while'", "'for'", 
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
			"ELSE_IF", "ELSE", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", 
			"SWITCH", "BREAK", "CASE", "VAR", "LET", "D_APPEND", "D_REMOVE", "REPEATING", 
			"COUNT", "STRUCT", "MUTATING", "FUNC", "SELF_D", "PUNTO", "INOUT", "GUION_B", 
			"NUMBER", "STRING", "ID", "DIF", "IG_IG", "MAS_IGUAL", "MENOS_IGUAL", 
			"NOT", "OR", "AND", "SOLO_AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", 
			"MUL", "DIV", "ADD", "MODULO", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", 
			"LLAVEDER", "D_PTS", "CORIZQ", "CORDER", "COMA", "PUNTO_COMA", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT"
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
			setState(30);
			((SContext)_localctx).block = block();
			setState(31);
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
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(35); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(34);
				((BlockContext)_localctx).instruction = instruction();
				((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
				}
				}
				setState(37); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRINT) | (1L << IF) | (1L << VAR) | (1L << LET) | (1L << ID))) != 0) );

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
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(53);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				enterOuterAlt(_localctx, 1);
				{
				setState(41);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 2);
				{
				setState(44);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinst 
				}
				break;
			case VAR:
			case LET:
				enterOuterAlt(_localctx, 3);
				{
				setState(47);
				((InstructionContext)_localctx).declarationstmt = declarationstmt();
				 _localctx.inst = ((InstructionContext)_localctx).declarationstmt.dec 
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 4);
				{
				setState(50);
				asignacionstmt();

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
			setState(55);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(56);
			match(PARIZQ);
			setState(57);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(58);
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
		public TerminalNode ELSE_IF() { return getToken(SwiftGrammarParser.ELSE_IF, 0); }
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
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
			setState(61);
			match(IF);
			setState(62);
			expr(0);
			setState(63);
			match(LLAVEIZQ);
			setState(64);
			block();
			setState(65);
			match(LLAVEDER);
			setState(66);
			match(ELSE_IF);
			setState(67);
			match(ELSE);

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
			setState(77);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ELSE:
				enterOuterAlt(_localctx, 1);
				{
				setState(70);
				match(ELSE);
				setState(71);
				match(LLAVEIZQ);
				setState(72);
				block();
				setState(73);
				match(LLAVEDER);

				}
				break;
			case EOF:
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
			setState(90);
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
					setState(80);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(81);
					match(ELSE_IF);
					setState(82);
					expr(0);
					setState(83);
					match(LLAVEIZQ);
					setState(84);
					block();
					setState(85);
					match(LLAVEDER);

					}
					} 
				}
				setState(92);
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
			setState(107);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(93);
				reserv();
				setState(94);
				match(ID);
				setState(95);
				match(D_PTS);
				setState(96);
				types();
				setState(97);
				match(IG);
				setState(98);
				expr(0);

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(101);
				reserv();
				setState(102);
				match(ID);
				setState(103);
				match(IG);
				setState(104);
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
		enterRule(_localctx, 16, RULE_asignacionstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			match(ID);
			setState(110);
			ig_type();
			setState(111);
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
		enterRule(_localctx, 18, RULE_ig_type);
		try {
			setState(122);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IG:
				enterOuterAlt(_localctx, 1);
				{
				setState(114);
				match(IG);

				}
				break;
			case IG_IG:
				enterOuterAlt(_localctx, 2);
				{
				setState(116);
				match(IG_IG);

				}
				break;
			case MAS_IGUAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(118);
				match(MAS_IGUAL);

				}
				break;
			case MENOS_IGUAL:
				enterOuterAlt(_localctx, 4);
				{
				setState(120);
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
		enterRule(_localctx, 20, RULE_reserv);
		try {
			setState(128);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(124);
				match(VAR);

				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(126);
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
		enterRule(_localctx, 22, RULE_types);
		try {
			setState(143);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(130);
				match(INT);
				 _localctx.ty = environment.INTEGER 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(132);
				match(FLOAT);
				 _localctx.ty = environment.FLOAT 
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(134);
				match(STR);
				 _localctx.ty = environment.STRING 
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(136);
				match(BOOL);
				 _localctx.ty = environment.BOOLEAN 
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 5);
				{
				setState(138);
				match(CHAR);

				}
				break;
			case CORIZQ:
				enterOuterAlt(_localctx, 6);
				{
				setState(140);
				match(CORIZQ);
				setState(141);
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
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PARIZQ:
				{
				setState(146);
				match(PARIZQ);
				setState(147);
				((ExprContext)_localctx).expr = expr(0);
				setState(148);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case ID:
				{
				setState(151);
				((ExprContext)_localctx).list = listArray(0);
				 _localctx.e = ((ExprContext)_localctx).list.p
				}
				break;
			case CORIZQ:
				{
				setState(154);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(155);
				((ExprContext)_localctx).listParams = listParams(0);
				setState(156);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), ((ExprContext)_localctx).listParams.l) 
				}
				break;
			case NUMBER:
				{
				setState(159);
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
				setState(161);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case TRU:
				{
				setState(163);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case FAL:
				{
				setState(165);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(206);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(204);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(169);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(170);
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
						setState(171);
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
						setState(174);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(175);
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
						setState(176);
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
						setState(179);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(180);
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
						setState(181);
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
						setState(184);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(185);
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
						setState(186);
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
						setState(189);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(190);
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
						setState(191);
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
						setState(194);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(195);
						((ExprContext)_localctx).op = match(AND);
						setState(196);
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
						setState(199);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(200);
						((ExprContext)_localctx).op = match(OR);
						setState(201);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(9);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(208);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
		int _startState = 26;
		enterRecursionRule(_localctx, 26, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(210);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(220);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
					setState(213);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(214);
					match(COMA);
					setState(215);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(222);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
		int _startState = 28;
		enterRecursionRule(_localctx, 28, RULE_listArray, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(224);
			((ListArrayContext)_localctx).ID = match(ID);
			 _localctx.p = expressions.NewCallVar((((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getCharPositionInLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(235);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
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
					setState(227);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(228);
					match(CORIZQ);
					setState(229);
					((ListArrayContext)_localctx).expr = expr(0);
					setState(230);
					match(CORDER);
					 _localctx.p = expressions.NewArrayAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, ((ListArrayContext)_localctx).expr.e) 
					}
					} 
				}
				setState(237);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 6:
			return else_ifstmt_sempred((Else_ifstmtContext)_localctx, predIndex);
		case 12:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 13:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 14:
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
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 14);
		case 2:
			return precpred(_ctx, 13);
		case 3:
			return precpred(_ctx, 12);
		case 4:
			return precpred(_ctx, 11);
		case 5:
			return precpred(_ctx, 10);
		case 6:
			return precpred(_ctx, 9);
		case 7:
			return precpred(_ctx, 8);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 8:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listArray_sempred(ListArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 9:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3D\u00f1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\3\2\3\2\3\2\3\2\3\3"+
		"\6\3&\n\3\r\3\16\3\'\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\5\48\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7P\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\7\b[\n\b\f\b\16\b^\13\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\5\tn\n\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\5\13}\n\13\3\f\3\f\3\f\3\f\5\f\u0083\n\f\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u0092\n\r\3\16\3\16\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\5\16\u00aa\n\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\7\16\u00cf\n\16\f\16\16\16\u00d2\13\16\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\7\17\u00dd\n\17\f\17\16\17\u00e0\13\17\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\7\20\u00ec\n\20\f\20\16"+
		"\20\u00ef\13\20\3\20\2\6\16\32\34\36\21\2\4\6\b\n\f\16\20\22\24\26\30"+
		"\32\34\36\2\7\3\2\64\65\4\2\66\6688\4\2\60\60\62\62\4\2\61\61\63\63\3"+
		"\2\'(\2\u0100\2 \3\2\2\2\4%\3\2\2\2\6\67\3\2\2\2\b9\3\2\2\2\n?\3\2\2\2"+
		"\fO\3\2\2\2\16Q\3\2\2\2\20m\3\2\2\2\22o\3\2\2\2\24|\3\2\2\2\26\u0082\3"+
		"\2\2\2\30\u0091\3\2\2\2\32\u00a9\3\2\2\2\34\u00d3\3\2\2\2\36\u00e1\3\2"+
		"\2\2 !\5\4\3\2!\"\7\2\2\3\"#\b\2\1\2#\3\3\2\2\2$&\5\6\4\2%$\3\2\2\2&\'"+
		"\3\2\2\2\'%\3\2\2\2\'(\3\2\2\2()\3\2\2\2)*\b\3\1\2*\5\3\2\2\2+,\5\b\5"+
		"\2,-\b\4\1\2-8\3\2\2\2./\5\n\6\2/\60\b\4\1\2\608\3\2\2\2\61\62\5\20\t"+
		"\2\62\63\b\4\1\2\638\3\2\2\2\64\65\5\22\n\2\65\66\b\4\1\2\668\3\2\2\2"+
		"\67+\3\2\2\2\67.\3\2\2\2\67\61\3\2\2\2\67\64\3\2\2\28\7\3\2\2\29:\7\n"+
		"\2\2:;\79\2\2;<\5\32\16\2<=\7:\2\2=>\b\5\1\2>\t\3\2\2\2?@\7\13\2\2@A\5"+
		"\32\16\2AB\7;\2\2BC\5\4\3\2CD\7<\2\2DE\7\f\2\2EF\7\r\2\2FG\b\6\1\2G\13"+
		"\3\2\2\2HI\7\r\2\2IJ\7;\2\2JK\5\4\3\2KL\7<\2\2LM\b\7\1\2MP\3\2\2\2NP\3"+
		"\2\2\2OH\3\2\2\2ON\3\2\2\2P\r\3\2\2\2Q\\\b\b\1\2RS\f\4\2\2ST\7\f\2\2T"+
		"U\5\32\16\2UV\7;\2\2VW\5\4\3\2WX\7<\2\2XY\b\b\1\2Y[\3\2\2\2ZR\3\2\2\2"+
		"[^\3\2\2\2\\Z\3\2\2\2\\]\3\2\2\2]\17\3\2\2\2^\\\3\2\2\2_`\5\26\f\2`a\7"+
		"&\2\2ab\7=\2\2bc\5\30\r\2cd\7/\2\2de\5\32\16\2ef\b\t\1\2fn\3\2\2\2gh\5"+
		"\26\f\2hi\7&\2\2ij\7/\2\2jk\5\32\16\2kl\b\t\1\2ln\3\2\2\2m_\3\2\2\2mg"+
		"\3\2\2\2n\21\3\2\2\2op\7&\2\2pq\5\24\13\2qr\5\32\16\2rs\b\n\1\2s\23\3"+
		"\2\2\2tu\7/\2\2u}\b\13\1\2vw\7(\2\2w}\b\13\1\2xy\7)\2\2y}\b\13\1\2z{\7"+
		"*\2\2{}\b\13\1\2|t\3\2\2\2|v\3\2\2\2|x\3\2\2\2|z\3\2\2\2}\25\3\2\2\2~"+
		"\177\7\27\2\2\177\u0083\b\f\1\2\u0080\u0081\7\30\2\2\u0081\u0083\b\f\1"+
		"\2\u0082~\3\2\2\2\u0082\u0080\3\2\2\2\u0083\27\3\2\2\2\u0084\u0085\7\3"+
		"\2\2\u0085\u0092\b\r\1\2\u0086\u0087\7\4\2\2\u0087\u0092\b\r\1\2\u0088"+
		"\u0089\7\6\2\2\u0089\u0092\b\r\1\2\u008a\u008b\7\5\2\2\u008b\u0092\b\r"+
		"\1\2\u008c\u008d\7\7\2\2\u008d\u0092\b\r\1\2\u008e\u008f\7>\2\2\u008f"+
		"\u0090\7?\2\2\u0090\u0092\b\r\1\2\u0091\u0084\3\2\2\2\u0091\u0086\3\2"+
		"\2\2\u0091\u0088\3\2\2\2\u0091\u008a\3\2\2\2\u0091\u008c\3\2\2\2\u0091"+
		"\u008e\3\2\2\2\u0092\31\3\2\2\2\u0093\u0094\b\16\1\2\u0094\u0095\79\2"+
		"\2\u0095\u0096\5\32\16\2\u0096\u0097\7:\2\2\u0097\u0098\b\16\1\2\u0098"+
		"\u00aa\3\2\2\2\u0099\u009a\5\36\20\2\u009a\u009b\b\16\1\2\u009b\u00aa"+
		"\3\2\2\2\u009c\u009d\7>\2\2\u009d\u009e\5\34\17\2\u009e\u009f\7?\2\2\u009f"+
		"\u00a0\b\16\1\2\u00a0\u00aa\3\2\2\2\u00a1\u00a2\7$\2\2\u00a2\u00aa\b\16"+
		"\1\2\u00a3\u00a4\7%\2\2\u00a4\u00aa\b\16\1\2\u00a5\u00a6\7\b\2\2\u00a6"+
		"\u00aa\b\16\1\2\u00a7\u00a8\7\t\2\2\u00a8\u00aa\b\16\1\2\u00a9\u0093\3"+
		"\2\2\2\u00a9\u0099\3\2\2\2\u00a9\u009c\3\2\2\2\u00a9\u00a1\3\2\2\2\u00a9"+
		"\u00a3\3\2\2\2\u00a9\u00a5\3\2\2\2\u00a9\u00a7\3\2\2\2\u00aa\u00d0\3\2"+
		"\2\2\u00ab\u00ac\f\20\2\2\u00ac\u00ad\t\2\2\2\u00ad\u00ae\5\32\16\21\u00ae"+
		"\u00af\b\16\1\2\u00af\u00cf\3\2\2\2\u00b0\u00b1\f\17\2\2\u00b1\u00b2\t"+
		"\3\2\2\u00b2\u00b3\5\32\16\20\u00b3\u00b4\b\16\1\2\u00b4\u00cf\3\2\2\2"+
		"\u00b5\u00b6\f\16\2\2\u00b6\u00b7\t\4\2\2\u00b7\u00b8\5\32\16\17\u00b8"+
		"\u00b9\b\16\1\2\u00b9\u00cf\3\2\2\2\u00ba\u00bb\f\r\2\2\u00bb\u00bc\t"+
		"\5\2\2\u00bc\u00bd\5\32\16\16\u00bd\u00be\b\16\1\2\u00be\u00cf\3\2\2\2"+
		"\u00bf\u00c0\f\f\2\2\u00c0\u00c1\t\6\2\2\u00c1\u00c2\5\32\16\r\u00c2\u00c3"+
		"\b\16\1\2\u00c3\u00cf\3\2\2\2\u00c4\u00c5\f\13\2\2\u00c5\u00c6\7-\2\2"+
		"\u00c6\u00c7\5\32\16\f\u00c7\u00c8\b\16\1\2\u00c8\u00cf\3\2\2\2\u00c9"+
		"\u00ca\f\n\2\2\u00ca\u00cb\7,\2\2\u00cb\u00cc\5\32\16\13\u00cc\u00cd\b"+
		"\16\1\2\u00cd\u00cf\3\2\2\2\u00ce\u00ab\3\2\2\2\u00ce\u00b0\3\2\2\2\u00ce"+
		"\u00b5\3\2\2\2\u00ce\u00ba\3\2\2\2\u00ce\u00bf\3\2\2\2\u00ce\u00c4\3\2"+
		"\2\2\u00ce\u00c9\3\2\2\2\u00cf\u00d2\3\2\2\2\u00d0\u00ce\3\2\2\2\u00d0"+
		"\u00d1\3\2\2\2\u00d1\33\3\2\2\2\u00d2\u00d0\3\2\2\2\u00d3\u00d4\b\17\1"+
		"\2\u00d4\u00d5\5\32\16\2\u00d5\u00d6\b\17\1\2\u00d6\u00de\3\2\2\2\u00d7"+
		"\u00d8\f\4\2\2\u00d8\u00d9\7@\2\2\u00d9\u00da\5\32\16\2\u00da\u00db\b"+
		"\17\1\2\u00db\u00dd\3\2\2\2\u00dc\u00d7\3\2\2\2\u00dd\u00e0\3\2\2\2\u00de"+
		"\u00dc\3\2\2\2\u00de\u00df\3\2\2\2\u00df\35\3\2\2\2\u00e0\u00de\3\2\2"+
		"\2\u00e1\u00e2\b\20\1\2\u00e2\u00e3\7&\2\2\u00e3\u00e4\b\20\1\2\u00e4"+
		"\u00ed\3\2\2\2\u00e5\u00e6\f\4\2\2\u00e6\u00e7\7>\2\2\u00e7\u00e8\5\32"+
		"\16\2\u00e8\u00e9\7?\2\2\u00e9\u00ea\b\20\1\2\u00ea\u00ec\3\2\2\2\u00eb"+
		"\u00e5\3\2\2\2\u00ec\u00ef\3\2\2\2\u00ed\u00eb\3\2\2\2\u00ed\u00ee\3\2"+
		"\2\2\u00ee\37\3\2\2\2\u00ef\u00ed\3\2\2\2\17\'\67O\\m|\u0082\u0091\u00a9"+
		"\u00ce\u00d0\u00de\u00ed";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}