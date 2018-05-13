package gohubspot

import (
	"fmt"
	"net/http"
	"net/url"
)

type APIKeyAuth struct {
	apiKey string
}

const (
	apiKeyParam    = "hapikey=%s"
	singleParam    = "/?"
	multipleParams = "&"
)

// NewAPIKeyAuth create new API KEY Authenticator
func NewAPIKeyAuth(apikey string) APIKeyAuth {
	return APIKeyAuth{apiKey: apikey}
}

// Authenticate set auth
func (auth APIKeyAuth) Authenticate(request *http.Request) error {
	base_url, err := url.Parse(request.URL.String())
	if err != nil {
		return err
	}

	// Add the API key parameter but be mindful of params that already exist on the url
	appender := singleParam
	if base_url.RawQuery != "" {
		appender = multipleParams
	}
	urlStr := request.URL.String() + appender + fmt.Sprintf(apiKeyParam, auth.apiKey)
	url, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	request.URL = url

	return nil
}
