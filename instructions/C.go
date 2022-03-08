package instructions

import "strings"

type C struct {
	Dest string
	Comp string
	Jump string
}

func (c C) Translate() string {
	res := strings.Builder{}
	if c.Dest != "null" {
		res.WriteString(c.Dest)
	}
	res.WriteString(" = " + c.Comp)
	if c.Jump != "null" {
		res.WriteString(";" + c.Jump)
	}
	return res.String()
}
