package main

import "fmt"

// Generator: Functions that return a channel
func main() {
	c := boring("boring!")

	for i := 0; i < 5; i++ {
		fmt.Printf("you say: %q/n", <-c)
	}

	fmt.Println("You are boring, I'm leaving")
}

func boring(msg string) <-chan string {

	c := make(chan string)

	go func ()  {
		for i := 0; ;i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)
		}
	}()

	return c
}