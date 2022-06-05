### Flame Compiler to Go

This is the Flame language:

Semantics of the language: [SEMANTICS](SEMANTICS.md)

#### Encoding
The Flame compiler supports UTF-8 encoded files, but you can only use ASCII when declaring variables.
The supported line termination sequence is ASCII LF (\n). (CRLF (\r\n) is not seen as a newline, only \n).

#### Indentation, and whitespace
The Flame language is whitespace-insensitive, meaning whitespace is not important. 

#### Comments
Comments start with ```//```. Multi-line comments are not supported yet.

#### Identifiers
Identifiers are one or more characters of ASCII letters, or ```_```. ({A..Z, a..z, _})
Identifiers are case-sensitive (e.g ```firstname```, and ```firstName``` is not the same variable.).
Keywords cannot be used as identifiers.

#### Keywords 
TODO

#### String literals
String literals are enclosed in ```"``` (double quote) pairs. You can interpolate strings using ```${}``` notation. (e.g Assume world variable stores "WORLD" value. "Hello ${world}" == "Hello WORLD")

Escape sequences are not yet supported. (TODO)

#### Numeric literals
There are three types of numerals: 
    - Signed integers (```int```, ```int32```),
    - Unsigned integers (```uint```, ```uint32```), and
    - Floating point numbers (```float64```, ```float32```)

Octal (```0o```), hexadecimal (```0x```), and binary notations (```0b```) are not yet supported.

#### Operators
There are three types of operators:
    - Prefix,
    - Infix, and
    - Postfix
```
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
<<              // left shift operator (bitwise)
>>              // right shift operator 
&&              // logical and
||              // logical or
&               // bitwise and
|               // bitwise or 
^               // bitwise xor
'               // indexing operator

// postfix
--              // decrement-by-one
++              // increment-by-one
```

### Context-Free Grammar (Extended Backus-Naur Form)

TODO

#### Variable declaration
```
string name = "Jennifer"
uint age = 44
```
#### Constant declaration
```go
#string name = "Jennifer"
#uint age = 44
```
#### Types
```rust
void
string
bool
int i32
uint u32
f64 f32
```
- Slices
```rust
#[string] fruits = ["Kiwi", "Orange", "Apple", "Banana"]
fruits'0 // indexing
fruits'-1 // last element

#string popped = fruits.pop() 
popped // Banana
fruits.push("Strawberry")
```
- Maps
```rust
#{string:u32} people = {"Jennifer": 44}
people."Mehmet" = 77
u32 jenniferAge = people'"Jennifer" 
jenniferAge // 44
people.delete("Mehmet")
```
- structs
```rust
pub struct A {
    pub string x
    pub struct B {
        u32 y
    }
}

#A a = A.new(x: "Hello", B: B.new(y: 45))
a.x // "Hello"

#[A] AX = [a, a, a]
AX'1 // a
```
- Pointers
TODO

#### Functions
```rust
void->void helloWorld => {
    println("hello")
}

string x, bool y->[u32 z] X => {
    [42, 24, 66, 12, 3]
}
// alternative: omit curly braces and put the return value right after '=>'
// string x, bool y->[u32 z] X => [42, 24, 66, 12, 3]
// you can also add return if you want
// string x, bool y->[u32 z] X => {
//    return [42, 24, 66, 12, 3]
//}
```
##### Calling functions
```rust
helloWorld()
#uint32 z = X("yay", true)
println(z) //[42, 24, 66, 12, 3]
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
} elseif {} else {}

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

#### Methods

TODO

