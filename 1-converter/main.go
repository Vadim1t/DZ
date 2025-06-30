package main

import (
	"errors"
	"fmt"
)

const usdToEuro = 0.86
const usdToRub = 78.25
const euroToRub = usdToRub / usdToEuro

func main() {
	user()
	var amount float64
	var valutaIn, valutaTo string

	for {
		fmt.Print("Выберите исходнкю валюту для конвертации: USD, EUR, RUB: ")
		fmt.Scan(&valutaIn)
		ishodValuta, err := proverkaValuti(valutaIn)

		if err != nil {
			fmt.Println("Неправильно указана исходная валюта: ", valutaIn)
			continue
		}

		for {
			fmt.Print("Введите сумму для конвертации: ")
			n, err := fmt.Scan(&amount) // проверяем колличество введеных аргументов n
			if err != nil || n != 1 {
				fmt.Println("Ошибка ввода. Пожалуйста, введите число.")
				var discard string
				fmt.Scanln(&discard) // очищаем остаток строки
				continue
			}
			if amount < 0 {
				fmt.Println("Сумма не может быть отрицательной.")
				continue
			}
			break
		}

		for {
			valutaVibor1, valutaVibor2 := viborValuti(ishodValuta)
			fmt.Println("Выберите целевую валюту для конвертации: ", valutaVibor1, "/", valutaVibor2)
			fmt.Scan(&valutaTo)
			vibor, err := proverkaValuti(valutaTo)
			if err != nil {
				fmt.Println("Неправильно указана целевая валюта: ", vibor)
				continue
			}
			break
		}

		fmt.Printf("Итого: %.2f", convert(amount, valutaIn, valutaTo))
		break
	}
}

func user() string {
	var name string
	fmt.Printf("Здравствуйте! Представьтесь пожалуйста: ")
	fmt.Scan(&name)
	return name
}

func convert(amount float64, from string, to string) float64 {
	switch {
	case from == "USD" && to == "EUR":
		return amount * usdToEuro
	case from == "EUR" && to == "USD":
		return amount * (1 / usdToEuro)
	case from == "USD" && to == "RUB":
		return amount * usdToRub
	case from == "RUB" && to == "USD":
		return amount / usdToRub
	case from == "EUR" && to == "RUB":
		return amount * euroToRub
	case from == "RUB" && to == "EUR":
		return amount / euroToRub
	default:
		return amount
	}
}

func viborValuti(valuta string) (string, string) {
	switch valuta {
	case "USD":
		return "EUR", "RUB"
	case "EUR":
		return "USD", "RUB"
	default:
		return "USD", "EUR"
	}
}

func proverkaValuti(valuta string) (string, error) {
	if valuta != "USD" && valuta != "EUR" && valuta != "RUB" {
		return "", errors.New("no_params_error")
	} else {
		return valuta, nil
	}
}
