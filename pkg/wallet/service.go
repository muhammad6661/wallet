package wallet

import (
	"github.com/muhammad6661/wallet/pkg/types"
	"errors"
)


var ErrAccountNotFound =errors.New("Account Not Found")



type Service struct{
	accounts []types.Account
 }



func (service *Service) FindAccountById(AccountID int64)(*types.Account,error){
	
   for _,account:=range service.accounts{
	   if(account.ID==AccountID){
		   result:=&types.Account{
			   ID: account.ID,
			   Phone: account.Phone,
			   Balance: account.Balance,
		   }
		   return result,nil
	   }
   }
   return nil,ErrAccountNotFound
}