
## Interpreter for the Monkey programming language

Based on Thorsten Ball's *Writing an Interpreter in Go* book.

### Syntax

```js
let a = 2 + 3 * 2; // Arithmetic expressions are supported
let b = "Hello"!; // Strings work too
let c = true; // And booleans

// Hashmaps and arrays are supported as well
let hashmap = {
    "yes": true,
    "no": false
}
hashmap["maybe"] = "?"; // Assign values to keys
let x = hashmap["maybe"] + " " + b; // And access them

let y = [1, 2, true];
puts(y[1]); // 2



puts(x); // You can also print values!

// Conditionals, implicit and explicit returns are supported
// as well
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            return 1;
        } else {
            fibonacci(x - 1) + fibonacci(x-2);
        }
    }
}

// We can also create closures!
let newGreeter = fn(greeting) {
    fn(name) {greeting + " " + name + "!"};
};

let hello = newGreeter("Hello");
puts(hello("world")); // Hello world!
```

### Builtins

The interpreter provides a few builtin functions to interact with its objects:

* `len(String | Array)`
* `first(Array)`
* `last(Array)`
* `rest(Array)`
* `push(Array, Any)`
* `puts(Any)` (Each object provides its string implementation)

### Usage

* To run the repl:

```sh
go run main.go
```

* To pass a file to the interpreter:

```sh
go run main.go [file_path]
```
