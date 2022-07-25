package main

import (
	"fmt"
	"time"
)

// these program make Joe and Ann count in lockstep

func main() {
	joe := boring("joe")
	ann := boring("ann")

	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	fmt.Println("You are boring, I'm leaving")
}

func boring(msg string) <-chan string {

	c := make(chan string)

	go func ()  {
		for i := 0; ;i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)
			time.Sleep(time.Duration(time.Second * 1))
		}
	}()

	return c
}