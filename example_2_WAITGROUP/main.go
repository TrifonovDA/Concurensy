package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("goroutine #{x} \n")
		}(x)
	}
	wg.Wait()
}

/*
на одном процессоре проблем нет
на 4х процессорах два раза идет захват единички, нет нуля
*/
