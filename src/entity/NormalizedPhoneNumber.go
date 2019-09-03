package entity

type NormalizedPhoneNumber struct {
	PhoneNumber string
}

type NormalizedPhoneNumbersList struct {
	NormalizedPhoneNumbers[] NormalizedPhoneNumber
}
