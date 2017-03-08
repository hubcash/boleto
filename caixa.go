package boleto

import (
	"html/template"
	"net/http"
	"fmt"
	"strconv"
)

// CEF - Caixa econômica federal - Modelo SIGCB
// Source: (http://www.caixa.gov.br/Downloads/cobranca-caixa/ESP_COD_BARRAS_SIGCB_COBRANCA_CAIXA.pdf)
type Caixa struct {
	Agency                int
	Account               int
	Carteira              string
	Company               Company
}

// configCaixa is a global for this bank configs
var configCaixa = bankConfig{
	Id:           104,
	Aceite:       "N",
	Currency:     9,
	CurrencyName: "R$",
	AccountMaxSize: 6,
}

// Barcode Get the Barcode, creating a BarcodeNumber
func (b Caixa) Barcode(d Document) Barcode {

	// Verify if carteira is supported
	if b.Carteira != "RG" {
		panic("Bank.Carteira only supported is 'RG'")
	}

	// Complete the BankNumbers digits, by adding carteira rules according to the bank
	var bn string

	// Verify max of Document.OurNumber size
	ourNumberSize := len(strconv.Itoa(d.OurNumber))
	if ourNumberSize > 15 {
		panic("Document.OurNumber max of 15 digits")
	}

	// codeModule var defines the Document.OurNumber DV, using '14' as prefix,
	// means that, 1 is the identificator of registered billets and 4 is because the billet
	// was issued from the client/company
	codeModule := module11("14"+strconv.Itoa(d.OurNumber))

	// this code var is part of the BankNumbers,
	// we use it to generate another var with module10
	var code string
	code += fmt.Sprintf("%015d", d.OurNumber)
	code += strconv.Itoa(codeModule)

	// module11 with bank account
	codeModuleAccount := module11(strconv.Itoa(b.Account))

	// CCCCCCXHHHHJJJJPPPPPPPPPY
	// C = Bank account number int(6)
	// X = DV of Bank account number int(1)
	// H = 3-5 position of code and const int(4)
	// J = 6-8 position of code and const int(4)
	// P = 9-17 position int(9)
	// Y = module11 using previous digits int(1)
	bn += fmt.Sprintf("%0"+strconv.Itoa(configCaixa.AccountMaxSize)+"d", b.Account)
	bn += strconv.Itoa(codeModuleAccount)
	// 3-5 position of code, and add "1" as identificator of registered billets
	bn += code[2:5] + "1"
	// 6-8 position of code, and add "4" because issued from the client/company
	bn += code[5:8] + "4"
	// 9-17 position of code
	bn += code[8:17]

	// module11 with current bn and add to itself converting to string
	codeModuleBn := module11(bn)
	bn += strconv.Itoa(codeModuleBn)

	// Create a new Barcode
	var n Barcode = &BarcodeNumber{
		BankId:        configCaixa.Id,
		CurrencyId:    configCaixa.Currency,
		DateDueFactor: dateDueFactor(d.DateDue),
		Value:         d.Value,
		BankNumbers: fmt.Sprintf("%0"+strconv.Itoa(bankNumbersSize)+"s", bn),
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
	layout.ExecuteTemplate(w, "caixa", map[string]interface{}{
		"Barcode": barcode,
		"Document": d,
		"Bank": b,
	})
}
