package main

import (
	"fmt"
)

func main(){
	fmt.Printf("Sum: %d\n",sumNum(2,3))
	credential()
}

func sumNum(num1 int, num2 int) int {
	return num1+num2;
}

func credential() {
	username := "admin"
    var password = "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"
    fmt.Println("Doing something with: ", username, password)
}

//to test the client-side tool DeadCode
func subNum(num1 int, num2 int) int {
	return num2-num1;
}

//to test the client-side tool aligncheck
func subNum()  {
	config := new(Config) // not nil
	if config == nil {
		// then
	}
}