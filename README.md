# Core Go Programming, by Packt

## Module 01 - Functions

When it comes to functions, the first few videos of the module teach much of what I already knew.
You declare them using the `func` keyword, followed by the **name** of the function. In parentheses, we specify the **parameters** with a **name** and a **type**. After the closing parenthesis, we specify the **type of return values**, if it has any:

```
func func01 (param01 int, param02 string) float64 {}
```

The first video whose subject I didn't have complete dominion over was called "Returning Multiple Values from Function". To return more than one value, we specify their types in **parentheses**, separated by a **comma**:

```
fun func02 (param01 bool) (string, int) {}
```

In Go, it's possible to give names to the values we return from functions. There is nothing similar in other languages.
By setting a name to the return value, we are setting a new variable:

```
func estimateBrewTime (cups int, secondsPerCup int) (totalTimeSeconds int) {
    totalTimeSeconds = cups * secondsPerCup
    return totalTimeSeconds
}
```

But do we need an additional variable? Only if we want to simply write **return** without specifying the value. This is called **naked return**:

```
func estimateBrewTime (cups int, secondsPerCup int) (totalTimeSeconds int) {
    totalTimeSeconds = cups * secondsPerCup
    return
}
```

The **naked return** also works for more than one return value:

```
func estimateBrewTime (cups int, secondsPerCup int) (totalTimeSeconds int, info string) {
    totalTimeSeconds = cups * secondsPerCup
    info = "Coffee brew time"
    return
}
```

> Particularly, I don't think this is a good practice. I doesn't make much sense and makes the code harder to read.

### Named function literals

Inside a function, we can't declare a regular named function, so we have to use **named function literals**. That means assigning an unnamed function to a variable:

```
func main () {
    calculateTax := func (amount float64) float64 {
        return amount * 0.20
    }
}
```

This proves that functions in **Go** are simply values.

### Closures in functions

A closure is a **function value that references variables from outside its body**.
We can think of it as a function that carries around a backpack filled with variables it needs.
If I understand it correctly, we need **two functions** in order to have a closure function: The "mother" function, that is called once, to create the closure function, which is then called several times.

In the example in the `08_closures.go` file, the `createTemperatureAdjuster()` function creates the `originalTemperature` variable. It also returns another function that "closes over" (captures) `originalTemperature`. Each time we call te returned function `adjustTemperature`, it remembers the last value of `originalTemperature` and can update it. It's worth mentioning that outer function parameters are local variables, too.

## Module 02 - Pointers

To get the address of a variabçe, we use the `&` symbol. To display it using the `fmt` package, we use `%p`:

```
fmt.Printf("Address in memory is: %p", &var01)
```

We can also declare function parameters as pointers. That way, we pass an address (&) as argument and handle the variable directly in its memory address:

```
func applyDiscount(price *float64, discountRate float64) {
    *price = *price - (*price * discountRate)
}
```

## Module 03 - Arrays

If we just declare an array, without specifying its elements, all the elements are created of type zero.
An array of strings, for instance, is created with empty strings:

```
var coffeeSizes [3]string
// [  ]
```

To fill the array, we can specify the index and then the value of it:

```
coffeeSizes[0] = "Small"
coffeeSizes[1] = "Medium"
coffeeSizes[2] = "Large"
```

### Array Literals

When we declare a variable using `var`, we have to specify **either the type or value (or both)**.
And when we use `:=`, we _have to_ specify the value.
Array literals are arrays that are created and immediately assigned values to:

```
// Maneiras de declarar um array usando array literal
var countries [3]string = [2]string{"Brazil", "USA", "Canada"}
var countires = [3]string{"Brazil", "USA", "Canada"}
countries := [3]string{"Brazil", "USA", "Canada"}
```

### Slices

Because of their size flexibility, slices are more commonly used nowadays than arrays.
If we specify the type of a slice variable, the only difference is that **there is no number inside square brackets**:

```
var slice []string
```

So far, the only way he mentioned for assigning values to a slice is using `[ , ]`. But that soesn't work in this case:

```
var countries [6]string = [6]string{"Belgium", "Australia", "Japan", "New Zealand", "Germany", "Thailand"}
var asianCountries []string = countries[2:6]
// Asian countries should be 2 AND 5, and not every single one in between
```

### Relationship between array and slice

Modifying a slice also modifies the underlying array.
It also modifies any other slices that were created with that array, assuming the slices have common elements of the array.

### Different ways to create slices

We can create a slice literal without explicitly having to create an array first:

