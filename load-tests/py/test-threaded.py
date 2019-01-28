import random
import time
import threading
import requests
from totals import Totals


# Тестируемый URL
url: str = "http://localhost:5555/hello"

# время тестирования sec
test_time: int = 10

# задержка между запросами sec
request_interval: float = 0.1

# количество клиентов
num_clients: int = 500


totals = Totals()
sum_totals = Totals()
start_time = time.time()


def request():
    global totals
    while True:
        time.sleep(request_interval)
        try:
            r: requests.Response = requests.get(url)
        except requests.RequestException as e:
            print("EXP ", str(e))
            totals.inc_errors(1)
            continue

        totals.inc_count(1)
        totals.inc_bytes(len(r.content))
        if r.status_code != 200:
            print(f'status_code={r.status_code}')
        r.close()

        if time.time() - start_time > test_time:
            return


def printTotals():
    while True:
        time.sleep(1)
        print(totals)
        updateTotals()

        if time.time() - start_time > test_time:
            return


def updateTotals():
    sum_totals.inc_bytes(totals.bytes)
    sum_totals.inc_count(totals.count)
    sum_totals.inc_errors(totals.errors)
    totals.bytes = 0
    totals.count = 0
    totals.errors = 0


def main():
    threads = []

    threads.append(threading.Thread(target=printTotals))

    for i in range(num_clients):
        threads.append(threading.Thread(target=request))

    for t in threads:
        t.start()

    for t in threads:
        t.join()


    print('------------------------------------')
    updateTotals()
    print(totals)
    print(sum_totals)


if __name__ == "__main__":
    main()
