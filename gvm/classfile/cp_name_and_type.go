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

type ConstantNameAndTypeInfo struct {
	nameIndex uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader){
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}