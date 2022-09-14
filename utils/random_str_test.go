package utils

import "testing"

func TestRandStr(t *testing.T) {
	testCases := []struct {
		name        string
		inputLength uint8
	}{
		{
			name:        "it shoud return exact length, when given input length",
			inputLength: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := RandStr(testCase.inputLength)

			if len(got) != int(testCase.inputLength) {
				t.Fatalf("got %d, want %d", len(got), testCase.inputLength)
			}
		})
	}
}
