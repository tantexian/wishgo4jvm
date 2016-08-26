package loads

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

// Load long from local variable
type LLOAD struct{ base.Index8Instruction }

func (self *LLOAD) Execute(frame *runtimedata.Frame) {
	_lload(frame, uint(self.Index))
}

type LLOAD_0 struct{ base.NoOperandsInstruction }

func (self *LLOAD_0) Execute(frame *runtimedata.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOperandsInstruction }

func (self *LLOAD_1) Execute(frame *runtimedata.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOperandsInstruction }

func (self *LLOAD_2) Execute(frame *runtimedata.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOperandsInstruction }

func (self *LLOAD_3) Execute(frame *runtimedata.Frame) {
	_lload(frame, 3)
}

func _lload(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
