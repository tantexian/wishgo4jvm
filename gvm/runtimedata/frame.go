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

/*
    Description: 结构：{Frame [ReturnValue] [LocalVariables[][][][]...] [OperandStack [][][]...] [ConstPoolRef] }
		每次方法调用均会创建一个对应的Frame，方法执行完毕或者异常终止，Frame被销毁。
		一个方法A调用另一个方法B时，A的frame停止，新的frame被创建赋予B，
		执行完毕后，把计算结果传递给A，A继续执行。

    Author: tantexian
    Since:  2016/8/25
*/
type Frame struct {
	next         *Frame        // 用来实现链表数据结构
	localVars    LocalVars     // 保存局部变量表指针
	operandStack *OperandStack // 保存操作数栈指针
	thread       *Thread
	nextPC       int // the next instruction after the call
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func NewFrameNoThread(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
