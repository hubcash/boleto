package boleto

// Defines a company type,
// holds the data of the emissor
// @LegalName the 'Razão social' of your company
type Company struct {
	Name      string
	LegalName string
	Document  string
	Address   string
	Contact   string
}
