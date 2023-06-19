package main

import "fmt"

func main(){
    // pointers and dereferencing

    var a int = 67
    var b *int = &a // b is a pointer to a memory address where an int is stored

    fmt.Println(a,b)
    fmt.Println(&a,b) // &a is the memory address of a
    fmt.Println(a,*b) // dereferencing
}

