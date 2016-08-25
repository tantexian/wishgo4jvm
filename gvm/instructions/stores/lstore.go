/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package stores

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description:

   Author: tantexian
   Since: 2016/8/25
*/

// Store long into local variable
type LSTORE struct{ base.Index8Instruction }

func (self *LSTORE) Execute(frame *runtimedata.Frame) {
	_lstore(frame, uint(self.Index))
}

type LSTORE_0 struct{ base.NoOperandsInstruction }

func (self *LSTORE_0) Execute(frame *runtimedata.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct{ base.NoOperandsInstruction }

func (self *LSTORE_1) Execute(frame *runtimedata.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct{ base.NoOperandsInstruction }

func (self *LSTORE_2) Execute(frame *runtimedata.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct{ base.NoOperandsInstruction }

func (self *LSTORE_3) Execute(frame *runtimedata.Frame) {
	_lstore(frame, 3)
}

func _lstore(frame *runtimedata.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
