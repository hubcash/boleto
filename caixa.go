package goboleto

import (
	"encoding/base64"
)

// CEF - Caixa econômica federal
// Source: (http://www.caixa.gov.br/Downloads/cobranca-caixa/ESP_COD_BARRAS_SIGCB_COBRANCA_CAIXA.pdf)
type Caixa struct {
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

// configCaixa is a global for this bank configs
var configCaixa = bankConfig{
	Id: 104,
	Aceite: "N",
	Currency: 9,
	CurrencyName: "R$",
}

// Barcode Get the barcode, return string, it may contain dots and spaces
func (b Caixa) Barcode(d Document) string {
	return "1001.011011.1 123002 2"
}

// BarcodeImage return a image/base64, using a document
func (b Caixa) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

// Transference Return the transference file (arquivo de remessa)
func (b Caixa) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Caixa) Layout(d Document) {
	// TODO
}