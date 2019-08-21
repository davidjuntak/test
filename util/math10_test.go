
package util

import (
    "testing"
    "reflect"
)

func TestAdd10_0(t *testing.T) {
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
        result := Add10_0(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd10_1(t *testing.T) {
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
        result := Add10_1(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd10_2(t *testing.T) {
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
        result := Add10_2(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd10_3(t *testing.T) {
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
        result := Add10_3(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd10_4(t *testing.T) {
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
        result := Add10_4(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd10_5(t *testing.T) {
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
        result := Add10_5(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd10_6(t *testing.T) {
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
        result := Add10_6(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd10_7(t *testing.T) {
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
        result := Add10_7(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd10_8(t *testing.T) {
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
        result := Add10_8(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}

func TestAdd10_9(t *testing.T) {
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
        result := Add10_9(test.a, test.b)
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test failed")
        }
    }
}
