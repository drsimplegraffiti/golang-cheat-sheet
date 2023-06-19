package main

import (
    "fmt"
    "time"
)

func greet(){
    fmt.Println("hello there!")
}

func main(){
  go greet() // this is a goroutine and it runs concurrently with the main function i.e it runs in the background


  time.Sleep(10 * time.Second) // this is to make the main function wait for 10 seconds before exiting
}
