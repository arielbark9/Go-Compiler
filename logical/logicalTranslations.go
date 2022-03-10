package logical

import (
	. "github.com/arielbark9/Go-Compiler/instructions"
)

func Eq() []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "D-M", Jump: ""})
	ifTrue0, _ := NewLabel(IfTrue)
	res = append(res, A{Label: ifTrue0})
	res = append(res, C{Dest: "", Comp: "D", Jump: "JEQ"})
	res = append(res, C{Dest: "D", Comp: "0", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	ifFalse0, _ := NewLabel(IfFalse)
	res = append(res, A{Label: ifFalse0})
	res = append(res, C{Dest: "", Comp: "0", Jump: "JMP"})
	res = append(res, ifTrue0)
	res = append(res, C{Dest: "D", Comp: "-1", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, ifFalse0)
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func Lt() []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M-D", Jump: ""})
	ifTrue0, _ := NewLabel(IfTrue)
	res = append(res, A{Label: ifTrue0})
	res = append(res, C{Dest: "", Comp: "D", Jump: "JLT"})
	res = append(res, C{Dest: "D", Comp: "0", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	ifFalse0, _ := NewLabel(IfFalse)
	res = append(res, A{Label: ifFalse0})
	res = append(res, C{Dest: "", Comp: "0", Jump: "JMP"})
	res = append(res, ifTrue0)
	res = append(res, C{Dest: "D", Comp: "-1", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, ifFalse0)
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}
func Gt() []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "D", Comp: "M-D", Jump: ""})
	ifTrue0, _ := NewLabel(IfTrue)
	res = append(res, A{Label: ifTrue0})
	res = append(res, C{Dest: "", Comp: "D", Jump: "JGT"})
	res = append(res, C{Dest: "D", Comp: "0", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	ifFalse0, _ := NewLabel(IfFalse)
	res = append(res, A{Label: ifFalse0})
	res = append(res, C{Dest: "", Comp: "0", Jump: "JMP"})
	res = append(res, ifTrue0)
	res = append(res, C{Dest: "D", Comp: "-1", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "A", Comp: "M-1", Jump: ""})
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D", Jump: ""})
	res = append(res, ifFalse0)
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}
