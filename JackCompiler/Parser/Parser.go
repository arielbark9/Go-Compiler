package Parser

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var currentToken int
var tokens []token

var ClassSymbolTable *SymbolTable
var SubroutineSymbolTable *SymbolTable
var currentClassName string
var writer VMWriter

var subroutineNames []string
var CommaRegex = regexp.MustCompile(",")

var opMap = map[string]string{
	"+":     "add",
	"-":     "sub",
	"*":     "call Math.multiply 2",
	"/":     "call Math.divide 2",
	"&":     "and",
	"&amp;": "and",
	"|":     "or",
	"<":     "lt",
	"&lt;":  "lt",
	">":     "gt",
	"&gt;":  "gt",
	"=":     "eq",
}

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
		writer.ifCount = 0
		writer.whileCount = 0
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
		res.children = append(res.children, subroutineBody(name.nValue, s.nValue))
	}
	return res
}

func subroutineBody(name string, kind string) *node {
	res := &node{
		nType:    "subroutineBody",
		nValue:   "",
		children: []*node{},
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("\\{")))
	for c := varDec(); c != nil; c = varDec() {
		res.children = append(res.children, c)
	}
	writer.writeFunction(currentClassName, name, SubroutineSymbolTable.varCount("VAR"))

	switch kind {
	case "constructor":
		writer.writePush("constant", ClassSymbolTable.varCount("FIELD"))
		writer.writeCall("Memory.alloc", 1)
		writer.writePop("pointer", 0)
		break
	case "method":
		writer.writePush("argument", 0)
		writer.writePop("pointer", 0)
		break
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
	} else {
		writer.writePush("constant", 0)
	}
	writer.writeReturn()
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
	res.children = append(res.children, subroutineCall().children...)
	res.children = append(res.children, matchToken(regexp.MustCompile(";")))
	writer.writePop("temp", 0)
	return res
}

func whileStatement() *node {
	res := &node{
		nType:    "whileStatement",
		nValue:   "",
		children: []*node{},
	}
	// for nested whiles
	currentWhileCount := writer.whileCount
	writer.whileCount++
	res.children = append(res.children, matchToken(regexp.MustCompile("while")))
	res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
	writer.writeLabel("WHILE_EXP" + strconv.Itoa(currentWhileCount))
	res.children = append(res.children, expression())
	writer.writeArithmetic("not")
	writer.writeIf("WHILE_END" + strconv.Itoa(currentWhileCount))
	res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
	res.children = append(res.children, matchToken(regexp.MustCompile("\\{")))
	res.children = append(res.children, statements())
	res.children = append(res.children, matchToken(regexp.MustCompile("}")))
	writer.writeGoto("WHILE_EXP" + strconv.Itoa(currentWhileCount))
	writer.writeLabel("WHILE_END" + strconv.Itoa(currentWhileCount))
	//writer.whileCount--
	return res
}

func ifStatement() *node {
	res := &node{
		nType:    "ifStatement",
		nValue:   "",
		children: []*node{},
	}
	ifElse := false
	res.children = append(res.children, matchToken(regexp.MustCompile("if")))
	res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
	res.children = append(res.children, expression())
	res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
	// for nested ifs
	currentIfCount := writer.ifCount
	writer.ifCount++
	writer.writeIf("IF_TRUE" + strconv.Itoa(currentIfCount))
	writer.writeGoto("IF_FALSE" + strconv.Itoa(currentIfCount))
	writer.writeLabel("IF_TRUE" + strconv.Itoa(currentIfCount))
	res.children = append(res.children, matchToken(regexp.MustCompile("\\{")))
	res.children = append(res.children, statements())
	res.children = append(res.children, matchToken(regexp.MustCompile("}")))
	if tokens[currentToken].tValue == "else" {
		ifElse = true
		writer.writeGoto("IF_END" + strconv.Itoa(currentIfCount))
	}
	writer.writeLabel("IF_FALSE" + strconv.Itoa(currentIfCount))
	if s := matchToken(regexp.MustCompile("else")); s != nil {
		res.children = append(res.children, s)
		res.children = append(res.children, matchToken(regexp.MustCompile("\\{")))
		res.children = append(res.children, statements())
		res.children = append(res.children, matchToken(regexp.MustCompile("}")))
	}
	if ifElse {
		writer.writeLabel("IF_END" + strconv.Itoa(currentIfCount))
	}
	//writer.ifCount--
	return res
}

