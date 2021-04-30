package rgx

import (
	"regexp"

	log "github.com/sirupsen/logrus"
)

const RegexForPhoneNumbers = `\+[0-9]{5,15}`

var Regex = regexp.MustCompile(RegexForPhoneNumbers)

func GetPhoneNumbers(pageContent []byte) []string {
	numbers := Regex.FindAllString(string(pageContent), -1)
	if numbers == nil {
		log.Info("No numbers found")
		return nil
	}

	return numbers
}
