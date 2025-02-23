package main

import "fmt"

func main() {

	var a float64
	var b float64
	var operator string

	fmt.Println("== KALKULATOR SEDERHANA ==")
	fmt.Println("Masukan angka pertama: ")
	fmt.Scanln(&a)

	fmt.Println("Masukan operator: (+, -, *, /) ")
	fmt.Scanln(&operator)

	fmt.Println("Masukan angka kedua: ")
	fmt.Scanln(&b)

	var hasil float64

	switch operator {
	case "+":
		hasil = a + b

	case "-":
		hasil = a - b

	case "*":
		hasil = a * b

	case "/":
		if b != 0 {
			hasil = a / b

		} else {
			fmt.Println("Error: tidak bisa membagi dengan nol!")
		}
	default:
		fmt.Println("Operator tidak valid")
	}
	fmt.Println("Hasil:", hasil)
}
