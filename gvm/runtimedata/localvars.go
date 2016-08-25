/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package runtimedata

import "math"
/*
   Description:
   		关于boolean、byte、short和char类型存取方法，这些类型的值都可以转换成int值类来处理。

   Author: tantexian
   Since: 2016/8/24
*/

/*
    Description: 局部变量表的大小在编译期就被确定。
    		基元类型数据以及引用和返回地址（returnAddress）占用一个局部变量大小，
    		long/double需要两个。

    Author: tantexian
    Since:  2016/8/25
 */
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (self LocalVars)SetInt(index uint, val int32) {
	self[index].num = val
}

func (self LocalVars)GetInt(index uint) int32 {
	return self[index].num
}

// float需要先转换为int，再按照int处理
func (self LocalVars)SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

func (self LocalVars)GetFloat(index uint) float32 {
	uint32bits := uint32(self[index].num)
	return math.Float32frombits(uint32bits)
}

// long需要拆为两个int变量
func (self LocalVars)SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index + 1].num = int32(val >> 32)
}

func (self LocalVars)GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index + 1].num)
	return int64(high) << 32 | int64(low)
}

// double需要先转化为long，然后按照long处理
func (self LocalVars)SetDouble(index uint, val float64) {
	float64Val := math.Float64bits(val)
	self.SetLong(index, int64(float64Val))
}

func (self LocalVars)GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

// 引用值处理
func (self LocalVars)SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

func (self LocalVars)GetRef(index uint) *Object {
	return self[index].ref
}