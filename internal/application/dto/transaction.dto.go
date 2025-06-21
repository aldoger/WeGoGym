package dto

type TransactionRequestDto struct {
	MembershipId string  `json:"membership_id"`
	Kode         string  `json:"kode"`
	Price        float64 `json:"harga"`
}

type TransactionResponseDto struct {
	TokenTransaksi string `json:"token"`
	UrlRedirect    string `json:"url_redirect"`
}
