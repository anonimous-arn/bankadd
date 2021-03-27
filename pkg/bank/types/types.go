package types

// Money представляет собой денежную сумму  в максимальных и минимальнвх единицах
type Money int64

// Currency представляет код валюты 
type Currency string 

// Коды валюты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN представляет номер карты
type PAN string

// Card предстовляет информацию о платёжной карте
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
	MinBalance    Money
}
// Payment...
type Payment struct{
	ID int
	Amount Money
}
