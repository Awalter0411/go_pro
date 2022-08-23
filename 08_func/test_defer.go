package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("step 1")
	defer fmt.Println("step 2")
	defer fmt.Println("step 3")
	fmt.Println("end")

}
