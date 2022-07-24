package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var memoryAccess sync.Mutex

	var value int

	go func ()  {
		fmt.Println("go routine")
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
		fmt.Println(value)

	}()

	memoryAccess.Lock()
	if value == 0 {
		fmt.Println("if else")
		fmt.Println(value)
	} else {
		fmt.Println(value)
	}
	memoryAccess.Unlock()
	
	time.Sleep(time.Second * 2)
}