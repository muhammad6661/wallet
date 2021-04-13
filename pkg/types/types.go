package types

type Money int64
type Phone string

type PaymentCategory string 


//PaymentStatus string
type PaymentStatus string

//Statuses
const(
  PaymentStatusOk PaymentStatus = "OK"
  PaymentStatusFail PaymentStatus = "FAIL"
  PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

type Payment struct {
	ID string
	AccountID int64
	Amount Money
    Category PaymentCategory
	Status PaymentStatus
}


type Account struct{
	ID int64
	Phone Phone
	Balance Money
}

type Favorites struct{
	ID string
    AccountID int64
	Name string
	Amount Money
	Category PaymentCategory
}