package main

import "fmt"

func main(){
	fmt.Printf("Sum: %d\n",sumNum(2,3))
}

func sumNum(num1 int, num2 int) int {
	return num1+num2;
}
//to test the client-side tool DeadCode
// func subNum(num1 int, num2 int) int {
// 	return num1-num2;
// }