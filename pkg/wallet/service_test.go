package wallet

import (
	"testing"
	"fmt"
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
	pay,_:=sv.Pay(1,10,"Alif");fmt.Println(pay)
	err:=sv.Reject("2")


if err == nil {
	t.Errorf("\ngot > %v \nwant > %v", err,nil)
} 
}	

func TestService_Reject_faild(t *testing.T){
	sv:=Service{}
	_,_=sv.Pay(1,10,"Alif");
	err:=sv.Reject("20")

if err == nil {
	t.Errorf("\ngot > %v \nwant > %v", err,ErrPaymentNotFound)
} 
}	
func TestService_FindPaymentById_found(t *testing.T){
	sv:=Service{}
	_,_=sv.Pay(1,10,"Alif");
	_,err:=sv.FindPaymentById("5")


if err == nil {
	t.Errorf("\ngot > %v \nwant > %v", err,nil)
} 
}	
func TestService_FindPaymentById_faild(t *testing.T){
	sv:=Service{}
	_,_=sv.Pay(1,10,"Alif");
	_,err:=sv.FindPaymentById("10")

if err == nil {
	t.Errorf("\ngot > %v \nwant > %v", err,ErrPaymentNotFound)
} 
}	
