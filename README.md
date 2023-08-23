


# monkey
Implementation of the Monkey language from [Writing an interpreter in go](https://interpreterbook.com/) and [Writing a compiler in go](https://compilerbook.com).
Monkey is a programming language created by Thorsten Ball.


## TESTS

`go test $(go list ./... | grep -v /playground)` - run all tests except `playground`

## Build
`go build -o monkey`

## Build for Playground
This also contains a web `Playground`.
To build for the `Playground` run: `GOARCH=wasm GOOS=js go build -o ./playground/monkey.wasm ./playground/main.go`

## Running the playground
Use any static file-server.
Example: `cd playground && python -m http.server` and then navigate to `http://localhost:8000`

## Running as REPL
`go run main.go -repl`
or
`go build -o monkey && ./monkey -repl`

## Running monkey programs
To execute a monkey program saved in a file

`go run main.go ./examples/conditionals.monkey`
or
`./monkey ./examples/conditionals.monkey`

## Language Features

### Supported data types
Monkey supports following datatypes natively:
- integers
- strings
- booleans
- arrays
- dictionary or hash
- functions

### Variable bindings

```
let x = 10;
let str = "hello, world!";
let bool = x == 10;
```

### Arithmetic operators
Monkey supports following arithmetic operators which are similar to other programming languages:

- \+ (addition) 
- \- (subtraction)
- \* (multiplication)
- \/ (division)

### Comparision operators
Monkey supports following comparison operators
- == (equals)
- != (not equals)
- \> (greater than)
- < (less than)

### Creating string literals
```
let s = "hello world";
```

### Creating array literals
Monkey arrays can contains objects of any type supported by monkey

```
let add = fn(a, b) { a + b};
let arr = [1, "two", "3", true, false, add]
```

We can access arrays by using positive indices from `0` to `n - 1` (`n` being the length of the array)
```
let first_v = arr[0]
let last_v = arr[len(arr) - 1]
```

### Creating monkey dictionaries
We can use integers, strings and boolean types as keys in dictionaries in monkey
```
let d = {"foo": 1, "bar": 2, true: 3, false: 4, 1: 5};
```

Values can be accessed from dictionaries as follows:

```
let one = d["foo"];
let two = d["bar"];
let three = d[true];
let five = d[1];
```

### Functions
```
let factorial = fn(n) {
    if (n == 0) {
        return 1;
    }
    return n * factorial(n - 1);
}
```

Calling functions
```
>> factorial(3)
6
```

### Higher order functions

```
let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            accumulated
        } else {
            iter(rest(arr), push(accumulated, f(first(arr))))
        }
    };
    iter(arr, []);
};
```

```
>> let a = [1, 2, 3, 4];
>> let double = fn(x) { x * 2 };
>> map(a, double);
[2, 4, 6, 8]
```

### if expressions
`if` statements can produce values in monkey.
```
let x = 10;
let y = 0;
if (x == 10) {
    x
} else {
    y
}
```

### Builtin functions

**len**
`len` returns the length of the object, it is supported for strings, arrays and dictionaries

```
>> let s = "hello";
>> 
5

>> let arr = [1, 2, 3];
>> len(arr)
3
```

**first**

`first` returns the first element from an array, it's not supported for other types

```
>> let arr = [1, 2, 3]
>> first(arr)
1
```

**last**

`last` returns the last element from an array, it's not supported for other types
```
>> let arr = [1, 2, 3]
>> last(arr)
3
```

**rest**

`rest` returns a new array which is a copy of the given array except its first element

```
>> let arr = [1, 2, 3, 4]
>> rest(arr)
[2, 3, 4]

>> rest(arr(arr))
[3, 4]
```

**push**

`push` returns a new array by copying the old array and adds new element at its end

```
>> let arr = [1, 2, 3]
>> push(arr, 4)
[1, 2, 3, 4]
```

**puts**

`puts` prints the value of a monkey object on stdout

```
>> let s = "hello";
>> puts(s)
hello

>> let arr = [1, 2, 3];
>> puts(arr)
[1, 2, 3]
```