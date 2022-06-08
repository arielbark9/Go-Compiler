package Parser

import (
	"os"
	"strconv"
	"strings"
)

type token struct {
	tType  string
	tValue string
}

type node struct {
	nType    string
	nValue   string
	children []*node
}

func (root *node) ToXml(tabCount int) string {
	if root == nil {
		return ""
	}
	tabs := ""
	for i := 0; i < tabCount; i++ {
		tabs += "  "
	}
	res := tabs + "<" + root.nType + ">"
	for _, child := range root.children {
		c := child.ToXml(tabCount + 1)
		if c != "" {
			res += "\n" + c
		}
	}
	if root.nValue != "" {
		res += " " + root.nValue + " " + "</" + root.nType + ">"
	} else {
		res += "\n" + tabs + "</" + root.nType + ">"
	}
	return res
}

type SymbolTableEntry struct {
	name  string
	kind  string
	type_ string
	index int
}

type SymbolTable struct {
	entries     []*SymbolTableEntry
	argIndex    int
	varIndex    int
	staticIndex int
	fieldIndex  int
}

func newSymbolTable() *SymbolTable {
	return &SymbolTable{[]*SymbolTableEntry{}, 0, 0, 0, 0}
}

func (table *SymbolTable) startSubroutine() {
	*table = SymbolTable{[]*SymbolTableEntry{}, 0, 0, 0, 0}
}

func (table *SymbolTable) define(name string, type_ string, kind string) {
	var entry SymbolTableEntry

	if kind == "ARGUMENT" {
		entry = SymbolTableEntry{name, kind, type_, table.argIndex}
		table.argIndex++
	} else if kind == "VAR" {
		entry = SymbolTableEntry{name, kind, type_, table.varIndex}
		table.varIndex++
	} else if kind == "STATIC" {
		entry = SymbolTableEntry{name, kind, type_, table.staticIndex}
		table.staticIndex++
	} else if kind == "FIELD" {
		entry = SymbolTableEntry{name, kind, type_, table.fieldIndex}
		table.fieldIndex++
	}

	table.entries = append(table.entries, &entry)
}

func (table *SymbolTable) varCount(kind string) int {
	count := 0
	for _, entry := range table.entries {
		if entry.kind == kind {
			count++
		}
	}
	return count
}

func (table *SymbolTable) kindOf(name string) string {
	for _, entry := range table.entries {
		if entry.name == name {
			return strings.ToLower(entry.kind)
		}
	}
	return ""
}

func (table *SymbolTable) typeOf(name string) string {
	for _, entry := range table.entries {
		if entry.name == name {
			return entry.type_
		}
	}
	return ""
}

func (table *SymbolTable) indexOf(name string) int {
	for _, entry := range table.entries {
		if entry.name == name {
			return entry.index
		}
	}
	return -1
}

type VMWriter struct {
	oFile      *os.File
	whileCount int
	ifCount    int
}

func (writer *VMWriter) open(fileName string) {
	writer.oFile, _ = os.Create(fileName)
	writer.whileCount = 0
	writer.ifCount = 0
}

func (writer *VMWriter) writePush(segment string, index int) {
	if segment == "var" {
		segment = "local"
	} else if segment == "field" {
		segment = "this"
	}
	writer.oFile.WriteString("push " + segment + " " + strconv.Itoa(index) + "\n")
}

func (writer *VMWriter) writePop(segment string, index int) {
	if segment == "var" {
		segment = "local"
	} else if segment == "field" {
		segment = "this"
	}
	writer.oFile.WriteString("pop " + segment + " " + strconv.Itoa(index) + "\n")
}

func (writer *VMWriter) writeArithmetic(command string) {
	writer.oFile.WriteString(command + "\n")
}

func (writer *VMWriter) writeLabel(label string) {
	writer.oFile.WriteString("label " + label + "\n")
}

func (writer *VMWriter) writeGoto(label string) {
	writer.oFile.WriteString("goto " + label + "\n")
}

func (writer *VMWriter) writeIf(label string) {
	writer.oFile.WriteString("if-goto " + label + "\n")
}

func (writer *VMWriter) writeCall(name string, nArgs int) {
	writer.oFile.WriteString("call " + name + " " + strconv.Itoa(nArgs) + "\n")
}

func (writer *VMWriter) writeFunction(className string, name string, nLocals int) {
	writer.oFile.WriteString("function " + className + "." + name + " " + strconv.Itoa(nLocals) + "\n")
}

func (writer *VMWriter) writeReturn() {
	writer.oFile.WriteString("return\n")
}

func (writer *VMWriter) close() {
	writer.oFile.Close()
}
