package digit

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)
	
	

// Ввести с экрана число и посчитать сколько раз встречалась каждая цифра этого числа, вывести на экран
func DigitCount() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите число: ")
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	digitCount := [10]int{}

	for input != 0 {
		digit := input % 10
		digitCount[digit]++
		input /= 10
	}

	for i, count := range digitCount {
		if count > 0 {
			fmt.Printf("Цифра %d встречается %d раз\n", i, count)
		}
	}
}