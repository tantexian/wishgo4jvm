/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package conversions

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description: 类型转换

   Author: tantexian
   Since: 2016/8/26
*/

// Convert double to float
type D2F struct{ base.NoOperandsInstruction }

func (self *D2F) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct{ base.NoOperandsInstruction }

func (self *D2I) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOperandsInstruction }

func (self *D2L) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
