## Go Concurrency Quizzes

### Quiz 1

<details>
 <summary>Mutex code</summary>
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

** Question **

- A: can't compile
- B: output `main --> A --> B --> C`
- C: output `main`
- D: panic