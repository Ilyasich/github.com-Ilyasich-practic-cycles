package summ

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Ввести с экрана число и посчитать сумму первой и последней цифр числа
func SummNumber() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите число: ")
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	lastDigit := input % 10
	var firstDigit int

	for input != 0 {
		firstDigit = input % 10
		input /= 10
	}

	sum := firstDigit + lastDigit

	fmt.Printf("Сумма первой и последней цифры: %d\n", sum)
}
