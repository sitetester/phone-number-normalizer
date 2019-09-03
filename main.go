package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"phone-number-normalizer/src/entity"
	"phone-number-normalizer/src/importer"
	"phone-number-normalizer/src/normalizer"
	"phone-number-normalizer/src/parser"
)

func main()  {
	db, err := gorm.Open("sqlite3", "db/phoneNumbers.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// delete if already exists
	db.Delete(&entity.PhoneNumber{})

	// load schema in DB
	db.AutoMigrate(&entity.PhoneNumber{})

	// import sample data in DB
	var importer importer.Importer

	var fileParser parser.FileParser
	importer.ImportSampleData(db, fileParser)

	var n normalizer.PhoneNumbersNormalizer
	normalizedPhoneNumbers := n.Normalize(db)

	fmt.Println("normalizedPhoneNumbers: ")
	fmt.Println(normalizedPhoneNumbers)
}

