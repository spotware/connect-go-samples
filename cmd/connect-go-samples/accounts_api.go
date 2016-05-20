package main

import (
	"net/http"
	"net/url"
	"io"
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://sandbox-api.spotware.com/connect/"
	userAgent = "connect-go-samples/" + libraryVersion
	mediaType = "application/json"
)

type AccountsAPI struct {
	// HTTP client used to communicate with the API.
	client *http.Client
	// Base URL for Accounts API requests.  Defaults to the Sandbox Open API URL
	BaseURL *url.URL
	// User agent used when communicating with the API.
	UserAgent string
}


type Message struct {
	Data  interface{}  `json:"data,omitempty"`
	Error *Error       `json:"error,omitempty"`
}

type Error struct {
	ErrorCode    *string  `json:"errorCode,omitempty"`
	Description  *string  `json:"description,omitempty"`
}

func (e *Error) Error() string { return *e.ErrorCode + " " + *e.Description }

func NewAccountsAPI(httpClient *http.Client) *AccountsAPI {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &AccountsAPI{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *AccountsAPI) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaType)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response.
func (c *AccountsAPI) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}
	return resp, err
}



// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
