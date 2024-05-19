package main

import (

	"fmt"
	"practic/christmas"
	"practic/count"
	"practic/digit"
	"practic/perfect"
	"practic/prime"
	"practic/rand"
	"practic/remove"
	"practic/summ"
	"practic/syracuse"
	"practic/turn"
	
)

func main() {

	christmas.Christmas()
	rand.RandNumsZerro()
	count.CountNums()
	summ.SummNumber()
	digit.DigitCount()
	turn.TurnNums()
	fmt.Println(remove.RemoveDigit(123, 1))
	fmt.Println(prime.IsPrime(7))
	perfect.FindPerfectNumbers(10000)
	syracuse.Syr(20)
	syracuse.Syracuse(30)
}
