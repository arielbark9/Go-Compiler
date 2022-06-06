package Parser

import (
	"regexp"
	"strings"
)

var currentToken int
var tokens []token

var ClassSymbolTable *SymbolTable
var SubroutineSymbolTable *SymbolTable
var currentClassName string

var subroutineNames []string
var COMMA_REGEX = regexp.MustCompile(",")

func matchToken(that *regexp.Regexp) *node {
	if that.MatchString(tokens[currentToken].tValue) {
		res := node{
			nType:    tokens[currentToken].tType,
			nValue:   tokens[currentToken].tValue,
			children: nil,
		}
		currentToken++
		return &res
	}
	return nil
}

func class() *node {
	res := &node{
		nType:    "class",
		nValue:   "",
		children: []*node{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("class")))
	name := className()
	currentClassName = name.nValue
	res.children = append(res.children, name)
	res.children = append(res.children, matchToken(regexp.MustCompile("\\{")))
	for c := classVarDec(); c != nil; c = classVarDec() {
		res.children = append(res.children, c)
	}
	for c := subroutineDec(); c != nil; c = subroutineDec() {
		res.children = append(res.children, c)
		SubroutineSymbolTable.startSubroutine()
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("\\}")))
	return res
}

func typeDefinition() *node {
	if s := matchToken(regexp.MustCompile("int|char|boolean")); s != nil {
		return s
	} else if IDENTIFIER_REGEX.Match([]byte(tokens[currentToken].tValue)) {
		return className()
	} else {
		return nil
	}
}

func subroutineDec() *node {
	var res *node = nil
	if s := matchToken(regexp.MustCompile("(constructor|function|method)")); s != nil {
		res = &node{
			nType:    "subroutineDec",
			nValue:   "",
			children: []*node{s},
		}

		if s.nValue == "method" {
			SubroutineSymbolTable.define("this", currentClassName, "ARGUMENT")
		}

		if s := matchToken(regexp.MustCompile("void")); s != nil {
			res.children = append(res.children, s)
		} else {
			res.children = append(res.children, typeDefinition())
		}
		name := subroutineName()
		res.children = append(res.children, name)
		subroutineNames = append(subroutineNames, name.nValue)
		res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
		res.children = append(res.children, parameterList())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
		res.children = append(res.children, subroutineBody())
	}
	return res
}

func subroutineBody() *node {
	res := &node{
		nType:    "subroutineBody",
		nValue:   "",
		children: []*node{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("\\{")))
	for c := varDec(); c != nil; c = varDec() {
		res.children = append(res.children, c)
	}
	if s := statements(); s != nil {
		res.children = append(res.children, s)
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("\\}")))
	return res
}

func statements() *node {
	if s := statement(); s != nil {
		res := &node{
			nType:    "statements",
			nValue:   "",
			children: []*node{s},
		}
		for s := statement(); s != nil; s = statement() {
			res.children = append(res.children, s)
		}
		return res
	}
	return &node{
		nType:    "statements",
		nValue:   "",
		children: nil,
	}
}

func statement() *node {
	val := tokens[currentToken].tValue
	if val == "let" {
		return letStatement()
	} else if val == "if" {
		return ifStatement()
	} else if val == "while" {
		return whileStatement()
	} else if val == "do" {
		return doStatement()
	} else if val == "return" {
		return returnStatement()
	}
	return nil
}

func returnStatement() *node {
	res := &node{
		nType:    "returnStatement",
		nValue:   "",
		children: []*node{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("return")))
	if s := expression(); s != nil {
		res.children = append(res.children, s)
	}
	res.children = append(res.children, matchToken(regexp.MustCompile(";")))
	return res
}

func doStatement() *node {
	res := &node{
		nType:    "doStatement",
		nValue:   "",
		children: []*node{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("do")))
	//res.children = append(res.children, subroutineCall())
	res.children = append(res.children, subroutineCall().children...)
	res.children = append(res.children, matchToken(regexp.MustCompile(";")))
	return res
}

func whileStatement() *node {
	res := &node{
		nType:    "whileStatement",
		nValue:   "",
		children: []*node{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("while")))
	res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
	res.children = append(res.children, expression())
	res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
	res.children = append(res.children, matchToken(regexp.MustCompile("\\{")))
	res.children = append(res.children, statements())
	res.children = append(res.children, matchToken(regexp.MustCompile("}")))
	return res
}

