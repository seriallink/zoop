package zoop

type TransactionParams struct {
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	//PaymentType
}
