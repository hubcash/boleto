package goboleto

import (
	"encoding/base64"
)

/*
Bank is defined as an interface,
then we force to implement these functions:

@Barcode Get the barcode, as string, because it may contain dots and spaces
@BarcodeImage Return a image/base64, based in barcode
@Layout the HTML template with all the data

 */
type Bank interface {
	Barcode(Document) string
	BarcodeImage(Document) base64.Encoding
	Layout(Document)
}