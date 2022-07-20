package main

import (
	"sync"
	"sync/atomic"
	"time"
)

func main() {

}

func CalcSum(i int, j int) int {
	return i + j
}

func MutexCounter() int {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	gorotinesCount := 0

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			gorotinesCount++
			m.Unlock()
			time.Sleep(time.Microsecond)
			m.Lock()
			gorotinesCount--
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return gorotinesCount
}

func AtomicCounter() int32 {
	goroutinesCount := int32(0)
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&goroutinesCount, 1)
			time.Sleep(time.Microsecond)
			atomic.AddInt32(&goroutinesCount, -1)
			wg.Done()
		}()
	}
	wg.Wait()
	return goroutinesCount
}

/*
должно быть наоборот, но получился странный результат. на 1000 итераций выиграл почему то мьютекс, на 10000 выиграл атомик
*/
