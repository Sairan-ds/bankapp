package card

import (
	"bank/pkg/bank/types"
	"fmt"
)
func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
	}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 10_000_00, Active: true}, 15_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
	}

func ExampleWithdraw_inActive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
	}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 30_000_00, Active: true}, 21_000_00)
	fmt.Println(result.Balance)
	// Output: 3000000
	}


func ExampleDeposit_positive() {
	card := types.Card{Balance: 30_000_00, Active: true}
	Deposit(&card, 21_000_00)
	fmt.Println(card.Balance)
	// Output: 5100000
	}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 30_000_00, Active: false}
	Deposit(&card, 21_000_00)
	fmt.Println(card.Balance)
	// Output: 3000000
	}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 30_000_00, Active: true}
	Deposit(&card, 51_000_00)
	fmt.Println(card.Balance)
	// Output: 3000000
	}

func ExampleDeposit_negative() {
	card := types.Card{Balance: -30_000_00, Active: true}
	Deposit(&card, 40_000_00)
	fmt.Println(card.Balance)
	// Output: 1000000
	}


func ExampleAddBonus_positive() {
	card := types.Card{Balance: 10_000_00, Active: true, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1002465
	}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 10_000_00, Active: false, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1000000
	}
func ExampleAddBonus_negative() {
	card := types.Card{Balance: -10_000_00, Active: true, MinBalance: -10_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -1000000
	}
func ExampleAddBonus_extraBonus() {
	card := types.Card{Balance: 600_000_00, Active: true, MinBalance: 600_000_00}
	AddBonus(&card, 100, 30, 365)
	fmt.Println(card.Balance)
	// Output: 60000000
	}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
	}

	fmt.Println(Total(cards))
	//Output: 3000000
	
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			ID: 01,
			Balance: 10000,
			Active: true,
		},
		{
			ID: 02,
			Balance: 30000,
			Active: true,
		},
		{
			ID: 03,
			Balance: -20000,
			Active: true,
		},
	}
	sources := PaymentSources(cards)
	for _,source := range sources {
		fmt.Println(source.Number)
	}
	//Output: 
	// 1
	// 2
}