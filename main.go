package main

import (
	"unicode"
)

func isValidPassword(password string) bool {

	if len(password) < 5 || len(password) > 12 {
		return false

	}

	hasDigit := false
	hasCapital := false
	for _, c := range password {
		if unicode.IsDigit(c) {
			hasDigit = true
		}

		if unicode.IsUpper(c){
			hasCapital = true
		}

		if hasCapital && hasDigit{
			return true
		}

	}

	return false

}

