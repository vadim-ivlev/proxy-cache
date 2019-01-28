<?php
/**
 *  Totals class accumulates results of responses
 * from the system being tested.
 */
class Totals extends Thread {

    public $count = 0;
    public $errors = 0;
    public $bytes = 0;

    public function inc_count(int $v) {
        $this->lock();
        $this->count += $v;
        $this->unlock();
    }

    public function inc_errors(int $v) {
        $this->lock();
        $this->errors += $v;
        $this->unlock();
    }

    public function inc_bytes(int $v) {
        $this->lock();
        $this->bytes += $v;
        $this->unlock();
    }

    public function prnt() {
        print("count:{$this->count} errors:{$this->errors} bytes:{$this->bytes}");
    }

}
