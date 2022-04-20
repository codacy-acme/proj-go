package main

import "fmt"

func main(){
	fmt.Println("Hello, World!")
	fmt.Println(sumNum(2,3))
}

func sumNum(num1 uint64, num2 uint64) uint64{
	return num1+num2;
}