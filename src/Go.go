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

// Mapping a function over a collection
array := []int { 1, 2, 3, 4, 5 }
f := func(x int) int { return x + 1 }
for i, v := range array {
  array[i] = f(v) // Modifies original array!!
}

// Multiple return values
func g(x int) (int, int) {
  return x, x+1
}

// Unpacking multiple return values
x1, y1 := g(10)

// ----------------------------------------------------------------------------
// Classes (structures in Go, not classes)

type Point struct {
  X int // Upper case makes the variable public
  Y int
}

p := Point{1, 2}
p.X // => 1
p.Y // => 2






// Embedding and overridding
type A struct {}

// A method that returns a number
func (a *A) F() int {
  return 1
}

type B struct {
  *A // Embed A, just to show how overriding would work
}

// A method that returns a number
func (b *B) F() int {
  return 2
}

var a = &A{}
var b = &B{}

a.F() // => 1
b.F() // => 2

// ----------------------------------------------------------------------------
// Data structures

// Array literals
shoppingList := []string {"Eggs", "Milk"}

// Iteration
for _, item := range shoppingList {
    fmt.Println(item)
}

// Tuples
// Go doesn't have tuples




// Dictionary literals
lookup := map[string]int{ "x": 24, "y": 25, "z": 26 }
x, xIsPresent := lookup["x"] // => 24, true
n, nIsPresent := lookup["n"] // => 0, false


// ----------------------------------------------------------------------------
// Generics

// Go doesn't have generics. Deliberately: http://golang.org/doc/faq#generics
type IntToFloat64 func(int) float64
type IntToInt func(int) int

// ... and so on for each combination of types

func MapIntToFloat64(sequence []int, f IntToFloat64) []float64 {
  var result := make([]float64, len(sequence))
  for index, value := range sequence {
    result[index] = f(value)
  }
}

// ... and so on for each combination of types
