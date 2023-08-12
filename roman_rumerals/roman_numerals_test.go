package roman_rumerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var tests = []struct {
	num   uint16
	roman string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestToRoman(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d to %q", tt.num, tt.roman), func(t *testing.T) {
			if got := ToRoman(tt.num); got != tt.roman {
				t.Errorf("ToRoman() = %v, roman %v", got, tt.roman)
			}
		})
	}
}

func TestFromRoman(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q to %d", tt.roman, tt.num), func(t *testing.T) {
			if got := FromRoman(tt.roman); got != tt.num {
				t.Errorf("FromRoman() = %v, num %v", got, tt.num)
			}
		})
	}
}
func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(num uint16) bool {
		if num > 3999 {
			return true
		}
		roman := ToRoman(num)
		fromRoman := FromRoman(roman)
		return fromRoman == num
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
