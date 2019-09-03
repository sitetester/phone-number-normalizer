package parser

import (
	"bufio"
	"log"
	"os"
	"phone-number-normalizer/src/entity"
)

type FileParserInterface interface {
	ParseFile() entity.PhoneNumbersList
}

type FileParser struct{}

func (fileParser FileParser) ParseFile() entity.PhoneNumbersList {
	var phoneNumbers []entity.PhoneNumber

	file, err := os.Open("src/importer/resources/import.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	var phoneNumber entity.PhoneNumber
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		phoneNumber.PhoneNumber = scanner.Text()

		phoneNumbers = append(phoneNumbers, phoneNumber)
	}

	var phoneNumbersList entity.PhoneNumbersList
	phoneNumbersList.PhoneNumbers = phoneNumbers

	return phoneNumbersList
}
