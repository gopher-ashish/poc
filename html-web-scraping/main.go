package main

import (
	"github.com/PersonalGithubAccount/poc/html-web-scapping/crawl"

	log "github.com/sirupsen/logrus"
)

func main() {
	numbers, err := crawl.CrawlingPage("{WEBSITE_URL_FROM_WERE_YOU_WANT_TO_RETRIVE_NUMBERS}") //example := https://sms24.me/numbers-1
	if err != nil {
		log.Error("Error: ", err)
		return
	}

	log.Infof(" result \n %v ", numbers)
}
