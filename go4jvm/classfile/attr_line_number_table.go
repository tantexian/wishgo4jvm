/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classfile
/*
   Description:

   Author: tantexian
   Since: 2016/8/22
*/
type LineNumberTableAttribute struct {
	LineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLen := reader.readUint16()
	self.LineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLen)
	for i := range self.LineNumberTable {
		self.LineNumberTable[i] = &LineNumberTableEntry{
			startPc: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}