func ifStatement() *node {
	res := &node{
		nType:    "ifStatement",
		nValue:   "",
		children: []*node{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("if")))
	res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
	res.children = append(res.children, expression())
	res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
	res.children = append(res.children, matchToken(regexp.MustCompile("\\{")))
	res.children = append(res.children, statements())
	res.children = append(res.children, matchToken(regexp.MustCompile("}")))
	if s := matchToken(regexp.MustCompile("else")); s != nil {
		res.children = append(res.children, s)
		res.children = append(res.children, matchToken(regexp.MustCompile("\\{")))
		res.children = append(res.children, statements())
		res.children = append(res.children, matchToken(regexp.MustCompile("}")))
	}
	return res
}

func letStatement() *node {
	res := &node{
		nType:    "letStatement",
		nValue:   "",
		children: []*node{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("let")))
	res.children = append(res.children, varName())
	if s := matchToken(regexp.MustCompile("[\\[\\(]")); s != nil {
		res.children = append(res.children, s)
		res.children = append(res.children, expression())
		res.children = append(res.children, matchToken(regexp.MustCompile("[\\]\\)]")))
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("=")))
	res.children = append(res.children, expression())
	res.children = append(res.children, matchToken(regexp.MustCompile(";")))
	return res
}

func expression() *node {
	res := &node{
		nType:    "expression",
		nValue:   "",
		children: []*node{},
	}
	if s := term(); s != nil {
		res.children = append(res.children, s)
		for s := op(); s != nil; s = op() {
			res.children = append(res.children, s)
			res.children = append(res.children, term())
		}
		return res
	}
	return nil
}

func op() *node {
	if s := matchToken(regexp.MustCompile("[+\\-*/&|<>=]")); s != nil {
		return s
	}
	return nil
}

func term() *node {
	res := &node{
		nType:    "term",
		nValue:   "",
		children: []*node{},
	}
	tokenCount := len(tokens)
	if s := matchToken(INTEGER_REGEX); s != nil {
		res.children = append(res.children, s)
	} else if s := matchToken(STRING_REGEX); s != nil {
		res.children = append(res.children, s)
	} else if s := keywordConstant(); s != nil {
		res.children = append(res.children, s)
	} else if tokens[currentToken+1].tValue != "[" &&
		((currentToken < tokenCount-1 && tokens[currentToken+1].tValue == "(") ||
			((currentToken < tokenCount-1 && tokens[currentToken+1].tValue == ".") &&
				currentToken < tokenCount-3 && tokens[currentToken+3].tValue == "(")) {
		if s := subroutineCall(); s != nil {
			//res.children = append(res.children, s)
			res.children = append(res.children, s.children...)
		} else if s := varName(); s != nil {
			res.children = append(res.children, s)
			if s := matchToken(regexp.MustCompile("[\\[\\(]")); s != nil {
				res.children = append(res.children, s)
				res.children = append(res.children, expression())
				res.children = append(res.children, matchToken(regexp.MustCompile("[\\]\\)]")))
			}
		} else if s := matchToken(regexp.MustCompile("\\(")); s != nil {
			res.children = append(res.children, s)
			res.children = append(res.children, expression())
			res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
		} else if s := unaryOp(); s != nil {
			res.children = append(res.children, s)
			res.children = append(res.children, term())
		} else {
			return nil
		}
	} else if s := varName(); s != nil {
		res.children = append(res.children, s)
		if s := matchToken(regexp.MustCompile("[\\[\\(]")); s != nil {
			res.children = append(res.children, s)
			res.children = append(res.children, expression())
			res.children = append(res.children, matchToken(regexp.MustCompile("[\\]\\)]")))
		}
	} else if s := matchToken(regexp.MustCompile("\\(")); s != nil {
		res.children = append(res.children, s)
		res.children = append(res.children, expression())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
	} else if s := unaryOp(); s != nil {
		res.children = append(res.children, s)
		res.children = append(res.children, term())
	} else {
		return nil
	}
	return res
}

func keywordConstant() *node {
	if s := matchToken(regexp.MustCompile("true|false|null|this")); s != nil {
		return s
	}
	return nil
}

func unaryOp() *node {
	if s := matchToken(regexp.MustCompile("[~\\-]")); s != nil {
		return s
	}
	return nil
}

