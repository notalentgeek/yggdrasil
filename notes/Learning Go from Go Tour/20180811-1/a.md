* This is a note of me Learning Go from Go Tour from this link: [https://tour.golang.org/welcome/1](https://tour.golang.org/welcome/1).
* Here is the codes from the first Go Tour:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
``` 

* Here are my codes:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

* I expect that I can done at least 5 examples before I go to sleep.
* What happened if I change the double quote into single quote?
    * There will be an error.
* Below is the error caused when using single quote on multi-character string:

```
➜  tour-1-1 git:(feature/creating-python-walk-function) ✗ go run main.go
# command-line-arguments
./main.go:6:14: invalid character literal (more than one character)
```

* Apparently in Go single quote only works with single character.

```go
// Single quote will work on:
fmt.Println('a')

// Single quote will not work on:
fmt.Println('Hello World')
```

* Apparently, the 2nd to 3rd tour are just explanations.
* Here are the codes for the 4th tour:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
}
```

* Below are my adjusted codes:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Printf("The time is %s", time.Now())
}
```

* Please look carefully on `Printf()`. `Printf()` is a function to print formatted string.
* The tour for the first section ends here.
* Next, I would like to do another Go Tour that starts from here: [https://tour.golang.org/basics/1](https://tour.golang.org/basics/1).