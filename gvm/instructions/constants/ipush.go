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
   Description:

   Author: tantexian
   Since: 2016/8/25
*/

/*
   Description: bipush指令从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶。

   Author: tantexian
   Since:  2016/8/25
*/
type BIPUSH struct {
	val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *runtimedata.Frame) {
	int32Val := int32(self.val)
	frame.OperandStack().PushInt(int32Val)
}

/*
   Description: sipush指令从操作数中获取一个short型整数，扩展成int型，然后推入栈顶。

   Author: tantexian
   Since:  2016/8/25
*/
type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *runtimedata.Frame) {
	int32Val := int32(self.val)
	frame.OperandStack().PushInt(int32Val)
}
