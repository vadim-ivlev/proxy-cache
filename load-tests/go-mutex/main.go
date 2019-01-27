package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Тестируемый URL
var url = "http://localhost:5555/"

// время тестирования sec
var test_time float64 = 10.0

// задержка между запросами sec
var request_interval = time.Second / 100

// количество клиентов
var num_clients = 100

var totals = new(Totals)
var sum_totals = new(Totals)
var start_time = time.Now()
var wg sync.WaitGroup

func request() {
	for {
		time.Sleep(request_interval)
		r, err := http.Get(url)
		if err != nil {
			// fmt.Println("Exp ", err.Error())
			totals.inc_errors(1)
			continue
		}

		totals.inc_count(1)
		totals.inc_bytes(r.ContentLength)

		if r.StatusCode != 200 {
			fmt.Printf("status_code=%d", r.StatusCode)
		}

		r.Body.Close()

		if (time.Now()).Sub(start_time).Seconds() > test_time {
			wg.Done()
			return
		}
	}
}

func printTotals() {
	for {
		time.Sleep(time.Second)
		totals.print()
		updateTotals()

		if (time.Now()).Sub(start_time).Seconds() > test_time {
			wg.Done()
			return
		}
	}
}

func updateTotals() {
	sum_totals.inc_bytes(totals.bytes)
	sum_totals.inc_count(totals.count)
	sum_totals.inc_errors(totals.errors)
	totals.bytes = 0
	totals.count = 0
	totals.errors = 0
}

func main() {
	fmt.Println("hello from main")

	wg.Add(1)
	go printTotals()

	for i := 0; i < num_clients; i++ {
		wg.Add(1)
		go request()
	}

	wg.Wait()

	fmt.Print("------------------------------------")
	updateTotals()
	totals.print()
	sum_totals.print()
}
