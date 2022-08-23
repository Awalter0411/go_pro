package main

import (
	"fmt"
	"mod_demo/service"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("\"hello\": %v\n", "hello")
	service.TestUserService()
	r := gin.Default()
	fmt.Printf("r: %v\n", r)
}
