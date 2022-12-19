# pager

Page through a slice.

```go
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/a-h/pager"
)

func main() {
	data := []string{"a", "b", "c", "d", "e", "f", "g"}
	pageSize := 2
	err := pager.Func(data, pageSize, func(page []string) error {
		fmt.Println(strings.Join(page, ", "))
		return nil
	})
	if err != nil {
		log.Fatal("failed to page through data: %v", err)
	}
}
```

Output:

```
a, b
c, d
e, f
g
```
