
package main

import "fmt"

func doSomething() {
    fmt.Println("Doing something")
}

func main() {
    // Defer the execution of a function until the surrounding function returns
    defer fmt.Println("Deferred statement")
    defer fmt.Println("Another deferred statement")
    defer doSomething()

    fmt.Println("Hello, World!")
}
