//Nomer 1 Gradient
package main

import "fmt"

func main() {
	var xA1, yA1, xA2, yA2 int
	var xB1, yB1, xB2, yB2 int
	var xC1, yC1, xC2, yC2 int
	var g1, g2, g3 float64

	fmt.Scanln(&xA1, &yA1, &xA2, &yA2)
	fmt.Scanln(&xB1, &yB1, &xB2, &yB2)
	fmt.Scanln(&xC1, &yC1, &xC2, &yC2)

	g1 = float64(yA1-yA2) / float64(xA1-xA2)
	g2 = float64(yB1-yB2) / float64(xB1-xB2)
	g3 = float64(yC1-yC2) / float64(xC1-xC2)

	fmt.Println(g1)
	fmt.Println(g2)
	fmt.Println(g3)
}
