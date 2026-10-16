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

When we declare a variable using `var`, we have to specify _either the type or value (or both)_.
And when we use `:=`, we _have to_ specify the value.
Array literals are arrays that are created and immediately assigned values to:

```
// Maneiras de declarar um array usando array literal
var countries [3]string = [2]string{"Brazil", "USA", "Canada"}
var countires = [3]string{"Brazil", "USA", "Canada"}
countries := [3]string{"Brazil", "USA", "Canada"}
```
