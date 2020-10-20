package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	runtime.GOMAXPROCS(3)
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 10)

	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			chan1 <- i
			fmt.Println("chan1")
		}
	}()

	go func () {
		defer wg.Done()
		for i := 100; i < 200; i++ {
			chan2 <- i
			fmt.Println("chan2")
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case c := <-chan1:
				fmt.Printf("select from chan1 %d \n", c)
			case c := <-chan2:
				fmt.Printf("select from chan2 %d \n", c)

			}
		}
	}()

	wg.Wait()
}
