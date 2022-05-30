package Parser

type token struct {
	tType string
	value string
}

type xmlElement struct {
	xType    string
	value    string
	children []*xmlElement
}

func (x *xmlElement) ToXml(tabCount int) string {
	if x == nil {
		return ""
	}
	tabs := ""
	for i := 0; i < tabCount; i++ {
		tabs += "  "
	}
	res := tabs + "<" + x.xType + ">"
	for _, child := range x.children {
		c := child.ToXml(tabCount + 1)
		if c != "" {
			res += "\n" + c
		}
	}
	if x.value != "" {
		res += " " + x.value + " " + "</" + x.xType + ">"
	} else {
		res += "\n" + tabs + "</" + x.xType + ">"
	}
	return res
}
