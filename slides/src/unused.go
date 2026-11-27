// vet: expect "declared and not used: x" - teaching: the unused variable and import are the lesson OMIT
package main

import (
	"fmt"
	"math"
)

func main() {
	x := 1
	fmt.Println("Hello!")
}
