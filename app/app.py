from bottle import Bottle, run
import time
import random
import os
import string
from pathlib import Path

log_file = 'app.log'

n = 0
app = Bottle()


def recreate_log(f):
    if os.path.exists(f):
        os.remove(f)
    Path(f).touch()



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



recreate_log(log_file)

@app.route('/')
@app.route('/hello')
def hello():
    time.sleep(1)
    s = generate_line()
    write_log(log_file, s)
    return s + generate_text()


run(app, host='0.0.0.0', port=8080, quiet=True)
