package boleto

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// BB - Banco do Brasil
// Source: (http://www.bb.com.br/docs/pub/emp/mpe/espeboletobb.pdf)
type BB struct {
	Agency                int
	Account               int
	Convenio              int
	Carteira              int
	Company               Company
}

// configBB is a global for this bank configs
var configBB = bankConfig{
	Id:             001,
	Aceite:         "N",
	Currency:       9,
	CurrencyName:   "R$",
	AgencyMaxSize:  4,
	AccountMaxSize: 8,
}

// Barcode Get the Barcode, creating a BarcodeNumber
func (b BB) Barcode(d Document) Barcode {

	// Complete the BankNumbers digits, by adding convenio rules according to the bank
	var bn string
	convenioSize := len(strconv.Itoa(b.Convenio))
	ourNumberSize := len(strconv.Itoa(d.OurNumber))
	if convenioSize == 4 {
		// For Convenio size 4: CCCCNNNNNNN-X
		// C = Convenio number int(4)
		// N = OurNumber int(7)
		// X = DV, calculated by module11 int(1)
		if ourNumberSize > 7 {
			panic("Document.OurNumber max of 7 digits")
		}

		bn += strconv.Itoa(b.Convenio)
		bn += fmt.Sprintf("%07d", d.OurNumber)
		bn += fmt.Sprintf("%0"+strconv.Itoa(configBB.AgencyMaxSize)+"d", b.Agency)
		bn += fmt.Sprintf("%0"+strconv.Itoa(configBB.AccountMaxSize)+"d", b.Account)
		bn += strconv.Itoa(b.Carteira)

	} else if convenioSize == 6 {
		// For Convenio size 6: CCCCCCNNNNN-X
		// C = Convenio number int(6)
		// N = OurNumber int(5)
		// X = DV, calculated by module11 int(1)
		if ourNumberSize > 5 {
			panic("Document.OurNumber max of 5 digits")
		}

		bn += strconv.Itoa(b.Convenio)
		bn += fmt.Sprintf("%05d", d.OurNumber)
		bn += fmt.Sprintf("%0"+strconv.Itoa(configBB.AgencyMaxSize)+"d", b.Agency)
		bn += fmt.Sprintf("%0"+strconv.Itoa(configBB.AccountMaxSize)+"d", b.Account)
		bn += strconv.Itoa(b.Carteira)

	} else if convenioSize == 7 {
		// For Convenio size 7: CCCCCCCNNNNNNNNNN
		// C = Convenio number int(7)
		// N = OurNumber int(9)
		if ourNumberSize > 9 {
			panic("Document.OurNumber max of 9 digits")
		}

		bn += fmt.Sprintf("%013d", b.Convenio)
		bn += fmt.Sprintf("%09d", d.OurNumber)
		bn += strconv.Itoa(b.Carteira)

	} else {
		panic("Invalid Bank.Convenio")
	}

	// Create a new Barcode
	var n Barcode = &BarcodeNumber{
		BankId:        configBB.Id,
		CurrencyId:    configBB.Currency,
		DateDueFactor: dateDueFactor(d.DateDue),
		Value:         d.Value,
		BankNumbers:   fmt.Sprintf("%0"+strconv.Itoa(bankNumbersSize)+"s", bn),
	}
	n.verification()
	return n
}

// Transference Return the transference file (arquivo de remessa)
func (b BB) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b BB) Layout(w http.ResponseWriter, d Document) {
	var barcode Barcode = b.Barcode(d)
	layout, _ := template.ParseFiles("templates/bb.html")
	layout.ExecuteTemplate(w, "bb", map[string]interface{}{
		"Barcode": barcode,
		"Document": d,
		"Bank": b,
	})
}
