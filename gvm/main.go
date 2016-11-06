/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package main

import (
	"fmt"
	"strings"
	"wishgo4jvm/gvm/classfile"
	"wishgo4jvm/gvm/classpath"
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description:
   	在Go语言中，main是一个特殊的包，这个包所在的目录（可以叫作任何名字）会被编译为可执行文件。
   	Go程序的入口也是main（）函数，但是不接收任何参数，也不能有返回值。

   	注意：1/若希望使用idea的run/debug功能，则选择+ Go Application
   	     2/然后Package设置为wishgo4jvm/gvm即可
   	     3/启动参数放置到Program arguments项


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
		// 启动java虚拟机，将对应的class加载到内存中
		// startJVM(cmd)

		// 启动JVM并解析class文件，打印class对于信息
		//startJVMAndPrint(cmd)

		// 启动JVM，模拟操作数栈
		//startJVMWithRuntimeDataArea(cmd)

		// 启动JVM，解析class文件为对应指令集
		startJVMWithInstructions(cmd)
	}
}

/*
    Description: 此函数启动jvm，将对应的class加载到内存中
			Cmd 启动参数 eg：gvm.exe -Xjre "C:\Program Files\Java\jdk1.8.0_92\jre" java.lang.String

    Author: tantexian
    Since:  2016/8/25
*/
func startJVM(cmd *Cmd) {
	// 解析所有classpath变量，其中cp为Classpath结构体，保存了bootClasspath、extClasspath、userClasspath路径信息
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// 打印命令行参数
	fmt.Printf("classpath：%v \nclass：%v \n\nargs：%v\n\n", cp, cmd.class, cmd.args)

	//fmt.Println(cp.String())

	// 将class的所有.符号替换为/的文件目录路径
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Can't find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data：％v\n", classData)
}

/*
    Description: 此函数根据java虚拟机规范，解析对应的class文件，并打印出关键信息
		代码编译：go build wishgo4jvm\gvm
		Cmd 启动参数 eg: gvm.exe -Xjre "C:\Program Files\Java\jdk1.8.0_92\jre" java.lang.String

    Author: tantexian
    Since:  2016/8/25
*/
func startJVMAndPrint(cmd *Cmd) {
	// 解析所有classpath变量，其中cp为Classpath结构体，保存了bootClasspath、extClasspath、userClasspath路径信息
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// 打印命令行参数
	fmt.Printf("classpath：%v \nclass：%v \n\nargs：%v\n\n", cp, cmd.class, cmd.args)

	// 将class的所有.符号替换为/的文件目录路径
	className := strings.Replace(cmd.class, ".", "/", -1)

	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	// 根据className在jre中寻找对于的class文件，读取到内存中为classDate
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Print("Classpath ReadClass err ==> ")
		panic(err)
	}
	// 根据java虚拟机规范，解析classDate（class类的byte数组）
	cf, err := classfile.Parse(classData)
	if err != nil {
		fmt.Print("Java class file Parse err ==> ")
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
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

/*
    Description: 此函数根据java虚拟机规范，模拟runtime data area 的栈及栈帧相关使用
		代码编译：go build wishgo4jvm\gvm
		Cmd 启动参数 eg: gvm.exe -Xjre "C:\Program Files\Java\jdk1.8.0_92\jre" java.lang.String

    Author: tantexian
    Since:  2016/8/25
*/
func startJVMWithRuntimeDataArea(cmd *Cmd) {
	frame := runtimedata.NewFrameNoThread(100, 100)
	println("Test LocalVars set and get methods:")
	tesLocalVars(frame.LocalVars())
	println("\nTest OperandStack push and pop methods:")
	testOperandStack(frame.OperandStack())
}

func tesLocalVars(localvals runtimedata.LocalVars) {
	// int 、float占据一个索引位置
	localvals.SetInt(0, 99)
	localvals.SetInt(1, -99)
	localvals.SetFloat(2, 3.1415926)
	// long、double各占据两个索引位置
	localvals.SetLong(3, 1234)
	localvals.SetLong(5, -214748364)
	localvals.SetDouble(7, -3.14159265358979)
	obj := &runtimedata.Object{
		Name: "ttx",
	}
	localvals.SetRef(8, obj)

	println(localvals.GetInt(0))
	println(localvals.GetInt(1))
	println(localvals.GetFloat(2))
	println(localvals.GetLong(3))
	println(localvals.GetLong(5))
	println(localvals.GetDouble(7))
	println(localvals.GetRef(8).Name)

}

func testOperandStack(operandStack *runtimedata.OperandStack) {
	operandStack.PushInt(100)
	operandStack.PushInt(-100)
	operandStack.PushFloat(3.14159)
	operandStack.PushDouble(3.1415926535)
	obj := &runtimedata.Object{
		Name: "ttx",
	}
	operandStack.PushRef(obj)

	// 依次逆序反着对应类型出栈
	println(operandStack.PopRef().Name)
	println(operandStack.PopDouble())
	println(operandStack.PopFloat())
	println(operandStack.PopInt())
	println(operandStack.PopInt())
}

/*
    Description: 此函数根据java虚拟机规范，加入java指令集Instructions解析功能
		代码编译：go build wishgo4jvm\gvm
		Cmd 启动参数 eg: gvm.exe -Xjre "C:\Program Files\Java\jdk1.8.0_92\jre" -cp "./resource" com.hello.HelloWord
		注：1、使用javac将resource目录下的HelloWord.java编程成class文件
		    2、将1中编译得到的HelloWord.class文件放置到路径./resource/com/hello目录下

    Author: tantexian
    Since:  2016/8/25
*/
func startJVMWithInstructions(cmd *Cmd) {
	// 解析所有classpath变量，其中cp为Classpath结构体，保存了bootClasspath、extClasspath、userClasspath路径信息
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classFile := loadClass(className, cp)
	mainMethod := getMainMethod(classFile)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class ％s\n", cmd.class)
	}
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
