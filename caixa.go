package goboleto

// CEF - Caixa econômica federal
// Source: (http://www.caixa.gov.br/Downloads/cobranca-caixa/ESP_COD_BARRAS_SIGCB_COBRANCA_CAIXA.pdf)
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

// configCaixa is a global for this bank configs
var configCaixa = bankConfig{
	Id: 104,
	Aceite: "N",
	Currency: 9,
	CurrencyName: "R$",
}

// Barcode Get the barcode
func (b Caixa) Barcode(d Document) BarcodeNumber {
	n := BarcodeNumber{
		BankId: configCaixa.Id,
		CurrencyId: configCaixa.Currency,
		DateDueFactor: dateDueFactor(d.DateDue),
		Value: formatValue(d.Value),
	}
	// TODO, bank numbers (nosso numero, de acordo com a carteira e convenio)
	return n
}

// Transference Return the transference file (arquivo de remessa)
func (b Caixa) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Caixa) Layout(d Document) {
	// TODO
}