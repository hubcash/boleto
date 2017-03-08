package boleto

import (
	"time"
	"fmt"
)

func init() {
	BilletBB()
}

func BilletBB() {

	// static data, you should keep this configured in somewhere
	var bank Bank = BB{
		Account:            88888888,
		Agency:             4444,
		Carteira:           55,
		Convenio:           4321,
		Company: Company{
			Name:      "Nome da empresa",
			LegalName: "Razao social",
			Address:   "Endereço",
			Contact:   "Email e telefone",
			Document:  "CNPJ/CPF",
		},
	}

	// dynamic data, you should have this data coming from a database
	var document = Document{
		Id:            1111,
		Value:         99999,
		ValueTax:      100,
		ValueDiscount: 0,
		ValueForfeit:  19999,
		OurNumber:     111111,
		FebrabanType:  "md",
		Date:          time.Now(),
		DateDue:       time.Now().AddDate(0, 0, 4),
		Instructions: [6]string{
			"Não receber após o vencimento",
			"Após vencimento, receber apenas no meu banco",
		},
		Payer: Payer{
			Name:    "Nome completo",
			Address: "Endereço",
			Contact: "Email e telefone",
		},
	}

	// Optional, to use in your backend,
	// then you can save the barcode digitable number, or save the image separately
	var barcode Barcode = bank.Barcode(document)
	fmt.Println(barcode.toString())
	fmt.Println(barcode.Digitable())

}
