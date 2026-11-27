// vet: expect "expected operand" - the dots stand for code left out OMIT
package main

import "math/rand" // HL
import myrand "other/rand" // HL
import . "fmt" // HL

func main() {
	rand.Seed(...)
	myrand.Something(...)
	Println("Hello, world!")
}
