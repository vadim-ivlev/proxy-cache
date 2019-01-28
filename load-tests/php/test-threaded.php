<?php
require 'totals.php';



// Тестируемый URL
$url = 'http://localhost:5555/';

// время тестирования sec
$test_time  = 10.0;

// задержка между запросами sec
$request_interval = 0.01 ;

// количество клиентов
$num_clients = 100;

$totals = new Totals();
$sum_totals = new Totals();
$start_time = microtime();
// var wg sync.WaitGroup

function request() {
	while (true) {
        usleep(1000000 * $request_interval );

        $sh = curl_init($url);
        curl_setopt($sh, CURLOPT_RETURNTRANSFER, 1); 
        $r = curl_exec($ch);
        curl_close($ch);

        if($errno = curl_errno($ch)) {
            $error_message = curl_strerror($errno);
            echo "cURL error ({$errno}):\n {$error_message}";
            $totals -> inc_errors(1);
            continue;
        }
        

        
		$totals -> inc_count(1);
		$totals -> inc_bytes(strlen($r));

		// if r.StatusCode != 200 {
		// 	fmt.Printf("status_code=%d", r.StatusCode)
		// }

		if (microtime() - $start_time > $test_time ) {
			return;
		}
	}
}

function printTotals() {
	while (true) {
		sleep(1);
		$totals ->prnt();
		updateTotals();

		if (microtime() - $start_time > $test_time ) {
			return;
		}
	}
}

function updateTotals() {
	$sum_totals->inc_bytes($totals->bytes);
	$sum_totals->inc_count($totals->count);
    $sum_totals->inc_errors($totals->errors);

	$totals->bytes = 0;
	$totals->count = 0;
	$totals->errors = 0;
}

function main() {
	print("hello from main");

	// wg.Add(1)
	// go printTotals()

	for ($i = 0; $i < $num_clients; $i++) {
		// wg.Add(1)
		// go request()
	}
    
	// wg.Wait()

	print("------------------------------------");
	updateTotals();
	$totals -> prnt();
	$sum_totals->print();
}

