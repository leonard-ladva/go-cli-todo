package main

import (
	"flag"
	"fmt"
)

func main() {
	todos := []string{"eat something"}	
	newTodo := flag.String("new", "blank", "Todo content")	

	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)
}