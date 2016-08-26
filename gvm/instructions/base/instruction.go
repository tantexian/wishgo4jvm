/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package base

import (
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description:

   Author: tantexian
   Since: 2016/8/25
*/
type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// 执行指令逻辑
	Execute(frame *runtimedata.Frame)
}

type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// 由于是没有操作数的指令，因此不做任何事情
}

// 跳转指令
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/*
   Description: 存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出。
   		把这类指令抽象成Index8Instruction结构体，用Index字段表示局部变量表索引。

   Author: tantexian
   Since:  2016/8/25
*/
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadInt8())
}

/*
   Description: 有一些指令需要访问运行时常量池，常量池索引由两字节操作数给出。
   		把这类指令抽象成Index16Instruction结构体，用Index字段表示常量池索引。

   Author: tantexian
   Since:  2016/8/25
*/
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadInt16())
}
