/*
Задача №1.
Программа получает на вход последовательность из 5 целых чисел.

Вам нужно определить вид последовательности:
 - возрастающая
 - убывающая
 - случайная
 - постоянная

В качестве ответа следуют выдать прописными латинскими буквами тип последовательности:
1. ASCENDING (строго возрастающая)
2. WEAKLY ASCENDING (нестрого возрастающая, то есть неубывающая)
3. DESCENDING (строго убывающая)
4. WEAKLY DESCENDING (нестрого убывающая, то есть невозрастающая)
5. CONSTANT (постоянная)
7. RANDOM (случайная)

Примеры входных и выходных данных:
In: 11 9 4 2 -1  Out: DESCENDING
In: 3 8 8 11 12  Out: WEAKLY ASCENDING
In: 2 -1 7 21 1  Out: RANDOM
In: 5 5 5 5 5     Out: CONSTANT

Подсказка: используем метод строки strings.Split()
*/

package main

import (
	"fmt"
)

func main() {
	var numbers [5]int
	fmt.Print("Введите последовательность из 5 чисел: ")

	//for i := 0; i < 5; i++ {
	//	fmt.Scan(&numbers[i])
	//}
	fmt.Scan(&numbers[0], &numbers[1], &numbers[2], &numbers[3], &numbers[4])

	var flag bool

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] >= numbers[i+1] {
			flag = false
			break
		}
		flag = true
	}

	if flag {
		fmt.Print("ASCENDING")
		return
	}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] != numbers[i+1] {
			flag = false
			break
		}
		flag = true
	}

	if flag {
		fmt.Print("CONSTANT")
		return
	}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			flag = false
			break
		}
		flag = true
	}

	if flag {
		fmt.Print("WEAKLY ASCENDING")
		return
	}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			flag = true
		} else {
			flag = false
			break
		}
	}

	if flag {
		fmt.Print("DESCENDING")
		return
	}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] >= numbers[i+1] {
			flag = true
		} else {
			flag = false
			break
		}
	}

	if flag {
		fmt.Print("WEAKLY DESCENDING")
		return
	}

	fmt.Print("RANDOM")

}
