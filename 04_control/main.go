package main

import "fmt"


func main(){

    // if else
    age := 4
    if age < 3{
        fmt.Println("under age")
    } else if age > 4{
        fmt.Println("fit to run")
    } else{
        fmt.Println("you entered an invali integer")
    }


    // switch
    height := 1

    switch height{
        case 1:
            fmt.Println("you chose 1")
        case 2:
            fmt.Println("you chose 2")
        default:
            fmt.Println("no number chosen")
    }


    for i:=0; i <= 10; i++{
        fmt.Println(i)
    }
}
