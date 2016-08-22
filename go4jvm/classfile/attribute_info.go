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

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributeCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}


func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo{
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code" : return &CodeAttribute{cp : cp}
	case "ConstantValue" : return &ConstantValueAttribute{}
	case "Deprecated" : return &DeprecatedAttribute{}
	case "Exceptions" : return &ExceptionsAttribute{}
	case "LineNumberTable" : return &LineNumberTableAttribute{}
	case "LocalVariableTable" : return &LocalVariableTableAttribute{}
	default : return &UnpasedAttribute{attrName, attrLen, nil}
	}
}
