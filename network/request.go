package network

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IRequest :
type IRequest interface {
	Execute(Method string, URL string, Headers map[string][]string, Payload string) (int, string, error)
}

// Request :
type Request struct{}

// Execute :
func (r Request) Execute(Method string, URL string, Headers map[string][]string, Payload string) (int, string, error) {
	client := http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	validurl, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	request := &http.Request{
		URL:    validurl,
		Body:   ioutil.NopCloser(strings.NewReader(Payload)),
		Header: Headers,
		Method: Method,
	}
	resp, err := client.Do(request)
	if err != nil {
		return 500, "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 500, "", err
	}
	return resp.StatusCode, string(body), nil
}
