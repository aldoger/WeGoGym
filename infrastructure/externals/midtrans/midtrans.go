package midtrans

import (
	"os"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

const KODE_REFERAL = "RPLGACORR"

type MidtransClient struct {
	Client snap.Client
}

func (m *MidtransClient) CreateTransaction(price float64, email string, kode string) (string, error) {
	m.Client.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	orderID := uuid.New().String()

	var PriceInt = int64(price)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: PriceInt,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: email,
		},
	}

	snapResp, err := m.Client.CreateTransaction(req)
	if err != nil {
		return "", err
	}

	return snapResp.Token, nil
}
