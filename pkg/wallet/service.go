package wallet

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/muhammad6661/wallet/pkg/types"
)

//ErrPhoneRegistered -- phone already registred
var ErrPhoneRegistered = errors.New("phone already registred")
//ErrAmountMustBePositive -- amount must be greater than zero
var ErrAmountMustBePositive = errors.New("amount must be greater than zero")
//ErrAccountNotFound -- account not found
var ErrAccountNotFound = errors.New("account not found")
//ErrPaymentNotFound --payment not found
var ErrPaymentNotFound = errors.New("payment not found")
//ErrNotEnoughtBalance -- account not found
var ErrNotEnoughtBalance = errors.New("account not enought balance")
//ErrFavoriteNotFound --payment not found
var ErrFavoriteNotFound = errors.New("favorite not found")

//Service model
type Service struct{
  nextAccountID int64
  accounts []*types.Account
  payments []*types.Payment
  favorites []*types.Favorites
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
}
//Deposit method
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

//Reject method for error payments
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

       account,err:=s.FindAccountByID(payment.AccountID)

       if err!=nil{
         return err
       }
       fmt.Println(account)
       fmt.Println(payment)

     payment.Status=types.PaymentStatusFail
     account.Balance+=payment.Amount
     fmt.Println(account)
     fmt.Println(payment)
   return nil
}

//method Find Payment By Id
func (s *Service)FindPaymentByID(paymentID string)(*types.Payment,error){
  for _,i:=range s.payments{
    if(i.ID==paymentID){
      return i ,nil
    }
  }
  return nil, ErrPaymentNotFound
}

// method Find Account By ID
func (service *Service) FindAccountByID(AccountID int64)(*types.Account,error){
	
   for _,account:=range service.accounts{
	   if(account.ID==AccountID){
		   return account,nil
	   }
   }
   return nil,ErrAccountNotFound
}






//method Repeat Function for payments
func (service *Service)Repeat(paymentID string)(*types.Payment,error){
  payment,_:=service.FindPaymentByID(paymentID)
  if(payment==nil){
    return nil,ErrPaymentNotFound
  }
   new_pay,err:=service.Pay(payment.AccountID,payment.Amount,payment.Category)
   if err!=nil{
     return nil,err
   }

  return new_pay,nil
 
}


//Favorite method

func (s *Service)FavoritePayment(paymentID string,name string)(*types.Favorites,error){
  
   payment,_:=s.FindPaymentByID(paymentID)

   if(payment==nil){
     return nil,ErrPaymentNotFound
   }
   
   
   favorite:=&types.Favorites{
     ID : "1",
     AccountID : payment.AccountID,
     Amount: payment.Amount,
     Name: name,
     Category: payment.Category,
   }
      s.favorites=append(s.favorites, favorite)
   return favorite,nil

}

//method PayFromFavorite for payment

func (s *Service)PayFromFavorite(favoriteID string)(*types.Payment,error){
 
  favorite,_:=s.FindFavoriteByID(favoriteID)

  if favorite==nil{
    return nil,ErrFavoriteNotFound
  }

  payment,_:=s.Pay(favorite.AccountID,favorite.Amount,favorite.Category)
   
  if payment==nil{
    return nil,ErrFavoriteNotFound
  }
   return payment,nil
}

//method for Find Favorite By Id

func (s *Service)FindFavoriteByID(favoriteID string) (*types.Favorites,error){
 
  for _,i:=range s.favorites{
    if i.ID==favoriteID{
      return i,nil
    }
  }
   return nil,ErrFavoriteNotFound
}



//Method for export Account to file

func (s *Service)ExportToFile(path string) error{

    
  file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if err!=nil{
    return err
  }

  defer file.Close()


  var str string

 for _,item_account:=range s.accounts{
  str+=fmt.Sprint(item_account.ID)+";"+string(item_account.Phone)+fmt.Sprint(item_account.Balance)+"|"
 }
  _,err=file.WriteString(str)

  if err!=nil{
    return err
  }

  fmt.Println(path)
    return nil
}