package instructions

type Instruction interface {
	Translate() string
}
