//
// Swift
//

// ----------------------------------------------------------------------------
// Variables

// A type inferred variable declaration
var a = 5

// An explicitly typed variable declaration
var a: Int = 5

// A constant
let b = 6

// An explicitly typed constant declaration
let c: Double = 5.0

// ----------------------------------------------------------------------------
// Functions

// A function that takes a value and returns a value
func f(x: Int) -> Int {
  return x + 1
}

// Multiple return values
func g(x: Int) {
  return (x, x+1)
}

// Unpacking multiple return values
let (x1, y1) = g(10)

// A function that returns a function closure
func makeFunc(value: Int) -> (Int) -> Int {
  return { (otherValue: Int) -> Int in
    return otherValue + value;
  }
}

makeFunc(2)(1) // => 3

// Passing a function to a function
func callFunc(value: Int, f: (Int) -> Int) -> Int {
  return f(value)
}

// Passing functions in-place
callFunc(1, { (number:Int) -> Int in number + 2 }) // => 3
callFunc(1, { $0 + 2 }) // => 3



// Pass a function as a block
callFunc(1) { number in number + 2 } // => 3
callFunc(1) { $0 + 2 } // => 3





// ----------------------------------------------------------------------------
// Classes

class Point {
  let x: Int
  let y: Int

  init(x: Int, y: Int) {
    self.x = x
    self.y = y
  }
}

let p = Point(x: 1, y: 2)
p.x // => 1
p.y // => 2

// Inheriting and overriding
class A {
  // A method that returns a number
  func f() -> Int {
    return 1
  }
}

class B: A {
  override func f() -> Int {
    return 2
  }
}

let a = A()
let b = B()

a.f() // => 1
b.f() // => 2

// ----------------------------------------------------------------------------
// Data structures

// Array literals
let shoppingList = ["Eggs", "Milk"]

// Iteration
for item in shoppingList {
    println(item)
}

// Mapping a function over a collection
(1 ... 5).map { $0 + 1 } // Terse inplace function syntax

// Tuples
let things = ("Some numbers", 1, 2, 3)
println(things.1)
println(things.2)
let (first, second, third, fourth) = things

// Dictionary literals
let lookup = [ "x": 24, "y": 25, "z": 26 ]
lookup["x"] // => Int? = {Some 24}
lookup["n"] // => Int? = nil

// ----------------------------------------------------------------------------
// Generics

func map<T, U>(var array: [T], f:(T) -> U) -> [U] {
  for i in 0 ..< array.count {
    array[i] = f(array[i])
  }

  return array
}
