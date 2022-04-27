package main

import "fmt"

func main(){
	fmt.Printf("Sum: %d\n",sumNum(2,3))
}

func sumNum(num1 int, num2 int) int {
	return num1+num2;
}

func subNum(num1 int, num2 int) int {
	return num1-num2;
}