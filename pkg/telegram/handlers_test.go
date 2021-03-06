package telegram

import "testing"

type validationDateTest struct {
	name     string
	date     string
	expected string
}

var validationDateTests = []validationDateTest{
	{"Sanity", "02/01/06", ""},
	{"With error", "02-01-06", "problem parsing date"},
	{"Empty", "", "problem parsing date"},
	{"Future", "02/01/55", "error from future"},
}

func TestValidationDate(t *testing.T) {
	for _, test := range validationDateTests {
		err := validationDate(test.date)
		if err != nil && err.Error() != test.expected {
			t.Errorf("test %s failed. Got: %v, want: %v", test.name, err.Error(), test.expected)
		}
	}
}
