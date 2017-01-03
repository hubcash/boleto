package goboleto

import (
	"encoding/base64"
	"html/template"
)

/*
Billet is defined as an interface,
then Billet has forced to implement these functions:

@Barcode Get the barcode, as string, because it may contain dots and spaces
@BarcodeImage Return a image/base64, based in barcode
@Layout the HTML template generate with all the data

 */
type Billet interface {
	Barcode(Document) string
	BarcodeImage(Document) base64.Encoding
	Layout(Document) template.HTML
}