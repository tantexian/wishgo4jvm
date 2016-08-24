/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classfile

import (
	"fmt"
)

/*
   Description: 类文件数据结构

   Author: tantexian
   Since: 2016/8/21
*/
/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type ClassFile struct {
	// magic uint32（以该0xCAFEBABE魔数开头，代表了是java的class文件）
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []*MemberInfo
}

// 将byte数组数据解析成ClassFile结构体数据
func Parse(classData []byte) (classFile *ClassFile, err error) {
	// panic的函数并不会立刻返回，而是先defer，再返回
	// 一旦panic，逻辑就会调用defer 定义的func函数
	// recover函数将会捕获到当前的panic（如果有的话）
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)

			}
		}
	}()
	// 赋值class内存byte数组数据到ClassReader结构体
	cr := &ClassReader{classData}
	classFile = &ClassFile{}
	// 使用classFile中的read方法解析class的内存byte数据到ClassFile结构体中
	classFile.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	// 读取魔法数字即class的标识，验证是否为class文件
	self.readAndCheckMagic(reader)
	// 读取版本
	self.readAndCheckVersion(reader)
	// 读取常量池数据到ClassFile结构的constantpool变量中
	self.constantPool = readConstantPool(reader)
	// 读取权限访问标志
	self.accessFlags = reader.readUint16()
	// 获取当前类名信息
	self.thisClass = reader.readUint16()
	// 获取父类信息
	self.superClass = reader.readUint16()
	// 获取接口信息
	self.interfaces = reader.readUint16s()
	// 获取fields字段
	self.fields = readMembers(reader, self.constantPool)
	// 获取所有方法
	self.methods = readMembers(reader, self.constantPool)
	// 获取所有属性
	self.attributes = readMembers(reader, self.constantPool)
}

// 获取class文件的magic number
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 获取并检查版本
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	// 在J2SE 1.2之前是45，从1.2开始，每次有大的Java版本发布都增加1
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError！")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {

	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只有java.lang.Object没有超类
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
