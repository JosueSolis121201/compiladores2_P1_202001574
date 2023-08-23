// Generated from c:\Users\josue\Desktop\compiladores2_P1_202001574\Server2\SwiftLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, BOOL=3, STR=4, CHAR=5, TRU=6, FAL=7, PRINT=8, IF=9, ELSE=10, 
		WHILE=11, FOR=12, IN=13, GUARD=14, CONTINUE=15, RETURN=16, SWITCH=17, 
		BREAK=18, CASE=19, VAR=20, LET=21, D_APPEND=22, D_REMOVE=23, REPEATING=24, 
		COUNT=25, STRUCT=26, MUTATING=27, FUNC=28, SELF_D=29, PUNTO=30, INOUT=31, 
		GUION_B=32, NUMBER=33, STRING=34, ID=35, DIF=36, IG_IG=37, MAS_IGUAL=38, 
		MENOS_IGUAL=39, NOT=40, OR=41, AND=42, SOLO_AND=43, IG=44, MAY_IG=45, 
		MEN_IG=46, MAYOR=47, MENOR=48, MUL=49, DIV=50, ADD=51, MODULO=52, SUB=53, 
		PARIZQ=54, PARDER=55, LLAVEIZQ=56, LLAVEDER=57, D_PTS=58, CORIZQ=59, CORDER=60, 
		COMA=61, PUNTO_COMA=62, WHITESPACE=63, COMMENT=64, LINE_COMMENT=65;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "STR", "CHAR", "TRU", "FAL", "PRINT", "IF", "ELSE", 
			"WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "SWITCH", "BREAK", 
			"CASE", "VAR", "LET", "D_APPEND", "D_REMOVE", "REPEATING", "COUNT", "STRUCT", 
			"MUTATING", "FUNC", "SELF_D", "PUNTO", "INOUT", "GUION_B", "NUMBER", 
			"STRING", "ID", "DIF", "IG_IG", "MAS_IGUAL", "MENOS_IGUAL", "NOT", "OR", 
			"AND", "SOLO_AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "MODULO", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"D_PTS", "CORIZQ", "CORDER", "COMA", "PUNTO_COMA", "WHITESPACE", "COMMENT", 
			"LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'", 
			"'false'", "'print'", "'if'", "'else'", "'while'", "'for'", "'in'", "'guard'", 
			"'continue'", "'return'", "'switch'", "'break'", "'case'", "'var'", "'Let'", 
			"'.append'", "'.remove'", "'repeating'", "'count'", "'STRUCT'", "'mutating'", 
			"'func'", "'self.'", "'.'", "'inout'", "'_'", null, null, null, "'!='", 
			"'=='", "'+='", "'-='", "'!'", "'||'", "'&&'", "'&'", "'='", "'>='", 
			"'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'%'", "'-'", "'('", "')'", 
			"'{'", "'}'", "':'", "'['", "']'", "','", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "CHAR", "TRU", "FAL", "PRINT", "IF", 
			"ELSE", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", "SWITCH", 
			"BREAK", "CASE", "VAR", "LET", "D_APPEND", "D_REMOVE", "REPEATING", "COUNT", 
			"STRUCT", "MUTATING", "FUNC", "SELF_D", "PUNTO", "INOUT", "GUION_B", 
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


	public SwiftLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SwiftLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2C\u01c0\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26"+
		"\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\36\3\36"+
		"\3\36\3\36\3\36\3\36\3\37\3\37\3 \3 \3 \3 \3 \3 \3!\3!\3\"\6\"\u0144\n"+
		"\"\r\"\16\"\u0145\3\"\3\"\6\"\u014a\n\"\r\"\16\"\u014b\5\"\u014e\n\"\3"+
		"#\3#\7#\u0152\n#\f#\16#\u0155\13#\3#\3#\3$\3$\7$\u015b\n$\f$\16$\u015e"+
		"\13$\3%\3%\3%\3&\3&\3&\3\'\3\'\3\'\3(\3(\3(\3)\3)\3*\3*\3*\3+\3+\3+\3"+
		",\3,\3-\3-\3.\3.\3.\3/\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3"+
		"\64\3\64\3\65\3\65\3\66\3\66\3\67\3\67\38\38\39\39\3:\3:\3;\3;\3<\3<\3"+
		"=\3=\3>\3>\3?\3?\3@\6@\u019f\n@\r@\16@\u01a0\3@\3@\3A\3A\3A\3A\7A\u01a9"+
		"\nA\fA\16A\u01ac\13A\3A\3A\3A\3A\3A\3B\3B\3B\3B\7B\u01b7\nB\fB\16B\u01ba"+
		"\13B\3B\3B\3C\3C\3C\3\u01aa\2D\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61"+
		"a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085\2"+
		"\3\2\t\3\2\62;\3\2$$\4\2C\\c|\6\2\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2"+
		"\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u01c6\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3"+
		"\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2"+
		"\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35"+
		"\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)"+
		"\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2"+
		"\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2"+
		"A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3"+
		"\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2"+
		"\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2"+
		"g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3"+
		"\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3"+
		"\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\3\u0087\3\2\2\2\5\u008b\3\2\2\2"+
		"\7\u0091\3\2\2\2\t\u0096\3\2\2\2\13\u009d\3\2\2\2\r\u00a7\3\2\2\2\17\u00ac"+
		"\3\2\2\2\21\u00b2\3\2\2\2\23\u00b8\3\2\2\2\25\u00bb\3\2\2\2\27\u00c0\3"+
		"\2\2\2\31\u00c6\3\2\2\2\33\u00ca\3\2\2\2\35\u00cd\3\2\2\2\37\u00d3\3\2"+
		"\2\2!\u00dc\3\2\2\2#\u00e3\3\2\2\2%\u00ea\3\2\2\2\'\u00f0\3\2\2\2)\u00f5"+
		"\3\2\2\2+\u00f9\3\2\2\2-\u00fd\3\2\2\2/\u0105\3\2\2\2\61\u010d\3\2\2\2"+
		"\63\u0117\3\2\2\2\65\u011d\3\2\2\2\67\u0124\3\2\2\29\u012d\3\2\2\2;\u0132"+
		"\3\2\2\2=\u0138\3\2\2\2?\u013a\3\2\2\2A\u0140\3\2\2\2C\u0143\3\2\2\2E"+
		"\u014f\3\2\2\2G\u0158\3\2\2\2I\u015f\3\2\2\2K\u0162\3\2\2\2M\u0165\3\2"+
		"\2\2O\u0168\3\2\2\2Q\u016b\3\2\2\2S\u016d\3\2\2\2U\u0170\3\2\2\2W\u0173"+
		"\3\2\2\2Y\u0175\3\2\2\2[\u0177\3\2\2\2]\u017a\3\2\2\2_\u017d\3\2\2\2a"+
		"\u017f\3\2\2\2c\u0181\3\2\2\2e\u0183\3\2\2\2g\u0185\3\2\2\2i\u0187\3\2"+
		"\2\2k\u0189\3\2\2\2m\u018b\3\2\2\2o\u018d\3\2\2\2q\u018f\3\2\2\2s\u0191"+
		"\3\2\2\2u\u0193\3\2\2\2w\u0195\3\2\2\2y\u0197\3\2\2\2{\u0199\3\2\2\2}"+
		"\u019b\3\2\2\2\177\u019e\3\2\2\2\u0081\u01a4\3\2\2\2\u0083\u01b2\3\2\2"+
		"\2\u0085\u01bd\3\2\2\2\u0087\u0088\7K\2\2\u0088\u0089\7p\2\2\u0089\u008a"+
		"\7v\2\2\u008a\4\3\2\2\2\u008b\u008c\7H\2\2\u008c\u008d\7n\2\2\u008d\u008e"+
		"\7q\2\2\u008e\u008f\7c\2\2\u008f\u0090\7v\2\2\u0090\6\3\2\2\2\u0091\u0092"+
		"\7D\2\2\u0092\u0093\7q\2\2\u0093\u0094\7q\2\2\u0094\u0095\7n\2\2\u0095"+
		"\b\3\2\2\2\u0096\u0097\7U\2\2\u0097\u0098\7v\2\2\u0098\u0099\7t\2\2\u0099"+
		"\u009a\7k\2\2\u009a\u009b\7p\2\2\u009b\u009c\7i\2\2\u009c\n\3\2\2\2\u009d"+
		"\u009e\7E\2\2\u009e\u009f\7j\2\2\u009f\u00a0\7c\2\2\u00a0\u00a1\7t\2\2"+
		"\u00a1\u00a2\7c\2\2\u00a2\u00a3\7e\2\2\u00a3\u00a4\7v\2\2\u00a4\u00a5"+
		"\7g\2\2\u00a5\u00a6\7t\2\2\u00a6\f\3\2\2\2\u00a7\u00a8\7v\2\2\u00a8\u00a9"+
		"\7t\2\2\u00a9\u00aa\7w\2\2\u00aa\u00ab\7g\2\2\u00ab\16\3\2\2\2\u00ac\u00ad"+
		"\7h\2\2\u00ad\u00ae\7c\2\2\u00ae\u00af\7n\2\2\u00af\u00b0\7u\2\2\u00b0"+
		"\u00b1\7g\2\2\u00b1\20\3\2\2\2\u00b2\u00b3\7r\2\2\u00b3\u00b4\7t\2\2\u00b4"+
		"\u00b5\7k\2\2\u00b5\u00b6\7p\2\2\u00b6\u00b7\7v\2\2\u00b7\22\3\2\2\2\u00b8"+
		"\u00b9\7k\2\2\u00b9\u00ba\7h\2\2\u00ba\24\3\2\2\2\u00bb\u00bc\7g\2\2\u00bc"+
		"\u00bd\7n\2\2\u00bd\u00be\7u\2\2\u00be\u00bf\7g\2\2\u00bf\26\3\2\2\2\u00c0"+
		"\u00c1\7y\2\2\u00c1\u00c2\7j\2\2\u00c2\u00c3\7k\2\2\u00c3\u00c4\7n\2\2"+
		"\u00c4\u00c5\7g\2\2\u00c5\30\3\2\2\2\u00c6\u00c7\7h\2\2\u00c7\u00c8\7"+
		"q\2\2\u00c8\u00c9\7t\2\2\u00c9\32\3\2\2\2\u00ca\u00cb\7k\2\2\u00cb\u00cc"+
		"\7p\2\2\u00cc\34\3\2\2\2\u00cd\u00ce\7i\2\2\u00ce\u00cf\7w\2\2\u00cf\u00d0"+
		"\7c\2\2\u00d0\u00d1\7t\2\2\u00d1\u00d2\7f\2\2\u00d2\36\3\2\2\2\u00d3\u00d4"+
		"\7e\2\2\u00d4\u00d5\7q\2\2\u00d5\u00d6\7p\2\2\u00d6\u00d7\7v\2\2\u00d7"+
		"\u00d8\7k\2\2\u00d8\u00d9\7p\2\2\u00d9\u00da\7w\2\2\u00da\u00db\7g\2\2"+
		"\u00db \3\2\2\2\u00dc\u00dd\7t\2\2\u00dd\u00de\7g\2\2\u00de\u00df\7v\2"+
		"\2\u00df\u00e0\7w\2\2\u00e0\u00e1\7t\2\2\u00e1\u00e2\7p\2\2\u00e2\"\3"+
		"\2\2\2\u00e3\u00e4\7u\2\2\u00e4\u00e5\7y\2\2\u00e5\u00e6\7k\2\2\u00e6"+
		"\u00e7\7v\2\2\u00e7\u00e8\7e\2\2\u00e8\u00e9\7j\2\2\u00e9$\3\2\2\2\u00ea"+
		"\u00eb\7d\2\2\u00eb\u00ec\7t\2\2\u00ec\u00ed\7g\2\2\u00ed\u00ee\7c\2\2"+
		"\u00ee\u00ef\7m\2\2\u00ef&\3\2\2\2\u00f0\u00f1\7e\2\2\u00f1\u00f2\7c\2"+
		"\2\u00f2\u00f3\7u\2\2\u00f3\u00f4\7g\2\2\u00f4(\3\2\2\2\u00f5\u00f6\7"+
		"x\2\2\u00f6\u00f7\7c\2\2\u00f7\u00f8\7t\2\2\u00f8*\3\2\2\2\u00f9\u00fa"+
		"\7N\2\2\u00fa\u00fb\7g\2\2\u00fb\u00fc\7v\2\2\u00fc,\3\2\2\2\u00fd\u00fe"+
		"\7\60\2\2\u00fe\u00ff\7c\2\2\u00ff\u0100\7r\2\2\u0100\u0101\7r\2\2\u0101"+
		"\u0102\7g\2\2\u0102\u0103\7p\2\2\u0103\u0104\7f\2\2\u0104.\3\2\2\2\u0105"+
		"\u0106\7\60\2\2\u0106\u0107\7t\2\2\u0107\u0108\7g\2\2\u0108\u0109\7o\2"+
		"\2\u0109\u010a\7q\2\2\u010a\u010b\7x\2\2\u010b\u010c\7g\2\2\u010c\60\3"+
		"\2\2\2\u010d\u010e\7t\2\2\u010e\u010f\7g\2\2\u010f\u0110\7r\2\2\u0110"+
		"\u0111\7g\2\2\u0111\u0112\7c\2\2\u0112\u0113\7v\2\2\u0113\u0114\7k\2\2"+
		"\u0114\u0115\7p\2\2\u0115\u0116\7i\2\2\u0116\62\3\2\2\2\u0117\u0118\7"+
		"e\2\2\u0118\u0119\7q\2\2\u0119\u011a\7w\2\2\u011a\u011b\7p\2\2\u011b\u011c"+
		"\7v\2\2\u011c\64\3\2\2\2\u011d\u011e\7U\2\2\u011e\u011f\7V\2\2\u011f\u0120"+
		"\7T\2\2\u0120\u0121\7W\2\2\u0121\u0122\7E\2\2\u0122\u0123\7V\2\2\u0123"+
		"\66\3\2\2\2\u0124\u0125\7o\2\2\u0125\u0126\7w\2\2\u0126\u0127\7v\2\2\u0127"+
		"\u0128\7c\2\2\u0128\u0129\7v\2\2\u0129\u012a\7k\2\2\u012a\u012b\7p\2\2"+
		"\u012b\u012c\7i\2\2\u012c8\3\2\2\2\u012d\u012e\7h\2\2\u012e\u012f\7w\2"+
		"\2\u012f\u0130\7p\2\2\u0130\u0131\7e\2\2\u0131:\3\2\2\2\u0132\u0133\7"+
		"u\2\2\u0133\u0134\7g\2\2\u0134\u0135\7n\2\2\u0135\u0136\7h\2\2\u0136\u0137"+
		"\7\60\2\2\u0137<\3\2\2\2\u0138\u0139\7\60\2\2\u0139>\3\2\2\2\u013a\u013b"+
		"\7k\2\2\u013b\u013c\7p\2\2\u013c\u013d\7q\2\2\u013d\u013e\7w\2\2\u013e"+
		"\u013f\7v\2\2\u013f@\3\2\2\2\u0140\u0141\7a\2\2\u0141B\3\2\2\2\u0142\u0144"+
		"\t\2\2\2\u0143\u0142\3\2\2\2\u0144\u0145\3\2\2\2\u0145\u0143\3\2\2\2\u0145"+
		"\u0146\3\2\2\2\u0146\u014d\3\2\2\2\u0147\u0149\7\60\2\2\u0148\u014a\t"+
		"\2\2\2\u0149\u0148\3\2\2\2\u014a\u014b\3\2\2\2\u014b\u0149\3\2\2\2\u014b"+
		"\u014c\3\2\2\2\u014c\u014e\3\2\2\2\u014d\u0147\3\2\2\2\u014d\u014e\3\2"+
		"\2\2\u014eD\3\2\2\2\u014f\u0153\7$\2\2\u0150\u0152\n\3\2\2\u0151\u0150"+
		"\3\2\2\2\u0152\u0155\3\2\2\2\u0153\u0151\3\2\2\2\u0153\u0154\3\2\2\2\u0154"+
		"\u0156\3\2\2\2\u0155\u0153\3\2\2\2\u0156\u0157\7$\2\2\u0157F\3\2\2\2\u0158"+
		"\u015c\t\4\2\2\u0159\u015b\t\5\2\2\u015a\u0159\3\2\2\2\u015b\u015e\3\2"+
		"\2\2\u015c\u015a\3\2\2\2\u015c\u015d\3\2\2\2\u015dH\3\2\2\2\u015e\u015c"+
		"\3\2\2\2\u015f\u0160\7#\2\2\u0160\u0161\7?\2\2\u0161J\3\2\2\2\u0162\u0163"+
		"\7?\2\2\u0163\u0164\7?\2\2\u0164L\3\2\2\2\u0165\u0166\7-\2\2\u0166\u0167"+
		"\7?\2\2\u0167N\3\2\2\2\u0168\u0169\7/\2\2\u0169\u016a\7?\2\2\u016aP\3"+
		"\2\2\2\u016b\u016c\7#\2\2\u016cR\3\2\2\2\u016d\u016e\7~\2\2\u016e\u016f"+
		"\7~\2\2\u016fT\3\2\2\2\u0170\u0171\7(\2\2\u0171\u0172\7(\2\2\u0172V\3"+
		"\2\2\2\u0173\u0174\7(\2\2\u0174X\3\2\2\2\u0175\u0176\7?\2\2\u0176Z\3\2"+
		"\2\2\u0177\u0178\7@\2\2\u0178\u0179\7?\2\2\u0179\\\3\2\2\2\u017a\u017b"+
		"\7>\2\2\u017b\u017c\7?\2\2\u017c^\3\2\2\2\u017d\u017e\7@\2\2\u017e`\3"+
		"\2\2\2\u017f\u0180\7>\2\2\u0180b\3\2\2\2\u0181\u0182\7,\2\2\u0182d\3\2"+
		"\2\2\u0183\u0184\7\61\2\2\u0184f\3\2\2\2\u0185\u0186\7-\2\2\u0186h\3\2"+
		"\2\2\u0187\u0188\7\'\2\2\u0188j\3\2\2\2\u0189\u018a\7/\2\2\u018al\3\2"+
		"\2\2\u018b\u018c\7*\2\2\u018cn\3\2\2\2\u018d\u018e\7+\2\2\u018ep\3\2\2"+
		"\2\u018f\u0190\7}\2\2\u0190r\3\2\2\2\u0191\u0192\7\177\2\2\u0192t\3\2"+
		"\2\2\u0193\u0194\7<\2\2\u0194v\3\2\2\2\u0195\u0196\7]\2\2\u0196x\3\2\2"+
		"\2\u0197\u0198\7_\2\2\u0198z\3\2\2\2\u0199\u019a\7.\2\2\u019a|\3\2\2\2"+
		"\u019b\u019c\7=\2\2\u019c~\3\2\2\2\u019d\u019f\t\6\2\2\u019e\u019d\3\2"+
		"\2\2\u019f\u01a0\3\2\2\2\u01a0\u019e\3\2\2\2\u01a0\u01a1\3\2\2\2\u01a1"+
		"\u01a2\3\2\2\2\u01a2\u01a3\b@\2\2\u01a3\u0080\3\2\2\2\u01a4\u01a5\7\61"+
		"\2\2\u01a5\u01a6\7,\2\2\u01a6\u01aa\3\2\2\2\u01a7\u01a9\13\2\2\2\u01a8"+
		"\u01a7\3\2\2\2\u01a9\u01ac\3\2\2\2\u01aa\u01ab\3\2\2\2\u01aa\u01a8\3\2"+
		"\2\2\u01ab\u01ad\3\2\2\2\u01ac\u01aa\3\2\2\2\u01ad\u01ae\7,\2\2\u01ae"+
		"\u01af\7\61\2\2\u01af\u01b0\3\2\2\2\u01b0\u01b1\bA\2\2\u01b1\u0082\3\2"+
		"\2\2\u01b2\u01b3\7\61\2\2\u01b3\u01b4\7\61\2\2\u01b4\u01b8\3\2\2\2\u01b5"+
		"\u01b7\n\7\2\2\u01b6\u01b5\3\2\2\2\u01b7\u01ba\3\2\2\2\u01b8\u01b6\3\2"+
		"\2\2\u01b8\u01b9\3\2\2\2\u01b9\u01bb\3\2\2\2\u01ba\u01b8\3\2\2\2\u01bb"+
		"\u01bc\bB\2\2\u01bc\u0084\3\2\2\2\u01bd\u01be\7^\2\2\u01be\u01bf\t\b\2"+
		"\2\u01bf\u0086\3\2\2\2\13\2\u0145\u014b\u014d\u0153\u015c\u01a0\u01aa"+
		"\u01b8\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}