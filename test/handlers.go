package test

import (
	"errors"
	"time"
)

func validationDate(s string) error {
	today := time.Now()
	d, err := time.Parse("02/01/06", s)
	if err != nil {
		return errors.New("error from validationDate")
	}
	if today.Before(d) {
		return errors.New("error from validationDate")
	}
	return nil
}
