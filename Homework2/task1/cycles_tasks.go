/*
=======
Задачи:
=======

3.1 Пользователь вводит числа a и b (b больше a).

	Вывести все целые числа, расположенные между ними.

3.2 Доработать предыдущую задачу так, чтобы выводились только числа,

	делящиеся на 5 без остатка.

3.3 Пользователь вводит число. Вывести таблицу умножения на это число (от 1 до 10)

3.4 В цикле получать от пользователя оценки по четырём экзаменам.

	Вывести сумму набранных им баллов.
	Функцию fmt.Scan() в коде используем только один раз.

3.5 В бесконечном цикле приложение запрашивает у пользователя числа.

	Ввод завершается, как только пользователь вводит число "-1".
	После завершения ввода приложение выводит сумму чисел.
	Используем конструкцию:
	for {
	  // body
	}

3.6 Получить от пользователя натуральное число, посчитать сумму цифр в нём.

	Решить с помощью индексов, т.е. работаем с числом, как со строкой.

3.7 Вводим строку без знаков препинания(5 слов).

	Найти самое длинное слово в строке и вывести сколько в нем букв.

Пример:
In: Скажи как дела в учебе
Out: учебе, 5

In: Закрепляем циклы в языке Golang
Out: Закрепляем, 10
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	var a int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Введите второе число больше первого: ")
	fmt.Scan(&b)

	if a >= b {
		fmt.Print("Некорректные данные \n")
	} else {

		fmt.Print("задача 3.1 \n")
		fmt.Print("Целые числа \n")

		for i := a + 1; i < b; i++ {
			fmt.Println(i)
		}

		fmt.Print("задача 3.2 \n")
		fmt.Print("Целые числа, делящиеся на 5 без остатка \n")

		for i := a + 1; i < b; i++ {
			if i%5 == 0 {
				fmt.Println(i)
			}
		}
	}

	fmt.Print("задача 3.3 \n")
	var c int
	fmt.Print("Введите число: ")
	fmt.Scan(&c)

	fmt.Print("Таблица умножения \n")

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d \n", i, c, i*c)
	}

	fmt.Print("задача 3.4 \n")
	var sum int
	var s int

	for i := 1; i <= 4; i++ {
		fmt.Printf("Введите оценку за %d экзамен \n", i)
		fmt.Scan(&s)
		sum += s
	}

	fmt.Printf("Итого %d баллов \n", sum)

	fmt.Print("задача 3.5 \n")
	var d int
	var currentSum int

	for {
		fmt.Print("Введите число (-1 для выхода) \n")
		fmt.Scan(&d)

		if d == -1 {
			break
		}
		currentSum += d
	}

	fmt.Printf("Сумма чисел %d  \n", currentSum)

	fmt.Print(" задача 3.6 \n")
	var e string
	var summ int

	fmt.Print("Введите натуральное число: ")
	fmt.Scan(&e)

	for i := 0; i < len(e); i++ {
		// Преобразуем символ в rune и затем в строку, чтобы получить цифру
		ch := e[i]

		// Преобразуем символ в число
		digit, err := strconv.Atoi(string(ch))
		if err != nil {
			fmt.Println("Ошибка при преобразовании символа в число:", err)
			return
		}
		summ += digit
	}

	fmt.Printf("Сумма цифр в числе: %d \n", summ)

	fmt.Print(" задача 3.7 \n")

	var f string

	fmt.Println("Введите строку без знаков препинания (5 слов):")
	fmt.Scan(&f)

	words := strings.Split(f, " ")
	fmt.Println(words)
	longestWord := ""

	for _, word := range words {
		fmt.Printf("Слово : %s", word)
		if utf8.RuneCountInString(word) > utf8.RuneCountInString(longestWord) {
			longestWord = word
		}
	}

	// Вывод результата
	fmt.Printf("Самое длинное слово: %s (длина: %d букв)\n", longestWord, utf8.RuneCountInString(longestWord))
}
