/*
Задача №2

Вход:
Пользователь должен ввести "правильный пароль", состоящий из:
цифр, букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что пароль правильный и он принят.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8

Реализацию оформить через функцию:
1. checkPassword(pass string) (bool, errors <- на усмотрение)
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

out:
	for i := 0; i < 5; i++ {
		fmt.Printf("Попытка No%v\n", i+1)
		fmt.Println("Введите пароль:")

		str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		result := checkPassword(str)

		if result == true {
			fmt.Println("Пароль правильный и он принят")
			break out
		} else if i == 4 {
			fmt.Println("Попытки закончились. Пароль не принят")
			break out
		}
	}

}

func checkPassword(pass string) bool {

	checkStatus := 0

	if checkPasswordChars(pass, 0) {
		checkStatus += 1
	} else {
		fmt.Println("В пароле нет цифр")
	}

	if checkPasswordChars(pass, 1) {
		checkStatus += 1
	} else {
		fmt.Println("В пароле нет маленьких букв")
	}

	if checkPasswordChars(pass, 2) {
		checkStatus += 1
	} else {
		fmt.Println("В пароле нет больших букв")
	}

	if checkPasswordChars(pass, 3) {
		checkStatus += 1
	} else {
		fmt.Println("В пароле нет спец символов")
	}

	if checkLenPassword(pass) {
		checkStatus += 1
	} else {
		fmt.Println("Длина пароля должна быть от 8 до 15 символов")
	}

	return checkStatus == 5
}

func checkPasswordChars(password string, checkType int) bool {

	var symbols = []string{"0123456789", "abcdefghijklmnopqrstuvwxyz", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "_!@#$%^&"}

	for _, character := range password {
		if strings.Contains(symbols[checkType], string(character)) {
			return true
		}
	}
	return false
}

func checkLenPassword(pass string) bool {

	checking := false
	if utf8.RuneCountInString(pass) > 8 && utf8.RuneCountInString(pass) <= 16 {
		checking = true
	}
	return checking
}
