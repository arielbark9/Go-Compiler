package instructions

import "strconv"

type A struct {
	Num   int
	Label label
}

func (a A) Translate() string {
	if a.Label.Name != Undefined {
		if a.Label.ID != -1 {
			return "@" + string(a.Label.Name) + "." + strconv.Itoa(a.Label.ID)
		} else {
			return "@" + string(a.Label.Name)
		}
	}
	return "@" + strconv.Itoa(a.Num)
}
