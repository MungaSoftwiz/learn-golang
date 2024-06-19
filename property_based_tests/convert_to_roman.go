package convert_roman

import "strings"

func ConvertToRoman(arabic int) string {
	// A Builder is used to efficiently build a string using
	// Write methods. It minimizes memory copying.
	var result strings.Builder

	for i := arabic; i > 0; i-- {
		if i == 5 {
			result.WriteString("V")
			break
		}
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
