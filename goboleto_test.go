package goboleto

import (
	"time"
	"fmt"
)

func init() {
	BilletBB()
}

func BilletBB() {

	// static data
	var bank Bank = BB {
		Account: 8888,
		Agency: 99999,
		Contrato: 12312351,
		Carteira: 15,
		Convenio: 7,
		FormatacaoConvenio: 1,
		FormatacaoNossoNumero: 1,
		VariacaoCarteira: 6,
		Company: &Company{
			Name: "Nome da empresa",
			LegalName: "Razao social",
			Address: "Endereço",
			Contact: "Email e telefone",
			Document: "CNPJ/CPF",
		},
	}

	// dynamic data
	var document = Document{
		Id: 123,
		NossoNumero: 123,
		Value: 999.99,
		ValueTax: 1.00,
		ValueDiscount: 0.00,
		ValueForfeit: 199.99,
		FebrabanType: "md",
		Date: time.Now(),
		DateDue: time.Now().AddDate(0, 0, 4),
		Instructions: [6]string{
			"Não receber após o vencimento",
			"Após vencimento, receber apenas no meu banco",
		},
		Payer: &Payer{
			Name: "Nome completo",
			Address: "Endereço",
			Contact: "Email e telefone",
		},
	}

	fmt.Println(bank.Barcode(document))
	
}