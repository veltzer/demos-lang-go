package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Println("Press Enter to continue...")

    // Wait for the user to press Enter
    reader := bufio.NewReader(os.Stdin)
    _, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("User pressed Enter")
}
