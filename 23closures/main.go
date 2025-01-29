package main

import "fmt"

func returnGiftCard(discount int) func() int {
	amount := 1000
	giftcard := func() int {
		amount -= discount
		return amount
	}
	return giftcard
}

func main() {
	giftcard1 := returnGiftCard(20)
	giftcard2 := returnGiftCard(50)

	fmt.Println(giftcard1())
	fmt.Println(giftcard2())
}
