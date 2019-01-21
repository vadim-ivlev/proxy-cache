import time
import requests



# Тестируемый URL

url = 'http://localhost:5555/hello'

# время тестирования в секундах
test_time = 30.0

# задержка между запросами в секундах
sleep_time = 0.1




i = 0
start_time = time.time()

while True:
    time.sleep(sleep_time)
    t0 = time.time()
    try:
        r = requests.get(url)
    except Exception as e:
        print(str(e))
        continue
    t1 = time.time()

    i += 1
    dt = t1-t0
    t = t1 - start_time
    cache_status = r.headers['X-Cache-Status']
 
    print(f'{i:>4} {r.status_code} time:{t:>8.2f} sec {cache_status:10} req time:{dt*1000:>5.2f}  Content length:{len(r.content):>7} bytes')

    if t > test_time: 
        break
