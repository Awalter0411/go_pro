package main

import "fmt"

func main() {
	var ip *int
	var i int = 100
	ip = &i
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %v\n", *ip)
}
