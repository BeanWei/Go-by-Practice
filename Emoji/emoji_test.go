package emoji

import (
	"bytes"
	"testing"
)

const (
	beerKey  = ":beer:"
	beerText = " ビール!!!"
	flag     = ":flag-us:"
	likebeer = ":beer: Beer!!!"
)

var testFText = "test " + emojize(beerKey) + beerText
var testText = emojize(beerKey) + beerText

func TestPrintln(t *testing.T) {
	_, err := Println(beerKey, beerText, likebeer)
	if err != nil {
		t.Error("Println ", err)
	}
}

func TestMultiColons(t *testing.T) {
	var buf bytes.Buffer
	_, err := Fprint(&buf, "A :smile: and another: :smile:")
	if err != nil {
		t.Error("Fprint ", err)
	}

	testCase := "A " + emojize(":smile:") + " and another: " + emojize(":smile:")
	if buf.String() != testCase {
		t.Error("Fprint ", buf.String(), "!=", testCase)
	}
}
