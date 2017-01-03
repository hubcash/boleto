package boleto

import "time"

type Document struct {
	Id 		int
	Number 		int
	Date		time.Time
	DateDue		time.Time
	Value 		float64
	ValueTax 	float64
	FebrabanType	string
	Instructions 	[6]string
	Client		Client
	Company		Company
}

type Client struct {
	Name 		string
	Address 	string
	Contact 	string
}

type Company struct {
	Name 		string
	LegalName 	string
	Document 	string
	Address 	string
	Contact 	string
}