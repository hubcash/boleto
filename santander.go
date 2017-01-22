package goboleto

// Santander
// Source: (https://www.santander.com.br/document/wps/sl-tabela-de-tarifas-cobranca.pdf)
type Santander struct {
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

// configSantander is a global for this bank configs
var configSantander = bankConfig{
	Id: 033,
	Aceite: "N",
	Currency: 9,
	CurrencyName: "R$",
}

// Barcode Get the Barcode, creating a BarcodeNumber
func (b Santander) Barcode(d Document) Barcode {
	
	// TODO, bank numbers (nosso numero, de acordo com a carteira e convenio)
	
	// Create a new Barcode
	var n Barcode = &BarcodeNumber{
		BankId: configSantander.Id,
		CurrencyId: configSantander.Currency,
		DateDueFactor:  dateDueFactor(d.DateDue),
		Value: formatValue(d.Value),
		//BankNumbers: fmt.Sprintf("%0"+strconv.Itoa(bankNumbersSize)+"s", bn),
	}
	n.verification()
	return n
}

// Transference Return the transference file (arquivo de remessa)
func (b Santander) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Santander) Layout(d Document) {
	// TODO
}