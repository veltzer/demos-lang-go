package main

import "fmt"

// Dog represents a simple dog
type Dog struct {
	name  string
	breed string
	age   int
}

// NewDog is a constructor function for creating Dog instances
func NewDog(name string, breed string, age int) *Dog {
	return &Dog{
		name:  name,
		breed: breed,
		age:   age,
	}
}

// Bark makes the dog bark
func (d *Dog) Bark() {
	fmt.Printf("%s (a %s) says: Woof!\n", d.name, d.breed)
}

// GetName returns the dog's name
func (d *Dog) GetName() string {
	return d.name
}

// GetAge returns the dog's age
func (d *Dog) GetAge() int {
	return d.age
}

func main() {
	// Create a dog using the constructor
	myDog := NewDog("Rex", "German Shepherd", 3)
	
	// Use the dog methods
	fmt.Printf("My dog's name is %s and he is %d years old\n", myDog.GetName(), myDog.GetAge())
	myDog.Bark()
	
	// Create another dog
	otherDog := NewDog("Buddy", "Golden Retriever", 2)
	otherDog.Bark()
}
