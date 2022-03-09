package instructions

type Comment struct {
	Text string
}

func (c Comment) Translate() string {
	return "// " + c.Text
}
