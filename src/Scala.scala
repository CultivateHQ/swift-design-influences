//
// Scala
//

// ----------------------------------------------------------------------------
// Variables

// A type inferred variable declaration
var a = 5

// An explicitly typed variable declaration
var a: Int = 5

// A constant
val b = 6

// An explicitly typed constant declaration
val c: Double = 5.0

// ----------------------------------------------------------------------------
// Functions

// A function that takes a value and returns a value
def f(x: Int): Int = {
  x + 1
}

// Multiple return values
def g(x: Int) = {
  (x, x+1)
}

// Unpacking multiple return values
val (x1, y1) = g(10)

// A function that returns a function closure
def makeFunc(value: Int): Int => Int = {
  (otherValue: Int) => otherValue + value
}

makeFunc(2)(1) // => 3



// Passing a function to a function
def callFunc(value: Int, f: Int => Int): Int = {
  return f(value)
}

// Passing functions in-place
callFunc(1, (number: Int) => number + 2)
callFunc(1, number => number + 2)
callFunc(1, { _ + 2 })
callFunc(1, _ + 2)

// Pass a function as a block (need to define receiver differently)
def callBlock(value: Int)(f: Int => Int): Int = {
  return f(value)
}

callBlock(1)(_ + 2)
callBlock(1) { _ + 2 }

// ----------------------------------------------------------------------------
// Classes

class Point(val x: Int, val y: Int)

// ^^ that's all you need







val p = new Point(1, 2)
p.x // => 1
p.y // => 2

// Inheritance and interface implementation syntax
class A {
  // A method that returns a number
  def f() = 1
}



class B extends A {
  override def f() = 2
}



val a = new A
val b = new B

a.f() // => 1
b.f() // => 2

// ----------------------------------------------------------------------------
// Data structures

// In Scala Lists are preferred over Arrays (although they are available)
val shoppingList = List("Eggs", "Milk")

// Iteration
for (item <- shoppingList) {
  println(item)
}

// Mapping a function over a collection
(1 to 5).map { _ + 1 } // Terse inplace function syntax

// Tuples
val things = ("Some numbers", 1, 2, 3)
println(things._1)
println(things._2)
val (first, second, third, fourth) = things

// Map literals
val lookup = Map("x" -> 24, "y" -> 25, "z" -> 26)
lookup.get("x") // => Option[Int] = Some(24)
lookup.get("n") // => Option[Int] = None

// ----------------------------------------------------------------------------
// Generics

def map[T, U](list: List[T], f: T => U): List[U] = {
  for (v <- list) yield f(v)
}




