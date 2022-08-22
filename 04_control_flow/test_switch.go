package main

import "fmt"

func main() {
	day := 3
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("work day")
	case 6, 7:
		fmt.Println("rest day")
	}

	// case 也可以是表达式
	// fallthrough
	score := 90
	switch {
	case score > 60:
		fmt.Printf("score: %v\n", score)
		fallthrough
	case score >= 80 && score < 90:
		fmt.Printf("score: %v\n", score)
	case score > 90 && score <= 100:
		fmt.Printf("score: %v\n", score)
	}

}
