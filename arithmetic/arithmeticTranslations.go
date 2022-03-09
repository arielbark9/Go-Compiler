package arithmetic

import (
	ist "github.com/arielbark9/Go-Compiler/instructions"
	"strconv"
)

var GetFirstVar = []ist.Instruction{
	ist.A{Label: SpLabel},
	ist.C{Dest: "A", Comp: "M-1", Jump: ""},
	ist.C{Dest: "D", Comp: "M", Jump: ""},
}

var SpLabel, _ = ist.NewLabel(ist.SP)

func Add() []ist.Instruction {
	var res []ist.Instruction
	res = append(res, GetFirstVar...)
	res = append(res, ist.C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, ist.C{Dest: "M", Comp: "D+M", Jump: ""})
	res = append(res, ist.A{Label: SpLabel})
	res = append(res, ist.C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func PushConstant(n int) []ist.Instruction {
	var res []ist.Instruction
	res = append(res, ist.A{Num: strconv.Itoa(n)})
	res = append(res, ist.C{Dest: "D", Comp: "A", Jump: ""})
	res = append(res, ist.A{Label: SpLabel})
	res = append(res, ist.C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, ist.C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, ist.A{Label: SpLabel})
	res = append(res, ist.C{Dest: "M", Comp: "M+1", Jump: ""})
	return res
}

// TODO: lt, gt, or, not Achikam
// TODO: eq, and, neg, sub Ariel
