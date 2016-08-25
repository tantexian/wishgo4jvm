/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package runtimedata

import "math"

/*
   Description: OperandStack操作数栈

   Author: tantexian
   Since: 2016/8/24
*/

/*
    Description: Frame被创建时，操作栈是空的。操作栈的每个项可以存放JVM的各种类型数据，包括long/double。
			操作栈有个栈深，long/double贡献两个栈深。
			操作栈调用其它有返回结果的方法时，会把结果push到栈上。
			更多详情请参考博文：http://my.oschina.net/tantexian/blog/737697

    Author: tantexian
    Since:  2016/8/25
*/
type OperandStack struct {
	size  uint // 记录栈顶位置
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

// 栈顶放置一个int
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

// long 使用两个slot
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}

// double 使用两个slot
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}
func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	// 置为nil，使得go垃圾回收器及时回收该空间
	self.slots[self.size].ref = nil
	return ref
}

// 栈指令直接对操作数栈进行操作
func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}
