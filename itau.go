package goboleto

import (
	"encoding/base64"
)

// Itau
// Source: (http://download.itau.com.br/bankline/cobranca_cnab240.pdf)
type Itau struct {
	Agency 			int
	Account			int
	Convenio		int
	Contrato		int
	Carteira		int
	VariacaoCarteira	int
	FormatacaoConvenio	int
	FormatacaoNossoNumero	int
	Company			*Company
}

// configItau is a global for this bank configs
var configItau = bankConfig{
	Id: 341,
	Aceite: "N",
	Currency: 9,
	CurrencyName: "R$",
}

// Barcode Get the barcode, return string, it may contain dots and spaces
func (b Itau) Barcode(d Document) string {
	return "1001.011011.1 123002 2"
}

// BarcodeImage return a image/base64, using a document
func (b Itau) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

// Transference Return the transference file (arquivo de remessa)
func (b Itau) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Itau) Layout(d Document) {
	// TODO
}