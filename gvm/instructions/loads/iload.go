/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package loads

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description: 加载指令从局部变量表获取变量，然后推入操作数栈顶。

   Author: tantexian
   Since: 2016/8/25
*/

// Load int from local variable
type ILOAD struct{ base.Index8Instruction }

type ILOAD_0 struct{ base.Index8Instruction }

type ILOAD_1 struct{ base.Index8Instruction }

type ILOAD_2 struct{ base.Index8Instruction }

type ILOAD_3 struct{ base.Index8Instruction }

func _iload(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (self *ILOAD) Execute(frame *runtimedata.Frame) {
	_iload(frame, uint(self.Index))
}

func (self *ILOAD_0) Execute(frame *runtimedata.Frame) {
	_iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *runtimedata.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *runtimedata.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *runtimedata.Frame) {
	_iload(frame, 3)
}
