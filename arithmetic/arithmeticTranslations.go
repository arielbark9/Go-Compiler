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

// only need to calc once because no labels

var AddSet = add()
var AndSet = and()
var OrSet = or()
var NegSet = neg()
var SubSet = sub()
var NotSet = not()

func add() []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D+M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res

}

func and() []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D&M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func or() []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "D|M", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func sub() []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar...)
	res = append(res, C{Dest: "A", Comp: "A-1", Jump: ""})
	res = append(res, C{Dest: "M", Comp: "M-D", Jump: ""})
	res = append(res, A{Label: SpLabel})
	res = append(res, C{Dest: "M", Comp: "M-1", Jump: ""})
	return res
}

func neg() []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar[:2]...)
	res = append(res, C{Dest: "M", Comp: "-M", Jump: ""})
	return res
}

func not() []Instruction {
	var res []Instruction
	res = append(res, GetFirstVar[:2]...)
	res = append(res, C{Dest: "M", Comp: "!M", Jump: ""})
	return res
}
