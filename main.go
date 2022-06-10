package main

import (
	"fmt"
	"time"
)

func main() {
	

	// gouroutine 1
	c := make(chan interface{})

	go func() {

		for i := 0; i <= 10; i++ {

			
			time.Sleep(time.Second * 3)
			c <- i
		}
		
		// goroutine 2
	}()


	for i := 0; i <= 10; i++ {
		fmt.Println("blocking!")
		data := <- c
		fmt.Println(data)
	}

}