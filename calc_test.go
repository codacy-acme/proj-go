package main

import (
    "testing"
)

func TestHelloName(t *testing.T) {
    num1 := 2
	num2 := 3
    sum := sumNum(num1,num2)
    if sum != 5 {
        t.Fatalf("Sum is invalid. It should be %d but the sum is %d", 5, sum)
    }
}