package main

import "fmt"

// Speaker is a simple interface with one method
type Speaker interface {
	Speak() string
}

// Dog implements the Speaker interface
type Dog struct {
	Name string
}

// Speak method for Dog
func (d Dog) Speak() string {
	return d.Name + " says: Woof!"
}

// Cat implements the Speaker interface
type Cat struct {
	Name string
}

// Speak method for Cat
func (c Cat) Speak() string {
	return c.Name + " says: Meow!"
}

// MakeSpeakTwice takes any Speaker and makes it speak twice
func MakeSpeakTwice(s Speaker) {
	fmt.Println(s.Speak())
	fmt.Println(s.Speak())
}

func main() {
	// Create instances
	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "Whiskers"}
	
	// Use the interface
	fmt.Println("Making animals speak:")
	MakeSpeakTwice(dog)
	fmt.Println()
	MakeSpeakTwice(cat)
}
