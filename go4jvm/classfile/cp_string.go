/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classfile

/*
   Description:

   Author: tantexian
   Since: 2016/8/21
*/

type ConstantStringInfo struct {
	cp         ConstantPool
	stingIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader){
	self.stingIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stingIndex)
}