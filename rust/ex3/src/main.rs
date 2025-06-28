/*
 * Yes! You solved it! A literal like "Alex" is of type 
 * &str and the struct Person was expecting a String object,
 * so you used "Alex".to_string() instead! Also congrats 
 * on fixing the meet() method... just adding a simple 
 * *---------------------------*
 * | println!("{}", self.name);|
 * *---------------------------*
 * did the trick... :)
 *
 * Let's complicate stuff... a lot... Generics...
 * Say you have a struct Person, and a method eat() for 
 * person, that takes a food object and calls it's ate()
 * method. To add some complexity we add many types of foods...
 * each implementing their own ate() method (apples and bread)
 * But all these new foods implement this one ate() method...
 * Having this in common we've also made an interface (or how
 * rust calls it - "trait") called Food that has one method (ate()).
 * This makes it easier to define the generics of the eat method.
 *
 * Try to understand the code below... And expand it... Add new foods.
 *
 * Hint: <T: Food> means that the function expects a generic type
 * T that implements the trait (or interface) Food.
 * */


trait Food {
    fn ate(&self);
}

struct Apple;
struct Bread;

impl Food for Apple {
    fn ate(&self) {
        println!("Ate an apple!");
    }
}

impl Food for Bread {
    fn ate(&self) {
        println!("Ate some bread!");
    }
}

struct Person;

impl Person {
    pub fn eat<T: Food>(&self, food: T) {
        food.ate();
    }
}

fn main() {
    let alex = Person;
    let apple = Apple;
    let bread = Bread;

    alex.eat(apple);
    alex.eat(bread);
}
