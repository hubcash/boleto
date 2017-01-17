package goboleto

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

// Barcode is defined as an interface,
// then we force to implement these functions:
// @Image Return a image/base64, using a BarcodeNumber
// @Digitable Get the barcode digitable, it may contain dots and spaces
type Barcode interface {
	Image() 	base64.Encoding
	Digitable() 	string
	Dv()
}

// Defines a barcode number type,
// holds numbers of the barcode
type BarcodeNumber struct {
	// Codigo do banco int(3)
	BankId 		int
	// Codigo da moeda int(1)
	CurrencyId 	int
	// Fator de vencimento int(4)
	DateDueFactor 	int
	// Valor formatado int(10)
	Value 		int
	// Campo livre, numeros do banco com nosso numero int(24)
	BankNumbers	int
	// Digito verificador do codigo de barras int(1)
	dv 		int
}

// Dv returns the BarcodeNumber verification number using module11
func (n *BarcodeNumber) Dv() {
	n.dv = module11(n.toString())
}

// Image return a image/base64, using a BarcodeNumber
func (n BarcodeNumber) Image() base64.Encoding {
	// TODO
	return base64.Encoding{}
}

// Digitable mount the barcode digitable number,
// taking all fields together:
// Field 1: AAABC.CCCCX
// A = FEBRABAN Bank identifier
// B = the currency identifier
// C = 20-24 barcode numbers
// X = DV
//
// Field 2: DDDDD.DDDDDX
// D = 25-34 barcode numbers
// X = DV
//
// Field 3: EEEEE.EEEEEX
// E = 35-44 barcode numbers
// X = DV
//
// Field 4: X
// X = DV
//
// Field 5: UUUUVVVVVVVVVV
// U = Due date factor
// V = Value, as integer
//
// return AAABC.CCCCX DDDDD.DDDDDX EEEEE.EEEEEX X UUUUVVVVVVVVVV
func (n BarcodeNumber) Digitable() string {
	// TODO
	return "AAABC.CCCCX DDDDD.DDDDDX EEEEE.EEEEEX X UUUUVVVVVVVVVV"
}

// toString takes numbers of the barcode, and converts to a string,
// including pad numbers and left zeros
func (n *BarcodeNumber) toString() string {
	var s = fmt.Sprintf("%0"+strconv.Itoa(bankMinSize)+"d", n.BankId)
	s = s + strconv.Itoa(n.CurrencyId)
	s = s + strconv.Itoa(n.DateDueFactor)
	s = s + strconv.Itoa(n.DateDueFactor)
	s = s + fmt.Sprintf("%0"+strconv.Itoa(valueMinSize)+"d", n.Value)
	s = s + strconv.Itoa(n.BankNumbers)
	return s
}