package goboleto

import (
	"encoding/base64"
)

// BB - Banco do Brasil
// Source: (http://www.bb.com.br/docs/pub/emp/mpe/espeboletobb.pdf)
type BB struct {
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

// configBB is a global for this bank configs
var configBB = bankConfig{
	Id: 001,
	Aceite: "N",
	Currency: 9,
	CurrencyName: "R$",
}

// Barcode Get the barcode, return string, it may contain dots and spaces
func (b BB) Barcode(d Document) string {
	var barcode = Barcode{}
	barcode.Field1 = &Field1{
		Bank: configBB.Id,
		Currency: configBB.Currency,
		Numbers: 4444,
		Dv: 8,
	}
	barcode.Field2 = &Field2{
		Numbers: 999999999,
		Dv: 8,
	}
	barcode.Field3 = &Field3{
		Numbers: 999999999,
		Dv: 8,
	}
	barcode.Field4 = &Field4{
		Dv: 8,
	}
	barcode.Field5 = &Field5{
		DateDue: dateDueFactor(d.DateDue),
		Value: formatValue(d.Value),
	}
	return generateBarcode(barcode)
}

// BarcodeImage return a image/base64, using a document
func (b BB) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

// Transference Return the transference file (arquivo de remessa)
func (b BB) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b BB) Layout(d Document) {
	// TODO
}