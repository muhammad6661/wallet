package wallet

import (
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
	sv:=Service{}
	sv.RegisterAccount("901605036")
	pay,_:=sv.Pay(1,10,"Alif");
	err:=sv.Reject(pay.ID)
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
