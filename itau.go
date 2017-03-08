package boleto

import (
	"html/template"
	"net/http"
	"fmt"
	"strconv"
)

// Itau
// Source: (http://download.itau.com.br/bankline/cobranca_cnab240.pdf)
type Itau struct {
	Agency                int
	Account               int
	Carteira              int
	ClientCode            int
	Company               Company
}

// configItau is a global for this bank configs
var configItau = bankConfig{
	Id:           341,
	Aceite:       "N",
	Currency:     9,
	CurrencyName: "R$",
	AgencyMaxSize: 4,
	AccountMaxSize: 5,
}

// Barcode Get the Barcode, creating a BarcodeNumber
func (b Itau) Barcode(d Document) Barcode {

	// Complete the BankNumbers digits, by adding carteira rules according to the bank
	var bn string

	// Verify max of Document.OurNumber size
	ourNumberSize := len(strconv.Itoa(d.OurNumber))
	if ourNumberSize > 8 {
		panic("Document.OurNumber max of 8 digits")
	}

	// Verify max of Document.Id size
	IdSize := len(strconv.Itoa(d.Id))
	if IdSize > 7 {
		panic("Document.Id max of 7 digits")
	}

	// Add rules to carteira equals to 107, 122, 142, 143, 196, 198
	if b.Carteira == 107||b.Carteira == 122||b.Carteira == 142||
	   b.Carteira == 143||b.Carteira == 196||b.Carteira == 198 {
		// CCCNNNNNNNNLLLLLLLDDDDDX0
		// C = Carteira number int(3)
		// N = OurNumber int(8)
		// L = Document number int(7)
		// D = Client code int(5)
		// X = DV, calculated by module10 int(1)

		// Verify max Bank.ClientCode size
		clientCodeSize := len(strconv.Itoa(b.ClientCode))
		if clientCodeSize > 5 {
			panic("Bank.ClientCode max of 5 digits")
		}

		// this code var is part of the BankNumbers,
		// we use it to generate another var with module10
		var code string
		code += strconv.Itoa(b.Carteira)
		code += fmt.Sprintf("%08d", d.OurNumber)
		code += fmt.Sprintf("%07d", d.Id)
		code += fmt.Sprintf("%05d", b.ClientCode)

		// module10 with code created above
		codeModule := strconv.Itoa(module10(code, 2))
		bn += code + codeModule + "0"

	} else {
		// CCCNNNNNNNNXAAAATTTTTY000
		// C = Carteira number int(3)
		// N = OurNumber int(8)
		// X = DV, calculated by module10 int(1)
		// A = Agency number int(4)
		// T = Account number int(5)
		// Y = DV, calculated by module10 int(1)

		// Add rules to carteira number equals to 126, 131, 146, 150, 168
		var codeModuleCarteira int
		if b.Carteira == 126||b.Carteira == 131||b.Carteira == 146||
		   b.Carteira == 150||b.Carteira == 168 {
			// module10 with bank agency, account, carteira, and OurNumber
			m := fmt.Sprintf("%0"+strconv.Itoa(configItau.AgencyMaxSize)+"d", b.Agency)
			m += fmt.Sprintf("%0"+strconv.Itoa(configItau.AccountMaxSize)+"d", b.Account)
			m += strconv.Itoa(b.Carteira)
			m += strconv.Itoa(d.OurNumber)
			codeModuleCarteira = module10(m, 2)
		} else {
			// module10 with carteira and OurNumber
			m := strconv.Itoa(b.Carteira)
			m += strconv.Itoa(d.OurNumber)
			codeModuleCarteira = module10(m, 2)
		}

		// module10 with bank agency and account
		codeModuleAccount := module10(strconv.Itoa(b.Account)+strconv.Itoa(b.Agency), 2)

		bn += strconv.Itoa(b.Carteira)
		bn += fmt.Sprintf("%08d", d.OurNumber)
		bn += strconv.Itoa(codeModuleCarteira)
		bn += fmt.Sprintf("%0"+strconv.Itoa(configItau.AgencyMaxSize)+"d", b.Agency)
		bn += fmt.Sprintf("%0"+strconv.Itoa(configItau.AccountMaxSize)+"d", b.Account)
		bn += strconv.Itoa(codeModuleAccount)
		bn += "000"

	}

	// Create a new Barcode
	var n Barcode = &BarcodeNumber{
		BankId:        configItau.Id,
		CurrencyId:    configItau.Currency,
		DateDueFactor: dateDueFactor(d.DateDue),
		Value:         d.Value,
		BankNumbers:   fmt.Sprintf("%0"+strconv.Itoa(bankNumbersSize)+"s", bn),
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
