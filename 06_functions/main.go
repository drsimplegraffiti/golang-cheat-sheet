package main

import "fmt"

func main(){
   result := add(3,4)
   fmt.Println(result)

   // return multiple values
   quotient,remainder := divide(11,5)
   fmt.Println(quotient, remainder)
}

func add(x,y int) int{
    return x + y
}


func divide(x,y int)(int, int){
   quotient := x /y
   remainder := x % y

   return quotient, remainder
}
