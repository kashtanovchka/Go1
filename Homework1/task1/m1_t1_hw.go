/*
Задача №1
Вход:

	расстояние(50 - 10000 км),
	расход в литрах (5-25 литров) на 100 км и
	стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*Проверка условий по желанию
*/
package main

import "fmt"

func main() {
	var distance float64
	var consumption float64
	const price = 48

	fmt.Println("Введите расстояние")
	fmt.Scan(&distance)

	fmt.Println("Введите расход в литрах")
	fmt.Scan(&consumption)

	var total = distance / 100 * consumption * price

	fmt.Printf("Стоимость поездки %v в рублях", total)
}
