package dto

type TransactionMemberRequestDto struct {
	MembershipId string `json:"membership_id"`
	Kode         string `json:"kode,omitempty" binding:"omitempty"`
}

type TransactionPersonalTrainerRequestDto struct {
	Harga int `json:"harga"`
	Sesi  int `json:"sesi"`
}

type TransactionResponseDto struct {
	TokenTransaksi string `json:"token"`
	RedirectUrl    string `json:"redirect_url"`
}
