package main

import (
	"strconv"
)

type Card struct {
	CardNumber string `json:"CardNumber"`
}



func ValidateCard(c Card) bool{

	sum := 0
	isSecond := false
	length := len(c.CardNumber)


	for i := length - 1 ; i>= 0; i--  {
		digit, _ := strconv.Atoi(string(c.CardNumber[i]))

		if isSecond {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }

		sum += digit
        isSecond = !isSecond

	}

	return sum % 10 == 0

}


