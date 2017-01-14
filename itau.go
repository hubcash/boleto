package goboleto

// Itau
// Source: (http://download.itau.com.br/bankline/cobranca_cnab240.pdf)
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

// configItau is a global for this bank configs
var configItau = bankConfig{
	Id: 341,
	Aceite: "N",
	Currency: 9,
	CurrencyName: "R$",
}

// Barcode Get the barcode
func (b Itau) Barcode(d Document) BarcodeNumber {
	n := BarcodeNumber{
		BankId: configItau.Id,
		CurrencyId: configItau.Currency,
		DateDueFactor: dateDueFactor(d.DateDue),
		Value: formatValue(d.Value),
	}
	// TODO, bank numbers (nosso numero, de acordo com a carteira e convenio)
	return n
}

// Transference Return the transference file (arquivo de remessa)
func (b Itau) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Itau) Layout(d Document) {
	// TODO
}