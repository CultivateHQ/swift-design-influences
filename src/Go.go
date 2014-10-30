//
// Go-Lang
//

// ----------------------------------------------------------------------------
// Variables

// A type inferred variable declaration
a := 5

// An explicitly typed variable declaration
var a int = 5

// A constant
const b = 6

// An explicitly typed constant declaration
const c float64 = 5.0

// ----------------------------------------------------------------------------
// Functions

// A function that takes a value and returns a value
func f(x int) int {
  return x + 1
}

// Multiple return values
func g(x int) (int, int) {
  return x, x+1
}

// Unpacking multiple return values
x1, y1 := g(10)

// A function that returns a function closure
func makeFunc(value int) func(int) int {
  return func(otherValue int) int {
    return otherValue + value
  }
}

makeFunc(2)(1) // => 3

// Passing a function to a function
func callFunc(value int, f func(int) int) int {
  return f(value)
}

// Passing functions in-place
callFunc(1, func(number int) int { return number + 2 }) // => 3




// Pass a function as a block
// Not a feature of Go






// ----------------------------------------------------------------------------
// Classes

// Go uses a completely different model for objects based on structures. This
// is sufficiently different that it is not likely to have been an influence on
// Swift. For more info on Go structures see: http://tour.golang.org/#26































// ----------------------------------------------------------------------------
// Data structures

// Array literals
shoppingList := []string {"Eggs", "Milk"}

// Iteration
for _, item := range shoppingList {
    fmt.Println(item)
}

// Mapping a function over a collection
// Go doesn't support map/reduce etc, so iterate and apply using a for loop

// Tuples
// Go doesn't have tuples




// Dictionary literals
lookup := map[string]int{ "x": 24, "y": 25, "z": 26 }
x, xIsPresent := lookup["x"] // => 24, true
n, nIsPresent := lookup["n"] // => 0, false

// ----------------------------------------------------------------------------
// Generics

// Go doesn't have generics at the moment (http://golang.org/doc/faq#generics)






