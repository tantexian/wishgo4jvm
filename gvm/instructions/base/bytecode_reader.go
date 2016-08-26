/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package base

/*
   Description:

   Author: tantexian
   Since: 2016/8/25
*/

type BytecodeReader struct {
	code []byte // 存放字节码
	pc   int    // 记录读取位置
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	val := self.code[self.pc]
	self.pc++
	return val
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

// 连续读取两个字节
func (self *BytecodeReader) ReadUint16() uint16 {
	high := uint16(self.ReadInt8())
	low := uint16(self.ReadInt8())
	return (high << 8) | low
}

// 连续读取两个字节
func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// 连续读取四个字节
func (self *BytecodeReader) ReadUint32() uint32 {
	byte1 := uint32(self.ReadInt8())
	byte2 := uint32(self.ReadInt8())
	byte3 := uint32(self.ReadInt8())
	byte4 := uint32(self.ReadInt8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// used by lookupswitch and tableswitch
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

// used by lookupswitch and tableswitch
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}
