package main

import "fmt"

func cycles() {
	// for loop
	for i := 0; i < 10; i++ {
		println(i)
	}

	// for while
	counter := 0
	for counter < 10 {
		println(counter)
		counter++
	}

	// for forever
	counter = 0
	for {
		println(counter)
		counter++
		if counter == 10 {
			break
		}
	}

	// for range
	arreglo := [9]int{2, 5, 6, 7, 8, 0, 3}
	for i, j := range arreglo {
		fmt.Printf("indice i: %d tiene como valor #%d\n", i, j)
	}
}
