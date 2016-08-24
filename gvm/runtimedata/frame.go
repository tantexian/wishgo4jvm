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

type Frame struct {
	next *Frame // 用来实现链表数据结构
	localVars LocalVars // 保存局部变量表指针
	operandStack *OperandStack // 保存操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}