package entity

// CardPayload is credit card info
type CardPayload struct {
	Number     string
	ExpMonth   int
	ExpYear    int
	HolderName string
	CVC        int
}

// NewCardPayload constructs card information
func NewCardPayload(
	number string, expMonth, expYear int, holderName string, cvc int,
) CardPayload {
	return CardPayload{
		Number:     number,
		ExpMonth:   expMonth,
		ExpYear:    expYear,
		HolderName: holderName,
		CVC:        cvc,
	}
}
