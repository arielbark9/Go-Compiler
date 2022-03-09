package logical

import (
	"github.com/arielbark9/Go-Compiler/arithmetic"
	ist "github.com/arielbark9/Go-Compiler/instructions"
)

func Eq() []ist.Instruction {
	var res []ist.Instruction
	res = append(res, arithmetic.GetFirstVar...)
	res = append(res, ist.C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, ist.C{Dest: "D", Comp: "D-M", Jump: ""})
	ifTrue0, _ := ist.NewLabel(ist.IfTrue)
	res = append(res, ist.A{Label: ifTrue0})
	res = append(res, ist.C{Dest: "", Comp: "D", Jump: "JEQ"})
	res = append(res, ist.C{Dest: "D", Comp: "0", Jump: ""})
	res = append(res, ist.A{Label: arithmetic.SpLabel})
	res = append(res, ist.C{Dest: "A", Comp: "M-1", Jump: ""})
	res = append(res, ist.C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, ist.C{Dest: "M", Comp: "D", Jump: ""})
	ifFalse0, _ := ist.NewLabel(ist.IfFalse)
	res = append(res, ist.A{Label: ifFalse0})
	res = append(res, ist.C{Dest: "", Comp: "0", Jump: "JMP"})
	res = append(res, ifTrue0)
	res = append(res, ist.C{Dest: "D", Comp: "-1", Jump: ""})
	res = append(res, ist.A{Label: arithmetic.SpLabel})
	res = append(res, ist.C{Dest: "A", Comp: "M-1", Jump: ""})
	res = append(res, ist.C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, ist.C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, ifFalse0)
	res = append(res, ist.A{Label: arithmetic.SpLabel})
	res = append(res, ist.C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}
