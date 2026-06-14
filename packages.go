package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var first int = 10
	var second int = 20

	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now we have %g problems.\n", math.Sqrt(7))

	fmt.Println("Add function:", add(first, second))

	resultado := func(a, b int) int {
		return a + b
	}(first, 1)

	fmt.Println("Instant function: ", resultado)

	a := "meu ponteiro"
	p := &a

	//*p returns the p address value
	fmt.Println(a, p, *p)

	*p = "outro valor"
	fmt.Println(a, p, *p)
}

func add(x int, y int) int {
	return x + y
}
