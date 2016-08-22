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

type ConstantClassInfo struct {
	cp ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader){
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}