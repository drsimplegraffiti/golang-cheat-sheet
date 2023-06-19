package main

import (
    "fmt"
    "math/rand"
)

func main(){
    randomnumber := rand.Intn(999)
    fmt.Println("randomNumber", randomnumber)
}
