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
   Description: iinc指令给局部变量表中的int变量增加常量值，局部变量表索引和常量值都由指令的操作数提供。

   Author: tantexian
   Since: 2016/8/26
*/

// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *runtimedata.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
