
package main

import "fmt"

func recoverFunc() {
    if r := recover(); r != nil { // r is of type interface{}
        fmt.Println("Recovered:", r)
    }
}

func doSomething() {
    defer recoverFunc()
    fmt.Println("Doing something")
    panic("Something went wrong") // panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array 
}

func main() {
    doSomething()
    fmt.Println("Program continues executing")
}
