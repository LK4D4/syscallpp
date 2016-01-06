Syscall Pretty Printer
======================

Helps you to convert syscall numbers to syscall names and back. Also can be
helpful for parsing ptrace.

# Example

```go
package main

import (
	"fmt"

	"github.com/LK4D4/syscallpp"
)

func main() {
	name := syscallpp.GetName(0)
	num := syscallpp.GetNum("readahead")
	fmt.Printf("%s: %d\n", name, 0)
	fmt.Printf("%s: %d\n", "readahead", num)
}
```
