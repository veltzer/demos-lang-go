package main

import "fmt"

func main() {
    fmt.Println("Press Enter to continue...")

    // Wait for the user to press Enter
    _, err := fmt.Scanln()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("User pressed Enter")
}
