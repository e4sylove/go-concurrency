package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	type value struct {
		m sync.Mutex
		value int
	}

	var	wg sync.WaitGroup

	printSum := func (v1, v2 *value)  {
		defer wg.Done()

		v1.m.Lock()
		defer v1.m.Unlock()
		
		time.Sleep(time.Second * 2)

		v2.m.Lock()
		defer v2.m.Unlock()

		fmt.Printf("sum=%v\n", v1.value + v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b) 
	go printSum(&b, &a) 
	wg.Wait()	


}