/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classfile
/*
   Description:

   Author: tantexian
   Since: 2016/8/22
*/

type DeprecatedAttribute struct { MarkerAttribute }

type SyntheticAttribute struct { MarkerAttribute }

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {

}