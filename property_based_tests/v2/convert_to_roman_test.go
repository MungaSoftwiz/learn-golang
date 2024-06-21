package convert_roman

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases[:4] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

/* There have been a few rules in the domain of Roman Numerals that we have worked with
 * 1.Can't have more than 3 consecutive symbols
 * 2.Only I (1), X (10) and C (100) can be "subtractors"
 * 3.Taking the result of ConvertToRoman(N) and passing it to ConvertToArabic should return us N
 
 * The tests we have written so far can be described as "example" based tests where we provide
 * examples for the tooling to verify.
 
 * What if we could take these rules that we know about our domain and somehow exercise them
 * against our code?
 *
 * Property based tests help you do this by throwing random data at your code and verifying
 * the rules you describe always hold true.
 * The real challenge about property based tests is having a good understanding of your domain so
 * you can write these properties.
 */

 func TestPropertiesOfConversion(t *testing.T) {
	// int gives a flawed output for arabic values
	// You can't do negative numbers with Roman Numerals
	// Given our rule of a max of 3 consecutive symbols we can't
	// represent a value greater than 3999 (well, kinda) and int
	// has a much higher maximum value than 3999.

	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	// quick.Check will generate random values and pass them to
	// the assertion function. The default checks is 100.
	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}

	// You can change the number of checks by passing a config
	/* if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	*/

 }

 /* resources:
  *		testing/quick: https://pkg.go.dev/testing/quick
  */