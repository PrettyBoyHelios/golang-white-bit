package models

type WithdrawParams struct {
	Ticker   string `json:"ticker"`
	Amount   string `json:"amount"`
	Address  string `json:"address"`
	Memo     string `json:"memo"`
	UniqueID string `json:"uniqueId"`
	Request  string `json:"request"`
	Nonce    string `json:"nonce"`
}
