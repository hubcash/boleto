package banks

import "goboleto"

/*
@id FEBRABAN bank identifier
@aceite if the payer accept the billet
@currency the currency identifier
@currencyName the currency name
 */
const id = 237
const aceite = "N"
const currency = 9
const currencyName = "R$"

/*
Bradesco
Source: (https://banco.bradesco/assets/pessoajuridica/pdf/4008-524-0121-08-layout-cobranca-versao-portuguesSS28785.pdf)
 */
type Bradesco struct {
	Agency 			int
	Account			int
	Convenio		int
	Contrato		int
	Carteira		int
	VariacaoCarteira	int
	FormatacaoConvenio	int
	FormatacaoNossoNumero	int
}

func (b Bradesco) Barcode(d goboleto.Document) string {
	return "1001.011011.1 123002 2"
}