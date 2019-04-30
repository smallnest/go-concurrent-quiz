## Go Concurrency Quizzes

### Quiz 1

<details>
 <summary><strong>Mutex quiz</strong></summary>

```go
package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var chain string

func main() {
	chain = "main"
	A()
	fmt.Println(chain)
}

func A() {
	mu.Lock()
	defer mu.Lock()
	chain = chain + " --> A"
	B()
}

func B() {
	chain = chain + " --> B"
	C()
}

func C() {
	mu.Lock()
	defer mu.Lock()
	chain = chain + " --> C"
}
```

</details>

**Question**

- [ ] A: can't compile
- [ ] B: output `main --> A --> B --> C`
- [ ] C: output `main`
- [ ] D: panic


### Quiz 2

<details>
 <summary><strong>RWMutex quiz</strong></summary>

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var count int

func main() {
	go A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}

func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}

func B() {
	time.Sleep(5 * time.Second)
	C()
}

func C() {
	mu.RLock()
	defer mu.RUnlock()
}
```

</details>

**Question**

- [ ] A: can't compile
- [ ] B: output `1`
- [ ] C: the program is hung
- [ ] D: panic

### Quiz 3

<details>
 <summary><strong>Waitgroup quiz</strong></summary>

```go
package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}
```

</details>

**Question**

- [ ] A: can't compile
- [ ] B: no output. exit as normal
- [ ] C: the program is hung
- [ ] D: panic

### Quiz 4

<details>
 <summary><strong>Double-checking quiz</strong></summary>

```go
package doublecheck

import (
	"sync"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}
```

</details>

**Question**

- [ ] A: can't compile
- [ ] B: can run that implemented the singleton pattern
- [ ] C: can run but has not implemented the singleton pattern
- [ ] D: panic

### Quiz 5

<details>
 <summary><strong>Mutex quiz</strong></summary>

```go
package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	sync.Mutex
}

func main() {
	var mu MyMutex

	mu.Lock()
	var mu2 = mu
	mu.count++
	mu.Unlock()

	mu2.Lock()
	mu2.count++
	mu2.Unlock()

	fmt.Println(mu.count, mu2.count)
}
```

</details>

**Question**

- [ ] A: can't compile
- [ ] B: output `1, 1`
- [ ] C: output `1, 2`
- [ ] D: panic

### Quiz 6

<details>
 <summary><strong>Pool quiz</strong></summary>

```go
package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func main() {
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}

	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}

func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}
```

</details>

**Question**

- [ ] A: can't compile
- [ ] B: can run and works as normal
- [ ] C: can run but has memory leak issue
- [ ] D: panic

### Quiz 7

<details>
 <summary><strong>Channel quiz</strong></summary>

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()

	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
```

</details>

**Question**

- [ ] A: can't compile
- [ ] B: after a while always output `#goroutines: 1`
- [ ] C: after a while always output `#goroutines: 2`
- [ ] D: panic

### Quiz 8

<details>
 <summary><strong>Channel quiz</strong></summary>

```go
package main

import "fmt"

func main() {
	var ch chan int
	var count int

	go func() {
		ch <- 1
	}()

	go func() {
		count++
		close(ch)
	}()

	<-ch

	fmt.Println(count)
}
```

</details>

**Question**

- [ ] A: can't compile
- [ ] B: output `1`
- [ ] C: output `0`
- [ ] D: panic


### Quiz 9

<details>
 <summary><strong>sync.Map quiz</strong></summary>

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")

	fmt.Println(m.Len())
}
```

</details>

**Question**

- [ ] A: can't compile
- [ ] B: output `1`
- [ ] C: output `0`
- [ ] D: panic

### Quiz 10

<details>
 <summary><strong>sync.Map quiz</strong></summary>

```go
package main

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}
func main() {
	go f()
	c <- 0
	print(a)
}
```

</details>

**Question**

- [ ] A: can't compile
- [ ] B: output `1`
- [ ] C: output `0`
- [ ] D: panic

### Quiz 11

<details>
 <summary><strong>concurrent map quiz</strong></summary>

```go
package main

import "sync"

type Map struct {
	m map[int]int
	sync.Mutex
}

func (m *Map) Get(key int) (int, bool) {
	m.Lock()
	defer m.Unlock()

	i, ok := m.m[key]
	return i, ok
}

func (m *Map) Put(key, value int) {
	m.Lock()
	defer m.Unlock()
	m.m[key] = value
}

func (m *Map) Len() int {
	return len(m.m)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	m := Map{m: make(map[int]int)}
	go func() {
		for i := 0; i < 10000000; i++ {
			m.Put(i, i)
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 10000000; i++ {
			m.Len()
		}

		wg.Done()
	}()

	wg.Wait()
}
```

run `go run quiz.go` to start this program.
</details>

**Question**

- [ ] A: can't compile
- [ ] B: can run but has no race issue
- [ ] C: can run but has race issue
- [ ] D: panic


### Quiz 12

<details>
 <summary><strong>slice quiz</strong></summary>

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var ints = make([]int, 0, 1000)

	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(len(ints))
}
```

run `go run quiz.go` to start this program.
</details>

**Question**

- [ ] A: can't compile
- [ ] B: output `2000`
- [ ] C: output a number but the number may not be `2000`
- [ ] D: panic



## Answer

<details>
 <summary><strong>Answer</strong></summary>

<p>
1. D 
2. D
3. D
4. C
5. D
6. C
7. C
8. D
9. A
10. B
11. C
12. C
</details>