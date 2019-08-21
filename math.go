package main

import (
    "fmt"
    "os"
    "log"
)

func main() {
    for j := 0; j < 200; j++ {
        dummyFunction := `
package util
`
        template := `
func Add%d_%d(a, b int) int {
    return a + b
}
`
        dummyFunction2 := `
package util

import (
    "testing"
    "reflect"
)
`
        template2 := `
func TestAdd%d_%d(t *testing.T) {
    tests := []struct {
        a    int
        b    int
        expected int
    }{
        {
            a: 1,
            b: 2,
            expected: 3,
        },
        {
            a: 3,
            b: 2,
            expected: 5,
        },
    }

    for _, test := range tests {
        result := Add%d_%d(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}
`
        for i := 0; i < 100; i++ {
            dummyFunction += fmt.Sprintf(template, j, i)
            dummyFunction2 += fmt.Sprintf(template2, j, i, j, i)
        }

        f, err := os.OpenFile(fmt.Sprintf("util/math%d.go", j), os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte(dummyFunction)); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }

        f, err = os.OpenFile(fmt.Sprintf("util/math%d_test.go", j), os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := f.Write([]byte(dummyFunction2)); err != nil {
            log.Fatal(err)
        }
        if err := f.Close(); err != nil {
            log.Fatal(err)
        }

        dummyFunction = ""
        dummyFunction2 = ""
    }
}

func Add(a, b int) int {
    return a + b
}

func Minus(a, b int) int {
    return a - b
}

func Multiple(a, b int) int {
    return a * b
}
