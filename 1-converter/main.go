package main

import (
	"errors"
	"fmt"
)

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

		rates := getRatesMap()
		result, err := convert(amount, valutaIn, valutaTo, rates)
		if err != nil {
			fmt.Println("Ошибка конвертации:", err)
			return
		}

		fmt.Printf("%.2f %s = %.2f %s\n", amount, valutaIn, result, valutaTo)
		break
	}
}

func user() string {
	var name string
	fmt.Printf("Здравствуйте! Представьтесь пожалуйста: ")
	fmt.Scan(&name)
	return name
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

func getRatesMap() *map[string]float64 {
	rates := map[string]float64{
		"USDEUR": 0.86,
		"EURUSD": 1 / 0.86,
		"USDRUB": 78.25,
		"RUBUSD": 1 / 78.25,
		"EURRUB": 78.25 / 0.86,
		"RUBEUR": 0.86 / 78.25,
	}
	return &rates
}

func convert(amount float64, valutaIn, valutaTo string, rates *map[string]float64) (float64, error) {
	key := valutaIn + valutaTo
	rate, ok := (*rates)[key]
	if !ok {
		if valutaIn == valutaTo {
			return amount, nil
		}
		return 0, errors.New("курс конвертации не найден")
	}
	return amount * rate, nil
}
