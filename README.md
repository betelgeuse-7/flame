### Flame Compiler to Go

This is the Flame language:

(The language will certainly change. A specification must be written.)

#### Variable declaration
```go
string name = "Jennifer"
uint age = 44
```
#### Constant declaration
```go
const string name = "Jennifer"
const uint age = 44
```
#### Types
```go
void
string
bool
int int32
uint uint32
float64 float32
```
- slices
- maps
- structs
- pointers

### Functions
```go
const void helloWorld = () => {
    println("hello world")
}
const string greetMe = (string name) => "hello ${name}"

println(greetMe("betelgeuse-7")) // hello betelgeuse-7

// a function that returns a function that takes in a string and returns an int, and a bool.
const ((string->int), bool) x = () => {
    return int (string a) => {
        match a {
            with "A":
                return 5
            with "B":
                return 18
            else:
                return -1
        }
    }, true
}
```
#### Control flow
```go
if true {
    // ...
} else {
    // ...
}

if a == 5 {
    // ...
} else if {} else {}

if true 5 else 6

// incremented variable is i, repeat 6 times. equivalent to: for(int i = 0; i < 6; i++) {//...} in C
repeat 6, i {
    println(i)
}
// infinite loop
forever {

}

foreach index, item in <sequence> {
    // do sth with item
}

foreach _, item in <sequence> {}

while a > 10 {}

break, continue

match {
    with <val>:
        // ...
    ...
    else:
        // ...
}
```
#### Compound data types
```go
// BUILT-IN
// slice
const [string] fruits = ["Kiwi", "Orange", "Apple", "Banana"]
const {string:uint32} people = {"Jennifer": 44, "SomeGuy": 26}

fruits[0] // Kiwi
fruits[-1] // "Banana"
fruits[1:3] // ["Orange", "Apple"]
people["Jennifer"] // 44
set(people, "AnotherGuy", 32)
append(fruits, "Peach")
delete(people, "Jennifer")
string popped = pop(fruits)
println(popped) // "Peach"

// STRUCTS
struct Person embeds [Human] {
    string name, city
    uint age

    void speak = () => println("Hello I am ${self.name}, I live in ${self.city}, and I am ${self.age} years old.")
    void wasBornIn2 = () => println(self.birthYear) // or self.Human.birthYear
}   

// structs implicitly get a 'new' method, which is basically a constructor.
const Person p1 = Person.new(name: "Hasan", city: "Ankara", age: 55, birthYear: 1967)
p1.speak() // Hello I am Hasan, ...
p1.wasBornIn() // I was born in 1967

struct Human {
    uint birthYear

    void wasBornIn = () => println("I was born in ${self.birthYear}")
}

// access modifiers
// every global variable or struct or struct field/method, is private by default.
// to make them public, add pub prefix.
pub struct A {
    pub string x
} 

// package system has the same semantics as Go.
// declaration: 
pkg codegen

//
import codegen

codegen.compileCInstruction()
```
#### Operators
```go
// prefix
-
--
++
&
*

// infix
+ - * / 
+= -= *= /= 
!= == << >>
&& ||
& | ^
// postfix
--
++
```
So on ...