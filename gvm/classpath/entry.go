/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classpath

import (
	"os"
	"strings"
)

/*
   Description: xxx

   Author: tantexian
   Since: 2016/8/18
*/

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 读取类文件
	readClass(className string) ([]byte, Entry, error)
	// 类似java的toString方法
	String() string
}

func newEntry(path string) Entry {
	// 如果path路径中包括路径分隔符，即多个路径用";"分开，则创建组合entry
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	// 如果path路径中包括*，则创建通配符entry
	if strings.HasPrefix(path, "*") {
		return newWildcardEntry(path)
	}

	// 如果path路径中包括路.jar,.zip，则创建zip的entry
	if strings.HasPrefix(path, ".jar") || strings.HasPrefix(path, ".JAR") ||
		strings.HasPrefix(path, ".zip") || strings.HasPrefix(path, ".ZIP") {
		return newZipEntry(path)
	}

	// 如果上述条件都不满足，则创建dir的entry
	return newDirEntry(path)
}
