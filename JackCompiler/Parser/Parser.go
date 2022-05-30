package Parser

import (
	"regexp"
)

var currentToken int
var tokens []token
var classNames []string
var varNames []string
var subroutineNames []string
var COMMA_REGEX = regexp.MustCompile(",")

func matchToken(that *regexp.Regexp) *xmlElement {
	if that.MatchString(tokens[currentToken].value) {
		res := xmlElement{
			xType:    tokens[currentToken].tType,
			value:    tokens[currentToken].value,
			children: nil,
		}
		currentToken++
		return &res
	}
	return nil
}

func class() *xmlElement {
	res := &xmlElement{
		xType:    "class",
		value:    "",
		children: []*xmlElement{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("class")))
	name := className()
	res.children = append(res.children, name)
	classNames = append(classNames, name.value)
	res.children = append(res.children, matchToken(regexp.MustCompile("\\{")))
	for c := classVarDec(); c != nil; c = classVarDec() {
		res.children = append(res.children, c)
	}
	for c := subroutineDec(); c != nil; c = subroutineDec() {
		res.children = append(res.children, c)
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("\\}")))
	return res
}

func typeDefinition() *xmlElement {
	if s := matchToken(regexp.MustCompile("int|char|boolean")); s != nil {
		return s
	} else if IDENTIFIER_REGEX.Match([]byte(tokens[currentToken].value)) {
		return className()
	} else {
		return nil
	}
}

func subroutineDec() *xmlElement {
	var res *xmlElement = nil
	if s := matchToken(regexp.MustCompile("(constructor|function|method)")); s != nil {
		res = &xmlElement{
			xType:    "subroutineDec",
			value:    "",
			children: []*xmlElement{s},
		}
		if s := matchToken(regexp.MustCompile("void")); s != nil {
			res.children = append(res.children, s)
		} else {
			res.children = append(res.children, typeDefinition())
		}
		name := subroutineName()
		res.children = append(res.children, name)
		subroutineNames = append(subroutineNames, name.value)
		res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
		res.children = append(res.children, parameterList())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
		res.children = append(res.children, subroutineBody())
	}
	return res
}

