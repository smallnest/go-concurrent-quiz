package doublecheck

import (
	"sync"
)

type dummyObject struct {
	d int
}
type Singleton struct {
	a, b, c int
	dummy   *dummyObject
}

type Once struct {
	m    sync.Mutex
	done *Singleton
}

func (o *Once) Do(f func()) {
	if o.done != nil {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.done == nil {
		o.done = &Singleton{
			a:     1,
			b:     2,
			c:     3,
			dummy: &dummyObject{4},
		}
		f()
	}
}
