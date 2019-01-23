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
var total Totals

func main() {
	summaryChannel = make(chan Summary, 100)
	defer close(summaryChannel)
	total := Totals{}

	start_time = time.Now()

	go analyze(&total)

	for i := 0; i < n_clients; i++ {
		go request(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Plain ++++++++++++++++++++++++++++++")

	time.Sleep(test_time)
	PrintTotals(total)
	fmt.Println("Good bye")
}

func analyze(total *Totals) {
	i := 0
	seconds := 0
	total1 := Totals{}

	for {
		i++
		s := <-summaryChannel
		t := (time.Now()).Sub(start_time).Seconds()

		total.bytes += s.contentLength
		total.count += 1

		total1.bytes += s.contentLength
		total1.count += 1

		if s.statusCode != 200 {
			total.errors += 1
			total1.errors += 1

			// fmt.Printf("%5d time:%5.2f  id_count=%3d_%04d  %4d  %8s  %5.2f ms  %4d b \n",
			// 	i, t, s.requesterId, s.count, s.statusCode, s.cacheStatus, s.responseTime, s.contentLength)
		}

		if int(t) != seconds {
			PrintTotals(total1)
			seconds = int(t)
			total1 = Totals{}
		}

	}

	fmt.Printf("%d requests analysed", i)

}

func PrintTotals(t Totals) {
	// fmt.Print("count: %d  errors, t.bytes", t.count, t.errors, t.bytes)
	fmt.Printf("%#v\n", t)
}

func request(id int) {
	i := 0
	for true {
		i++
		time.Sleep(sleep_time)

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
	}
}
