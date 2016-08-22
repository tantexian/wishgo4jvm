/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/*
   Description: 扫描带有通配符*的类加载路径

   Author: tantexian
   Since: 2016/8/19
*/

func newWildcardEntry(path string) CompositeEntry {
	// 去掉路径中的通配符*
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	// filepath.Walk回调函数
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		// 如果是文件目录，且为base目录，(返回SkipDir错误,未处理则相当于直接跳过)
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		// 如果为jar包则直接添加到compositeEntry中
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			// 所有的jar或者zip文件路径，都保存在ZipEntry结构体中
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// Walk遍历当前目录下所有文件及目录，回调wakkFn函数
	filepath.Walk(baseDir, walkFn)
	// 返回通配符路径下面所有的jar包加载路径
	return compositeEntry
}
