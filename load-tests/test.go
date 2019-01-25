package main

import (
	"fmt"
	"net/http"
	"time"
)

// # Тестируемый URL

var url = "http://localhost:5555/"

// # время тестирования в секундах
var test_time = 30.0

// # задержка между запросами в секундах
var sleep_time = 100 * time.Millisecond

func main() {

	i := 0
	start_time := time.Now()

	for true {
		time.Sleep(sleep_time)

		t0 := time.Now()
		r, err := http.Get(url)
		if err != nil {
			continue
		}
		t1 := time.Now()

		i++
		dt := t1.Sub(t0).Seconds() * 1000
		t := t1.Sub(start_time).Seconds()
		cache_status := r.Header["X-Cache-Status"][0]

		fmt.Printf("%4d %4d time:%6.2f sec %10s  Req time:%5.2f ms  Content length:%4d bytes \n",
			i, r.StatusCode, t, cache_status, dt, r.ContentLength)

		if t > test_time {
			break
		}

	}

}
