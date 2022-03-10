package instructions

type Instruction interface {
	Translate() string
}

var SpLabel, _ = NewLabel(SP)
var LclLabel, _ = NewLabel(LCL)
var ArgLabel, _ = NewLabel(ARG)
var ThisLabel, _ = NewLabel(THIS)
var ThatLabel, _ = NewLabel(THAT)

var GetFirstVar = []Instruction{
	A{Label: SpLabel},
	C{Dest: "A", Comp: "M-1", Jump: ""},
	C{Dest: "D", Comp: "M", Jump: ""},
}
