package main

import (
	"fmt"
	"sync"
)

// 如果用*sync.Once的话需要初始化
var o *sync.Once

func init()  {
	o = new(sync.Once)
}

func SyncOnce(){
	for i:=0; i < 10 ; i++{
		o.Do(func() {
			fmt.Println(1)
		})
	}
}
