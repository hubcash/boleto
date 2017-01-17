package goboleto

import (
	"time"
	"strings"
	"fmt"
	"strconv"
)

const (
	// The min size of a bankId
	bankMinSize = 3
	
	// The min size of the value formated
	valueMinSize = 10
	
	// Max multiplier for module 11
	maxModule11 = 9
	
	// min multiplier for module 11
	minModule11 = 2
)

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
	// TODO
	return 1
}

// module11 takes a number and returns his verifier digit (spect an string
// because it may contain left zeros and pad numbers)
// Each digit that makes up our number is multiplied by his multiplier weight,
// the multipliers range from 9 to 2, from right to left
// Multiplication results are summed and divided by eleven
func module11(s string) int {
	// Create a slice with the numbers
	numbers := make([]int, len(s))
	for i, r := range s {
		c := string(r)
		n, _ := strconv.Atoi(c)
		numbers[i] = n
	}
	numbersLen := len(numbers)
	
	// initial multiplier weight
	var p = maxModule11
	if numbersLen > 11 {
		p = minModule11
	}

	// Inverse the numbers creating for loop
	// Multiply all numbers using multiplier weight
	total := 0
	for i := len(numbers)-1; i >= 0; i-- {
		n := numbers[i]
		total += (n*p)
		
		// If the numbers length is higher than 11,
		// we need to inverse the min and max
		if numbersLen > 11 {
			p++
			// if the multiplier weight is higher then max
			if p > maxModule11 {
				p = minModule11
			}
			continue
		}
		
		p--
		// if the multiplier weight is lower then minimal
		if p < minModule11 {
			p = maxModule11
		}
		
	}
	
	// If the numbers length is higher than 11,
	// we need to divide also by 11
	if numbersLen > 11 {
		dv := total % 11
		dv = 11 - dv
		// If the verifier digit is equal 0, 10, 11,
		// need to be always 1
		if dv == 0 || dv == 10 || dv == 11 {
			dv = 1
		}
		return dv
	}
	
	// End by dividing
	return total % 11
}
