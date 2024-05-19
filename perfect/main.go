package perfect

import "fmt"

/* Совершенные числа
Число совершенно, если оно равно сумме всех своих делителей, кроме самого себя.  Найдите все совершенные числа от `1` до `10000` и выведите их на экран.
Пример: `6 = 1 + 2 + 3`. `6` делится на `1`, на `2` , на `3` и на `6` без остатка. Дальше мы находим сумму всех этих чисел, пропуская `6`.
Должно получиться `4` таких числа*/

func IsPerfect(n int) bool {
    sum := 0
    for i := 1; i <= n/2; i++ {
        if n%i == 0 {
            sum += i
        }
    }
    return sum == n
}

func FindPerfectNumbers(limit int) {
    for i := 1; i <= limit; i++ {
        if IsPerfect(i) {
            fmt.Println(i)
        }
    }
}

func main() {
    FindPerfectNumbers(10000)
}