/*
 * Yes! In Go, ":=" always declares at least one new variable, so
 * re-declaring "b" with ":=" a second time in the same scope is a
 * compile error. Using "b = 2" (a plain assignment) fixes it.
 *
 * Now let's try structs... Structs are like classes in a language
 * like C++ or Java (but much simpler...). In the next example we
 * ask you to implement a method for the struct Person that prints
 * "Hi! My name is <person's name>!". This method should be called
 * Meet() and be declared similar to the already implemented method
 * Greet().
 *
 * There's also a sneaky bug in the main() function... see if the
 * compiler helps you with that...
 * */

package main

import "fmt"

type Person struct {
	Name string
	Age  uint64
}

func (p Person) Greet() {
	fmt.Printf("I am %d years old!\n", p.Age)
}

func (p Person) Meet() {
	// implement: print "Hi! My name is <p.Name>!"
}

func main() {
	alex := Person{Name: "Alex", Age: "17"}
	alex.Meet()
	alex.Greet()
}
