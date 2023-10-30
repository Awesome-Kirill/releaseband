package domain

import "fmt"

type PayoutSymbol struct {
	Symbol string `json:"symbol"`
	Payout [5]int `json:"payout"`
}

type Payouts []PayoutSymbol

// Validate Payouts
func (p Payouts) Validate() error {
	for _, payoutSymbol := range p {
		_, ok := Alphabet[payoutSymbol.Symbol]
		if !ok {
			return fmt.Errorf("eroor : symbol:%v not in alphabet", payoutSymbol)
		}
	}
	return nil
}
