
package util

import (
    "testing"
    "reflect"
)

func TestAdd0_0(t *testing.T) {
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
        result := Add0_0(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd0_1(t *testing.T) {
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
        result := Add0_1(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd0_2(t *testing.T) {
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
        result := Add0_2(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

