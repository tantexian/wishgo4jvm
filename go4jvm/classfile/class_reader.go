/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classfile

import "encoding/binary"
/*
   Description: 读取class文件

   Author: tantexian
   Since: 2016/8/20
*/

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	// 由于只是一个字节，直接读取即可，没有大端小端区别
	// 每次从ClassReader结构体data中读取数据，然后将data去掉已被读取的数据
	val := self.data[0]
	self.data =self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {
	// class文件已大端方式存储，因此使用大端方式读取uint16，即两个字节
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// 相同类型的多条数据一般按表（table）的形式存储在class文件中。
// 表由表头和表项（item）构成，表头是u2或u4整数。假设表头是n，后面就紧跟着n个表项数据。
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

// 读取n个byte数据
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
