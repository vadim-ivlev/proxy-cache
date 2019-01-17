from bottle import Bottle, run
import time
import random

n = 0

app = Bottle()



@app.route('/')
@app.route('/hello')
def hello():
    global n

    # n = random.randint(1,100)
    n += 1
    t = time.strftime("%Y-%m-%d %H:%M:%S", time.time())
    s1 = f'{n:30>}  {t} \n'
    time.sleep(10)
    return s1 


run(app, host='0.0.0.0', port=8080)
