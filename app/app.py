from bottle import Bottle, run
import time
import random
import os
import string

log_file = 'app.log'
n = 0

if os.path.exists(log_file):
    os.remove(log_file)

app = Bottle()


def generate_line():
    global n

    n += 1
    t = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
    return f'{n:<10}  {t} \n'


def generate_text(size=100, chars=string.ascii_uppercase):
    s = ''
    for _ in range(random.randint(5, 30)):
        s += ''.join(random.choice(chars) for _ in range(size)) + '\n'
    return s


def write_log(log_file, s):
    f = open(log_file, "a+")
    f.write(s)
    f.close()


@app.route('/')
@app.route('/hello')
def hello():
    time.sleep(5)
    s = generate_line()
    write_log(log_file, s)
    return s + generate_text()


run(app, host='0.0.0.0', port=8080)
