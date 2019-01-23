package main

import (
	"fmt"
	"net/http"
	"time"
)

// # Тестируемый URL
var url = "http://localhost:5555/"

// # время тестирования в секундах
var test_time = 10 * time.Second

// # задержка между запросами
var sleep_time = 10 * time.Millisecond

// количество клиентов
var n_clients = 10

var start_time time.Time

type Summary struct {
	requesterId   int
	count         int
	statusCode    int
	cacheStatus   string
	responseTime  float64
	contentLength int64
}

type Totals struct {
	count  int
	errors int
	bytes  int64
}

var summaryChannel chan Summary

func main() {
	summaryChannel = make(chan Summary, 100)
	defer close(summaryChannel)

	start_time = time.Now()

	go analyze()

	for i := 0; i < n_clients; i++ {
		go request(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Plain ++++++++++++++++++++++++++++ n_clients = ", n_clients)
	time.Sleep(test_time)

}

func analyze() {
	i := 0
	seconds := 0.0
	total := Totals{}

	for {
		i++
		s := <-summaryChannel
		t := (time.Now()).Sub(start_time).Seconds()

		total.bytes += s.contentLength
		total.count += 1

		if s.statusCode != 200 {
			total.errors += 1
		}

		if t-seconds > 1.0 {
			fmt.Printf("%#v\n", total)
			seconds = t
			total = Totals{}
		}

	}

	fmt.Printf("%d requests analysed", i)

}

func request(id int) {
	i := 0
	for true {
		i++
		t0 := time.Now()
		r, err := http.Get(url)
		t1 := time.Now()

		rs := Summary{
			requesterId:  id,
			count:        i,
			responseTime: t1.Sub(t0).Seconds() * 1000,
		}
		if err == nil {
			rs.statusCode = r.StatusCode
			rs.cacheStatus = r.Header["X-Cache-Status"][0]
			rs.contentLength = r.ContentLength
		} else {
			// fmt.Println(err.Error())
		}
		r.Body.Close()
		summaryChannel <- rs
		time.Sleep(sleep_time)
	}
}
