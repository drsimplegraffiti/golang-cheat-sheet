package main

import "fmt"

func main(){
    // method map 1
    map1 := make(map[string]string)
    map1["name"] = "John"
    map1["cgpa"] = "3.14"

    fmt.Println(map1)


    map2 := map[int]string{1:"jude"}
    fmt.Println(map2)
}
