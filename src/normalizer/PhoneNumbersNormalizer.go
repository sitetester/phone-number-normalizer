package normalizer

import (
	"github.com/jinzhu/gorm"
	"log"
	"phone-number-normalizer/src/entity"
	"regexp"
)

type PhoneNumbersNormalizer struct {
}

func (phoneNumbersNormalizer PhoneNumbersNormalizer) Normalize(db *gorm.DB) entity.NormalizedPhoneNumbersList {
	var result []entity.PhoneNumber
	db.Table("phone_numbers").Scan(&result)

	var normalizedPhoneNumbersList entity.NormalizedPhoneNumbersList
	var normalizedPhoneNumber entity.NormalizedPhoneNumber

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	for _, phoneNumber := range result {
		tmpPhoneNumber := reg.ReplaceAllString(phoneNumber.PhoneNumber, "")

		if !stringInNormalizedPhoneNumbersList(tmpPhoneNumber, normalizedPhoneNumbersList) {
			normalizedPhoneNumber.PhoneNumber = tmpPhoneNumber
			normalizedPhoneNumbersList.NormalizedPhoneNumbers = append(normalizedPhoneNumbersList.NormalizedPhoneNumbers, normalizedPhoneNumber)
		}
	}

	return normalizedPhoneNumbersList
}

func stringInNormalizedPhoneNumbersList(str string, normalizedPhoneNumbersList entity.NormalizedPhoneNumbersList) bool {
	for _, ph := range normalizedPhoneNumbersList.NormalizedPhoneNumbers {
		if ph.PhoneNumber == str {
			return true
		}
	}
	return false
}
