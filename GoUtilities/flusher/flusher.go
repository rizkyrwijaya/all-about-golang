package flusher

import (
	"sync"
	"time"
)

type Keys interface {
	string | int
}

type Values interface {
	string | int
}

type data[K Keys, V Values] struct {
	k K
	v V
}

type Flush[K Keys, V Values] struct {
	buffer map[K]V

	cycle   time.Duration
	rwLock  sync.Mutex
	inputCH chan data[K, V]
}

func (f *Flush[K, V]) AddData(k K, v V) {
	f.rwLock.Lock()
	defer f.rwLock.Unlock()
	f.inputCH <- data[K, V]{k, v}
}

func (f *Flush[K, V]) FlushData() {

}
