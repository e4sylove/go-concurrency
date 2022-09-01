package main

import (
	"fmt"
	"sync"
)

func main() {
	
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, v := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, v), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("This is something", &wg)

}


func printSomething(word string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(word)
}