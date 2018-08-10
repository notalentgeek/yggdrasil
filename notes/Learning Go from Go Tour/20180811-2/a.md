* Code:

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("My favorite number is %x \n", rand.Intn(10))
}
```

* If there is a `%` inside the output in the standard output, it should be because there is lack of line break. This is the same case with PHP if there are no `PHP_EOL`:

```
➜  tour-2-1-1 git:(feature/learning-go-from-go-tour) ✗ go run main.go
My favorite number is 1%➜  tour-2-1-1 git:(feature/learning-go-from-go-tour) ✗ go run main.go
My favorite number is 1%➜  tour-2-1-1 git:(feature/learning-go-from-go-tour) ✗ go run main.go
My favorite number is 1
➜  tour-2-1-1 git:(feature/learning-go-from-go-tour) ✗
```

* That problem above  (`1%`) can be solved with adding `\n` in the output to standard output.