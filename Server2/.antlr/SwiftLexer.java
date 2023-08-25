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
		INT=1, FLOAT=2, BOOL=3, STR=4, CHAR=5, AT_DD=6, TRU=7, FAL=8, PRINT=9, 
		IF=10, ELSE=11, ELSE_IF=12, WHILE=13, FOR=14, IN=15, GUARD=16, CONTINUE=17, 
		RETURN=18, SWITCH=19, BREAK=20, CASE=21, VAR=22, LET=23, D_APPEND=24, 
		D_REMOVE=25, D_REMOVELAST=26, D_ISEMPTY=27, D_COUNT=28, REPEATING=29, 
		DEFAULT=30, COUNT=31, STRUCT=32, MUTATING=33, FUNC=34, SELF_D=35, PUNTO=36, 
		INOUT=37, GUION_B=38, NUMBER=39, STRING=40, ID=41, DIF=42, IG_IG=43, MAS_IGUAL=44, 
		MENOS_IGUAL=45, NOT=46, OR=47, AND=48, SOLO_AND=49, IG=50, MAY_IG=51, 
		MEN_IG=52, MAYOR=53, MENOR=54, MUL=55, DIV=56, ADD=57, MODULO=58, SUB=59, 
		PARIZQ=60, PARDER=61, LLAVEIZQ=62, LLAVEDER=63, D_PTS=64, CORIZQ=65, CORDER=66, 
		COMA=67, PUNTO_COMA=68, WHITESPACE=69, COMMENT=70, LINE_COMMENT=71;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "STR", "CHAR", "AT_DD", "TRU", "FAL", "PRINT", 
			"IF", "ELSE", "ELSE_IF", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", 
			"SWITCH", "BREAK", "CASE", "VAR", "LET", "D_APPEND", "D_REMOVE", "D_REMOVELAST", 
			"D_ISEMPTY", "D_COUNT", "REPEATING", "DEFAULT", "COUNT", "STRUCT", "MUTATING", 
			"FUNC", "SELF_D", "PUNTO", "INOUT", "GUION_B", "NUMBER", "STRING", "ID", 
			"DIF", "IG_IG", "MAS_IGUAL", "MENOS_IGUAL", "NOT", "OR", "AND", "SOLO_AND", 
			"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "MODULO", 
			"SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", 
			"CORDER", "COMA", "PUNTO_COMA", "WHITESPACE", "COMMENT", "LINE_COMMENT", 
			"ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'at:'", 
			"'true'", "'false'", "'print'", "'if'", "'else'", "'else if'", "'while'", 
			"'for'", "'in'", "'guard'", "'continue'", "'return'", "'switch'", "'break'", 
			"'case'", "'var'", "'let'", "'.append'", "'.remove'", "'.removelast'", 
			"'.IsEmpty'", "'.count'", "'repeating'", "'default'", "'count'", "'STRUCT'", 
			"'mutating'", "'func'", "'self.'", "'.'", "'inout'", "'_'", null, null, 
			null, "'!='", "'=='", "'+='", "'-='", "'!'", "'||'", "'&&'", "'&'", "'='", 
			"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'%'", "'-'", "'('", 
			"')'", "'{'", "'}'", "':'", "'['", "']'", "','", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "CHAR", "AT_DD", "TRU", "FAL", "PRINT", 
			"IF", "ELSE", "ELSE_IF", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "RETURN", 
			"SWITCH", "BREAK", "CASE", "VAR", "LET", "D_APPEND", "D_REMOVE", "D_REMOVELAST", 
			"D_ISEMPTY", "D_COUNT", "REPEATING", "DEFAULT", "COUNT", "STRUCT", "MUTATING", 
			"FUNC", "SELF_D", "PUNTO", "INOUT", "GUION_B", "NUMBER", "STRING", "ID", 
			"DIF", "IG_IG", "MAS_IGUAL", "MENOS_IGUAL", "NOT", "OR", "AND", "SOLO_AND", 
			"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "MODULO", 
			"SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", 
			"CORDER", "COMA", "PUNTO_COMA", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2I\u01fc\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7"+
		"\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r"+
		"\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\20\3\20\3\20"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\27\3\27"+
		"\3\27\3\27\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36"+
		"\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 "+
		"\3 \3 \3 \3!\3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3#"+
		"\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3%\3%\3&\3&\3&\3&\3&\3&\3\'\3\'\3(\6(\u0180"+
		"\n(\r(\16(\u0181\3(\3(\6(\u0186\n(\r(\16(\u0187\5(\u018a\n(\3)\3)\7)\u018e"+
		"\n)\f)\16)\u0191\13)\3)\3)\3*\3*\7*\u0197\n*\f*\16*\u019a\13*\3+\3+\3"+
		"+\3,\3,\3,\3-\3-\3-\3.\3.\3.\3/\3/\3\60\3\60\3\60\3\61\3\61\3\61\3\62"+
		"\3\62\3\63\3\63\3\64\3\64\3\64\3\65\3\65\3\65\3\66\3\66\3\67\3\67\38\3"+
		"8\39\39\3:\3:\3;\3;\3<\3<\3=\3=\3>\3>\3?\3?\3@\3@\3A\3A\3B\3B\3C\3C\3"+
		"D\3D\3E\3E\3F\6F\u01db\nF\rF\16F\u01dc\3F\3F\3G\3G\3G\3G\7G\u01e5\nG\f"+
		"G\16G\u01e8\13G\3G\3G\3G\3G\3G\3H\3H\3H\3H\7H\u01f3\nH\fH\16H\u01f6\13"+
		"H\3H\3H\3I\3I\3I\3\u01e6\2J\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25"+
		"\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32"+
		"\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a"+
		"\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087"+
		"E\u0089F\u008bG\u008dH\u008fI\u0091\2\3\2\t\3\2\62;\3\2$$\4\2C\\c|\6\2"+
		"\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2"+
		"\u0202\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2"+
		"\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3"+
		"\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2"+
		"\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2"+
		"/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2"+
		"\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2"+
		"G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3"+
		"\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2"+
		"\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2"+
		"m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3"+
		"\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2"+
		"\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2"+
		"\u008d\3\2\2\2\2\u008f\3\2\2\2\3\u0093\3\2\2\2\5\u0097\3\2\2\2\7\u009d"+
		"\3\2\2\2\t\u00a2\3\2\2\2\13\u00a9\3\2\2\2\r\u00b3\3\2\2\2\17\u00b7\3\2"+
		"\2\2\21\u00bc\3\2\2\2\23\u00c2\3\2\2\2\25\u00c8\3\2\2\2\27\u00cb\3\2\2"+
		"\2\31\u00d0\3\2\2\2\33\u00d8\3\2\2\2\35\u00de\3\2\2\2\37\u00e2\3\2\2\2"+
		"!\u00e5\3\2\2\2#\u00eb\3\2\2\2%\u00f4\3\2\2\2\'\u00fb\3\2\2\2)\u0102\3"+
		"\2\2\2+\u0108\3\2\2\2-\u010d\3\2\2\2/\u0111\3\2\2\2\61\u0115\3\2\2\2\63"+
		"\u011d\3\2\2\2\65\u0125\3\2\2\2\67\u0131\3\2\2\29\u013a\3\2\2\2;\u0141"+
		"\3\2\2\2=\u014b\3\2\2\2?\u0153\3\2\2\2A\u0159\3\2\2\2C\u0160\3\2\2\2E"+
		"\u0169\3\2\2\2G\u016e\3\2\2\2I\u0174\3\2\2\2K\u0176\3\2\2\2M\u017c\3\2"+
		"\2\2O\u017f\3\2\2\2Q\u018b\3\2\2\2S\u0194\3\2\2\2U\u019b\3\2\2\2W\u019e"+
		"\3\2\2\2Y\u01a1\3\2\2\2[\u01a4\3\2\2\2]\u01a7\3\2\2\2_\u01a9\3\2\2\2a"+
		"\u01ac\3\2\2\2c\u01af\3\2\2\2e\u01b1\3\2\2\2g\u01b3\3\2\2\2i\u01b6\3\2"+
		"\2\2k\u01b9\3\2\2\2m\u01bb\3\2\2\2o\u01bd\3\2\2\2q\u01bf\3\2\2\2s\u01c1"+
		"\3\2\2\2u\u01c3\3\2\2\2w\u01c5\3\2\2\2y\u01c7\3\2\2\2{\u01c9\3\2\2\2}"+
		"\u01cb\3\2\2\2\177\u01cd\3\2\2\2\u0081\u01cf\3\2\2\2\u0083\u01d1\3\2\2"+
		"\2\u0085\u01d3\3\2\2\2\u0087\u01d5\3\2\2\2\u0089\u01d7\3\2\2\2\u008b\u01da"+
		"\3\2\2\2\u008d\u01e0\3\2\2\2\u008f\u01ee\3\2\2\2\u0091\u01f9\3\2\2\2\u0093"+
		"\u0094\7K\2\2\u0094\u0095\7p\2\2\u0095\u0096\7v\2\2\u0096\4\3\2\2\2\u0097"+
		"\u0098\7H\2\2\u0098\u0099\7n\2\2\u0099\u009a\7q\2\2\u009a\u009b\7c\2\2"+
		"\u009b\u009c\7v\2\2\u009c\6\3\2\2\2\u009d\u009e\7D\2\2\u009e\u009f\7q"+
		"\2\2\u009f\u00a0\7q\2\2\u00a0\u00a1\7n\2\2\u00a1\b\3\2\2\2\u00a2\u00a3"+
		"\7U\2\2\u00a3\u00a4\7v\2\2\u00a4\u00a5\7t\2\2\u00a5\u00a6\7k\2\2\u00a6"+
		"\u00a7\7p\2\2\u00a7\u00a8\7i\2\2\u00a8\n\3\2\2\2\u00a9\u00aa\7E\2\2\u00aa"+
		"\u00ab\7j\2\2\u00ab\u00ac\7c\2\2\u00ac\u00ad\7t\2\2\u00ad\u00ae\7c\2\2"+
		"\u00ae\u00af\7e\2\2\u00af\u00b0\7v\2\2\u00b0\u00b1\7g\2\2\u00b1\u00b2"+
		"\7t\2\2\u00b2\f\3\2\2\2\u00b3\u00b4\7c\2\2\u00b4\u00b5\7v\2\2\u00b5\u00b6"+
		"\7<\2\2\u00b6\16\3\2\2\2\u00b7\u00b8\7v\2\2\u00b8\u00b9\7t\2\2\u00b9\u00ba"+
		"\7w\2\2\u00ba\u00bb\7g\2\2\u00bb\20\3\2\2\2\u00bc\u00bd\7h\2\2\u00bd\u00be"+
		"\7c\2\2\u00be\u00bf\7n\2\2\u00bf\u00c0\7u\2\2\u00c0\u00c1\7g\2\2\u00c1"+
		"\22\3\2\2\2\u00c2\u00c3\7r\2\2\u00c3\u00c4\7t\2\2\u00c4\u00c5\7k\2\2\u00c5"+
		"\u00c6\7p\2\2\u00c6\u00c7\7v\2\2\u00c7\24\3\2\2\2\u00c8\u00c9\7k\2\2\u00c9"+
		"\u00ca\7h\2\2\u00ca\26\3\2\2\2\u00cb\u00cc\7g\2\2\u00cc\u00cd\7n\2\2\u00cd"+
		"\u00ce\7u\2\2\u00ce\u00cf\7g\2\2\u00cf\30\3\2\2\2\u00d0\u00d1\7g\2\2\u00d1"+
		"\u00d2\7n\2\2\u00d2\u00d3\7u\2\2\u00d3\u00d4\7g\2\2\u00d4\u00d5\7\"\2"+
		"\2\u00d5\u00d6\7k\2\2\u00d6\u00d7\7h\2\2\u00d7\32\3\2\2\2\u00d8\u00d9"+
		"\7y\2\2\u00d9\u00da\7j\2\2\u00da\u00db\7k\2\2\u00db\u00dc\7n\2\2\u00dc"+
		"\u00dd\7g\2\2\u00dd\34\3\2\2\2\u00de\u00df\7h\2\2\u00df\u00e0\7q\2\2\u00e0"+
		"\u00e1\7t\2\2\u00e1\36\3\2\2\2\u00e2\u00e3\7k\2\2\u00e3\u00e4\7p\2\2\u00e4"+
		" \3\2\2\2\u00e5\u00e6\7i\2\2\u00e6\u00e7\7w\2\2\u00e7\u00e8\7c\2\2\u00e8"+
		"\u00e9\7t\2\2\u00e9\u00ea\7f\2\2\u00ea\"\3\2\2\2\u00eb\u00ec\7e\2\2\u00ec"+
		"\u00ed\7q\2\2\u00ed\u00ee\7p\2\2\u00ee\u00ef\7v\2\2\u00ef\u00f0\7k\2\2"+
		"\u00f0\u00f1\7p\2\2\u00f1\u00f2\7w\2\2\u00f2\u00f3\7g\2\2\u00f3$\3\2\2"+
		"\2\u00f4\u00f5\7t\2\2\u00f5\u00f6\7g\2\2\u00f6\u00f7\7v\2\2\u00f7\u00f8"+
		"\7w\2\2\u00f8\u00f9\7t\2\2\u00f9\u00fa\7p\2\2\u00fa&\3\2\2\2\u00fb\u00fc"+
		"\7u\2\2\u00fc\u00fd\7y\2\2\u00fd\u00fe\7k\2\2\u00fe\u00ff\7v\2\2\u00ff"+
		"\u0100\7e\2\2\u0100\u0101\7j\2\2\u0101(\3\2\2\2\u0102\u0103\7d\2\2\u0103"+
		"\u0104\7t\2\2\u0104\u0105\7g\2\2\u0105\u0106\7c\2\2\u0106\u0107\7m\2\2"+
		"\u0107*\3\2\2\2\u0108\u0109\7e\2\2\u0109\u010a\7c\2\2\u010a\u010b\7u\2"+
		"\2\u010b\u010c\7g\2\2\u010c,\3\2\2\2\u010d\u010e\7x\2\2\u010e\u010f\7"+
		"c\2\2\u010f\u0110\7t\2\2\u0110.\3\2\2\2\u0111\u0112\7n\2\2\u0112\u0113"+
		"\7g\2\2\u0113\u0114\7v\2\2\u0114\60\3\2\2\2\u0115\u0116\7\60\2\2\u0116"+
		"\u0117\7c\2\2\u0117\u0118\7r\2\2\u0118\u0119\7r\2\2\u0119\u011a\7g\2\2"+
		"\u011a\u011b\7p\2\2\u011b\u011c\7f\2\2\u011c\62\3\2\2\2\u011d\u011e\7"+
		"\60\2\2\u011e\u011f\7t\2\2\u011f\u0120\7g\2\2\u0120\u0121\7o\2\2\u0121"+
		"\u0122\7q\2\2\u0122\u0123\7x\2\2\u0123\u0124\7g\2\2\u0124\64\3\2\2\2\u0125"+
		"\u0126\7\60\2\2\u0126\u0127\7t\2\2\u0127\u0128\7g\2\2\u0128\u0129\7o\2"+
		"\2\u0129\u012a\7q\2\2\u012a\u012b\7x\2\2\u012b\u012c\7g\2\2\u012c\u012d"+
		"\7n\2\2\u012d\u012e\7c\2\2\u012e\u012f\7u\2\2\u012f\u0130\7v\2\2\u0130"+
		"\66\3\2\2\2\u0131\u0132\7\60\2\2\u0132\u0133\7K\2\2\u0133\u0134\7u\2\2"+
		"\u0134\u0135\7G\2\2\u0135\u0136\7o\2\2\u0136\u0137\7r\2\2\u0137\u0138"+
		"\7v\2\2\u0138\u0139\7{\2\2\u01398\3\2\2\2\u013a\u013b\7\60\2\2\u013b\u013c"+
		"\7e\2\2\u013c\u013d\7q\2\2\u013d\u013e\7w\2\2\u013e\u013f\7p\2\2\u013f"+
		"\u0140\7v\2\2\u0140:\3\2\2\2\u0141\u0142\7t\2\2\u0142\u0143\7g\2\2\u0143"+
		"\u0144\7r\2\2\u0144\u0145\7g\2\2\u0145\u0146\7c\2\2\u0146\u0147\7v\2\2"+
		"\u0147\u0148\7k\2\2\u0148\u0149\7p\2\2\u0149\u014a\7i\2\2\u014a<\3\2\2"+
		"\2\u014b\u014c\7f\2\2\u014c\u014d\7g\2\2\u014d\u014e\7h\2\2\u014e\u014f"+
		"\7c\2\2\u014f\u0150\7w\2\2\u0150\u0151\7n\2\2\u0151\u0152\7v\2\2\u0152"+
		">\3\2\2\2\u0153\u0154\7e\2\2\u0154\u0155\7q\2\2\u0155\u0156\7w\2\2\u0156"+
		"\u0157\7p\2\2\u0157\u0158\7v\2\2\u0158@\3\2\2\2\u0159\u015a\7U\2\2\u015a"+
		"\u015b\7V\2\2\u015b\u015c\7T\2\2\u015c\u015d\7W\2\2\u015d\u015e\7E\2\2"+
		"\u015e\u015f\7V\2\2\u015fB\3\2\2\2\u0160\u0161\7o\2\2\u0161\u0162\7w\2"+
		"\2\u0162\u0163\7v\2\2\u0163\u0164\7c\2\2\u0164\u0165\7v\2\2\u0165\u0166"+
		"\7k\2\2\u0166\u0167\7p\2\2\u0167\u0168\7i\2\2\u0168D\3\2\2\2\u0169\u016a"+
		"\7h\2\2\u016a\u016b\7w\2\2\u016b\u016c\7p\2\2\u016c\u016d\7e\2\2\u016d"+
		"F\3\2\2\2\u016e\u016f\7u\2\2\u016f\u0170\7g\2\2\u0170\u0171\7n\2\2\u0171"+
		"\u0172\7h\2\2\u0172\u0173\7\60\2\2\u0173H\3\2\2\2\u0174\u0175\7\60\2\2"+
		"\u0175J\3\2\2\2\u0176\u0177\7k\2\2\u0177\u0178\7p\2\2\u0178\u0179\7q\2"+
		"\2\u0179\u017a\7w\2\2\u017a\u017b\7v\2\2\u017bL\3\2\2\2\u017c\u017d\7"+
		"a\2\2\u017dN\3\2\2\2\u017e\u0180\t\2\2\2\u017f\u017e\3\2\2\2\u0180\u0181"+
		"\3\2\2\2\u0181\u017f\3\2\2\2\u0181\u0182\3\2\2\2\u0182\u0189\3\2\2\2\u0183"+
		"\u0185\7\60\2\2\u0184\u0186\t\2\2\2\u0185\u0184\3\2\2\2\u0186\u0187\3"+
		"\2\2\2\u0187\u0185\3\2\2\2\u0187\u0188\3\2\2\2\u0188\u018a\3\2\2\2\u0189"+
		"\u0183\3\2\2\2\u0189\u018a\3\2\2\2\u018aP\3\2\2\2\u018b\u018f\7$\2\2\u018c"+
		"\u018e\n\3\2\2\u018d\u018c\3\2\2\2\u018e\u0191\3\2\2\2\u018f\u018d\3\2"+
		"\2\2\u018f\u0190\3\2\2\2\u0190\u0192\3\2\2\2\u0191\u018f\3\2\2\2\u0192"+
		"\u0193\7$\2\2\u0193R\3\2\2\2\u0194\u0198\t\4\2\2\u0195\u0197\t\5\2\2\u0196"+
		"\u0195\3\2\2\2\u0197\u019a\3\2\2\2\u0198\u0196\3\2\2\2\u0198\u0199\3\2"+
		"\2\2\u0199T\3\2\2\2\u019a\u0198\3\2\2\2\u019b\u019c\7#\2\2\u019c\u019d"+
		"\7?\2\2\u019dV\3\2\2\2\u019e\u019f\7?\2\2\u019f\u01a0\7?\2\2\u01a0X\3"+
		"\2\2\2\u01a1\u01a2\7-\2\2\u01a2\u01a3\7?\2\2\u01a3Z\3\2\2\2\u01a4\u01a5"+
		"\7/\2\2\u01a5\u01a6\7?\2\2\u01a6\\\3\2\2\2\u01a7\u01a8\7#\2\2\u01a8^\3"+
		"\2\2\2\u01a9\u01aa\7~\2\2\u01aa\u01ab\7~\2\2\u01ab`\3\2\2\2\u01ac\u01ad"+
		"\7(\2\2\u01ad\u01ae\7(\2\2\u01aeb\3\2\2\2\u01af\u01b0\7(\2\2\u01b0d\3"+
		"\2\2\2\u01b1\u01b2\7?\2\2\u01b2f\3\2\2\2\u01b3\u01b4\7@\2\2\u01b4\u01b5"+
		"\7?\2\2\u01b5h\3\2\2\2\u01b6\u01b7\7>\2\2\u01b7\u01b8\7?\2\2\u01b8j\3"+
		"\2\2\2\u01b9\u01ba\7@\2\2\u01bal\3\2\2\2\u01bb\u01bc\7>\2\2\u01bcn\3\2"+
		"\2\2\u01bd\u01be\7,\2\2\u01bep\3\2\2\2\u01bf\u01c0\7\61\2\2\u01c0r\3\2"+
		"\2\2\u01c1\u01c2\7-\2\2\u01c2t\3\2\2\2\u01c3\u01c4\7\'\2\2\u01c4v\3\2"+
		"\2\2\u01c5\u01c6\7/\2\2\u01c6x\3\2\2\2\u01c7\u01c8\7*\2\2\u01c8z\3\2\2"+
		"\2\u01c9\u01ca\7+\2\2\u01ca|\3\2\2\2\u01cb\u01cc\7}\2\2\u01cc~\3\2\2\2"+
		"\u01cd\u01ce\7\177\2\2\u01ce\u0080\3\2\2\2\u01cf\u01d0\7<\2\2\u01d0\u0082"+
		"\3\2\2\2\u01d1\u01d2\7]\2\2\u01d2\u0084\3\2\2\2\u01d3\u01d4\7_\2\2\u01d4"+
		"\u0086\3\2\2\2\u01d5\u01d6\7.\2\2\u01d6\u0088\3\2\2\2\u01d7\u01d8\7=\2"+
		"\2\u01d8\u008a\3\2\2\2\u01d9\u01db\t\6\2\2\u01da\u01d9\3\2\2\2\u01db\u01dc"+
		"\3\2\2\2\u01dc\u01da\3\2\2\2\u01dc\u01dd\3\2\2\2\u01dd\u01de\3\2\2\2\u01de"+
		"\u01df\bF\2\2\u01df\u008c\3\2\2\2\u01e0\u01e1\7\61\2\2\u01e1\u01e2\7,"+
		"\2\2\u01e2\u01e6\3\2\2\2\u01e3\u01e5\13\2\2\2\u01e4\u01e3\3\2\2\2\u01e5"+
		"\u01e8\3\2\2\2\u01e6\u01e7\3\2\2\2\u01e6\u01e4\3\2\2\2\u01e7\u01e9\3\2"+
		"\2\2\u01e8\u01e6\3\2\2\2\u01e9\u01ea\7,\2\2\u01ea\u01eb\7\61\2\2\u01eb"+
		"\u01ec\3\2\2\2\u01ec\u01ed\bG\2\2\u01ed\u008e\3\2\2\2\u01ee\u01ef\7\61"+
		"\2\2\u01ef\u01f0\7\61\2\2\u01f0\u01f4\3\2\2\2\u01f1\u01f3\n\7\2\2\u01f2"+
		"\u01f1\3\2\2\2\u01f3\u01f6\3\2\2\2\u01f4\u01f2\3\2\2\2\u01f4\u01f5\3\2"+
		"\2\2\u01f5\u01f7\3\2\2\2\u01f6\u01f4\3\2\2\2\u01f7\u01f8\bH\2\2\u01f8"+
		"\u0090\3\2\2\2\u01f9\u01fa\7^\2\2\u01fa\u01fb\t\b\2\2\u01fb\u0092\3\2"+
		"\2\2\13\2\u0181\u0187\u0189\u018f\u0198\u01dc\u01e6\u01f4\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}