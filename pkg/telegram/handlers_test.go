package telegram

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_validationDate(t *testing.T) {
	var testCases = []struct {
		input string
		want  error
	}{
		{"02/01/06", nil},
		{"02-01-06", errors.New("problem parsing date")},
		{"02/13/06", errors.New("problem parsing date")},
		{"02/08/22", errors.New("the date entered cannot be later than today")},
	}

	for _, tc := range testCases {
		fmt.Println(tc.want)
		err := validationDate(tc.input)
		if !reflect.DeepEqual(err, tc.want) {
			t.Errorf("\nunexpected error: %v\nexpected error: %v", err, tc.want)
		}
	}
}
