package instructions

func Goto(l label) []Instruction {
	var res []Instruction
	res = append(res, A{Label: l})
	res = append(res, C{Dest: "", Comp: "0", Jump: "JMP"})
	return res
}

func IfGoto(l label) []Instruction {
	var res []Instruction
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: l})
	res = append(res, C{Comp: "D", Jump: "JNE"})
	return res
}

func LabelDec(l label) []Instruction {
	var res []Instruction
	res = append(res, l)
	return res
}
