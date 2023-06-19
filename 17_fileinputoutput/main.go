package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main(){
 // write to a file
 content := []byte("This is a test.  Hello world.")
 err := ioutil.WriteFile("test.txt", content, 0644)
 if err != nil {
     fmt.Println("Error writing file: ", err)
 return 
 }

 // read from a file
 data, err := ioutil.ReadFile("test.txt")
 if err != nil {
     fmt.Println("Error reading file: ", err)
     return
 }
 fmt.Println("Contents of file: ", string(data))

 // close and remove a file
 err = os.Remove("test.txt")
 if err != nil {
     fmt.Println("Error removing file: ", err)
     return
 }
}
