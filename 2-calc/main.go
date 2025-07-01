package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа через запятую: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
		return
	}

	input = strings.TrimSpace(input)

	numberStrings := strings.Split(input, ",")

	numbers := make([]float64, 0, len(strings.Split(input, ",")))

	for _, numStr := range numberStrings {
		num, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64)
		if err != nil {
			fmt.Printf("Ошибка при преобразовании числа: %s\n", numStr)
			continue
		}
		numbers = append(numbers, num)
	}

	fmt.Print("Выберите метод расчета: AVG, SUM, MED: ")
	var inputMetod string
	fmt.Scan(&inputMetod)
	switch {
	case inputMetod == "AVG":
		fmt.Printf("Результат расчета: %.2f", sliceAVG(numbers))
	case inputMetod == "SUM":
		fmt.Printf("Результат расчета: %.2f", sliceSUM(numbers))
	case inputMetod == "MED":
		fmt.Printf("Результат расчета: %.2f", sliceMED(numbers))
	}
}

func sliceAVG(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func sliceSUM(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func sliceMED(numbers []float64) float64 {
	sort.Float64s(numbers)
	length := len(numbers)
	if length%2 == 0 {
		mid := length / 2
		return (numbers[mid-1] + numbers[mid]) / 2
	}
	return numbers[length/2]
}
