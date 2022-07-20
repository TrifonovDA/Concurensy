package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := map[string]int{"касса": 1, "касса1": 1}
	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func(i int) {
			mu.Lock()
			cs["касса"] += 1
			mu.Unlock()
		}(i)
	}

	for i := 0; i < 1000; i++ {
		go func(i int) {
			mu.Lock()
			cs["касса1"] += 1
			mu.Unlock()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}

/*
Без mutex не работает так как map не потокобезопасен.

sync.Mutex.Lock() - все потоки поочередно будут ждать
sync.Mutex.Unlock() - все потоки работают ассинхронно, параллельно

два раза подряд Lock - получим дедлок

Unlock незалоченого мьютекса - паника

Mutex есть внутри каналов при чтении и записи

Есть еще RWMutex - сокращенно от Read Write
*/
