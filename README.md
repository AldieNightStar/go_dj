# Go DJ
### Mix your dependencies

Some sort of DI Container for Golang.

```go
package main

import (
	"fmt"
	"github.com/AldieNightStar/go_dj"
)

type Resolver func (a, b int16) int16
type Math func () (a, b int16)

func main() {
        // Create container and register all the components
        // In our case it will be functions which is assigned to a types
	c := go_dj.NewContainer()

        // Register "math" which has "add", "sub" as a dependencies
        // Then return function which is our provider, which then
        //   will return what we expect
	c.Register("math", func(args ... go_dj.Any) go_dj.Any {
		return Math(func() (a, b int16){
			add := args[0].(Resolver)
			sub := args[1].(Resolver)
			return add(15, 5), sub(15, 5)
		})
	}, "add", "sub")

	c.Register("add", func(args ... go_dj.Any) go_dj.Any {
		return Resolver(func (a, b int16) int16 {
			return a + b
		})
	})

	c.Register("sub", func(args ... go_dj.Any) go_dj.Any {
		return Resolver(func (a, b int16) int16 {
			return a - b
		})
	})
    
        // Provide Math from the Container 
	math, _ := c.Provide("math") // Can return error when item is not available or dependencies are unmet
	m := math.(Math)
    
        // Do some operations with the Math object
	a, b := m()
	fmt.Printf("%d %d\n", a, b)
}
```
