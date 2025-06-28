/*
 * This Rust program is expected to print out this:
 * *-----------------------------*
 * | a = 2                       |
 * | b = 2                       |
 * *-----------------------------*
 * But instead it doesn't even compile... Try to fix this program
 * by reading the compile error that you get when you run 
 * *-----------------------------*
 * | $ cargo run                 |
 * *-----------------------------*
 * in your terminal.
 * */

fn main() {

    let mut a = 1;
    a += 1;
    println!("a = {}", a);
    
    let b = 1;
    b += 1;
    println!("b = {}", b);

}
