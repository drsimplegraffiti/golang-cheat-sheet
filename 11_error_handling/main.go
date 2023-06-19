package main

import (
    "fmt"
    "os"
)

func main(){
    f, err := os.Open("test.txt")
    if err != nil{
        fmt.Println("Error", err)
        return
    }
    defer f.Close()

}
