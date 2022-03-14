package instructions

import (
	"errors"
	"strconv"
)

var ifTrueID int = 0
var ifFalseID int = 0

type LabelType string

const (
	Undefined LabelType = ""
	SP                  = "SP"
	LCL                 = "LCL"
	ARG                 = "ARG"
	THIS                = "THIS"
	THAT                = "THAT"
	IfTrue              = "IF_TRUE"
	IfFalse             = "IF_FALSE"
)

var FileName LabelType

type label struct {
	Name LabelType
	ID   int
}

func NewLabel(name LabelType) (label, error) {
	l := label{Name: name}

	switch name {
	case IfTrue:
		l.ID = ifTrueID
		ifTrueID++
	case IfFalse:
		l.ID = ifFalseID
		ifFalseID++
	case Undefined:
		return label{}, errors.New("invalid label type")
	}

	return l, nil
}

func (l label) Translate() string {
	if l.Name != Undefined {
		return "(" + string(l.Name) + strconv.Itoa(l.ID) + ")"
	}
	return ""
}
