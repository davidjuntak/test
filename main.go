package main

import (
    "fmt"
    "github.com/davidjuntak/test/util"
)

func main() {
    result := Add(2, 3)
    fmt.Println(result)
    result = Minus(5, 3)
    fmt.Println(result)
    result = Add(3, 3)
    fmt.Println(result)
}
