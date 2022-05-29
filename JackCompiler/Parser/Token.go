package Parser

type token struct {
	tType string
	value string
}

type xmlElement struct {
	xType    string
	value    string
	children []xmlElement
}
