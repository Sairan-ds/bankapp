package payment

import "bank/pkg/bank/types"

// Max возвращает максимальный платеж
func Max(payments []types.Payment) types.Payment {
	var payment types.Payment = payments[0]

	for _,operation := range payments {
		if operation.Amount > payment.Amount {
			payment = operation

		}
	}
	return payment
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	
	var sources []types.PaymentSource
	for _, card := range cards{
		if card.Balance > 0 && card.Active  {
			sources = append(sources, types.PaymentSource{
				Type: "card",
				Number: string(card.PAN),


			})
		} 
	}
   }
   