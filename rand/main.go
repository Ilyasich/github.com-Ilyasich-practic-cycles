package rand

import (
	"fmt"
	"math/rand"
	"time"
)


// Сгенерировать случайное трехзначное число заканчивающееся на ноль
func RandNumsZerro() int {
	rand.NewSource(time.Now().Unix())
	number := rand.Intn(91)*10 + 100
	fmt.Println("Случайное трехзначное число заканчивающееся на ноль: ", number)
	return number
}