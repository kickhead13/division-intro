/*
 * The code below intends to print out the hash of the string
 * "Hello World!". The code itself looks good... it uses a function
 * "digest" from a module called Sha256, that comes from a library
 * called "sha2"... Even the library is imported using the "use" 
 * keyword (similar to "include" in C)... Oh! Wait! We didn't add
 * it to our dependencies! The library "sha2" is not part of the
 * Rust standard library... we must add it to the Cargo.toml file
 * which is the file that tracks the dependencies of a Rust project.
 *
 * 1. Try adding it manually by editing the file your self.
 * 2. Try using the "cargo add" command to add it automatically.
 * */


use sha2::{Sha256, Digest};

fn main() {
    let hash = Sha256::digest(b"Hello World!");
    println!("hash (as list of integers): {:?}", hash);
}
