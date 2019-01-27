package main

import (
	"fmt"
	"sync"
)

// """ Totals class accumulates results of responses """
type Totals struct {
	sync.Mutex
	count  int
	errors int
	bytes  int64
}

func (t *Totals) inc_count(v int) {
	t.Lock()
	t.count += v
	t.Unlock()
}

func (t *Totals) inc_errors(v int) {
	t.Lock()
	t.errors += v
	t.Unlock()
}

func (t *Totals) inc_bytes(v int64) {
	t.Lock()
	t.bytes += v
	t.Unlock()
}

func (t *Totals) print() {
	fmt.Printf(" count:%d errors:%d bytes:%d\n", t.count, t.errors, t.bytes)
}
