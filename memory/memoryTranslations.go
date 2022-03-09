package memory

import (
	. "github.com/arielbark9/Go-Compiler/arithmetic"
	. "github.com/arielbark9/Go-Compiler/instructions"
	"strconv"
)

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
