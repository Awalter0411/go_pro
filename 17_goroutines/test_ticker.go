package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	// counter := 1
	// for _ = range ticker.C {
	// 	counter++
	// 	if counter > 3 {
	// 		ticker.Stop()
	// 		break
	// 	}
	// 	fmt.Print("ticker ... \n")
	// }

	chanInt := make(chan int)

	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		sum += v
		fmt.Printf("v: %v\n", v)
		if sum > 10 {
			break
		}
	}
}
