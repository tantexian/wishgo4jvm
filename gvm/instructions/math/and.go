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
   Description: 布尔运算指令只能操作int和long变量，分为按位与（and）、按位或（or）、按位异或（xor）3种。

   Author: tantexian
   Since: 2016/8/26
*/

// Boolean AND int
type IAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// Boolean AND long
type LAND struct{ base.NoOperandsInstruction }

func (self *LAND) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
