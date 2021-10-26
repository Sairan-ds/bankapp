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