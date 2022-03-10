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

// only need to calc once because no labels or parameters

var AddSet = append(GetFirstVar, []Instruction{
	C{Dest: "A", Comp: "A-1", Jump: ""},
	C{Dest: "M", Comp: "D+M", Jump: ""},
	A{Label: SpLabel},
	C{Dest: "M", Comp: "M-1", Jump: ""},
}...)

var AndSet = append(GetFirstVar, []Instruction{
	C{Dest: "A", Comp: "A-1", Jump: ""},
	C{Dest: "M", Comp: "D&M", Jump: ""},
	A{Label: SpLabel},
	C{Dest: "M", Comp: "M-1", Jump: ""},
}...)

var OrSet = append(GetFirstVar, []Instruction{
	C{Dest: "M", Comp: "D|M", Jump: ""},
	A{Label: SpLabel},
	C{Dest: "A", Comp: "A-1", Jump: ""},
	C{Dest: "M", Comp: "M-1", Jump: ""},
}...)

var SubSet = append(GetFirstVar, []Instruction{
	C{Dest: "A", Comp: "A-1", Jump: ""},
	C{Dest: "M", Comp: "M-D", Jump: ""},
	A{Label: SpLabel},
	C{Dest: "M", Comp: "M-1", Jump: ""},
}...)

var NegSet = append(GetFirstVar[:2], C{Dest: "M", Comp: "-M", Jump: ""})

var NotSet = append(GetFirstVar[:2], C{Dest: "M", Comp: "!M", Jump: ""})
