package math

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

// Subtract double
type DSUB struct{ base.NoOperandsInstruction }

func (self *DSUB) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

// Subtract float
type FSUB struct{ base.NoOperandsInstruction }

func (self *FSUB) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

// Subtract int
type ISUB struct{ base.NoOperandsInstruction }

func (self *ISUB) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

// Subtract long
type LSUB struct{ base.NoOperandsInstruction }

func (self *LSUB) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
