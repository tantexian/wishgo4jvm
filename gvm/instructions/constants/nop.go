/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package constants

import (
	"wishgo4jvm/gvm/instructions/base"
	"wishgo4jvm/gvm/runtimedata"
)

/*
   Description: 常量指令把常量推入操作数栈顶。
   		常量可以来自三个地方：隐含在操作码里、操作数和运行时常量池。

   Author: tantexian
   Since: 2016/8/25
*/

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *runtimedata.Frame) {
	// 空操作符，什么也不做
}
