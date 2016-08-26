package math

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

// Negate double
type DNEG struct{ base.NoOperandsInstruction }

func (self *DNEG) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type FNEG struct{ base.NoOperandsInstruction }

func (self *FNEG) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type INEG struct{ base.NoOperandsInstruction }

func (self *INEG) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type LNEG struct{ base.NoOperandsInstruction }

func (self *LNEG) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
