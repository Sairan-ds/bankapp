package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 10000,
		},
		{
			ID: 2,
			Amount: 20000,
		},
		{
			ID: 3,
			Amount: 30000,
		},
		
	}
	max := Max(payments)
	fmt.Println(max.ID)
	//Output: 3
	
}


func ExamplePaymentSources() {
	cards := []types.Card{
		{
			ID: 1,
			Balance: 10000,
			Number: "1234 xxxx xxxx 8888",
			Active: true,
		},
		{
			ID: 2,
			Balance: 30000,
			Number: "1234 xxxx xxxx 7777",
			Active: true,
		},
		{
			ID: 3,
			Balance: -20000,
			Number: "1234 xxxx xxxx 9999",
			Active: true,
		},
	}
	sources := PaymentSources(cards)
	for _,source := range sources {
		fmt.Println(source.Number)
	}
	//Output: 
	// 1234 xxxx xxxx 8888
	// 1234 xxxx xxxx 7777
}