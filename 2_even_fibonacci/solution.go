package main

import "fmt"

func main() {
	suma := 0
	b1, b2 := 1, 2
	fibonnacis := []int{}
	limite := 4000000

	fibonnacis = append(fibonnacis, b2)

	for i := 0; i < limite; i++ {
		b3 := b1 + b2

		b1 = b2
		b2 = b3
		if b3%2 == 0 {
			fibonnacis = append(fibonnacis, b3)
		}
		if b3 > limite {
			break
		}
	}
	fmt.Println(fibonnacis)

	for _, value := range fibonnacis {
		fmt.Println("Value:", value, "Sum:", suma)
		suma += value
	}

	fmt.Println("Solution:", suma)

}
