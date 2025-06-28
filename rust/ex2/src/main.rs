/*
 * Yes! Rust doesn't allow you to change variables by default,
 * you must specify that you are going to change a variable 
 * when you initialize it!
 *
 * Now let's try structs... Structs are like classes in a 
 * language like C++ or Java (but much simpler...). In the
 * next example we ask you to implement a method for the struct
 * Person that prints "Hi! My name is <person's name>!". This 
 * method should be called meet() and be declared similar to the
 * already implemented method age().
 *
 * There's also a sneaky bug in the main() function... see if the
 * compiler helps you with that...
 * */

struct Person {
    name: String,
    age: u64
}

impl Person {
    pub fn age(&self) {
        println!("I am {} years old!", self.age);
    }
    pub fn meet(&self) {
        //implement
    }
}


fn main() {
    let alex = Person { name: "Alex", age: 17 };
    alex.meet();
    alex.age();
}
