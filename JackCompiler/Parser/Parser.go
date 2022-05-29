package Parser

import "regexp"

var currentToken int
var tokens []token
var TYPE_REGEX = regexp.MustCompile("^(int|char|boolean)")

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
	var class xmlElement
	class.xType = "class"
	matchToken(regexp.MustCompile("class"))
	class.children = append(class.children, *matchToken(IDENTIFIER_REGEX))
	class.children = append(class.children, *matchToken(regexp.MustCompile("{")))
	for c := classVarDec(); c != nil; {
		class.children = append(class.children, c)
	}
	for c := subRoutineDec(); c != nil; {
		class.children = append(class.children, c)
	}
	class.children = append(class.children, *matchToken(regexp.MustCompile("}")))
	return &class
}

func subRoutineDec() *xmlElement {
	var res *xmlElement = nil
	if s := matchToken(regexp.MustCompile("constructor|function|method")); s != nil {
		res = &xmlElement{
			xType:    "subroutineDec",
			value:    "",
			children: []xmlElement{*s},
		}
		res.children = append(res.children, *matchToken(regexp.MustCompile(
			regexp.MustCompile(TYPE_REGEX.String() + "|void"))))
		res.children = append(res.children, *matchToken(IDENTIFIER_REGEX))
		res.children = append(res.children, *matchToken(regexp.MustCompile("\\(")))
		res.children = append(res.children, *parameterList())
		res.children = append(res.children, *matchToken(regexp.MustCompile("\\)")))
		res.children = append(res.children, *subroutineBody())
	}
	return res
}

func classVarDec() *xmlElement {
	var res *xmlElement = nil
	if s := matchToken(regexp.MustCompile("static|field")); s != nil {
		res = &xmlElement{
			xType:    "classVarDec",
			value:    "",
			children: []xmlElement{*s},
		}
		res.children = append(res.children, *matchToken(regexp.MustCompile(TYPE_REGEX)))
		res.children = append(res.children, *matchToken(IDENTIFIER_REGEX))
		for s = matchToken(regexp.MustCompile(",")); s != nil; {
			res.children = append(res.children, *s)
			res.children = append(res.children, *matchToken(IDENTIFIER_REGEX))
		}
		res.children = append(res.children, *matchToken(regexp.MustCompile(";")))
	}
	return res
}

func ParseToXML(source string) string {
	tokens = GetTokens(source)
	currentToken = 0
	return class()
}
