package goboleto

// Bradesco
// Source: (https://banco.bradesco/assets/pessoajuridica/pdf/4008-524-0121-08-layout-cobranca-versao-portuguesSS28785.pdf)
type Bradesco struct {
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

// configBradesco is a global for this bank configs
var configBradesco = bankConfig{
	Id: 237,
	Aceite: "N",
	Currency: 9,
	CurrencyName: "R$",
}

// Barcode Get the barcode
func (b Bradesco) Barcode(d Document) BarcodeNumber {
	n := BarcodeNumber{
		BankId: configBradesco.Id,
		CurrencyId: configBradesco.Currency,
		DateDueFactor: dateDueFactor(d.DateDue),
		Value: formatValue(d.Value),
	}
	// TODO, bank numbers (nosso numero, de acordo com a carteira e convenio)
	return n
}

// Transference Return the transference file (arquivo de remessa)
func (b Bradesco) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Bradesco) Layout(d Document) {
	// TODO
}