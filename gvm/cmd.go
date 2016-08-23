/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

/*
   Description: java命令行工具实现

   Author: tantexian
   Since: 2016/8/17
*/

// 定义Cmd结构体,保存命令行信息
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

/*
   Description: 解析命令行参数

   Author: tantexian
   Since:  2016/8/17
*/
func parseCmd() *Cmd {
	cmd := &Cmd{}
	// 当发生错误，自动调用usage函数
	flag.Usage = printUsage
	// 从命令行参数中获取对应的值
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "jre path")
	flag.StringVar(&cmd.cpOption, "classpath", "", "class path")
	flag.StringVar(&cmd.cpOption, "cp", "", "class path") // cp为classpath的缩写，与classpath含义一致
	// 如果解析失败则调用flag.Usage，将命令的用法打印到控制台
	flag.Parse()
	// 如果上述解析成功，获取其他未解析参数
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	// 返回解析正常的命令行参数，数据结构变量
	return cmd
}

func printUsage() {
	fmt.Print("Usage : %s [-option] class [args...]\n", os.Args[0])
}
