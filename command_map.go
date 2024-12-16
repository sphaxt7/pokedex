package main

import "fmt"

func commandMap() error {
	fmt.Println("Dsiplaying map")
	for i := 0; i < 20; i++ {
		fmt.Printf("location: %d\n", i + 1)
	}

	return nil
}