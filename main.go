package main

import "fmt"

func main() {
	names := []string{"John", "James", "Tina", "Joey", "Han"}

    // for loop
    for i := 0; i < len(names); i++ {
        fmt.Println(names[i])
    }

    fmt.Println("------")
    // single value range
    for i := range names {
        fmt.Println(i)
    }

    fmt.Println("------")
    // double value range
    for i, name := range names {
        fmt.Println(i, name)
    }

    fmt.Println("------")
    // double value range, ignore index using "_"
    for _, name := range names {
        fmt.Println(name)
    }
}
