package banks

import "goboleto"

/*
@id FEBRABAN bank identifier
@aceite if the payer accept the billet
@currency the currency identifier
@currencyName the currency name
 */
const id = 104
const aceite = "N"
const currency = 9
const currencyName = "R$"

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
}

func (b Caixa) Barcode(d goboleto.Document) string {
	return "1001.011011.1 123002 2"
}