package dto

type TransactionDto struct {
	MembershipId string  `json:"membership_id"`
	Kode         string  `json:"kode"`
	Price        float64 `json:"harga"`
}
