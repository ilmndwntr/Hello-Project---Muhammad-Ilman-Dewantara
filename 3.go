// Nomer 3 Massa
package main

import "fmt"

func main() {
	var gram, kg, pons, ons float64
	var gram2, kg2, pons2, ons2 float64
	var gram3, kg3, pons3, ons3 float64
	var gram4, kg4, pons4, ons4 float64

	fmt.Scanln(&gram)
	fmt.Scanln(&gram2)
	fmt.Scanln(&gram3)
	fmt.Scanln(&gram4)

	kg = gram / 1000
	pons = gram / 453.592
	ons = gram / 28.3495

	kg2 = gram / 1000
	pons2 = gram / 453.592
	ons2 = gram / 28.3495

	kg3 = gram / 1000
	pons3 = gram / 453.592
	ons3 = gram / 28.3495

	kg4 = gram / 1000
	pons4 = gram / 453.592
	ons4 = gram / 28.3495
	fmt.Println(kg, pons, ons)
	fmt.Println(kg2, pons2, ons2)
	fmt.Println(kg3, pons3, ons3)
	fmt.Println(kg4, pons4, ons4)
}
