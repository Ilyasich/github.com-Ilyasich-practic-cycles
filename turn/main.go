package turn
	
import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

// Ввести с экрана число и “перевернуть” его, т.е. получить число из цифр исходного, идущих в обратном порядке
func TurnNums() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите число: ")
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	reversed := 0

	for input != 0 {
		digit := input % 10
		reversed = reversed*10 + digit
		input /= 10
	}

	fmt.Printf("Перевернутое число: %d\n", reversed)
}