func letStatement() *node {
	res := &node{
		nType:    "letStatement",
		nValue:   "",
		children: []*node{},
	}
	withArray := false
	res.children = append(res.children, matchToken(regexp.MustCompile("let")))
	name := varName()
	res.children = append(res.children, name)
	if s := matchToken(regexp.MustCompile("[\\[\\(]")); s != nil {
		withArray = true
		res.children = append(res.children, s)
		res.children = append(res.children, expression())
		res.children = append(res.children, matchToken(regexp.MustCompile("[\\]\\)]")))
		if SubroutineSymbolTable.indexOf(name.nValue) != -1 {
			writer.writePush(SubroutineSymbolTable.kindOf(name.nValue), SubroutineSymbolTable.indexOf(name.nValue))
		} else {
			writer.writePush(ClassSymbolTable.kindOf(name.nValue), ClassSymbolTable.indexOf(name.nValue))
		}
		writer.writeArithmetic("add")
	}
	res.children = append(res.children, matchToken(regexp.MustCompile("=")))
	res.children = append(res.children, expression())
	res.children = append(res.children, matchToken(regexp.MustCompile(";")))

	if !withArray {
		if SubroutineSymbolTable.indexOf(name.nValue) != -1 {
			writer.writePop(SubroutineSymbolTable.kindOf(name.nValue), SubroutineSymbolTable.indexOf(name.nValue))
		} else {
			writer.writePop(ClassSymbolTable.kindOf(name.nValue), ClassSymbolTable.indexOf(name.nValue))
		}
	} else {
		writer.writePop("temp", 0)
		writer.writePop("pointer", 1)
		writer.writePush("temp", 0)
		writer.writePop("that", 0)
	}

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
			writer.writeArithmetic(opMap[s.nValue])
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
	if s := matchToken(INTEGER_REGEX); s != nil {
		res.children = append(res.children, s)
		n, _ := strconv.Atoi(s.nValue)
		writer.writePush("constant", n)
	} else if tokens[currentToken].tType == "stringConstant" {
		s := matchToken(regexp.MustCompile("^*$"))
		res.children = append(res.children, s)
		writer.writePush("constant", len(s.nValue))
		writer.writeCall("String.new", 1)
		for _, c := range s.nValue {
			writer.writePush("constant", int(c))
			writer.writeCall("String.appendChar", 2)
		}
	} else if s := keywordConstant(); s != nil {
		res.children = append(res.children, s)
		switch s.nValue {
		case "true":
			writer.writePush("constant", 0)
			writer.writeArithmetic("not")
			break
		case "false":
			writer.writePush("constant", 0)
			break
		case "null":
			writer.writePush("constant", 0)
			break
		case "this":
			writer.writePush("pointer", 0)
		}
	} else if IDENTIFIER_REGEX.Match([]byte(tokens[currentToken].tValue)) {
		res.children = append(res.children, subroutineCall())
	} else if tokens[currentToken].tValue == "(" {
		res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
		res.children = append(res.children, expression())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
	} else if s := unaryOp(); s != nil {
		res.children = append(res.children, s)
		switch s.nValue {
		case "~":
			res.children = append(res.children, term())
			writer.writeArithmetic("not")
			break
		case "-":
			res.children = append(res.children, term())
			writer.writeArithmetic("neg")
			break
		}
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
	sName := subroutineName()
	name1 := sName.nValue
	count := 0
	res.children = append(res.children, sName)
	switch tokens[currentToken].tValue {
	case "(":
		res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
		writer.writePush("pointer", 0)
		count++
		res.children = append(res.children, expressionList())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
		if len(res.children[2].children) != 0 {
			count += len(res.children[2].children)/2 + 1
		}
		writer.writeCall(currentClassName+"."+name1, count)
		break
	case "[":
		res.children = append(res.children, matchToken(regexp.MustCompile("\\[")))
		res.children = append(res.children, expression())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\]")))
		if SubroutineSymbolTable.indexOf(name1) != -1 {
			writer.writePush(SubroutineSymbolTable.kindOf(name1), SubroutineSymbolTable.indexOf(name1))
		} else {
			writer.writePush(ClassSymbolTable.kindOf(name1), ClassSymbolTable.indexOf(name1))
		}
		writer.writeArithmetic("add")
		writer.writePop("pointer", 1)
		writer.writePush("that", 0)
		break
	case ".":
		res.children = append(res.children, matchToken(regexp.MustCompile("\\.")))
		name2 := subroutineName()
		res.children = append(res.children, name2)
		res.children = append(res.children, matchToken(regexp.MustCompile("\\(")))
		res.children = append(res.children, expressionList())
		res.children = append(res.children, matchToken(regexp.MustCompile("\\)")))
		if SubroutineSymbolTable.typeOf(name1) != "" {
			writer.writePush(SubroutineSymbolTable.kindOf(name1), SubroutineSymbolTable.indexOf(name1))
			count++
			name1 = SubroutineSymbolTable.typeOf(name1)
		} else if ClassSymbolTable.typeOf(name1) != "" {
			writer.writePush(ClassSymbolTable.kindOf(name1), ClassSymbolTable.indexOf(name1))
			count++
			name1 = ClassSymbolTable.typeOf(name1)
		}
		if len(res.children[4].children) != 0 {
			count += len(res.children[4].children)/2 + 1
		}
		writer.writeCall(name1+"."+name2.nValue, count)
		break
	default:
		if SubroutineSymbolTable.indexOf(name1) != -1 {
			writer.writePush(SubroutineSymbolTable.kindOf(name1), SubroutineSymbolTable.indexOf(name1))
		} else {
			writer.writePush(ClassSymbolTable.kindOf(name1), ClassSymbolTable.indexOf(name1))
		}
	}
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
		for s := matchToken(CommaRegex); s != nil; s = matchToken(CommaRegex) {
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
	for s := matchToken(CommaRegex); s != nil; s = matchToken(CommaRegex) {
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
		for s := matchToken(CommaRegex); s != nil; s = matchToken(CommaRegex) {
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

		for s = matchToken(CommaRegex); s != nil; s = matchToken(CommaRegex) {
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

func ParseToVM(file *os.File) {
	source, _ := os.ReadFile(file.Name())
	tokens = GetTokens(string(source))
	currentToken = 0
	ClassSymbolTable = newSymbolTable()
	SubroutineSymbolTable = newSymbolTable()
	writer.open(strings.TrimSuffix(file.Name(), ".jack") + ".vm")
	class()
	writer.close()
}
