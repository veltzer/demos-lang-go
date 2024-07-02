package main

import "fmt"

func main() {
    // For loop
    for i := 0; i < 5; i++ {
        fmt.Println("Iteration:", i)
    }

    // While loop (using for loop syntax)
    count := 0
    for count < 3 {
        fmt.Println("While loop iteration:", count)
        count++
    }
}
