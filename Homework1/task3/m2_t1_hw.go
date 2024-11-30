/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1 9 2
Выход: 1 2 9

Про sort мы пока ничего не знаем!
Это задача на условный оператор
*/
package main

import "fmt"

func main() {
	var a uint16
	fmt.Println("Введите первое число")
	fmt.Scan(&a)

	var b uint16
	fmt.Println("Введите второе число")
	fmt.Scan(&b)

	var c uint16
	fmt.Println("Введите третье число")
	fmt.Scan(&c)

	var first, second, third uint16

	if a <= b && a <= c {
		first = a
		if b <= c {
			second = b
			third = c
		} else {
			second = c
			third = b
		}
	} else if b <= a && b <= c {
		first = b
		if a <= c {
			second = a
			third = c
		} else {
			second = c
			third = a
		}
	} else {
		first = c
		if a <= b {
			second = a
			third = b
		} else {
			second = b
			third = a
		}
	}

	fmt.Printf("Числа в порядке возрастания: %d, %d, %d", first, second, third)
}
