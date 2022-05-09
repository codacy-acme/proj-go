package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Printf("Sum: %d\n",sumNum(2,3))
	credential()
	sleep()
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

//to test the client-side tool StaticCheck
func sleep() {
    time.Sleep(8)
    fmt.Println("Sleep Over.....")
}