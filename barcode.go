package boleto

import (
	"fmt"
	barcodeImage "github.com/boombuler/barcode"
	"github.com/boombuler/barcode/twooffive"
	"image"
	"strconv"
)

const (
	// The min size of a bankId
	bankMinSize = 3
	// The min size of the value formated
	valueMinSize = 10
	// The min size of a barcode
	barcodeNumberMinSize = 19
	// The max size of a barcode
	barcodeNumberMaxSize = 44
	// BarcodeNumber.BankNumbers size
	bankNumbersSize = 25
)

// Barcode is defined as an interface,
// then we force to implement these functions:
// @Image Return a image, using a BarcodeNumber
// @Digitable Get the barcode digitable, it may contain dots and spaces
type Barcode interface {
	Image() image.Image
	Digitable() string
	toString() string
	verification()
}

// Defines a barcode number type,
// holds numbers of the barcode
type BarcodeNumber struct {
	// Codigo do banco int(3)
	BankId int
	// Codigo da moeda int(1)
	CurrencyId int
	// Fator de vencimento int(4)
	DateDueFactor int
	// Valor formatado int(10)
	Value int
	// Campo livre, numeros do banco com nosso numero string(25)
	BankNumbers string
	// Digito verificador do codigo de barras int(1)
	Dv int
}

// Image return a image.Image, using a BarcodeNumber
func (n BarcodeNumber) Image() image.Image {
	e, _ := twooffive.Encode(n.toString(), true)
	img, _ := barcodeImage.Scale(e, 410, 50)
	return img
}

// Digitable mount the barcode digitable number,
// taking all BarcodeNumber fields together:
// Field 1: AAABC.CCCCX
// A = FEBRABAN Bank identifier
// B = the currency identifier
// C = 20-24 barcode numbers
// X = DV, using module10
//
// Field 2: DDDDD.DDDDDX
// D = 25-34 barcode numbers
// X = DV, using module10
//
// Field 3: EEEEE.EEEEEX
// E = 35-44 barcode numbers
// X = DV, using module10
//
// Field 4: X
// X = DV, BarcodeNumber.Dv
//
// Field 5: UUUUVVVVVVVVVV
// U = Due date factor
// V = Value
//
// return AAABC.CCCCX DDDDD.DDDDDX EEEEE.EEEEEX X UUUUVVVVVVVVVV
func (n BarcodeNumber) Digitable() string {
	s := n.toString()

	// Field 1
	var f1 = fmt.Sprintf("%0"+strconv.Itoa(bankMinSize)+"d", n.BankId)
	f1 += strconv.Itoa(n.CurrencyId)
	f1 += string(s[19]) + "." + s[20:24]
	f1 += strconv.Itoa(module10(f1, maxModule10))

	// Field 2
	var f2 = s[24:29] + "." + s[29:34]
	f2 += strconv.Itoa(module10(f2, minModule10))

	// Field 3
	var f3 = s[34:39] + "." + s[39:44]
	f3 += strconv.Itoa(module10(f3, minModule10))

	// Field 5
	var f4 = strconv.Itoa(n.Dv)

	// Field 5
	var f5 = strconv.Itoa(n.DateDueFactor)
	f5 += fmt.Sprintf("%0"+strconv.Itoa(valueMinSize)+"d", n.Value)

	// All fields together
	d := f1 + " " + f2 + " " + f3 + " " + f4 + " " + f5
	return d
}

// verification returns the BarcodeNumber verification number using module11
func (n *BarcodeNumber) verification() {
	s := n.toString()
	n.Dv = module11(s)
}

// toString takes BarcodeNumber, and converts to a string,
// including pad numbers and left zeros
func (n *BarcodeNumber) toString() string {
	var s = fmt.Sprintf("%0"+strconv.Itoa(bankMinSize)+"d", n.BankId)
	s += strconv.Itoa(n.CurrencyId)
	s += strconv.Itoa(n.Dv)
	s += strconv.Itoa(n.DateDueFactor)
	s += fmt.Sprintf("%0"+strconv.Itoa(valueMinSize)+"d", n.Value)
	s += n.BankNumbers

	if len(s) < barcodeNumberMinSize {
		panic("There are missing values in Bank and Document structures")
	}
	if len(s) > barcodeNumberMaxSize {
		panic("There are remaining values in Bank and Document structures")
	}
	return s
}
