/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package runtimedata

/*
   Description:

   Author: tantexian
   Since: 2016/8/24
*/

// 若java虚拟机栈大小有限制，
type Thread struct {
	pc    int // 程序技术器，指向下一条需要执行的代码
	stack *Stack // Stack（Java虚拟机栈)指针。
}

// 如果Java虚拟机栈有大小限制，且执行线程所需的栈空间超出了这个限制，
// 会导致StackOverflowError异常抛出。
// 如果Java虚拟机栈可以动态扩展，但是内存已经耗尽，会导致OutOfMemor
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// 直接调用stack相应的方法
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame(frame *Frame) {
	self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}