package wallet

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

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
func (s *Service) ExportToFile(path string) error {

  file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    return err
  }
  defer file.Close()

  var str string
  for _, v := range s.accounts {
    str += fmt.Sprint(v.ID) + ";" + string(v.Phone) + ";" + fmt.Sprint(v.Balance) + "|"
  }
  _, err = file.WriteString(str)

  if err != nil {
    return err
  }

  return nil
}

func (s *Service)ImportFromFile(path string) error{
 
   data,err:=ioutil.ReadFile(path)
   if err!=nil{
     return err
   }
   str:=strings.Split(string(data),"|")

   for i:=0; i<len(str)-1;i++{
      str_item:=strings.Split(str[i],";")
      id, _:= strconv.ParseInt(str_item[0], 10, 64)
      balance, _:= strconv.ParseInt(str_item[2], 10, 64)
       phone:=(str_item[1])
       s.accounts=append( s.accounts,&types.Account{
        ID: id,
        Phone: types.Phone(phone),
       Balance: types.Money(balance),
      })

   }
  return nil
}



//Method for Full-Export service;
func (s *Service) Export(dir string) error {

     dirAcounts:=dir+"/accounts.dump"
     dirPayments:=dir+"/payments.dump"
     dirFavorites:=dir+"/favorites.dump"
//File Accounts
  fmt.Println("PAs=",dir);

  var str string
  kA:=0
  for _, v := range s.accounts {
    kA++
    str += fmt.Sprint(v.ID) + ";" + string(v.Phone) + ";" + fmt.Sprint(v.Balance) + "\n"
  }

  if(kA!=0){
  fileAccounts, _ := os.OpenFile(dirAcounts, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
  defer fileAccounts.Close()

  _, _= fileAccounts.WriteString(str)
 
}



  //File Payments

 
  var strP string
  kP:=0
  for _, v := range s.payments {
    kP++
    strP += fmt.Sprint(v.ID) + ";" + fmt.Sprint(v.AccountID) + ";" + fmt.Sprint(v.Amount) +";"+ fmt.Sprint(v.Category) +";"+ fmt.Sprint(v.Status) + "\n"
  }
  if(kP!=0){
  filePayments, _:= os.OpenFile(dirPayments, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
  
  

  defer filePayments.Close()

  _, _= filePayments.WriteString(strP)


}


    //File Favorites
  
    var strF string
    kF:=0
    for _, v := range s.favorites {
      kF++
      strF += fmt.Sprint(v.ID) + ";" + fmt.Sprint(v.AccountID) + ";" + fmt.Sprint(v.Name) +";"+ fmt.Sprint(v.Amount) +";"+ fmt.Sprint(v.Category) + "\n"
    }
    if(kF!=0){
    fileFavorites, _:= os.OpenFile(dirFavorites, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
   
    defer fileFavorites.Close()
    _, _= fileFavorites.WriteString(strF)
  
  }


  return nil
}


//Method for Full-Import service
func (s *Service)Import(dir string) error{
  
	_, err := os.Stat(dir + "/accounts.dump")

	if err == nil {	
		content, err := ioutil.ReadFile(dir + "/accounts.dump")
		if err != nil {
			return err
		}

		strArray := strings.Split(string(content), "\n")
		if len(strArray) > 0 {
			strArray = strArray[:len(strArray)-1]
		}
		for _, v := range strArray {
			strArrAcount := strings.Split(v, ";")
			fmt.Println(strArrAcount)

			id, err := strconv.ParseInt(strArrAcount[0], 10, 64)
			if err != nil {
				return err
			}
			balance, err := strconv.ParseInt(strArrAcount[2], 10, 64)
			if err != nil {
				return err
			}
			flag := true
			for _, v := range s.accounts {
				if v.ID == id {
					v.Phone = types.Phone(strArrAcount[1])
					v.Balance = types.Money(balance)
					flag = false
				}
			}
			if flag {
				account := &types.Account{
					ID:      id,
					Phone:   types.Phone(strArrAcount[1]),
					Balance: types.Money(balance),
				}
				s.accounts = append(s.accounts, account)
			}
		}
	}

	_, err1 := os.Stat(dir + "/payments.dump")

	if err1 == nil {
		content, err := ioutil.ReadFile(dir + "/payments.dump")
		if err != nil {
			return err
		}

		strArray := strings.Split(string(content), "\n")
		if len(strArray) > 0 {
			strArray = strArray[:len(strArray)-1]
		}
		for _, v := range strArray {
			strArrAcount := strings.Split(v, ";")
			fmt.Println(strArrAcount)

			id := strArrAcount[0]
			if err != nil {
				return err
			}
			aid, err := strconv.ParseInt(strArrAcount[1], 10, 64)
			if err != nil {
				return err
			}
			amount, err := strconv.ParseInt(strArrAcount[2], 10, 64)
			if err != nil {
				return err
			}
			flag := true
			for _, v := range s.payments {
				if v.ID == id {
					v.AccountID = aid
					v.Amount = types.Money(amount)
					v.Category = types.PaymentCategory(strArrAcount[3])
					v.Status = types.PaymentStatus(strArrAcount[4])
					flag = false
				}
			}
			if flag {
				data := &types.Payment{
					ID:        id,
					AccountID: aid,
					Amount:    types.Money(amount),
					Category:  types.PaymentCategory(strArrAcount[3]),
					Status:    types.PaymentStatus(strArrAcount[4]),
				}
				s.payments = append(s.payments, data)
			}
		}
	}

	_, err2 := os.Stat(dir + "/favorites.dump")

	if err2 == nil {
		content, err := ioutil.ReadFile(dir + "/favorites.dump")
		if err != nil {
			return err
		}

		strArray := strings.Split(string(content), "\n")
		if len(strArray) > 0 {
			strArray = strArray[:len(strArray)-1]
		}
		for _, v := range strArray {
			strArrAcount := strings.Split(v, ";")
			fmt.Println(strArrAcount)

			id := strArrAcount[0]
			if err != nil {
				return err
			}
			aid, err := strconv.ParseInt(strArrAcount[1], 10, 64)
			if err != nil {
				return err
			}
			amount, err := strconv.ParseInt(strArrAcount[2], 10, 64)
			if err != nil {
				return err
			}
			flag := true
			for _, v := range s.favorites {
				if v.ID == id {
					v.AccountID = aid
					v.Amount = types.Money(amount)
					v.Category = types.PaymentCategory(strArrAcount[3])
					flag = false
				}
			}
			if flag {
				data := &types.Favorites{
					ID:        id,
					AccountID: aid,
					Amount:    types.Money(amount),
					Category:  types.PaymentCategory(strArrAcount[3]),
				}
				s.favorites = append(s.favorites, data)
			}
		}
	}

	return nil}

func(s*Service)FillAccountFromFile(path string) error{
  fileAccounts, err := os.Open(path)
  _,err=fileAccounts.Stat()

defer fileAccounts.Close()
if(err==nil){
  readerA:=bufio.NewReader(fileAccounts)
  for{
	line,err:=readerA.ReadString('\n')
	if err==io.EOF{
		fmt.Println(line)
		break
	}
	if err!=nil{
		fmt.Println(err)
		return err
	}
  str_item:=strings.Split(line,";")
  id, _:= strconv.ParseInt(str_item[0], 10, 64)
  balance, _:= strconv.ParseInt(str_item[2], 10, 64)
   phone:=(str_item[1])
   account,err:=s.FindAccountByID(id)
   if  err==nil {
    account.Balance=types.Money(balance)
    account.Phone=types.Phone(phone)
} else{
   s.accounts=append( s.accounts,&types.Account{
    ID: id,
    Phone: types.Phone(phone),
   Balance: types.Money(balance),
  })
}

}
}
 return nil
}

func(s*Service)FillPaymentsFromFile(path string) error{
  filePayments, err := os.Open(path)
  _,err=filePayments.Stat()

defer filePayments.Close()
if(err==nil){
  readerA:=bufio.NewReader(filePayments)
  for{
	line,err:=readerA.ReadString('\n')
	if err==io.EOF{
		fmt.Println(line)
		break
	}
	if err!=nil{
		fmt.Println(err)
		return err
	}
  str_item:=strings.Split(line,";")
  id:= str_item[0]
  AccountID, _:= strconv.ParseInt(str_item[1], 10, 64)
  Amount, _:= strconv.ParseInt(str_item[2], 10, 64)
  Category:=str_item[3]
  Status:=str_item[4]
   payment,err:=s.FindPaymentByID(id)
   if  err==nil {
    payment.Amount=types.Money(Amount)
    payment.Category=types.PaymentCategory(Category)
    payment.Status=types.PaymentStatus(Status)
} else{
   s.payments=append( s.payments,&types.Payment{
    ID: id,
    AccountID: AccountID,
    Amount: types.Money(Amount),
   Category: types.PaymentCategory(Category),
   Status: types.PaymentStatus(Status),
  })
}

}
}
 return nil
}


//Fill Favorites sllice service

func(s*Service)FillFavoritesFromFile(path string) error{
  fileFavorites, err := os.Open(path)
_,err=fileFavorites.Stat()

defer fileFavorites.Close()
if(err==nil){
  readerA:=bufio.NewReader(fileFavorites)
  for{
	line,err:=readerA.ReadString('\n')
	if err==io.EOF{
		fmt.Println(line)
		break
	}
	if err!=nil{
		fmt.Println(err)
		return err
	}
  str_item:=strings.Split(line,";")
  id:= str_item[0]
  AccountID, _:= strconv.ParseInt(str_item[1], 10, 64)
  Name:=str_item[2]
  Amount, _:= strconv.ParseInt(str_item[3], 10, 64)
  Category:=str_item[4]

   favorite,err:=s.FindFavoriteByID(id)
   if  err==nil {
    favorite.Amount=types.Money(Amount)
    favorite.Category=types.PaymentCategory(Category)
    favorite.Name=Name
    favorite.AccountID=AccountID  
} else{
   s.favorites=append( s.favorites,&types.Favorites{
    ID: id,
    AccountID: AccountID,
    Amount: types.Money(Amount),
   Category: types.PaymentCategory(Category),
   Name:Name,
  })
}

}
}
 return nil
}
