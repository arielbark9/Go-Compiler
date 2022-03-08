package arithmetic

import (
	ist "github.com/arielbark9/Go-Compiler/instructions"
	"strconv"
)

var getFirstVar = []ist.Instruction{
	ist.A{Val: "SP"},
	ist.C{Dest: "A", Comp: "M-1", Jump: ""},
	ist.C{Dest: "D", Comp: "M", Jump: ""},
}

func Add() []ist.Instruction {
	var res []ist.Instruction
	res = append(res, getFirstVar...)
	res = append(res, ist.C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, ist.C{Dest: "M", Comp: "D+M", Jump: ""})
	res = append(res, ist.A{Val: "SP"})
	res = append(res, ist.C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func PushConstant(n int) []ist.Instruction {
	var res []ist.Instruction
	res = append(res, ist.A{Val: strconv.Itoa(n)})
	res = append(res, ist.C{Dest: "D", Comp: "A", Jump: ""})
	res = append(res, ist.A{Val: "SP"})
	res = append(res, ist.C{Dest: "A", Comp: "M", Jump: ""})
	res = append(res, ist.C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, ist.A{Val: "SP"})
	res = append(res, ist.C{Dest: "M", Comp: "M+1", Jump: ""})
	return res
}

// TODO: lt, gt, or, not Achikam
// TODO: eq, and, neg, sub Ariel
