package payment

import (
	user "User/User"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type service struct {
}

/*

idtrans provides a comprehensive payment gateway solution that supports various payment methods,
including credit/debit cards, bank transfers, e-wallets (such as UPI, Google Pay, and PhonePe), and more.

*/

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {

	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {

	midclient := midtrans.NewClient()

	midclient.ServerKey = ""
	midclient.ClientKey = ""
	midclient.APIEnvType = midtrans.Sandbox // Sandbox is environment for intracting the third party api for testing
	//test various scenarios, such as successful transactions, declined payments, and error handling, without any financial risk

	snapGateway := midtrans.SnapGateway{

		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{

		CustomerDetail: &midtrans.CustDetail{

			Email: user.Email,
			FName: user.Name,
		},

		TransactionDetails: midtrans.TransactionDetails{

			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq) // GetToken generate the token

	if err != nil {

		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
