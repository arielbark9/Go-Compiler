package instructions

import "strconv"

type A struct {
	Num   string
	Label label
}

func (a A) Translate() string {
	if a.Label.Name != Undefined {
		return "@" + string(a.Label.Name) + strconv.Itoa(a.Label.ID)
	}
	return "@" + string(a.Num)
}
