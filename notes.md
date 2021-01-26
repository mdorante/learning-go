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
