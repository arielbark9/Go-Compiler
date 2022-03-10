package memory

import (
	. "github.com/arielbark9/Go-Compiler/arithmetic"
	. "github.com/arielbark9/Go-Compiler/instructions"
	"strconv"
)

var PushPointer0Set = []Instruction{}
var PushPointer1Set = []Instruction{}
var PopPointer0Set = []Instruction{}
var PopPointer1Set = []Instruction{}

func PushConstant(n int) []Instruction {
	var res []Instruction
	res = append(res, A{Num: strconv.Itoa(n)})
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
	return res
}

func PushArgument(n int) []Instruction {
	var res []Instruction
	return res
}

func PushThis(n int) []Instruction {
	var res []Instruction
	return res
}

func PushThat(n int) []Instruction {
	var res []Instruction
	return res
}

func PushTemp(n int) []Instruction {
	var res []Instruction
	return res
}

func PushStatic(n int) []Instruction {
	var res []Instruction
	return res
}

func PopLocal(n int) []Instruction {
	var res []Instruction
	return res
}

func PopArgument(n int) []Instruction {
	var res []Instruction
	return res
}

func PopThis(n int) []Instruction {
	var res []Instruction
	return res
}

func PopThat(n int) []Instruction {
	var res []Instruction
	return res
}

func PopTemp(n int) []Instruction {
	var res []Instruction
	return res
}

func PopStatic(n int) []Instruction {
	var res []Instruction
	return res
}
