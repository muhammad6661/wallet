package types

type Money int64
type Phone string

type PaymentCategory string 


type Payment struct {
	ID string
	AccountID int64
	Amount Money
    Category PaymentCategory
}


type Account struct{
	ID int64
	Phone Phone
	Balance Money
}