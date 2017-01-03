package banks

import "goboleto"

/*
@id FEBRABAN bank identifier
@aceite if the payer accept the billet
@currency the currency identifier
@currencyName the currency name
 */
const id = 341
const aceite = "N"
const currency = 9
const currencyName = "R$"

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
}

func (b Itau) Barcode(d goboleto.Document) string {
	return "1001.011011.1 123002 2"
}