package main

import "fmt"


func main(){

    // arrays
    var animals [2]string
    animals[0]= "monkey"
    animals[1]= "dog"
    fmt.Println(animals)


    //Ex 2
    var people = [2]string{"john", "lexy"}
    fmt.Println(people)


    // slices
    animals2 := []string{"rat", "donkey", "deer", "lion", "tiger"}
    fmt.Println(animals2)
    fmt.Println(animals2[1])
    fmt.Println(animals2[1:2])



}
