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

type ConstantPool []ConstantInfo

// 表头给出的常量池大小n比实际大1
// 0为无效索引
// CONSTANT_Long_info 及 CONSTANT_Double_info各占两个位置
func readConstantPool(reader *ClassReader) ConstantPool {
	constantPoolCount := int(reader.readUint16())
	constantPool := make([]ConstantInfo, constantPoolCount)
	for i := 1; i < constantPoolCount; i++ {// 索引从1开始
		readConstantInfo(reader, constantPool)
		switch constantPool[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ //占两个位置
		}
	}
	return constantPool
}

// 此处不能使用self *ConstantPool，因为ConstantPool本来就是一个数组类型
// 其中ConstantPool为数组本身，*ConstantPool为指向数组的指针
func (self ConstantPool) getConstantInfo(idnex uint16) ConstantInfo {
	if cpInfo := self[idnex]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index！")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	constInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(constInfo.nameIndex)
	_type := self.getUtf8(constInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
