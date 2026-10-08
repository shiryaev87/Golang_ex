package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const numRequests = 1000

var count int64 // atomic требует int64

func networkRequest() {
	time.Sleep(time.Millisecond)
	atomic.AddInt64(&count, 1) // Используем атомарный инкремент
}

func main() {
	defer timer()()

	var wg sync.WaitGroup
	wg.Add(numRequests) // Говорим WaitGroup, сколько горутин мы ждём

	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done() // Сообщаем, что горутина завершилась
			networkRequest()
		}()
	}

	wg.Wait() // Блокируем main, пока все горутины не завершатся
}

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("count %v took %v\n", atomic.LoadInt64(&count), time.Since(start)) // Читаем счётчик атомарно
	}
}
