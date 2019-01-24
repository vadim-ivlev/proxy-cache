import threading
import time
import random


class Totals:
    pass

counter = 0
start_time = time.time()

lock = threading.Lock()

def incCounter(v):
    global  counter
    time.sleep(0.5)
    lock.acquire()
    counter+=v
    lock.release()
    print(f'incCounter +{v}')



def worker(n=0):
    print("worker begin")
    for i in range(10):
        time.sleep(0.01)
        r = random.randint(1,9)
        incCounter(1)
    print("worker done")
        # print(f'worker:{n}  {r}')



def printTotals():
    global  counter
    for i in range(10):
        time.sleep(5)
        print(f'counter={counter}')
        lock.acquire()
        counter = 0
        lock.release()





def main():

    threading.Thread(target=printTotals).start()
    for i in range(1):
        threading.Thread(target=worker, args=(i,)).start()
    print("done")

    


       



main()