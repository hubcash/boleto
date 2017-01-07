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
	var barcode = Barcode{
		Field1: Field1{
			bank: configBB.Id,
			currency: configBB.Currency,
			numbers: 4444,
			dv: 8,
		},
		Field2: Field2{
			numbers: 999999999,
			dv: 8,
		},
		Field3: Field3{
			numbers: 999999999,
			dv: 8,
		},
		Field4: Field4{
			dv: 8,
		},
		Field5: Field5{
			dueDate: 4444,
			value: 0000000001,
		},
	}

	return generateBarcode(barcode)
}

// BarcodeImage return a image/base64, using a document
func (b BB) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

// Layout return a HTML template using a document
func (b BB) Layout(d Document) {
	// TODO
}