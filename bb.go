package goboleto

// BB - Banco do Brasil
// Source: (http://www.bb.com.br/docs/pub/emp/mpe/espeboletobb.pdf)
type BB struct {
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

// configBB is a global for this bank configs
var configBB = bankConfig{
	Id: 001,
	Aceite: "N",
	Currency: 9,
	CurrencyName: "R$",
}

// Barcode Get the barcode
func (b BB) Barcode(d Document) BarcodeNumber {
	n := BarcodeNumber{
		BankId: configBB.Id,
		CurrencyId: configBB.Currency,
		DateDueFactor:  dateDueFactor(d.DateDue),
		Value: formatValue(d.Value),
	}
	// TODO, bank numbers (nosso numero, de acordo com a carteira e convenio)
	return n
}

// Transference Return the transference file (arquivo de remessa)
func (b BB) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b BB) Layout(d Document) {
	// TODO
}