package stores

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

// Store reference into local variable
type ASTORE struct{ base.Index8Instruction }

func (self *ASTORE) Execute(frame *runtimedata.Frame) {
	_astore(frame, uint(self.Index))
}

type ASTORE_0 struct{ base.NoOperandsInstruction }

func (self *ASTORE_0) Execute(frame *runtimedata.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct{ base.NoOperandsInstruction }

func (self *ASTORE_1) Execute(frame *runtimedata.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct{ base.NoOperandsInstruction }

func (self *ASTORE_2) Execute(frame *runtimedata.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct{ base.NoOperandsInstruction }

func (self *ASTORE_3) Execute(frame *runtimedata.Frame) {
	_astore(frame, 3)
}

func _astore(frame *runtimedata.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
