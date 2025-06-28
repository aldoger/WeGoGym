package dto

type TransactionRequestDto struct {
	MembershipId string `json:"membership_id"`
	Kode         string `json:"kode,omitempty" binding:"omitempty"`
}

type TransactionResponseDto struct {
	TokenTransaksi string `json:"token"`
	RedirectUrl    string `json:"redirect_url"`
}
