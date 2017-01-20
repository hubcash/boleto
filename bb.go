package goboleto

import "strconv"

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

	// Complete the OurNumber digits, by adding convenio rules according to the bank
	convenioSize := len(strconv.Itoa(b.Convenio))
	if b.FormatacaoConvenio == 4 && convenioSize == 4 {
		// For Convenio size 4: CCCCNNNNNNN-X
		// C = Convenio number int(4)
		// N = OurNumber int(7)
		// X = DV, calculated by module11 int(1)
		// TODO, bank numbers according to the format above
		
	} else if b.FormatacaoConvenio == 6 && convenioSize == 6 {
		// For Convenio size 6: CCCCCCNNNNN-X
		// C = Convenio number int(6)
		// N = OurNumber int(5)
		// X = DV, calculated by module11 int(1)
		// TODO, bank numbers according to the format above
		
	} else if b.FormatacaoConvenio == 7 && convenioSize == 7 {
		// For Convenio size 7: CCCCCCCNNNNNNNNNN
		// C = Convenio number int(7)
		// N = OurNumber int(10)
		// TODO, bank numbers according to the format above
		
	} else {
		panic("Invalid Bank.FormatacaoConvenio and Bank.Convenio")
	}
	
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