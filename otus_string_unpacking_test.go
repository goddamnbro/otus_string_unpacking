package otus_string_unpacking

import "testing"

var samples = []map[string]string{
	{
		"input":    "a4bc2d5e",
		"expected": "aaaabccddddde",
	},
	{
		"input":    "abcd",
		"expected": "abcd",
	},
	{
		"input":    "45",
		"expected": "",
	},
	{
		"input":    "45",
		"expected": "",
	},
}

var advancedSamples = []map[string]string{
	{
		"input":    "qwe\4\5",
		"expected": "qwe45",
	},
	{
		"input":    "qwe\45",
		"expected": "qwe44444",
	},
	{
		"input":    "qwe\\5",
		"expected": `qwe\\\\\`,
	},
}

func TestSimpleUnpack(t *testing.T) {
	for _, sample := range samples {
		actual := UnpackString(sample["input"])
		expected := sample["expected"]

		if actual != expected {
			t.Errorf("Unpacked string is incorrect!\n "+
				"got:      %s \n "+
				"expected: %s",
				actual, expected)
		}
	}
}

