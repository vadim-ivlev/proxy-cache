import threading
import time
import random


class Totals:
    pass

def worker(n=0):
    for i in range(10):
        time.sleep(0.002)
        r = random.randint(1,9)
        print(f'worker:{n}  {r}')



def main():
    for i in range(3):
        threading.Thread(target=worker, args=(i,)).start()
    print("done")


       



main()