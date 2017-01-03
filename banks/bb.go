package banks

import "goboleto"

/*
@id FEBRABAN bank identifier
@aceite if the payer accept the billet
@currency the currency identifier
@currencyName the currency name
 */
const id = 001
const aceite = "N"
const currency = 9
const currencyName = "R$"

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
}

func (b BB) Barcode(d goboleto.Document) string {
	return "1001.011011.1 123002 2"
}