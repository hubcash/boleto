package goboleto

const barcodeLength = 44

// Barcode = Linha digitavel
// Barcode is an alias for "Linha digitavel"

/*
Defines a barcode type,
to organize the fields data
 */
type Barcode struct {
	Field1		Field1
	Field2		Field2
	Field3		Field3
	Field4		Field4
	Field5		Field5
}

/*
Field 1: AAABC.CCCCX
A = FEBRABAN Bank identifier
B = the currency identifier
C = 20-24 numbers
X = DV
 */
type Field1 struct {
	bank		int
	currency	int
	numbers		int
	dv		int
}

/*
Field 2: DDDDD.DDDDDX
D = 25-34 numbers
X = DV
 */
type Field2 struct {
	numbers		int
	dv		int
}

/*
Field 3: EEEEE.EEEEEX
E = 35-44 numbers
X = DV
 */
type Field3 struct {
	numbers		int
	dv		int
}

/*
Field 4: X
X = DV
 */
type Field4 struct {
	dv		int
}

/*
Field 5: UUUUVVVVVVVVVV
U = Due date factor
V = Value, two decimal
 */
type Field5 struct {
	dueDate		int
	value		int
}

/*
@generateBarcode mount the barcode numbers

Barcode format (all fields together):
AAABC.CCCCX DDDDD.DDDDDX EEEEE.EEEEEX X UUUUVVVVVVVVVV
 */
func generateBarcode(b Barcode) string {
	return "1001.011011.1 123002 2"
}


// TODO digitos verificadores
// TODO modulo10()