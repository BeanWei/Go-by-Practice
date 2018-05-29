package emoji

import (
	"testing"
)

const (
	beerKey  = ":beer:"
	beerText = " ビール!!!"
	flag     = ":flag-us:"
)

var testFText = "test " + emojize(beerKey) + beerText
var testText = emojize(beerKey) + beerText

func TestPrintln(t *testing.T) {
	_, err := Println(beerKey, beerText)
	if err != nil {
		t.Error("Println ", err)
	}
}
