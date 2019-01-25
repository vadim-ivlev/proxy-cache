package main

import (
	"fmt"
	"time"
)

// Тестируемый URL
var url = "http://localhost:5555/"

// время тестирования sec
var test_time = time.Second * 10

// задержка между запросами sec
var request_interval = time.Second / 10

// количество клиентов
var num_clients = 500

func main() {
	fmt.Println("hello from main")
}
