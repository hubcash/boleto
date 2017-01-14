package goboleto

import (
	"time"
	"strings"
	"fmt"
	"strconv"
)

// The size of the value formated
const valueMinSize = 10;

// Defines a document type,
// holds the data of the billet itself
// @Id identifier of your program orders/payments
// @FebrabanType is the document type according FEBRABAN,
// the default used is "DM" (Duplicata mercantil),
// Source: (http://www.bb.com.br/docs/pub/emp/empl/dwn/011DescrCampos.pdf)
// @Value valor do boleto
// @ValueTax taxa do boleto
// @ValueDiscount abatimento/desconto
// @ValueForfeit juros/multa
type Document struct {
	Id 		int
	Date		time.Time
	DateDue		time.Time
	Value 		float64
	ValueTax 	float64
	ValueDiscount	float64
	ValueForfeit	float64
	FebrabanType	string
	Instructions 	[6]string
	Payer		Payer
}

// dateDueFactor use a DateDue type time.Time to return a int,
// with is the quantity of days subsequents from 1997-10-07
func dateDueFactor(dateDue time.Time) int {
	var dateDueFixed = time.Date(1997, 10, 07, 0, 0, 0, 0, time.UTC)
	dif := dateDue.Sub(dateDueFixed);
	return int(dif.Hours()/24);
}

// formatValue format the Document Value,
// in order to replace dots and commas
func formatValue(v float64) int {
	s := fmt.Sprint(v)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	
	value, _ := strconv.Atoi(s)
	return value
}

func module10(n *BarcodeNumber) int {
	// To add pad to numbers: fmt.Sprintf("%010d", int)
	// TODO
	return 1
}

func module11(n *BarcodeNumber) int {
	// To add pad to numbers: fmt.Sprintf("%010d", int)
	// TODO
	return 1
}
