/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classpath

import (
	"io/ioutil"
	"path/filepath"
)

/*
   Description: 扫描路径形式的类加载路径

   Author: tantexian
   Since: 2016/8/18
*/

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	// 将参数转换成绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	// 直接从当前绝对路径中读取文件
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
