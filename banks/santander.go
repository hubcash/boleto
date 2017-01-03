package banks

import "goboleto"

/*
@id FEBRABAN bank identifier
@aceite if the payer accept the billet
@currency the currency identifier
@currencyName the currency name
 */
const id = 033
const aceite = "N"
const currency = 9
const currencyName = "R$"

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
}

func (b Santander) Barcode(d goboleto.Document) string {
	return "1001.011011.1 123002 2"
}