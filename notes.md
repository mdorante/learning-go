# Notes

## Variables

- Variable declaration and initialization

```go
package main

import "fmt"

// variables can also be assigned outside of a func, it just has to be done with long syntax
var decks int = 52

func main() {
  // creates a variable, explicitly specifying the data type
  var card string = "Ace of Spades"

  // also creates a variable, but the Go compiler infers the data type from whatever value you're assigning to the variable
  card2 := "Five of Diamonds"

  // NOTE: ":=" is only used at variable initialization, if you want to reassign another value, you can use "="
  card2 = "Queen of Hearts"

  // we can also declare a variable without assigning a value to it, it can be assigned later
  var deckSize int
  deckSize = 20
}
```

## Functions

```go
func funcName() returnType {
  // function body
}
```

## Data Structures

### Arrays and Slices

> Array: a fixed length list of things
> Slice: an array that can grow or shrink

**NOTE**: Every element in a slice (and array) must be of the same type

- Creating a slice:

```go
// syntax: []dataType{list, of, elements}

// example:
numbers := []int{3, 4, 2, 5, 6, 123}
```

- Adding an element:

```go
numbers = append(numbers, 200)
```

- Iterate through an array/slice

```go
// prints all indexes and elements in the slice
for index, number := range numbers {
  fmt.Println(index, number)
}
```

## OOP vs Go Approach

Go is **NOT** an Object Oriented Programming language, the closest thing you can do is create **custom types** which can be extended with functions that use type _instances_ as _receivers_ (kinda like methods for objects).

- Creating a custom type

```go
// Create a type which "extends" a slice of ints
type customType []int
```

- Creating **receiver functions** (like methods)

```go
//  instance references the actual object, it's similar to 'this' in JS or 'self' in Python
func (instance customType) printItems() {
  for index, item := range instance {
    fmt.Println(index, item)
  }
}
```

## Writing and Reading files

In order to write data to a file on disk, we have to convert our data into a byte slice `[]byte`.
And in order to read a file from the disk, we read it as a byte slice and then we have to convert it to the data type that we need.

## Testing

Go testing is not like selenium, mocha, or any other testing framework.
It is just plain Go code with logic that tests other Go code.

Test files are any file with a `_test.go` postfix.

In order to run the tests, we can run `go test` inside the package directory and all tests will be executed.
