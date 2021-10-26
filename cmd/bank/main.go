package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	card1 := types.Card{Balance: 10_000_00, Active: true, MinBalance: 10_000_00}
	card.AddBonus(&card1, 3, 30, 365)
	fmt.Println(card1.Balance)
}