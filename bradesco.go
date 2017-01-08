package goboleto

import (
	"encoding/base64"
)

// Bradesco
// Source: (https://banco.bradesco/assets/pessoajuridica/pdf/4008-524-0121-08-layout-cobranca-versao-portuguesSS28785.pdf)
type Bradesco struct {
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

// configBradesco is a global for this bank configs
var configBradesco = bankConfig{
	Id: 237,
	Aceite: "N",
	Currency: 9,
	CurrencyName: "R$",
}

// Barcode Get the barcode
func (b Bradesco) Barcode(d Document) string {
	return "12345678911111111112222222222333333333344444";
}

// Barcode Get the barcode digitable number (Linha digitavel), return string,
// it may contain dots and spaces
func (b Bradesco) BarcodeDigitable(d Document) string {
	return "1001.011011.1 123002 2"
}

// BarcodeImage return a image/base64, using a document
func (b Bradesco) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

// Transference Return the transference file (arquivo de remessa)
func (b Bradesco) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Bradesco) Layout(d Document) {
	// TODO
}