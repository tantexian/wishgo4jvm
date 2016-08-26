/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package math

import (
	"math"
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description: 求余（rem）指令

   Author: tantexian
   Since: 2016/8/26
*/

/*
   Description: 先从操作数栈中弹出两个int值，然后再将两个值取余，再将结果push到操作数栈中

   Author: tantexian
   Since:  2016/8/26
*/
type IREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushInt(result)
}

// Remainder double
type DREM struct{ base.NoOperandsInstruction }

func (self *DREM) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2) // todo
	stack.PushDouble(result)
}

// Remainder float
type FREM struct{ base.NoOperandsInstruction }

func (self *FREM) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	stack.PushFloat(result)
}

// Remainder long
type LREM struct{ base.NoOperandsInstruction }

func (self *LREM) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}
