package environment

import (
	"errors"
	"testing"
)

func TestConvertYesOrNoToBool(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output bool
		error  error
	}{
		{name: "supports yes as a value", input: "yes", output: true, error: nil},
		{name: "supports no as a value", input: "no", output: false, error: nil},
		{name: "returns an error otherwise", input: "whatever", output: false, error: errors.New("Provided string needs to be 'yes' or 'no'")},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, error := convertYesOrNoToBool(testCase.input)

			if error != nil && error.Error() != testCase.error.Error() {
				t.Errorf("error got: %v, want: %v", error, testCase.error)
			}

			if result != testCase.output {
				t.Errorf("result got: %v, want: %v", result, testCase.output)
			}
		})
	}
}
