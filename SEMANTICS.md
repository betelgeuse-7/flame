## Flame language semantics

>*The semantics will be mostly the same as the Go language. This is because, I am new to these stuff (building compilers). In fact, this is my first (?) compiler. I don't want to make this project a nightmare for me by being unnecessarily idealistic, and trying to come up with a unique language (in terms of semantics).*

---
#### Variable and constant declarations

- The contents of variables can change at any point in a Flame program.

```go
string name = "Jennifer"
println(name) // Jennifer
name = "Noam"
println(name) // Noam
```
Trying to assign a different type of value to an already initialized variable is a compilation-time error (The program won't compile.).
 
```rust
u32 age = 44
age = true
```
The above code will cause a comp-time error.

- Constant variables' contents cannot be changed at any point of program execution. Assignment to constant expression will cause a comp-time error.

```rust
#f32 PI = 3.14
PI = 6.28
```

The above program wouldn't compile.

- There's no uninitialized variables, or constants, in Flame. So there's not a ```nil``` value like in Go.

#### Arithmetic expressions
```
+    sum                    integers, floats
-    difference             integers, floats
*    product                integers, floats
/    quotient               integers, floats
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers

<<   left shift             integer << (integer >= 0)
>>   right shift            integer >> (integer >= 0)
```
> Copied and modified the above list from [Go specification#Arithmetic_operators](https://go.dev/ref/spec#Arithmetic_operators)

- Division by zero
  - TODO
  - I think, division by constant zero value can be a compile-time error, but, otherwise, you will be getting compile-time or runtime errors from the Go compiler, which compiles the compiled code ```:)```.

---
### Types

#### ```bool```

- ```true```, and ```false``` are the boolean constants

- A string with a length of 0 is falsy, otherwise truthy.
- A numeric value that is equal to 0, is falsy, otherwise true.
  - -1 -> true
  - 0 -> false
  - 1 -> true
  - 0.1 -> true
  - 0.0 -> false

```go
#string city = "Istanbul"
if city {
    println("X")
} else {
    println("Y")
}
```
'X' will be printed.

```go
#f32 temperature = 29.8
while temperature {} // this will run forever
////////////
#i32 z = 0
if z {} else { println("!Z") } // '!Z' will be printed
```

#### ```void```

- ```void``` keyword is used to denote that a function, or method, returns nothing.
- Cannot be used as variable types (comp-time error).
- Trying to assign the result of a function that returns void, to a variable will cause a compile-time error.

#### ```string```
TODO