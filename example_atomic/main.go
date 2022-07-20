package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	value := int32(1)
	atomic.AddInt32(&value, 1)
	fmt.Println(value)
}

/*
Atomic - атомарные операции. Очень быстрые, происходят за один такт процессора.
atomic -потокобезопасный
*/
