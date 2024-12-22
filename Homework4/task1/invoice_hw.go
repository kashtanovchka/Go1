/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес: индекс(ровно 6 цифр), город, улица, дом, квартира

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

Реализовать конструктор и несколько методов у типа "Накладная"

Пример:
invoice = NewInvoice()

или

order = NewOrder()
*/
package main

import (
	"fmt"
	"regexp"
)

type Invoice struct {
	ProductName  string
	Quantity     int
	CustomerName string
	ContactPhone string
	Address      Address
}

type Address struct {
	PostalCode string
	City       string
	Street     string
	House      int
	Apartment  int
}

func NewInvoice() *Invoice {
	return &Invoice{}
}

func (inv *Invoice) InputData() error {
	var err error

	inv.ProductName, err = inputString("Введите наименование товара (1-100 символов):", 1, 100)
	if err != nil {
		return err
	}

	inv.Quantity, err = inputInt("Введите количество (целое число):")
	if err != nil {
		return err
	}

	inv.CustomerName, err = inputString("Введите ФИО покупателя (только буквы):", 1, 100)
	if err != nil {
		return err
	}

	inv.ContactPhone, err = inputPhone("Введите контактный телефон (10 цифр):")
	if err != nil {
		return err
	}

	inv.Address, err = inputAddress()
	if err != nil {
		return err
	}

	return nil
}

func inputString(prompt string, minLen, maxLen int) (string, error) {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scanln(&input)

		if len(input) < minLen || len(input) > maxLen {
			fmt.Printf("Ошибка: длина должна быть от %d до %d символов.\n", minLen, maxLen)
			continue
		}

		if matched, _ := regexp.MatchString("^[а-яА-ЯёЁa-zA-Z ]+$", input); !matched {
			fmt.Printf("Ошибка: используйте буквы русского и латинского алфавита.\n")
			continue
		}
		return input, nil
	}
}

func inputInt(prompt string) (int, error) {
	var input int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Ошибка: введите целое число.")
			continue
		}
		return input, nil
	}
}

func inputPhone(prompt string) (string, error) {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scanln(&input)

		if matched, _ := regexp.MatchString("^[0-9]{10}$", input); !matched {
			fmt.Println("Ошибка: телефон должен содержать ровно 10 цифр.")
			continue
		}
		return input, nil
	}
}

func inputAddress() (Address, error) {
	var addr Address
	var err error

	for {
		fmt.Print("Введите индекс (ровно 6 цифр):")
		fmt.Scanln(&addr.PostalCode)

		if matched, _ := regexp.MatchString("^[0-9]{6}$", addr.PostalCode); !matched {
			fmt.Println("Ошибка: индекс должен состоять ровно из 6 цифр.")
			continue
		}
		break
	}
	addr.PostalCode = addr.PostalCode
	addr.City, err = inputString("Введите город:", 1, 50)

	if err != nil {
		return addr, err
	}

	addr.Street, err = inputString("Введите улицу:", 1, 50)
	if err != nil {
		return addr, err
	}
	addr.House, err = inputInt("Введите дом:")
	if err != nil {
		return addr, err
	}
	addr.Apartment, err = inputInt("Введите квартиру:")
	if err != nil {
		return addr, err
	}

	return addr, nil
}

func (inv *Invoice) Display() {
	fmt.Println("Данные о заказе:")
	fmt.Printf("Наименование товара: %s \n", inv.ProductName)
	fmt.Printf("Количество: %d \n", inv.Quantity)
	fmt.Printf("ФИО покупателя: %s \n", inv.CustomerName)
	fmt.Printf("Контактный телефон: %s \n", inv.ContactPhone)
	fmt.Println("Адрес:")
	fmt.Printf("Индекс: %s, Город: %s, Улица: %s, Дом: %v, Квартира: %v \n",
		inv.Address.PostalCode, inv.Address.City, inv.Address.Street, inv.Address.House, inv.Address.Apartment)
}

func main() {
	invoice := NewInvoice()

	err := invoice.InputData()

	if err != nil {
		fmt.Println("Ошибка ввода данных:", err)
		return

	}

	invoice.Display()
}
