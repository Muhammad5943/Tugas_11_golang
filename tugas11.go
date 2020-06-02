package main

import "fmt"
import "strconv"

func main() {

	var nilai1 string
	var nilai2 string
	fmt.Print("Input variable 1 : ")
	fmt.Scanln(&nilai1)
	fmt.Print("Input variable 2 : ")
	fmt.Scanln(&nilai2)

	var num1, err = strconv.Atoi(nilai1)
	var num2, err2 = strconv.Atoi(nilai2)

	if err == nil {
		fmt.Println("Nilai 1 : ", num1)
	}
	if err2 == nil {
		fmt.Println("Nilai 2 : ", num2)
	}
	fmt.Println("Hasil Penjumlahan :", num1+num2)
	fmt.Println("Hasil Pengurangan :", num1-num2)
	fmt.Println("Hasil Perkalian :", num1*num2)
	fmt.Println("Hasil Pembagian :", num1/num2)

}
