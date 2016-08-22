/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

/*
   Description: 扫描jar包或者zip包的类加载路径

   Author: tantexian
   Since: 2016/8/18
*/

type ZipEntry struct {
	// 存放jar或者zip的绝对路径
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) String() string {
	return self.absPath
}

// 从当前路径absPath所指向的zip或者jar包中读取className
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 获取zip或者jar包中的所有文件
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	// 编译当前zip或者jar包中的所有文件，是否与当前需要加载的className匹配
	for _, f := range r.File {
		println(f.Name)
		if f.Name == className {
			// 如果在类加载路径中，找到了需要加载的类，则直接打开该文件（即加载到内存中）
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			// ReadAll 读取 rc 中的所有数据
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil

		}
	}

	return nil, nil, errors.New("class not find:" + className)
}
