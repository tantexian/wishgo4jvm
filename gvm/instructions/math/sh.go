/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package math

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description: 先从操作数栈中弹出两个int变量v2和v1。
   		v1是要进行位移操作的变量，v2指出要移位多少比特。位移之后，把结果推入操作数栈。
   		这里注意两点：第一，int变量只有32位，所以只取v2的前5个比特就足够表表示位移位数了；
   		第二，Go语言位移操作符右侧必须是无符号整数，所以需要对v2进行类型转换。

   Author: tantexian
   Since: 2016/8/26
*/

// Shift left int
type ISHL struct{ base.NoOperandsInstruction }

func (self *ISHL) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// Arithmetic shift right int
type ISHR struct{ base.NoOperandsInstruction }

func (self *ISHR) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// Logical shift right int
type IUSHR struct{ base.NoOperandsInstruction }

func (self *IUSHR) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

// Shift left long
type LSHL struct{ base.NoOperandsInstruction }

func (self *LSHL) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// Arithmetic shift right long
type LSHR struct{ base.NoOperandsInstruction }

func (self *LSHR) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// Logical shift right long
type LUSHR struct{ base.NoOperandsInstruction }

func (self *LUSHR) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
