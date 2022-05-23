package instructions

var functionReturnIdCounter = 0

func Call(name string, nArgs int) []Instruction {
	var res []Instruction
	functionName := LabelType(name)
	functionLabel, _ := NewLabel(functionName)
	functionReturn := LabelType(name + ".ReturnAddress")
	functionReturnLabel, _ := NewLabel(functionReturn)
	functionReturnLabel.ID = functionReturnIdCounter
	functionReturnIdCounter++
	// push return address
	res = append(res, A{Label: functionReturnLabel})
	res = append(res, C{Dest: "D", Comp: "A", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M+1", Jump: ""})
	res = append(res, PushLabel(LclLabel)...)
	res = append(res, PushLabel(ArgLabel)...)
	res = append(res, PushLabel(ThisLabel)...)
	res = append(res, PushLabel(ThatLabel)...)
	// arg = sp - nArgs - 5
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	x := nArgs + 5
	res = append(res, A{Num: x})
	res = append(res, C{Dest: "D", Comp: "D-A", Jump: ""})
	res = append(res, A{Label: ArgLabel})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// lcl = sp
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: LclLabel})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// goto name
	res = append(res, A{Label: functionLabel})
	res = append(res, C{Dest: "", Comp: "0", Jump: "JMP"})
	// insert return label
	res = append(res, functionReturnLabel)
	return res
}

func FunctionDeclaration(name string, nArgs int) []Instruction {
	var res []Instruction
	functionName := LabelType(name)
	functionLabel, _ := NewLabel(functionName)
	// insert function label
	res = append(res, functionLabel)
	// initialize local variables
	res = append(res, A{Num: nArgs})
	res = append(res, C{Dest: "D", Comp: "A", Jump: ""})
	// construct loop
	loopEnd := LabelType(name + ".LoopEnd")
	loopStart := LabelType(name + ".LoopStart")
	loopEndLabel, _ := NewLabel(loopEnd)
	loopStartLabel, _ := NewLabel(loopStart)
	// jmp to end if 0 variables
	res = append(res, A{Label: loopEndLabel})
	res = append(res, C{Dest: "", Comp: "D", Jump: "JEQ"})
	res = append(res, loopStartLabel)
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "0", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M+1", Jump: ""})
	res = append(res, A{Label: loopStartLabel})
	res = append(res, C{Dest: "D", Comp: "D-1", Jump: "JNE"})
	res = append(res, loopEndLabel)
	return res
}

func Return() []Instruction {
	var res []Instruction
	// Frame = LCL
	res = append(res, A{Label: LclLabel})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	// ret = *(Frame-5)
	// ram[13] = return address
	res = append(res, A{Num: 5})
	res = append(res, C{Dest: "A", Comp: "D-A", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Num: 13})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// *ARG = pop()
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: ArgLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// SP = ARG + 1
	res = append(res, A{Label: ArgLabel})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "D+1", Jump: ""})
	// THAT = *(Frame-1)
	res = append(res, A{Label: LclLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: ThatLabel})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// THIS = *(Frame-2)
	res = append(res, A{Label: LclLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: ThisLabel})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// ARG = *(Frame-3)
	res = append(res, A{Label: LclLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: ArgLabel})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// LCL = *(Frame-4)
	res = append(res, A{Label: LclLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: LclLabel})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// goto ret
	res = append(res, A{Num: 13})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "", Comp: "0", Jump: "JMP"})
	return res
}