func subroutineBody() *xmlElement {
	res := &xmlElement{
		xType:    "subroutineBody",
		value:    "",
		children: []*xmlElement{},
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

func statements() *xmlElement {
	if s := statement(); s != nil {
		res := &xmlElement{
			xType:    "statements",
			value:    "",
			children: []*xmlElement{s},
		}
		for s := statement(); s != nil; s = statement() {
			res.children = append(res.children, s)
		}
		return res
	}
	return nil
}

func statement() *xmlElement {
	val := tokens[currentToken].value
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

func returnStatement() *xmlElement {
	res := &xmlElement{
		xType:    "returnStatement",
		value:    "",
		children: []*xmlElement{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("return")))
	if s := expression(); s != nil {
		res.children = append(res.children, s)
	}
	res.children = append(res.children, matchToken(regexp.MustCompile(";")))
	return res
}

func doStatement() *xmlElement {
	res := &xmlElement{
		xType:    "doStatement",
		value:    "",
		children: []*xmlElement{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("do")))
	//res.children = append(res.children, subroutineCall())
	res.children = append(res.children, subroutineCall().children...)
	res.children = append(res.children, matchToken(regexp.MustCompile(";")))
	return res
}

func whileStatement() *xmlElement {
	res := &xmlElement{
		xType:    "whileStatement",
		value:    "",
		children: []*xmlElement{},
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

func ifStatement() *xmlElement {
	res := &xmlElement{
		xType:    "ifStatement",
		value:    "",
		children: []*xmlElement{},
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

func letStatement() *xmlElement {
	res := &xmlElement{
		xType:    "letStatement",
		value:    "",
		children: []*xmlElement{},
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

func expression() *xmlElement {
	res := &xmlElement{
		xType:    "expression",
		value:    "",
		children: []*xmlElement{},
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

func op() *xmlElement {
	if s := matchToken(regexp.MustCompile("[+\\-*/&|<>=]")); s != nil {
		return s
	}
	return nil
}

func term() *xmlElement {
	res := &xmlElement{
		xType:    "term",
		value:    "",
		children: []*xmlElement{},
	}
	tokenCount := len(tokens)
	if s := matchToken(INTEGER_REGEX); s != nil {
		res.children = append(res.children, s)
	} else if s := matchToken(STRING_REGEX); s != nil {
		res.children = append(res.children, s)
	} else if s := keywordConstant(); s != nil {
		res.children = append(res.children, s)
	} else if tokens[currentToken+1].value != "[" &&
		(currentToken < tokenCount-1 && tokens[currentToken+1].value == "(") ||
		(currentToken < tokenCount-3 && tokens[currentToken+3].value == "(") {
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

func keywordConstant() *xmlElement {
	if s := matchToken(regexp.MustCompile("true|false|null|this")); s != nil {
		return s
	}
	return nil
}

func unaryOp() *xmlElement {
	if s := matchToken(regexp.MustCompile("[~\\-]")); s != nil {
		return s
	}
	return nil
}

func subroutineCall() *xmlElement {
	res := &xmlElement{
		xType:    "subroutineCall",
		value:    "",
		children: []*xmlElement{},
	}
	name := tokens[currentToken].value
	if contains(subroutineNames, name) {
		res.children = append(res.children, subroutineName())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
		res.children = append(res.children, expressionList())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
		return res
	} else if contains(varNames, tokens[currentToken].value) {
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

func expressionList() *xmlElement {
	res := &xmlElement{
		xType:    "expressionList",
		value:    "",
		children: []*xmlElement{},
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

func varName() *xmlElement {
	if s := matchToken(IDENTIFIER_REGEX); s != nil {
		return s
	}
	return nil
}

func className() *xmlElement {
	if s := matchToken(IDENTIFIER_REGEX); s != nil {
		return s
	}
	return nil
}

func subroutineName() *xmlElement {
	if s := matchToken(IDENTIFIER_REGEX); s != nil {
		return s
	}
	return nil
}

func varDec() *xmlElement {
	res := &xmlElement{
		xType:    "varDec",
		value:    "",
		children: []*xmlElement{},
	}
	if s := matchToken(regexp.MustCompile("var")); s != nil {
		res.children = append(res.children, s)
	} else {
		return nil
	}
	res.children = append(res.children, typeDefinition())
	res.children = append(res.children, varName())
	for s := matchToken(COMMA_REGEX); s != nil; s = matchToken(COMMA_REGEX) {
		res.children = append(res.children, s)
		res.children = append(res.children, varName())
	}
	res.children = append(res.children, matchToken(regexp.MustCompile(";")))
	return res
}

func parameterList() *xmlElement {
	res := &xmlElement{
		xType:    "parameterList",
		value:    "",
		children: []*xmlElement{},
	}
	if s := typeDefinition(); s != nil {
		res.children = append(res.children, s)
		res.children = append(res.children, varName())
		for s := matchToken(COMMA_REGEX); s != nil; s = matchToken(COMMA_REGEX) {
			res.children = append(res.children, s)
			res.children = append(res.children, typeDefinition())
			res.children = append(res.children, varName())
		}
	}
	return res
}

func classVarDec() *xmlElement {
	var res *xmlElement = nil
	if s := matchToken(regexp.MustCompile("static|field")); s != nil {
		res = &xmlElement{
			xType:    "classVarDec",
			value:    "",
			children: []*xmlElement{s},
		}
		res.children = append(res.children, typeDefinition())
		name := varName()
		res.children = append(res.children, name)
		varNames = append(varNames, name.value)

		for s = matchToken(COMMA_REGEX); s != nil; s = matchToken(COMMA_REGEX) {
			res.children = append(res.children, s)
			name = varName()
			res.children = append(res.children, name)
			varNames = append(varNames, name.value)
		}
		res.children = append(res.children, matchToken(regexp.MustCompile(";")))
	}
	return res
}

func ParseToXML(source string) string {
	tokens = GetTokens(source)
	currentToken = 0
	root := class()
	return root.ToXml(0)
}
