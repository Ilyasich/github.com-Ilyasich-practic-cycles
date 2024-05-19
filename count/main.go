package count

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

// Ввести с экрана число и посчитать количество четных и нечетных цифр в числе
func CountNums() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите число: ")
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	evenCount := 0
	oddCount := 0

	for input != 0 {
		digit := input % 10
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
		input /= 10
	}

	fmt.Printf("Количество четных цифр: %d\n", evenCount)
	fmt.Printf("Колличество нечетных цифр: %d\n", oddCount)
}
