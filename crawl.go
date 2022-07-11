package main

import (
	"fmt"
	"time"
)


func main() {
	urlCh := make(chan int, 10)

	for i := 0; i < 5; i++ {
		go crawl(i, urlCh)
	}

	for j := 0; j < 10000; j++ {
		urlCh <- j
	}
}

func crawl(w int, ch chan int) {

	for i := range ch {
		time.Sleep(time.Second * 1)
		fmt.Printf("Woker %v crawled url %v\n", w, i)
	}
}


