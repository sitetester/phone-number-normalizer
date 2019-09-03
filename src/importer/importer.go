package importer

import (
	"github.com/jinzhu/gorm"
	"phone-number-normalizer/src/parser"
)

type Importer struct {
}

func (importer *Importer) ImportSampleData(db *gorm.DB, fileParser parser.FileParserInterface) {
	var phoneNumbersList = fileParser.ParseFile()

	for _, phoneNumber := range phoneNumbersList.PhoneNumbers {
		db.Create(&phoneNumber)
	}
}
