package token

// from book "Writing An Interpreter In Go"

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	PIPE      = "|"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Comment
	QUART = "'"

	// type keywords
	ACTOR= "actor"
	PARTICIPANT= "participant"
	USECASE= "usecase"
	CLASS= "class"
	INTERFACE= "interface"
	ABSTRACT= "abstract"
	ENUM= "enum"
	COMPONENT= "component"
	STATE= "state"
	OBJECT= "object"
	ARTIFACT= "artifact"
	FOLDER= "folder"
	RECTANGLE= "rectangle"
	NODE= "node"
	FRAME= "frame"
	CLOUD= "cloud"
	DATABASE= "database"
	STORAGE= "storage"
	AGENT= "agent"
	STACK= "stack"
	BOUNDARY= "boundary"
	CONTROL= "control"
	ENTITY= "entity"
	CARD= "card"
	FILE= "file"
	PACKAGE= "package"
	QUEUE= "queue"
	ARCHIMATE= "archimate"
	DIAMOND= "diamond"
	DETACH= "detach"
	
	// Keywords
	STARTUML = "@startuml"
	ENDUML = "@enduml"
	STARTDOT = "@startdot"
	ENDDOT = "@enddot"
	STARTSALT = "@startsalt"
	ENDSALT = "@endsalt"
	STARTJSON = "@startjson"
	ENDJSON = "@endjson"
	AS = "as"
	ALSO = "also"
	AUTONUMBER = "autonumber"
	CAPTION = "caption"
	TITLE = "title"
	NEWPAGE = "newpage"
	BOX = "box"
	ALT = "alt"
	OPT = "opt"
	LOOP = "loop"
	PAR = "par"
	BREAK = "break"
	CRITICAL = "critical"
	NOTE = "note"
	RNOTE = "rnote"
	HNOTE = "hnote"
	LEGEND = "legend"
	GROUP = "group"
	LEFT = "left"
	RIGHT = "right"
	OF = "of"
	ON = "on"
	LINK = "link"
	OVER = "over"
	END = "end"
	ACTIVATE = "activate"
	DEACTIVATE = "deactivate"
	DESTROY = "destroy"
	CREATE = "create"
	FOOTBOX = "footbox"
	HIDE = "hide"
	SHOW = "show"
	SKINPARAM = "skinparam"
	SKIN = "skin"
	TOP = "top"
	TO ="to" // top to bottom direction
	BOTTOM = "bottom"
	DIRECTION = "direction" // top to bottom direction
	NAMESPACE = "namespace"
	PAGE = "page"
	UP = "up"
	DOWN = "down"
	IF = "if"
	THEN = "then"
	ELSE = "else"
	ELSEIF = "elseif"
	ENDIF = "endif"
	PARTITION = "partition"
	FOOTER = "footer"
	HEADER = "header"
	CENTER = "center"
	ROTATE = "rotate"
	REF = "ref"
	RETURN = "return"
	IS = "is"
	REPEAT = "repeat"
	START = "start"
	STOP = "stop"
	WHILE = "while"
	ENDWHILE = "endwhile"
	FORK = "fork"
	AGAIN = "again"
	KILL = "kill"
	ORDER = "order"
	ALLOW_MIXING = "allow_mixing"
	ALLOWMIXING = "allowmixing"
	MAINFRAME = "mainframe"
	ACROSS = "across"
	STEREOTYPE = "stereotype"
	SPLIT = "split"
	STYLE = "style"
	SPRITE = "sprite"
	CIRCLE = "circle"
	EMPTY = "empty"
	MEMBERS = "members"
	DESCRIPTION = "description"
	TRUE = "true"
	FALSE = "false"
	NORMAL = "normal"
	ITALIC = "italic"
	BOLD = "bold"
	PLAIN = "plain"
	COLOR = "color"
	DOTTED = "dotted"
	DASHED = "dashed"
	MAP = "map"

)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"actor": ACTOR,
	"participant": PARTICIPANT,
	"usecase": USECASE,
	"class": CLASS,
	"interface": INTERFACE,
	"abstract": ABSTRACT,
	"enum": ENUM,
	"component": COMPONENT,
	"state": STATE,
	"object": OBJECT,
	"artifact": ARTIFACT,
	"folder": FOLDER,
	"rectangle": RECTANGLE,
	"node": NODE,
	"frame": FRAME,
	"cloud": CLOUD,
	"database": DATABASE,
	"storage": STORAGE,
	"agent": AGENT,
	"stack": STACK,
	"boundary": BOUNDARY,
	"control": CONTROL,
	"entity": ENTITY,
	"card": CARD,
	"file": FILE,
	"package": PACKAGE,
	"queue": QUEUE,
	"archimate": ARCHIMATE,
	"diamond": DIAMOND,
	"detach": DETACH,

	"@startuml": STARTUML,
	"@enduml": ENDUML,
	"@startdot": STARTDOT,
	"@enddot": ENDDOT,
	"@startsalt": STARTSALT,
	"@endsalt": ENDSALT,
	"@startjson": STARTJSON,
	"@endjson": ENDJSON,
	"as": AS,
	"also": ALSO,
	"autonumber": AUTONUMBER,
	"caption": CAPTION,
	"title": TITLE,
	"newpage": NEWPAGE,
	"box": BOX,
	"alt": ALT,
	"opt": OPT,
	"loop": LOOP,
	"par": PAR,
	"break": BREAK,
	"critical": CRITICAL,
	"note": NOTE,
	"rnote": RNOTE,
	"hnote": HNOTE,
	"legend": LEGEND,
	"group": GROUP,
	"left": LEFT,
	"right": RIGHT,
	"of": OF,
	"on": ON,
	"link": LINK,
	"over": OVER,
	"end": END,
	"activate": ACTIVATE,
	"deactivate": DEACTIVATE,
	"destroy": DESTROY,
	"create": CREATE,
	"footbox": FOOTBOX,
	"hide": HIDE,
	"show": SHOW,
	"skinparam": SKINPARAM,
	"skin": SKIN,
	"top": TOP,
	"to": TO, // top to bottom direction
	"bottom": BOTTOM,
	"direction": DIRECTION, // top to bottom direction
	"namespace": NAMESPACE,
	"page": PAGE,
	"up": UP,
	"down": DOWN,
	"if": IF,
	"then": THEN,
	"else": ELSE,
	"elseif": ELSEIF,
	"endif": ENDIF,
	"partition": PARTITION,
	"footer": FOOTER,
	"header": HEADER,
	"center": CENTER,
	"rotate": ROTATE,
	"ref": REF,
	"return": RETURN,
	"is": IS,
	"repeat": REPEAT,
	"start": START,
	"stop": STOP,
	"while": WHILE,
	"endwhile": ENDWHILE,
	"fork": FORK,
	"again": AGAIN,
	"kill": KILL,
	"order": ORDER,
	"allow_mixing": ALLOW_MIXING,
	"allowmixing": ALLOWMIXING,
	"mainframe": MAINFRAME,
	"across": ACROSS,
	"stereotype": STEREOTYPE,
	"split": SPLIT,
	"style": STYLE,
	"sprite": SPRITE,
	"circle": CIRCLE,
	"empty": EMPTY,
	"members": MEMBERS,
	"description": DESCRIPTION,
	"true": TRUE,
	"false": FALSE,
	"normal": NORMAL,
	"italic": ITALIC,
	"bold": BOLD,
	"plain": PLAIN,
	"color": COLOR,
	"dotted": DOTTED,
	"dashed": DASHED,
	"map": MAP,

}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

















































