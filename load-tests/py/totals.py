import threading


class Totals:
    """ Totals class accumulates results of responses 
        from the system being tested.
    """

    def __init__(self):
        self.lock: threading.lock = threading.Lock()
        self.count: int = 0
        self.errors: int = 0
        self.bytes: int = 0
    
    def inc_count(self, v: int):
        self.lock.acquire()
        self.count += v
        self.lock.release()

    def inc_errors(self, v: int):
        self.lock.acquire()
        self.errors += v
        self.lock.release()

    def inc_bytes(self, v: int):
        self.lock.acquire()
        self.bytes += v
        self.lock.release()

    def __str__(self):
        return f' count:{self.count} errors:{self.errors} bytes:{self.bytes}'

