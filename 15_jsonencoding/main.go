package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

func main() {
    // Encode struct to JSON
    person := Person{Name: "John", Age: 30}
    jsonData, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("JSON data:", string(jsonData))

    // Decode JSON to struct
    jsonStr := `{"name":"Alice","age":25}`
    var decodedPerson Person
    err = json.Unmarshal([]byte(jsonStr), &decodedPerson)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Decoded JSON data: %+v\n", decodedPerson)
}



