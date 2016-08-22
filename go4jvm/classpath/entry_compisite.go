/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
*/
package classpath

import (
	"errors"
	"strings"
)

/*
   Description:

   Author: tantexian
   Since: 2016/8/19
*/

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compsiteEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compsiteEntry = append(compsiteEntry, entry)
	}
	return compsiteEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err != nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)

}
