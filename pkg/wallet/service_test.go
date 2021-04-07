package wallet

import (
	"testing"
)

func TestService_FindAccountById_success(t *testing.T){
  sv:=Service{}
// 	  accounts: []types.Account{
// 		{ID: 1, Phone: "901605036", Balance: 1_000_00},
// 		{ID: 2, Phone: "901605036", Balance: 2_000_00},
// 		{ID: 3, Phone: "901605036", Balance: 3_000_00},
// 		{ID: 4, Phone: "901605036", Balance: 4_000_00},
// 		{ID: 5, Phone: "901605036", Balance: 5_000_00},
// 		{ID: 6, Phone: "901605036", Balance: 6_000_00},
// 	   },
//   }
  sv.RegisterAccount("901605036")

_,err:=sv.FindAccountById(1)

if err != nil {
	t.Errorf("\ngot > %v \nwant > nil", err)
} 
}	

func TestService_Account_faild(t *testing.T){
	sv:=Service{}
	sv.RegisterAccount("901605036")

_,err:=sv.FindAccountById(7)

if err == nil {
	t.Errorf("\ngot > %v \nwant > %v", err,ErrAccountNotFound)
} 
}	
