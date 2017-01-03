package goboleto

import "time"

/*
Defines a document type, can be also
considered the billet data structure

@Id identifier of your program orders/payments

@FebrabanType is the document type according FEBRABAN
the default used is "DM" (Duplicata mercantil)
Source: (http://www.bb.com.br/docs/pub/emp/empl/dwn/011DescrCampos.pdf)

 */
type Document struct {
	Id 		int
	NossoNumero 	int
	Date		time.Time
	DateDue		time.Time
	Value 		float64
	ValueTax 	float64
	FebrabanType	string
	Instructions 	[6]string
	Payer		*Payer
	Company		*Company
}