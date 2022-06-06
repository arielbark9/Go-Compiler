package Parser

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
