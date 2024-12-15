/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через другую строку.

Задача №1.1
Реализовать и функцию зашифровки

codeToString(code) -> "???????'

stringToCode("hello") -> "??????"
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func codeToString(code string) (string, error) {

	var res strings.Builder

	for i := 0; i < len(code); i += 2 {

		num, err := strconv.Atoi(code[i : i+2]) //перевод кода в число

		if err != nil {
			return "Ошибка перевода кода в число", err
		}
		//Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел)
		if num >= 0 && num <= 25 {
			res.WriteByte(byte(num + 'a')) // 'a' имеет код 97
		} else if num == 26 {
			res.WriteByte(' ')
		} else {
			return "", fmt.Errorf("Ошибка перевода: %d", num)
		}
	}

	return res.String(), nil
}

func stringToCode(text string) (string, error) {

	var res strings.Builder

	for _, character := range text {
		if character == ' ' {
			res.WriteString("26") // Кодируем пробел
		} else if 'a' <= character && character <= 'z' {
			res.WriteString(fmt.Sprintf("%02d", character-'a')) // Кодируем букву
		} else {
			return "", fmt.Errorf("Ошибка кодировки %v", character)
		}
	}

	return res.String(), nil
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите код для расшифровки: ")

	if reader.Scan() {

		code := reader.Text()
		text, err := codeToString(code)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Расшифрованный текст: '%v'\n", text)
	}

	fmt.Print("Текст для шифрования латинскими буквами: ")

	if reader.Scan() {

		text := reader.Text()
		code, err := stringToCode(text)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Зашифрованный код: '%v'\n", code)
	}
}
