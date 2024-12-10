package main

import (
	"fmt"
	"example_init/cache"

)

func main() {

	c := cache.New()


	c.Set("userId", 42)


	userId := c.Get("userId")
	fmt.Println("userId:", userId)

	c.Delete("userId")

	userId = c.Get("userId")
	fmt.Println("userId after delete:", userId) 
}
