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
