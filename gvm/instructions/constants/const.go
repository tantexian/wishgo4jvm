/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package constants

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description: 这一系列指令把隐含在操作码中的常量值推入操作数栈顶。

   Author: tantexian
   Since: 2016/8/25
*/

// Push null
type ACONST_NULL struct{ base.NoOperandsInstruction }

func (self *ACONST_NULL) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushRef(nil)
}

// Push double
type DCONST_0 struct{ base.NoOperandsInstruction }

func (self *DCONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct{ base.NoOperandsInstruction }

func (self *DCONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// Push float
type FCONST_0 struct{ base.NoOperandsInstruction }

func (self *FCONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct{ base.NoOperandsInstruction }

func (self *FCONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct{ base.NoOperandsInstruction }

func (self *FCONST_2) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// Push int constant
type ICONST_M1 struct{ base.NoOperandsInstruction }

func (self *ICONST_M1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct{ base.NoOperandsInstruction }

func (self *ICONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct{ base.NoOperandsInstruction }

func (self *ICONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct{ base.NoOperandsInstruction }

func (self *ICONST_2) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct{ base.NoOperandsInstruction }

func (self *ICONST_3) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct{ base.NoOperandsInstruction }

func (self *ICONST_4) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct{ base.NoOperandsInstruction }

func (self *ICONST_5) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(5)
}

// Push long constant
type LCONST_0 struct{ base.NoOperandsInstruction }

func (self *LCONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct{ base.NoOperandsInstruction }

func (self *LCONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushLong(1)
}
