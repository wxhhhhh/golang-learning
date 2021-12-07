package main

import "sync"

type Obj struct {
	m sync.Locker
}

func NewObj() Obj {
	return Obj{
		m: &sync.Mutex{},
	}
}

func UseMu() {
	o := NewObj()
	o.m.Lock()
}
