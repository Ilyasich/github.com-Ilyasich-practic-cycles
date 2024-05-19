package christmas

import (
	"fmt"
	"time"
)

// Новогоднее поздравление )
func Christmas() {
	n := 0
	for n < 3 {
		fmt.Println("Ho!")
		time.Sleep(1 * time.Second)
		n += 1
	}
	fmt.Println("merry christmas!!!")
}
