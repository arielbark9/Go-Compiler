package instructions

type A struct {
	Val string
}

func (a A) Translate() string {
	return "@" + string(a.Val)
}
