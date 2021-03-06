package wallet

import (
	"fmt"
	"testing"
)

func TestService_FindAccountById_success(t *testing.T){
  sv:=Service{}

  sv.RegisterAccount("901605036")

_,err:=sv.FindAccountByID(1)

if err != nil {
	t.Errorf("\ngot > %v \nwant > nil", err)
} 
}	

func TestService_Account_faild(t *testing.T){
	sv:=Service{}
	sv.RegisterAccount("901605036")

_,err:=sv.FindAccountByID(7)

if err == nil {
	t.Errorf("\ngot > %v \nwant > %v", err,ErrAccountNotFound)
} 
}	
func TestService_Reject_found(t *testing.T){
	sv:=&Service{}
	sv.RegisterAccount("901605036")
	account,_:=sv.FindAccountByID(1)
	sv.Deposit(account.ID,200)
	pay,err:=sv.Pay(1,10,"Alif");
	if pay.Category!="Alif"{
		fmt.Println(pay,err)
	}
	err=sv.Reject(pay.ID)
if err != nil {
	t.Errorf("\ngot > %v \nwant > %v", err,nil)
} 
}	

func TestService_Reject_faild(t *testing.T){
	sv:=Service{}
	sv.RegisterAccount("901605036")
	_,_=sv.Pay(1,10,"Alif");
	err:=sv.Reject("20")

if err == nil {
	t.Errorf("\ngot > %v \nwant > %v", err,ErrPaymentNotFound)
} 
}	
func TestService_FindPaymentById_found(t *testing.T){
	sv:=Service{}
	sv.RegisterAccount("901605036")
	account,_:=sv.FindAccountByID(1)
	sv.Deposit(account.ID,200)
	pay,_:=sv.Pay(1,10,"Alif");
	_,err:=sv.FindPaymentByID(pay.ID)


if err != nil {
	t.Errorf("\ngot > %v \nwant > %v", err,nil)
} 
}	
func TestService_FindPaymentByID_faild(t *testing.T){
	sv:=Service{}
	sv.RegisterAccount("901605036")
	

	_,_=sv.Pay(1,10,"Alif");
	_,err:=sv.FindPaymentByID("10")

if err == nil {
	t.Errorf("\ngot > %v \nwant > %v", err,ErrPaymentNotFound)
} 
}	

func TestService_Deposit_correct(t *testing.T){
	sv:=Service{}
	account,err:=sv.RegisterAccount("901605036")

	err=sv.Deposit(account.ID,10)

if err != nil {
	t.Errorf("\ngot > %v \nwant > %v", err,nil)
} 
}	
func TestService_Deposit_incorrect(t *testing.T){
	sv:=Service{}
	_,_=sv.RegisterAccount("901605036")

	err:=sv.Deposit(4,10)

if err == nil {
	t.Errorf("\ngot > %v \nwant > %v", err,ErrAccountNotFound)
} 
}	


func Test_Repeat_success(t*testing.T){
sv:=Service{}
  account,_:=sv.RegisterAccount("901605036")
  
  _=sv.Deposit(account.ID,100_000)
  payment,_:=sv.Pay(account.ID,100,"alif")
   
  pay,_:=sv.FindPaymentByID(payment.ID)


 _,err:=sv.Repeat(pay.ID)

  if(err!=nil){
	t.Errorf("\ngot > %v \nwant > %v", err,ErrAccountNotFound)
  }

}



func TestFavorite_fsuccess(t *testing.T){

 sv:=Service{}
 account,_:=sv.RegisterAccount("901605036")

 _=sv.Deposit(account.ID,10_000)

 payment,_:=sv.Pay(account.ID,1000,"ALif")

 favorite,err:=sv.FavoritePayment(payment.ID,"Academy")

 if(err!=nil){
	t.Errorf("method PayFromFavorite returned not nil error, paymentFavorite => %v", favorite)

 }

 pay_favorite,err:=sv.PayFromFavorite(favorite.ID)
  
 if err!=nil{
	t.Errorf("method PayFromFavorite returned not nil error, payfromtFavorite => %v", pay_favorite)
 }
}


func TestService_Export_success_user(t *testing.T) {
	var svc Service

	svc.RegisterAccount("901605036")
	svc.RegisterAccount("901605037")
	svc.RegisterAccount("901605038")

	err := svc.ExportToFile("export.txt")
	if err != nil {
		t.Errorf("method ExportToFile returned not nil error, err => %v", err)
	}

}
func TestService_Import_success(t *testing.T) {
	var svc Service

	err := svc.ExportToFile("export.txt")
	if err != nil {
		t.Errorf("method ExportToFile returned not nil error, err => %v", err)
	}

   err=svc.ImportFromFile("export.txt")
   if err != nil {
	t.Errorf("method ImportToFile returned not nil error, err => %v", err)
} 
}


func TestService_Export(t *testing.T) {
	var svc Service

	account1,_:=svc.RegisterAccount("901605036")
	account2,_:=svc.RegisterAccount("901605037")
	account3,_:=svc.RegisterAccount("901605038")

	_=svc.Deposit(account1.ID,10_000)
	_=svc.Deposit(account2.ID,10_000)
	_=svc.Deposit(account3.ID,10_000)

	payment1,_:=svc.Pay(account1.ID,1000,"ALif1")
	payment2,_:=svc.Pay(account2.ID,1000,"ALif2")
	payment3,_:=svc.Pay(account3.ID,1000,"ALif3")
   
	_,_=svc.FavoritePayment(payment1.ID,"Academy1")
	_,_=svc.FavoritePayment(payment2.ID,"Academy2")
	_,_=svc.FavoritePayment(payment3.ID,"Academy3")


	err := svc.Export("D:/Alif-Academy/Golang/src/17/1")
	if err != nil {
		t.Errorf("method ExportToFile returned not nil error, err => %v", err)
	}

}


func TestService_Import(t *testing.T) {
	var svc Service


	err := svc.Import("D:/Alif-Academy/Golang/src/17/1")
	if err != nil {
		t.Errorf("method ImportFromFile returned not nil error, err => %v", err)
	}

}