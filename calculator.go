package main

import "fmt"

func main(){
	fmt.Printf("Sum: %d\n",sumNum(2,3))
	fmt.Printf("Credential: %s\n",credential())
}

func sumNum(num1 int, num2 int) int {
	return num1+num2;
}

func credential() string {
	return "312321321321321";
}
//to test the client-side tool DeadCode
// func subNum(num1 int, num2 int) int {
// 	return num1-num2;
// }