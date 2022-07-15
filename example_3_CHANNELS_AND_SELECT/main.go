package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(1)
	wg := new(sync.WaitGroup)

	ch := make([]chan int, 3)
	ch[0] = make(chan int)
	ch[1] = make(chan int)
	ch[2] = make(chan int)

	for i := 0; i < 3; i++ {
		go func(i int) {
			for y := 0; y < 5; y++ {

				wg.Add(1) // отмечаем открытие всех доп горутин (почему-то их 5)
				ch[i] <- y
			}
		}(i)
	}

	time.Sleep(time.Second * 1)
	go func() {
		for {
			//wg.Add(1)
			select {
			case x := <-ch[0]:
				fmt.Printf("<%v> -> %v\n", 0, x)
			case x := <-ch[1]:
				fmt.Printf("<%v> -> %v\n", 1, x)
			case x := <-ch[2]:
				fmt.Printf("<%v> -> %v\n", 2, x)
			}
			wg.Done()
		}
	}()
	//wg.Done()
	wg.Wait()

	//подумать, как сделать без таймслипа
}
