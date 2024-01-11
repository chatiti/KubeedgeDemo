package counter

import (
	"fmt"
	"math/rand"
	"time"
)

func generateLicensePlate(number int) string {
	rand.Seed(time.Now().UnixNano())

	var licensePlates []string
	for i := 0; i < number; i++ {
		licensePlate := generateRandomLicensePlate()
		licensePlates = append(licensePlates, licensePlate)
	}

	return joinWithSemicolon(licensePlates)
}

func generateRandomLicensePlate() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const numbers = "0123456789"

	licensePlate := ""
	for i := 0; i < 6; i++ {
		if i < 3 {
			licensePlate += string(letters[rand.Intn(len(letters))])
		} else {
			licensePlate += string(numbers[rand.Intn(len(numbers))])
		}
	}

	return licensePlate
}

func joinWithSemicolon(licensePlates []string) string {
	return fmt.Sprintf("%s", joinStrings(licensePlates, ";"))
}

func joinStrings(strList []string, separator string) string {
	returnStr := ""
	for i, str := range strList {
		if i > 0 {
			returnStr += separator
		}
		returnStr += str
	}
	return returnStr
}
