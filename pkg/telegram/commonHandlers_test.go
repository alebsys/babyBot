package telegram

import (
	"testing"
)

// TODO слайд 21. Переписать тесты на библиотеку `testify`, выделить ошибочные тесты в отдельную функцию
// добавить sub test через t.Run
func TestValidationDate(t *testing.T) {
	var testCases = []struct {
		caseName string
		date     string
		expected string
	}{
		{"Sanity", "02/01/06", ""},
		{"With error", "02-01-06", "problem parsing date"},
		{"Empty", "", "problem parsing date"},
		{"Future", "02/01/55", "error from future"},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run("check "+tc.caseName, func(t *testing.T) {
			err := validationDate(tc.date)
			if err != nil && err.Error() != tc.expected {
				t.Fatalf("test %s failed. Got: %v, want: %v", tc.caseName, err.Error(), tc.expected)
			}
		})
	}
}

func TestValidationWeigth(t *testing.T) {
	var testCases = []struct {
		caseName string
		value    string
		expected string
	}{
		{"Sanity", "80", ""},
		{"Sanity", "65.4", ""},
		{"With error", "60,3", "problem parsing weight value"},
		{"Empty", "", "problem parsing weight value"},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run("check "+tc.caseName, func(t *testing.T) {
			err := validationWeight(tc.value)
			if err != nil && err.Error() != tc.expected {
				t.Fatalf("test %s failed. Got: %v, want: %v", tc.caseName, err.Error(), tc.expected)
			}
		})
	}
}
