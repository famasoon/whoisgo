package function

import (
	"encoding/json"

	"github.com/undiabler/golang-whois"
)

type WhoisResult struct {
	Domain string `json:domain`
	Result string `json:result`
}

// Handle a serverless request
func Handle(req []byte) string {
	domain := string(req)
	result, err := whois.GetWhois(domain)
	if err != nil {
		panic(err)
	}

	whoisResult := WhoisResult{
		Domain: domain,
		Result: result,
	}

	s, _ := json.Marshal(whoisResult)
	return string(s)
}
