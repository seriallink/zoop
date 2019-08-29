package zoop

type Owner struct {
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phone_number"`
	TaxPayerId  string   `json:"taxpayer_id"`
	Birthdate   *Date    `json:"birthdate"`
	Address     *Address `json:"address"`
}
