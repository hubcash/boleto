package boleto

import (
	"html/template"
	"net/http"
)

// Itau
// Source: (http://download.itau.com.br/bankline/cobranca_cnab240.pdf)
type Itau struct {
	Agency                int
	Account               int
	Convenio              int
	Contrato              int
	Carteira              int
	VariacaoCarteira      int
	FormatacaoConvenio    int
	FormatacaoNossoNumero int
	Company               Company
}

// configItau is a global for this bank configs
var configItau = bankConfig{
	Id:           341,
	Aceite:       "N",
	Currency:     9,
	CurrencyName: "R$",
}

// Barcode Get the Barcode, creating a BarcodeNumber
func (b Itau) Barcode(d Document) Barcode {

	// TODO, bank numbers (nosso numero, de acordo com a carteira e convenio)

	// Create a new Barcode
	var n Barcode = &BarcodeNumber{
		BankId:        configItau.Id,
		CurrencyId:    configItau.Currency,
		DateDueFactor: dateDueFactor(d.DateDue),
		Value:         formatValue(d.Value),
		//BankNumbers: fmt.Sprintf("%0"+strconv.Itoa(bankNumbersSize)+"s", bn),
	}
	n.verification()
	return n
}

// Transference Return the transference file (arquivo de remessa)
func (b Itau) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Itau) Layout(w http.ResponseWriter, d Document) {
	var barcode Barcode = b.Barcode(d)
	layout, _ := template.ParseFiles("templates/itau.html")
	layout.ExecuteTemplate(w, "itau", map[string]interface{}{
		"Barcode": barcode,
		"Document": d,
		"Bank": b,
	})
}
