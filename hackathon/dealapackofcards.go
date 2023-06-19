package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	numPlayers := 4
	cardsPerPlayer := 12 / numPlayers

	for i := 0; i < numPlayers; i++ {
		startIndex := i * cardsPerPlayer
		endIndex := (i + 1) * cardsPerPlayer
		playerCards := deck[startIndex:endIndex]
		// if i != 3 {
		fmt.Printf("Player %d:", i+1)
		z01.PrintRune(' ')
		// z01.PrintRune(rune(playerCards[0] + '0'))
		fmt.Printf("%v", playerCards[0])
		z01.PrintRune(',')
		z01.PrintRune(' ')
		fmt.Printf("%v", playerCards[1])
		z01.PrintRune(',')
		z01.PrintRune(' ')
		fmt.Printf("%v", playerCards[2])

		// } else {
		// 	fmt.Printf("Player %d:", i+1)
		// 	z01.PrintRune(' ')
		// 	// z01.PrintRune(rune(playerCards[0]/10 + '0'))
		// 	fmt.Printf("%v", playerCards[0]/10)
		// 	// z01.PrintRune(rune(playerCards[0]%10 + '0'))
		// 	fmt.Printf("%v", playerCards[0]%10)
		// 	z01.PrintRune(',')
		// 	z01.PrintRune(' ')
		// 	// z01.PrintRune(rune(playerCards[1]/10 + '0'))
		// 	fmt.Printf("%v", playerCards[1]/10)
		// 	// z01.PrintRune(rune(playerCards[1]%10 + '0'))
		// 	fmt.Printf("%v", playerCards[1]%10)
		// 	z01.PrintRune(',')
		// 	z01.PrintRune(' ')
		// 	// z01.PrintRune(rune(playerCards[2]/10 + '0'))
		// 	fmt.Printf("%v", playerCards[2]/10)
		// 	z01.PrintRune(rune(playerCards[2]%10 + '0'))
		// }
		z01.PrintRune('\n')
	}
}
