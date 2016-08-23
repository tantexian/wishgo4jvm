/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classpath

import (
	"os"
	"path/filepath"
)

/*
   Description:

   Author: tantexian
   Since: 2016/8/19
*/

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption string, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 解析boot及ext的classpath
// 关于self Classpath 与 self *Classpath 使用区别：
// 具体请参考：https://github.com/golang/go/wiki/CodeReviewComments#Receiver_Type
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	// 获取jre的路径或者JAVA_HOME/jre
	jreDir := getJreDir(jreOption)

	// 解析jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	// 将JAVA_HOME/jre/lib/*路径下面所有jar包路径解析赋值给bootClasspath
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// 解析jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	// 将JAVA_HOME/jre/lib//ext/*路径下面所有jar包路径解析赋值给extClasspath
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// 解析用户classpath，用户classpath为命令行参数中的输入（例如java.lang.Object）
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	// bootClasspath在parseBootAndExtClasspath方法中被赋值为self.bootClasspath = newWildcardEntry(jreLibPath)
	// 即compositeEntry，因此调用compositeEntry的readClass方法
	// 如果没有错误即err == nil 则说明找到class了
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func getJreDir(jreOption string) string {
	// 如果jre路径存在则直接返回该路径
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 如果./jre文件目录存在，则返回.jre
	jreStr := "./jre"
	if exists(jreStr) {
		return jreStr
	}
	// 如果JAVA_HOME 环境变量存在则返回JAVA_HOME/jre路径
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	// 如果上述都不存在触发找不到panic
	panic("Can not found jre folder!")
}

// 判断文件是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) String() string {
	str := "\n[self.bootClasspath ==> " + self.bootClasspath.String() +
			"\nself.extClasspath ==> " + self.extClasspath.String() +
			"\nself.userClasspath ==> " + self.userClasspath.String() +
			"]\n"
	return str
}
