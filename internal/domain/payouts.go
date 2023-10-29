package domain

type PayoutSymbol struct {
	Symbol string `json:"symbol"`
	Payout [5]int `json:"payout"`
}

type Payouts []PayoutSymbol

func (p Payouts) Validate() error {
	return nil
}
