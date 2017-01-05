package goboleto

import "encoding/base64"

/*
Bank constants
@ID FEBRABAN bank identifier
@Aceite if the payer accept the billet
@Currency the currency identifier
@CurrencyName the currency name
 */
var configSantander = map[string]interface{}{
	"ID": 033,
	"Aceite": "N",
	"Currency": 9,
	"CurrencyName": "R$",
}

/*
Santander
Source: (https://www.santander.com.br/document/wps/sl-tabela-de-tarifas-cobranca.pdf)
 */
type Santander struct {
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

func (b Santander) Barcode(d Document) string {
	return "1001.011011.1 123002 2"
}

func (b Santander) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

func (b Santander) Layout(d Document) {
	// TODO
}