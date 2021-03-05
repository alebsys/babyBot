package test

import (
	"testing"
)

var tests = []struct {
	input, want string
}{
	{"02/01/06", ""},
	{"02-01-06", "error from validationDate"},
}

func TestValidationDate(t *testing.T) {
	for _, test := range tests {

		err := validationDate(test.input)

		if err.Error() != test.want {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
