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

func (m *MidtransClient) CreateTransaction(kode string) (*snap.Response, error) {
	m.Client.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	orderID := uuid.New().String()

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: 200000,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "Customer",
			Email: "customer@example.com",
		},
	}

	snapResp, err := m.Client.CreateTransaction(req)
	if err != nil {
		return nil, err
	}

	return snapResp, nil
}
