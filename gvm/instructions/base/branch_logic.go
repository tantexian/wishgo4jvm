/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package base

import "wishgo4jvm/gvm/runtimedata"

/*
   Description:

   Author: tantexian
   Since: 2016/8/26
*/

func Branch(frame *runtimedata.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
