package main

import (
    "testing"
    "reflect"
)

func TestAdd(t *testing.T) {
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
        result := Add(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed, expected: %d, result: %d", test.expected, result)
        }
    }
}

func TestMinus(t *testing.T) {
    tests := []struct {
        a    int
        b    int
        expected int
    }{
        {
            a: 3,
            b: 2,
            expected: 1,
        },
        {
            a: 4,
            b: 2,
            expected: 2,
        },
    }

    for _, test := range tests {
        result := Minus(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed, expected: %d, result: %d", test.expected, result)
        }
    }
}

func TestMultiple(t *testing.T) {
    tests := []struct {
        a    int
        b    int
        expected int
    }{
        {
            a: 3,
            b: 2,
            expected: 6,
        },
        {
            a: 4,
            b: 2,
            expected: 8,
        },
    }

    for _, test := range tests {
        result := Multiple(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed, expected: %d, result: %d", test.expected, result)
        }
    }
}
