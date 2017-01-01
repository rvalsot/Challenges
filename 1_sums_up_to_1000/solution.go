// Euler problem number 1 in go

package main

import "fmt"

func main() {

	suma := 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			suma = suma + i
		}
	}

	fmt.Println("Solution is:", suma)
}
