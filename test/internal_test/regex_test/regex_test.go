package regex_test

import (
	"steam-currency-parser/internal/regex"
	"testing"
)

func TestOnlyInt(t *testing.T) {
	testTable := []struct {
		text     string
		expected string
	}{
		{
			text:     "",
			expected: "",
		},
		{
			text:     "CHF 903.81",
			expected: "90381",
		},
		{
			text:     "24.600.664,20₫",
			expected: "2460066420",
		},
		{
			text:     "Market Price: 472 901,77₸",
			expected: "47290177",
		},
	}

	for _, testCase := range testTable {
		result := regex.OnlyInt(testCase.text)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %s, got %s",
				testCase.expected, result)
		}

	}
}
