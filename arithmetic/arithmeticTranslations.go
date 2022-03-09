package arithmetic

import (
	. "github.com/arielbark9/Go-Compiler/instructions"
)

var GetFirstVar = []Instruction{
	A{Label: SpLabel},
	C{Dest: "A", Comp: "M-1", Jump: ""},
	C{Dest: "D", Comp: "M", Jump: ""},
}

var SpLabel, _ = NewLabel(SP)

func Add() []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D+M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

// TODO: lt, gt, or, not Achikam
// TODO: eq, and, neg, sub Ariel
