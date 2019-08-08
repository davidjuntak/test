package main

import (
    "fmt"
    "github.com/davidjuntak/test/util"
)

func main() {
    result := util.Add(2, 3)
    fmt.Println(result)
    result = util.Minus(5, 3)
    fmt.Println(result)
    result = util.Add(3, 3)
    fmt.Println(result)
}
