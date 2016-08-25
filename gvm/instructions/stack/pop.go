/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package stack

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description:

   Author: tantexian
   Since: 2016/8/25
*/

type POP struct{ base.NoOperandsInstruction }

type POP2 struct{ base.NoOperandsInstruction }

// 栈顶变量弹出,pop指令只能用于弹出int、float等占用一个操作数栈位置的变量。
func (self *POP) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// double和long变量在操作数栈中占据两个位置，需要使用pop2指令弹出
func (self *POP2) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
