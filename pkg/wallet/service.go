package wallet

import (
	"github.com/muhammad6661/wallet/pkg/types"
	"github.com/google/uuid"
	"errors"
)

//ErrPhoneRegistered -- phone already registred
var ErrPhoneRegistered = errors.New("phone already registred")
//ErrAmountMustBePositive -- amount must be greater than zero
var ErrAmountMustBePositive = errors.New("amount must be greater than zero")
//ErrAccountNotFound -- account not found
var ErrAccountNotFound = errors.New("account not found")
//ErrPaymentNotFound --payment not found
var ErrPaymentNotFound = errors.New("account not found")
//ErrNotEnoughtBalance -- account not found
var ErrNotEnoughtBalance = errors.New("account not enought balance")


//Service model
type Service struct{
  nextAccountID int64
  accounts []*types.Account
  payments []*types.Payment
}

//RegisterAccount meth
func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error){
   for _, account := range s.accounts {
     if account.Phone == phone{
       return nil, ErrPhoneRegistered
     }
   }

   s.nextAccountID++
   account := &types.Account{
  ID: s.nextAccountID,
  Phone: phone,
  Balance: 200,
}
   s.accounts = append(s.accounts, account)

   return account, nil
}


//Pay method
func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory)(*types.Payment, error)  {
  
  if amount <= 0{
    return nil, ErrAmountMustBePositive
  }
  var account *types.Account
  for _, ac := range s.accounts {
    if ac.ID == accountID{
      account = ac
      break
    }
  }
  if account == nil{
    return nil, ErrAccountNotFound
  }
  if account.Balance < amount{
    return nil, ErrNotEnoughtBalance
  }
  account.Balance -= amount
  paymentID := uuid.New().String()
  payment := &types.Payment{
    ID: paymentID,
    AccountID: accountID,
    Amount: amount,
    Category: category,
    Status: types.PaymentStatusInProgress,
  }
  s.payments = append(s.payments, payment)
  return payment, nil
}//Deposit method
func (s *Service) Deposit(accountID int64, amount types.Money)error {
  
  if amount <= 0{
    return ErrAmountMustBePositive
  }
  var account *types.Account
  for _, ac := range s.accounts {
    if ac.ID == accountID{
      account = ac
      break
    }
  }
  if account == nil{
    return  ErrAccountNotFound
  }
  
  account.Balance += amount
 
  return nil
}


func (s *Service)Reject(paymentID string ) error{
   
  var payment *types.Payment

   for _,i:=range s.payments{
     if(i.ID==paymentID){
      payment=i
      break
         }
       }
  
       if payment==nil{
         return ErrPaymentNotFound
       }

      var account *types.Account
      for _,i:=range s.accounts{
        if i.ID==payment.AccountID{
          account=i
          break
        }
      }

       if account==nil{
         return ErrPaymentNotFound
       }
     payment.Stsfatus=types.PaymentStatusFail
     account.Balance+=payment.Amount
   
   return nil
}

func (s *Service)FindPaymentByID(paymentID string)(*types.Payment,error){
  for _,i:=range s.payments{
    if(i.ID==paymentID){
      return i ,nil
    }
  }
  return nil, ErrPaymentNotFound
}


func (service *Service) FindAccountByID(AccountID int64)(*types.Account,error){
	
   for _,account:=range service.accounts{
	   if(account.ID==AccountID){
		   account.Balance=0
		   return account,nil
	   }
   }
   return nil,ErrAccountNotFound
}

