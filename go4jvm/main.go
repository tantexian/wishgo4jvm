/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package main

import (
	"fmt"
	"strings"
	"wishjvm/go4jvm/classfile"
	"wishjvm/go4jvm/classpath"
)

/*
   Description:
   	在Go语言中，main是一个特殊的包，这个包所在的目录（可以叫作任何名字）会被编译为可执行文件。
   	Go程序的入口也是main（）函数，但是不接收任何参数，也不能有返回值。

   Author: tantexian
   Since: 2016/8/17
*/

// mian函数程序执行入口
func main() {
	// 解析命令行参数
	cmd := parseCmd()
	// 如果命令行参数中有-version，则打印当前版本
	if cmd.versionFlag {
		fmt.Println("version 1.0.0") //
	} else if cmd.helpFlag || cmd.class == "" {
		// 如果有-help参数或者没有指定具体的class则打印使用帮助
		printUsage()
	} else {
		// 启动java虚拟机
		//startJVM(cmd)
		startJVMAndPrint(cmd)
	}
}

// Cmd 启动参数 eg：go4jvm -Xjre "C:\Program Files\Java\jdk1.7.0_55\jre" java.lang.Object
func startJVM(cmd *Cmd) {
	// 解析所有classpath变量，其中cp为Classpath结构体，保存了bootClasspath、extClasspath、userClasspath路径信息
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// 打印命令行参数
	fmt.Printf("classpath：%v class：%v args：%v\n", cp, cmd.class, cmd.args)
	// 将class的所有.符号替换为/的文件目录路径
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Can't find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data：％v\n", classData)
}

func startJVMAndPrint(cmd *Cmd) {
	// 解析所有classpath变量，其中cp为Classpath结构体，保存了bootClasspath、extClasspath、userClasspath路径信息
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// 打印命令行参数
	fmt.Printf("classpath：%v class：%v args：%v\n", cp, cmd.class, cmd.args)
	// 将class的所有.符号替换为/的文件目录路径
	className := strings.Replace(cmd.class, ".", "/", -1)

	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile){
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}