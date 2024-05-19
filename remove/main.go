package remove

import (
	"fmt"
	"strconv"
	"strings"
)

// Написать функцию принимающую два числа, первое это исходное число, а второе цифра, которую надо удалить из исходного числа. Функция должна вернуть число int в котором не будет указанной цифры. Например 12313, 3 → 121
func RemoveDigit(num int, digit int) int {
	strNum := strconv.Itoa(num)
	strDigit := strconv.Itoa(digit)
	result := strings.Replace(strNum, strDigit, "", -1)
	newNum, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println(err)
	}
	return newNum
}