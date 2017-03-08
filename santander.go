package boleto

import (
	"html/template"
	"net/http"
	"fmt"
	"strconv"
)

// Santander
// Source: (https://www.santander.com.br/document/wps/sl-tabela-de-tarifas-cobranca.pdf)
type Santander struct {
	Agency                int
	Account               int
	Carteira              int
	IOS                   int
	Company               Company
}

// configSantander is a global for this bank configs
var configSantander = bankConfig{
	Id:           033,
	Aceite:       "N",
	Currency:     9,
	CurrencyName: "R$",
	AccountMaxSize: 7,
}

// Barcode Get the Barcode, creating a BarcodeNumber
func (b Santander) Barcode(d Document) Barcode {

	// Complete the BankNumbers digits
	var bn string
	bn += "9"
	bn += fmt.Sprintf("%0"+strconv.Itoa(configSantander.AccountMaxSize)+"d", b.Account)
	bn += fmt.Sprintf("%013d", d.OurNumber)
	bn += strconv.Itoa(b.IOS)
	bn += strconv.Itoa(b.Carteira)

	// Create a new Barcode
	var n Barcode = &BarcodeNumber{
		BankId:        configSantander.Id,
		CurrencyId:    configSantander.Currency,
		DateDueFactor: dateDueFactor(d.DateDue),
		Value:         d.Value,
		BankNumbers:   fmt.Sprintf("%0"+strconv.Itoa(bankNumbersSize)+"s", bn),
	}
	n.verification()
	return n
}

// Transference Return the transference file (arquivo de remessa)
func (b Santander) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Santander) Layout(w http.ResponseWriter, d Document) {
	var barcode Barcode = b.Barcode(d)
	layout, _ := template.ParseFiles("templates/santander.html")
	layout.ExecuteTemplate(w, "santander", map[string]interface{}{
		"Barcode": barcode,
		"Document": d,
		"Bank": b,
	})
}
