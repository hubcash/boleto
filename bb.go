package goboleto

import (
	"encoding/base64"
)

/*
Bank constants
@ID FEBRABAN bank identifier
@Aceite if the payer accept the billet
@Currency the currency identifier
@CurrencyName the currency name
 */
var configBB = map[string]interface{}{
	"ID": 001,
	"Aceite": "N",
	"Currency": 9,
	"CurrencyName": "R$",
}

/*
BB - Banco do Brasil
Source: (http://www.bb.com.br/docs/pub/emp/mpe/espeboletobb.pdf)
 */
type BB struct {
	Agency 			int
	Account			int
	Convenio		int
	Contrato		int
	Carteira		int
	VariacaoCarteira	int
	FormatacaoConvenio	int
	FormatacaoNossoNumero	int
	Company			Company
}

func (b BB) Barcode(d Document) string {
	var barcode = Barcode{
		Field1: Field1{
			bank: configBB["ID"].(int),
			currency: configBB["Currency"].(int),
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

func (b BB) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

func (b BB) Layout(d Document) {
	// TODO
}