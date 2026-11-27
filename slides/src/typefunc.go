// vet: expect "cannot use x" - teaching: the type error is the lesson OMIT
package main

// START OMIT
func myfunc(f float64) int {
	return int(f)
}

func main() {
	x := 3
	y := myfunc(x)
	println(y)
}

// END OMIT
