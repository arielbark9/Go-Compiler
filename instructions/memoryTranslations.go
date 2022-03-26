package instructions

var PushPointer0Set = []Instruction{
	A{Label: ThisLabel},
	C{Dest: "D", Comp: "M", Jump: ""},
	A{Label: SpLabel},
	C{Dest: "A", Comp: "M", Jump: ""},
	C{Dest: "M", Comp: "D", Jump: ""},
	A{Label: SpLabel},
	C{Dest: "M", Comp: "M+1", Jump: ""},
}

var PushPointer1Set = []Instruction{
	A{Label: ThatLabel},
	C{Dest: "D", Comp: "M", Jump: ""},
	A{Label: SpLabel},
	C{Dest: "A", Comp: "M", Jump: ""},
	C{Dest: "M", Comp: "D", Jump: ""},
	A{Label: SpLabel},
	C{Dest: "M", Comp: "M+1", Jump: ""},
}

var PopPointer0Set = []Instruction{
	A{Label: SpLabel},
	C{Dest: "A", Comp: "M-1", Jump: ""},
	C{Dest: "D", Comp: "M", Jump: ""},
	A{Label: ThisLabel},
	C{Dest: "M", Comp: "D", Jump: ""},
	A{Label: SpLabel},
	C{Dest: "M", Comp: "M-1", Jump: ""},
}

var PopPointer1Set = []Instruction{
	A{Label: SpLabel},
	C{Dest: "A", Comp: "M-1", Jump: ""},
	C{Dest: "D", Comp: "M", Jump: ""},
	A{Label: ThatLabel},
	C{Dest: "M", Comp: "D", Jump: ""},
	A{Label: SpLabel},
	C{Dest: "M", Comp: "M-1", Jump: ""},
}

func PushConstant(n int) []Instruction {
	var res []Instruction
	res = append(res, A{Num: n})
	res = append(res, C{Dest: "D", Comp: "A", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M+1", Jump: ""})
	return res
}

func PushLocal(n int) []Instruction {
	var res []Instruction
	res = append(res, A{Num: n})
	res = append(res, C{Dest: "D", Comp: "A", Jump: ""})
	res = append(res, A{Label: LclLabel})
	res = append(res, C{Dest: "A", Comp: "M+D", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M+1", Jump: ""})
	return res
}

func PushArgument(n int) []Instruction {
	var res []Instruction
	res = append(res, A{Num: n})
	res = append(res, C{Dest: "D", Comp: "A", Jump: ""})
	res = append(res, A{Label: ArgLabel})
	res = append(res, C{Dest: "A", Comp: "M+D", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M+1", Jump: ""})
	return res
}

func PushThis(n int) []Instruction {
	var res []Instruction
	res = append(res, A{Num: n})
	res = append(res, C{Dest: "D", Comp: "A", Jump: ""})
	res = append(res, A{Label: ThisLabel})
	res = append(res, C{Dest: "A", Comp: "M+D", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M+1", Jump: ""})
	return res
}

func PushThat(n int) []Instruction {
	var res []Instruction
	res = append(res, A{Num: n})
	res = append(res, C{Dest: "D", Comp: "A", Jump: ""})
	res = append(res, A{Label: ThatLabel})
	res = append(res, C{Dest: "A", Comp: "M+D", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M+1", Jump: ""})
	return res
}

func PushTemp(n int) []Instruction {
	var res []Instruction
	res = append(res, A{Num: n})
	res = append(res, C{Dest: "D", Comp: "A", Jump: ""})
	res = append(res, A{Num: 5})
	res = append(res, C{Dest: "A", Comp: "A+D", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M+1", Jump: ""})
	return res
}

func PushStatic(n int, fileName string) []Instruction {
	var res []Instruction
	FileName = LabelType(fileName)
	fileLabel, _ := NewLabel(FileName)
	fileLabel.ID = n
	res = append(res, A{Label: fileLabel})
	res = append(res, C{Dest: "D", Comp: "M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M+1", Jump: ""})
	return res
}

func PopLocal(n int) []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	// D now has top of stack val
	res = append(res, A{Label: LclLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	for i := 0; i < n; i++ { // increment until we reach LCL + n
		res = append(res, C{Dest: "A", Comp: "A+1", Jump: ""})
	}
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// decrement SP
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func PopArgument(n int) []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	// D now has top of stack val
	res = append(res, A{Label: ArgLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	for i := 0; i < n; i++ { // increment until we reach LCL + n
		res = append(res, C{Dest: "A", Comp: "A+1", Jump: ""})
	}
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// decrement SP
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func PopThis(n int) []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	// D now has top of stack val
	res = append(res, A{Label: ThisLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	for i := 0; i < n; i++ { // increment until we reach LCL + n
		res = append(res, C{Dest: "A", Comp: "A+1", Jump: ""})
	}
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// decrement SP
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func PopThat(n int) []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	// D now has top of stack val
	res = append(res, A{Label: ThatLabel})
	res = append(res, C{Dest: "A", Comp: "M", Jump: ""})
	for i := 0; i < n; i++ { // increment until we reach LCL + n
		res = append(res, C{Dest: "A", Comp: "A+1", Jump: ""})
	}
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// decrement SP
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func PopTemp(n int) []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	// D now has top of stack val
	res = append(res, A{Num: n + 5})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	// decrement SP
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func PopStatic(n int, fileName string) []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	FileName = LabelType(fileName)
	fileLabel, _ := NewLabel(FileName)
	fileLabel.ID = n
	res = append(res, A{Label: fileLabel})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}
