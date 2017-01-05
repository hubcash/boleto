package goboleto

import "encoding/base64"

/*
Bank constants
@ID FEBRABAN bank identifier
@Aceite if the payer accept the billet
@Currency the currency identifier
@CurrencyName the currency name
 */
var configItau = map[string]interface{}{
	"ID": 341,
	"Aceite": "N",
	"Currency": 9,
	"CurrencyName": "R$",
}

/*
Itau
Source: (http://download.itau.com.br/bankline/cobranca_cnab240.pdf)
 */
type Itau struct {
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

func (b Itau) Barcode(d Document) string {
	return "1001.011011.1 123002 2"
}

func (b Itau) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

func (b Itau) Layout(d Document) {
	// TODO
}