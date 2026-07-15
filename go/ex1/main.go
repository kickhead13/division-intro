/*
 * This Go program is expected to print out this:
 * *-----------------------------*
 * | a = 2                       |
 * | b = 2                       |
 * *-----------------------------*
 * But instead it doesn't even compile... Try to fix this program
 * by reading the compile error that you get when you run
 * *-----------------------------*
 * | $ go run main.go            |
 * *-----------------------------*
 * in your terminal.
 *
 * Hint: Go is very strict about how you declare variables. There
 * are (at least) two different bugs hiding below, and the compiler
 * error messages will point you right at them if you read carefully.
 * */

package main

import "fmt"

func main() {

	var a = 1
	a += 1
	fmt.Println("a =", a)

	b := 1
	b := 2
	fmt.Println("b =", b)

}
