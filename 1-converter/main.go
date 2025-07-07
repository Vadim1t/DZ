package main

import (
	"errors"
	"fmt"
)

// Глобальные переменные с курсами валют
var (
	rateUSDEUR float64 = 0.86
	rateEURUSD float64 = 1 / 0.86
	rateUSDRUB float64 = 78.25
	rateRUBUSD float64 = 1 / 78.25
	rateEURRUB float64 = 78.25 / 0.86
	rateRUBEUR float64 = 0.86 / 78.25
)

// Глобальный map с указателями на курсы
var valutaMAP = map[string]*float64{
	"USDEUR": &rateUSDEUR,
	"EURUSD": &rateEURUSD,
	"USDRUB": &rateUSDRUB,
	"RUBUSD": &rateRUBUSD,
	"EURRUB": &rateEURRUB,
	"RUBEUR": &rateRUBEUR,
}

func main() {
	user()

	// Берём указатель на глобальный map — создаётся один раз, вне цикла
	ratesPtr := &valutaMAP

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
			n, err := fmt.Scan(&amount)
			if err != nil || n != 1 {
				fmt.Println("Ошибка ввода. Пожалуйста, введите число.")
				var discard string
				fmt.Scanln(&discard)
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

		result, err := convert(amount, valutaIn, valutaTo, ratesPtr)
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
	}
	return valuta, nil
}

// convert принимает указатель на map[string]*float64
func convert(amount float64, valutaIn, valutaTo string, rates *map[string]*float64) (float64, error) {
	if valutaIn == valutaTo {
		return amount, nil
	}

	key := valutaIn + valutaTo
	ratePtr, ok := (*rates)[key]
	if !ok || ratePtr == nil {
		return 0, errors.New("курс конвертации не найден")
	}
	return amount * (*ratePtr), nil
}
