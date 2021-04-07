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