func subroutineCall() *node {
	res := &node{
		nType:    "subroutineCall",
		nValue:   "",
		children: []*node{},
	}
	name := tokens[currentToken].tValue
	if contains(subroutineNames, name) {
		res.children = append(res.children, subroutineName())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
		res.children = append(res.children, expressionList())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
		return res
	} else if SubroutineSymbolTable.indexOf(tokens[currentToken].tValue) != -1 {
		res.children = append(res.children, varName())
	} else if IDENTIFIER_REGEX.Match([]byte(name)) {
		res.children = append(res.children, className())
	} else {
		return nil
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("\\.")))
	res.children = append(res.children, subroutineName())
	res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
	res.children = append(res.children, expressionList())
	res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))

	return res
}

func expressionList() *node {
	res := &node{
		nType:    "expressionList",
		nValue:   "",
		children: []*node{},
	}
	if s := expression(); s != nil {
		res.children = append(res.children, s)
		for s := matchToken(COMMA_REGEX); s != nil; s = matchToken(COMMA_REGEX) {
			res.children = append(res.children, s)
			res.children = append(res.children, expression())
		}
	}
	return res
}

func varName() *node {
	if s := matchToken(IDENTIFIER_REGEX); s != nil {
		return s
	}
	return nil
}

func className() *node {
	if s := matchToken(IDENTIFIER_REGEX); s != nil {
		return s
	}
	return nil
}

func subroutineName() *node {
	if s := matchToken(IDENTIFIER_REGEX); s != nil {
		return s
	}
	return nil
}

func varDec() *node {
	res := &node{
		nType:    "varDec",
		nValue:   "",
		children: []*node{},
	}
	if s := matchToken(regexp.MustCompile("var")); s != nil {
		res.children = append(res.children, s)
	} else {
		return nil
	}
	typeDef := typeDefinition()
	res.children = append(res.children, typeDef)
	name := varName()
	res.children = append(res.children, name)
	SubroutineSymbolTable.define(name.nValue, typeDef.nValue, "VAR")
	for s := matchToken(COMMA_REGEX); s != nil; s = matchToken(COMMA_REGEX) {
		res.children = append(res.children, s)
		name = varName()
		res.children = append(res.children, name)
		SubroutineSymbolTable.define(name.nValue, typeDef.nValue, "VAR")
	}
	res.children = append(res.children, matchToken(regexp.MustCompile(";")))
	return res
}

func parameterList() *node {
	res := &node{
		nType:    "parameterList",
		nValue:   "",
		children: []*node{},
	}
	if typeDef := typeDefinition(); typeDef != nil {
		res.children = append(res.children, typeDef)
		name := varName()
		res.children = append(res.children, name)
		SubroutineSymbolTable.define(name.nValue, typeDef.nValue, "ARGUMENT")
		for s := matchToken(COMMA_REGEX); s != nil; s = matchToken(COMMA_REGEX) {
			res.children = append(res.children, s)
			typeDef = typeDefinition()
			res.children = append(res.children, typeDef)
			name = varName()
			res.children = append(res.children, name)
			SubroutineSymbolTable.define(name.nValue, typeDef.nValue, "ARGUMENT")
		}
	}
	return res
}

func classVarDec() *node {
	var res *node = nil
	if s := matchToken(regexp.MustCompile("static|field")); s != nil {
		res = &node{
			nType:    "classVarDec",
			nValue:   "",
			children: []*node{s},
		}
		varClassification := strings.ToUpper(s.nValue)
		varType := typeDefinition()
		res.children = append(res.children, varType)
		name := varName()
		res.children = append(res.children, name)
		ClassSymbolTable.define(name.nValue, varType.nValue, varClassification)

		for s = matchToken(COMMA_REGEX); s != nil; s = matchToken(COMMA_REGEX) {
			res.children = append(res.children, s)
			name = varName()
			res.children = append(res.children, name)

			ClassSymbolTable.define(name.nValue, varType.nValue, varClassification)
		}
		res.children = append(res.children, matchToken(regexp.MustCompile(";")))
	}
	return res
}

func ParseToXML(source string) string {
	tokens = GetTokens(source)
	currentToken = 0
	ClassSymbolTable = newSymbolTable()
	SubroutineSymbolTable = newSymbolTable()
	root := class()
	return root.ToXml(0)
}
