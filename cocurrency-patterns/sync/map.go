package main

import (
	"fmt"
	"sync"
)

var m *sync.Map

func NewSyncMap(){
	m = new(sync.Map)
	m.Store("a","b")
	c,_ := m.Load("a")
	m.LoadAndDelete("a")
	m.LoadOrStore()
	fmt.Println( c.(string))
	//m.LoadAndDelete()
}