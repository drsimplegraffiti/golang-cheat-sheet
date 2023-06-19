package main

import "fmt"


type Person struct{
    name string
    age  int
    isTall bool
    
}

func main(){
    p1 := Person{
        name: "James",
        age: 78,
        isTall: true,
    }

    fmt.Println(p1.name)
}

