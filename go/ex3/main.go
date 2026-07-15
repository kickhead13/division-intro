/*
 * Yes! You solved it! uint64 can't hold a string like "17", so you
 * either changed the field to a string or fixed the literal to
 * Age: 17. Also congrats on fixing Meet()... just adding a simple
 * *--------------------------------------------------*
 * | fmt.Printf("Hi! My name is %s!\n", p.Name)        |
 * *--------------------------------------------------*
 * did the trick... :)
 *
 * Let's complicate stuff... a lot... Generics...
 * Say you have a struct Person, and a method Eat() for Person,
 * that takes a food value and calls its Ate() method. To add some
 * complexity we add many types of foods... each implementing their
 * own Ate() method (apples and bread). But all these new foods
 * implement this one Ate() method... Having this in common we've
 * also made an interface called Food that has one method (Ate()).
 * This makes it easier to define the generics of the Eat method.
 *
 * Try to understand the code below... it won't compile as-is.
 *
 * Hint: unlike Rust (or Java/C++), Go methods are NOT allowed to
 * have their own type parameters -- only free-standing functions
 * can be generic. [T Food] on a method is illegal. You'll need to
 * restructure Eat so the generic type parameter lives on a function
 * instead of a method, while still calling it as naturally as
 * possible from main(). Once that compiles, add a new food (e.g.
 * Pizza) and make Person eat it too.
 * */

package main

import "fmt"

type Food interface {
	Ate()
}

type Apple struct{}
type Bread struct{}

func (Apple) Ate() { fmt.Println("Ate an apple!") }
func (Bread) Ate() { fmt.Println("Ate some bread!") }

type Person struct{}

func (p Person) Eat[T Food](food T) {
	food.Ate()
}

func main() {
	alex := Person{}
	apple := Apple{}
	bread := Bread{}

	alex.Eat(apple)
	alex.Eat(bread)
}
