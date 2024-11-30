/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 120, выход: 021
вход: 100, выход: 001
*/
package main

import "fmt"

func main() {
	var number uint16
	fmt.Println("Введите трехзначное число")
	fmt.Scan(&number)

	fmt.Printf("Реверс: %v%v%v", number%10, number/10%10, number/100)
}
