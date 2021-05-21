package Pinocchio

import "regexp"

const (
	regexEmailPattern       = `(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`
	regexStrictEmailPattern = `(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+` +
		`(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*` +
		`@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+` +
		`[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`
	regexUrlPattern          = `(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?`
	regexPhonePattern        = `^1([358][0-9]|4[579]|66|7[0135678]|9[89])[0-9]{8}$`
	regexUserIdentityPattern = `^([1-6][1-9]|50)\d{4}(18|19|20)\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`
)

// IsEmail validates string is an email address, if not return false
// basically validation can match 99% cases
func IsEmail(email string) bool {
	regexEmail := regexp.MustCompile(regexEmailPattern)
	return regexEmail.MatchString(email)
}

// IsEmailRFC validates string is an email address, if not return false
// this validation omits RFC 2822
func IsEmailRFC(email string) bool {
	regexStrictEmail := regexp.MustCompile(regexStrictEmailPattern)
	return regexStrictEmail.MatchString(email)
}

// IsUrl validates string is a url link, if not return false
// simple validation can match 99% cases
func IsUrl(url string) bool {
	regexUrl := regexp.MustCompile(regexUrlPattern)
	return regexUrl.MatchString(url)
}

// IsPhone validates string is a chinese phone number, if not return false
// simple validation can match 99% cases
func IsPhone(phone string) bool {
	regexPhone := regexp.MustCompile(regexPhonePattern)
	return regexPhone.MatchString(phone)
}

// IsIdentityCard validates string is a chinese user ID card, if not return false
// simple validation can match 99% cases
func IsIdentityCard(card string) bool {
	regexIdentityCard := regexp.MustCompile(regexUserIdentityPattern)
	return regexIdentityCard.MatchString(card)
}
