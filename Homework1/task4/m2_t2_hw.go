/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/
package main

import "fmt"

func main() {
	var number uint16
	fmt.Print("Введите четырехзначное число: ")
	fmt.Scan(&number)

	if number >= 1000 && number <= 9999 {
		if number/1000 == number%10 && number/100%10 == number/10%10 {
			fmt.Println(number, "является палиндромом.")
		} else {
			fmt.Println(number, "не является палиндромом.")
		}
	} else {
		fmt.Println("Пожалуйста, введите корректное четырехзначное число.")
	}
}
