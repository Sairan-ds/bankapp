package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.п)
type Money int64

//
type Currency string

//
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//
type PAN string

//
type Card struct {
	ID		 int
	PAN		 PAN
	Balance	 Money
	Currency Currency
	Color	 string
	Name	 string
	Active	 bool
	MinBalance Money
	Number string
}

// Payment представляет информацию о платеже
type Payment struct {
	ID int
	Amount Money
}


// PaymentSource - выбор карты
type PaymentSource struct {
	Type string //'card'
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}