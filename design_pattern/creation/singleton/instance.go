// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-22 23:19
package singleton

import (
	"sync"
	"sync/atomic"
)

var (
	instance Object
	once     Once
)

type Object struct{}

type Once struct {
	o int32
	sync.RWMutex
}

// GetInstance 获取实例
// doubleCheck
func GetInstance() Object {
	if atomic.LoadInt32(&once.o) == 0 {
		once.Lock()
		if atomic.LoadInt32(&once.o) == 0 {
			instance = Object{}
			atomic.StoreInt32(&once.o, 1)
		}
		once.Unlock()
	}
	return instance

}
