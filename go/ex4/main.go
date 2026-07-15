/*
 * The code below intends to print out a fresh random UUID. The
 * code itself looks good... it uses a function "New" from a
 * package called "uuid", that comes from a library called
 * "github.com/google/uuid" (imported with the "import" keyword,
 * similar to "use" in Rust or "include" in C)... Oh! Wait! We
 * didn't add it to our dependencies! The library "uuid" is not
 * part of the Go standard library... we must add it to the
 * go.mod file, which is the file that tracks the dependencies
 * of a Go module (the equivalent of Rust's Cargo.toml).
 *
 * 1. Try adding it manually by editing go.mod and go.sum yourself.
 * 2. Try using the "go get github.com/google/uuid" command to add
 *    it (and download it) automatically.
 * */

package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("uuid:", id.String())
}
