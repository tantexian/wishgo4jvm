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
type UnpasedAttribute struct {
	name string
	length uint32
	info []byte
}

func (self * UnpasedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}