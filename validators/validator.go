package validators

import (
	"strconv"
)

type CardNumber struct{ 
	Number string

}

func ( c *CardNumber) validateCard() bool{

	sum := 0
	isSecond := false
	length := len(c.Number)
	

	for i := length - 1 ; i>= 0; i--  {
		digit, _ := strconv.Atoi(string(c.Number[i]))

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