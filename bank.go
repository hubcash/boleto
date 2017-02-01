package boleto

import "net/http"

// Bank is defined as an interface,
// then we force to implement these functions:
// @Barcode Get the BarcodeNumber
// @Transference Return the transference file (arquivo de remessa)
// @Layout return a HTML template using a document
type Bank interface {
	Barcode(Document) Barcode
	Transference(Document)
	Layout(http.ResponseWriter, Document)
}

// Defines a bank configuration type,
// holds constants to use it:
// @Id FEBRABAN bank identifier
// @Aceite if the payer accepted the billet
// @Currency the currency identifier
// @CurrencyName the currency name
// @AgencyMaxSize the agency size
// @AccountMaxSize the account size
type bankConfig struct {
	Id             int
	Aceite         string
	Currency       int
	CurrencyName   string
	AgencyMaxSize  int
	AccountMaxSize int
}
