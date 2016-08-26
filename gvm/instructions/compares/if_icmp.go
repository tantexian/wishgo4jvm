/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package compares

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description:

   Author: tantexian
   Since: 2016/8/26
*/

// Branch if int comparison succeeds
type IF_ICMPEQ struct{ base.BranchInstruction }

func (self *IF_ICMPEQ) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPNE struct{ base.BranchInstruction }

func (self *IF_ICMPNE) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLT struct{ base.BranchInstruction }

func (self *IF_ICMPLT) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLE struct{ base.BranchInstruction }

func (self *IF_ICMPLE) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPGT struct{ base.BranchInstruction }

func (self *IF_ICMPGT) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPGE struct{ base.BranchInstruction }

func (self *IF_ICMPGE) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
}

func _icmpPop(frame *runtimedata.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}
