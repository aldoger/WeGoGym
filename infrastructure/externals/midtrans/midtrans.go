package midtrans

import (
	"go-kpl/internal/domain/models"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransClient struct {
	Client snap.Client
}

func NewMidtrans() *MidtransClient {
	var NewClient snap.Client
	NewClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
	return &MidtransClient{Client: NewClient}
}

func (m *MidtransClient) CreateMemberTransaction(userId string, email string, kode string, membershipDetail models.Membership) (*snap.Response, error) {

	var PriceInt = int64(membershipDetail.Price)

	var KODE_REFERAL = os.Getenv("KODE_REFERAL")

	if kode != "" && kode == KODE_REFERAL {
		Discount := (PriceInt * 20) / 100
		PriceInt -= Discount
	}

	items := []midtrans.ItemDetails{
		{
			ID:    membershipDetail.Id.String(),
			Name:  membershipDetail.Type + " Membership",
			Price: PriceInt,
			Qty:   1,
		},
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  userId,
			GrossAmt: PriceInt,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: email,
		},
		Items: &items,
	}

	snapResp, err := m.Client.CreateTransaction(req)
	if err != nil {
		return nil, err
	}

	return snapResp, nil
}
