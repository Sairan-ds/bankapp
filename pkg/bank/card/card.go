package card

import (
	"bank/pkg/bank/types"
)

// IssueCard - выпуск карты
func IssueCard(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
	}
}
// Withdrow - снятие средств
func Withdraw(card types.Card, amount types.Money) types.Card {
	if !card.Active {
		return card
	}
	if amount < 0  {
		return card
	}
	if amount > card.Balance  {
		return card
	}
	if amount > 2000000  {
		return card
	}
	card.Balance = card.Balance - amount
	return card
}


func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if amount > 5000000  {
		return
	}
	if amount < 0  {
		return
	}
	card.Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}
	if card.Balance < 0 || card.MinBalance < 0 {
		return
	}

	bonus := card.MinBalance * types.Money(percent) / 100 * types.Money(daysInMonth) / types.Money(daysInYear)

	if bonus > 500_000 {return}

	if bonus<=500_000 {	card.Balance  += bonus}

}
// Total - sum of all cards balance
func Total(cards []types.Card) types.Money {
	var sum types.Money
	for _,operation := range cards {
		if operation.Balance >0 && operation.Active {sum += operation.Balance}
	}
	return sum

}

// PaymentSources - выбор карты для оплаты
func PaymentSources(cards []types.Card) []types.PaymentSource {
	
	var sources []types.PaymentSource
	for _, card := range cards{
		if card.Balance > 0 && card.Active  {
			sources = append(sources, types.PaymentSource{
				Type: "card",
				Number: string(card.PAN),
				Balance: card.Balance,

			})
		} 
	}
	return sources
   }