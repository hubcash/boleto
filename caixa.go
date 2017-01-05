package goboleto

import "encoding/base64"

/*
Bank constants
@ID FEBRABAN bank identifier
@Aceite if the payer accept the billet
@Currency the currency identifier
@CurrencyName the currency name
 */
var configCaixa = map[string]interface{}{
	"ID": 104,
	"Aceite": "N",
	"Currency": 9,
	"CurrencyName": "R$",
}

/*
CEF - Caixa econômica federal
Source: (http://www.caixa.gov.br/Downloads/cobranca-caixa/ESP_COD_BARRAS_SIGCB_COBRANCA_CAIXA.pdf)
 */
type Caixa struct {
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

func (b Caixa) Barcode(d Document) string {
	return "1001.011011.1 123002 2"
}

func (b Caixa) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

func (b Caixa) Layout(d Document) {
	// TODO
}