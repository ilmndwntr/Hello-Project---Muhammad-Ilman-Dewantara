// Nomer 2 Teks
package main

import "fmt"

func main() {
	var string1, string2, string3 string
	var hasil bool

	fmt.Scanln(&string1, &string2, &string3)

	hasil = string1 == string2 || string1 == string3 || string2 == string3
	fmt.Println(hasil)
}
