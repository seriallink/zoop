package zoop

type Address struct {
	Line1        string `json:"line1"`
	Line2        string `json:"line2"`
	Line3        string `json:"line3"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postal_code"`
	CountryCode  string `json:"country_code"`
}
