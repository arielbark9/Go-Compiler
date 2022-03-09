package instructions

import "strconv"

type A struct {
	Num   string
	Label label
}

func (a A) Translate() string {
	if a.Label.Name != Undefined {
		switch a.Label.Name {
		case IfTrue, IfFalse:
			return "@" + string(a.Label.Name) + strconv.Itoa(a.Label.ID)
		default:
			return "@" + string(a.Label.Name)
		}
	}
	return "@" + string(a.Num)
}
