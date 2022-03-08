package arithmetic

import (
	ist "github.com/arielbark9/Go-Compiler/instructions"
)

var getFirstVar = []ist.Instruction{
	ist.A{Val: "SP"},
	ist.C{Dest: "A", Comp: "M-1", Jump: "null"},
	ist.C{Dest: "D", Comp: "M", Jump: "null"},
}

func Add() []ist.Instruction {
	var res []ist.Instruction
	res = append(res, getFirstVar...)
	res = append(res, ist.C{Dest: "A", Comp: "A-1", Jump: "null"})
	res = append(res, ist.C{Dest: "M", Comp: "D+M", Jump: "null"})
	res = append(res, ist.A{Val: "SP"})
	res = append(res, ist.C{Dest: "M", Comp: "M-1", Jump: "null"})
	return res
}

func PushConstant(n int) []ist.Instruction {
	var res []ist.Instruction
	res = append(res, ist.A{Val: "SP"})
	res = append(res, ist.C{Dest: "A", Comp: "M", Jump: "null"})
	res = append(res, ist.C{Dest: "M", Comp: string(n), Jump: "null"})
	res = append(res, ist.A{Val: "SP"})
	res = append(res, ist.C{Dest: "M", Comp: "M-1", Jump: "null"})
	return res
}
