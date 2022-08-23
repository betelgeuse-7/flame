### Operators
```c
// prefix
-               // negation operator
--              // decrement-by-one
++              // increment-by-one
&               // address-of operator
*               // dereferencing operator
!               // not operator

// infix
+ - * / %       // you know what these mean
+= -= *= /=     // these also
< > <= >=       // ...
!=              // not equal
==              // comparison
&&              // logical and
||              // logical or
'               // indexing operator

// postfix
--              // decrement-by-one
++              // increment-by-one
```
### Variable declaration
```
<type> <varname> = <expr>
```
### Constant declaration
```
#<type> <varname> = <expr>
```
### Primitive types

Integers (64-bit): 
```go
int
uint
```
Real numbers (64-bit)
```c
float
```
```go
string
bool
char
```

### Compound types
Slices:
```js
// Initialization
[int] nx = []
[string] names = ["A", "B", "C"]
// constant slice
#[string] y = ["y", "yy", "yyy"]

// Indexing
#string firstElementOfY = y'0

// Appending
names = push(names, "D")

// Deleting
names = delete(names, 1) // delete item with index 1
names = delete(names, -1) // -1 is the last index
```

Maps:
```js
// Initialization
{string:string} = {}
{string:uint} people = {"Jennifer": 44, "Mike": 26}
#{string:bool} x = {"t": false, "y": true}

// Get value with key
people'"Jennifer // 44

// Delete value with key
people = delete(people, "Jennifer")

// Update a field
people."Mike" = 77
```

### Control flow
If:
```go
if true {

} elseif false {

} else {

}
```
Loops:
```go
// incremented variable is i, repeat 6 times. equivalent to: for(int i = 0; i < 6; i++) {//...} in C
repeat 6, i {
    println(i) 
} // 0, 1, 2, 3, 4, 5

// infinite loop
forever {

}

// while loop
while a > 10 {}

break, continue
```

### Functions
```rust
// Declaration
phunc <func-name>: <newline>
[<args>]: <newline>
[<return-values>] {
    
}
```
```c
phunc greet:
[string name]:
[string] {
    return "Hello " + name
}

phunc add:
[int n, int n2]:
[int] {
    return n + n2
}

phunc helloWorld:
{
    println("Hello world")
}
```
```c
// Calling
#string greeted = greet("Abidin")
println(greeted) // Hello Abidin 
```