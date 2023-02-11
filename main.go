package main

import (
	"flag"
	"fmt"
)

func main() {

	// define the flags
	name := flag.String("name", "default_name", "Your name")
	age := flag.Int("age", 0, "Your age")
	isAdmin := flag.Bool("admin", false, "Is the user an administrator?")

	// Parse the flags
	flag.Parse()

	// Use the flags
	fmt.Printf("Hello, %s! Your age is %d\n", *name, *age)
	if *isAdmin {
		fmt.Println("User is an administrator")
	} else {
		fmt.Println("User is not an administrator")
	}

}
