package crawl

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	rgx "github.com/PersonalGithubAccount/poc/html-web-scraping/regex_operation"
)

func CrawlingPage(url string) ([]string, error) {
	resp, err := getAndCheckResponse(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	numbers, err := validateResponseAndRetriveNumbers(resp)
	if err != nil {
		return nil, err
	}

	return numbers, nil
}

func getAndCheckResponse(url string) (*http.Response, error) {
	resp, err := getResponseFromUrl(url)
	if err != nil {
		return nil, err
	}

	if err := CheckResponseCode(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func validateResponseAndRetriveNumbers(resp *http.Response) ([]string, error) {
	pageContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	numbers := rgx.GetPhoneNumbers(pageContent)
	if numbers == nil || len(numbers) < 1 {
		return nil, errors.New("no numbers found")
	}

	return numbers, nil
}

func getResponseFromUrl(url string) (*http.Response, error) {
	clinet := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := clinet.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func CheckResponseCode(res *http.Response) error {
	if res.StatusCode != http.StatusOK {
		err := errors.New("couldn't get the response")
		return err
	}
	return nil
}
