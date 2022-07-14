package main

import "fmt"

type wait struct {
	count uint
	ch    chan interface{}
}

func initWait(count uint) *wait {
	return &wait{
		count: count,
		ch:    make(chan interface{}, count),
	}
}

func (w *wait) Done() {
	w.ch <- 1
}

func (w *wait) Wait() {
	defer close(w.ch)
	for {
		<-w.ch
		w.count--
		if w.count <= 0 {
			return
		}
	}
}
func test(id int, w *wait) {
	defer w.Done()
	fmt.Printf("<%v> done\n", id)
}
func main() {
	w := initWait(5)
	for i := 0; i < 5; i++ {
		go test(i, w)
	}
	w.Wait()
	fmt.Println("Done")
}