```
var slice []int = []int{1, 2, 3}
slice := []int{1, 2, 3}
```

If we want to create a slice but don't yet know what elements to put in it, we can use the `make()` function:

```
slice = make([]string, 3)
```

### Adding elements to Slice and cap function

For array, the `len` and `cap` function return the same result. For slices, however, the `cap` function return the capacity of the underlying array, which can be different from the length of the slice.

To add elements to a **slice**, we can use the `append()` function:

```
dogs := []string{"Kira", "Beethoven", "Leia"}
dogs = append(dogs, "To be named")
```

By adding another element, the capacity of the underlying array **doubles in size**. Previously, it was **3**. Now, it is **6**.

What is the value of capacity (`cap`) when we use the `make()` function? It is the length of the slice. The length isn't zero because the slice is created and filled with zero values of that specific type.
However, we can add another argument to the `make()` function:

```
slice01 := make([]string, 0, 10)
// Length and capacity
```

In this case, we can't assign a value to specific index (`slice[0] = "Hello"`), because **no indexes exist**. We have to use `append()`.

### Removing elements from Slice

**There is no specific action or function** to remove an element from a slice. Instead, you have to **create a brand new slice**.
We can use `[:]` to get the elements before the index ogf the one we want to remove, and then specify all the other indexes we want to keep:

```
countries := []string{"Brazil", "Canada", "Austria"}
countries = append(countries[:1], countries[2])
```

Using the `...` notation, we could:

```
countries = append(countries[:indexToRemove], countries[indexToRemove+1:]...)
```

Instead of `append()`, we can use `slices.Delete()` from the `slices` package.

### Iterating with "FOR" loop

He explained the **FOR** loop using the `range` syntax.

## Module 04 - Control Flow

### IF Statement

The IF statement in Go is pretty much everything we are used to. Except **there are no parentheses** wrapping the condition.
Also, we **can't compare variables of different types**.

```
num01 := 1
num02 := 2.5

if num02 > num 01 { }
// Since one variable is float64 and the other is int, we can't compare them
```

## Module 06 - Map Type

Maps are sets of key-value pairs. To declare one, we use the `map` keyword, followed by the type of key and the type of value the keys will hold:

```
var prices map[string]float64 = map[string]float64 {
    "Bike": 125.50,
    "Car": 4500.00,
}
// Or...
prices := map[string]float64 {
    "Bike": 125.50,
    "Car": 4500.00,
}
```

- To access a value, we use the key name: `prices["Bike]`.
- To change a value, we also use the key name: `prices["Bike"] = 90.0`.
- We can also use the `len()` function to get the length of the `map`.
- To add another key-value pair, we simply assign a value to a new key: `prices["TV"] = 550.0`

### Checking key existence in map

To determine if a key is present in the `map`, we can do the following: Access the value of the key, and assign the result to two variables:

```
popul, exists := population["Nagasaki"]
// returns FALSE
```

### Deleting keys from map

We can use the `delete()` function:

```
var populations map[string]int = map[string]int {
    "Sao Paulo":   12000000,
	"Mexico City": 9000000,
	"Sacramento":  538753,
	"Detroit":     649095,
}
delete(population, "Detroit")
```

If we try to delete a value which is absent in the `map`, we don't get an error, but nothing changes.

### Iterating over map

We can use a `FOR ... RANGE` loop. In the case of `maps` there is no index, only key name and value:

```
for key, population := range populations {
	fmt.Printf("%s - %d\n", key, population)
}
```

### What is nil

To answer that question, we are going to create an empty `map` by using the `var` keyword:

```
var populations map[string]int
populations["Cairo"] = 10000000
// panic: assignment to entry in nil map
```

As we can see, we can't add keys to this empty `map`.
It's a runtime error, not a compilation error.
But what is "nil"? It is a special type in **Go** that represents the zero value for a pointer, channel, func, interface, map or slice type. More complex structures, not the zero value of strings or integers, for example.

## Using make function to create map

When we use the `make()` function to create an empty `map`, we can add values to it without any errors:

```
populations := make(map[string]int)
populations["Edinburgh"] = 531370
```

But this approach is more commonly used when we don't know the the keys and values, yet.

What are the differences between arrays, slices, and maps in Go, and when you might choose to use one over the others?
Arrays have a static size, while slices and maps can change size.
In arrays and slices, we can access values by index. In a map, we do it by the key name.
If we compare to JavaScript, `maps` can be thought of as arrays of objects. Although, in Go we can have array of structs, I suppose.
