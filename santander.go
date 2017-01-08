package goboleto

import (
	"encoding/base64"
)

// Santander
// Source: (https://www.santander.com.br/document/wps/sl-tabela-de-tarifas-cobranca.pdf)
type Santander struct {
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

// configSantander is a global for this bank configs
var configSantander = bankConfig{
	Id: 033,
	Aceite: "N",
	Currency: 9,
	CurrencyName: "R$",
}

// Barcode Get the barcode
func (b Santander) Barcode(d Document) string {
	return "12345678911111111112222222222333333333344444";
}

// Barcode Get the barcode digitable number (Linha digitavel), return string,
// it may contain dots and spaces
func (b Santander) BarcodeDigitable(d Document) string {
	return "1001.011011.1 123002 2"
}

// BarcodeImage return a image/base64, using a document
func (b Santander) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

// Transference Return the transference file (arquivo de remessa)
func (b Santander) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Santander) Layout(d Document) {
	// TODO
}