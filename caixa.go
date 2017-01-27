package goboleto

import (
	"html/template"
	"net/http"
)

// CEF - Caixa econômica federal
// Source: (http://www.caixa.gov.br/Downloads/cobranca-caixa/ESP_COD_BARRAS_SIGCB_COBRANCA_CAIXA.pdf)
type Caixa struct {
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

// configCaixa is a global for this bank configs
var configCaixa = bankConfig{
	Id:           104,
	Aceite:       "N",
	Currency:     9,
	CurrencyName: "R$",
}

// Barcode Get the Barcode, creating a BarcodeNumber
func (b Caixa) Barcode(d Document) Barcode {

	// TODO, bank numbers (nosso numero, de acordo com a carteira e convenio)

	// Create a new Barcode
	var n Barcode = &BarcodeNumber{
		BankId:        configCaixa.Id,
		CurrencyId:    configCaixa.Currency,
		DateDueFactor: dateDueFactor(d.DateDue),
		Value:         formatValue(d.Value),
		//BankNumbers: fmt.Sprintf("%0"+strconv.Itoa(bankNumbersSize)+"s", bn),
	}
	n.verification()
	return n
}

// Transference Return the transference file (arquivo de remessa)
func (b Caixa) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Caixa) Layout(w http.ResponseWriter, d Document) {
	var barcode Barcode = b.Barcode(d)
	layout, _ := template.ParseFiles("templates/caixa.html")
	layout.ExecuteTemplate(w, "caixa", barcode)
}
