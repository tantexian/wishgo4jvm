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

type Slot struct {
	num int32 // 存放整数
	ref *Object // 存放引用
}



