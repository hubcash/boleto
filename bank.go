package goboleto

// Bank is defined as an interface,
// then we force to implement these functions:
// @Barcode Get the BarcodeNumber
// @Transference Return the transference file (arquivo de remessa)
// @Layout return a HTML template using a document
type Bank interface {
	Barcode(Document) BarcodeNumber
	Transference(Document)
	Layout(Document)
}

// Defines a bank configuration type,
// holds constants to use it:
// @Id FEBRABAN bank identifier
// @Aceite if the payer accepted the billet
// @Currency the currency identifier
// @CurrencyName the currency name
type bankConfig struct {
	Id 		int
	Aceite 		string
	Currency 	int
	CurrencyName 	string
}