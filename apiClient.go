package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//APIClient represents HTTP API client
//APIClient reuse http client connections
type APIClient struct {
	client *http.Client

	//http headers for every request
	//By default contains only Content-Type : application/json
	headers map[string]string

	//func for detecting success response code status
	// if status success, client try to deserialize response body
	// if status not success, client return error with response status code
	//by default success status code only 200.
	isSuccessStatus func(statusCode int) bool
}

// NewAPIClient create new http client with request timeout
func NewAPIClient(timeoutMs int) *APIClient {
	return &APIClient{
		client: newHTTPClient(timeoutMs),
		headers: map[string]string{
			"Content-Type": "application/json",
		},
		isSuccessStatus: func(statusCode int) bool {
			return statusCode == http.StatusOK
		},
	}
}

// WithHeaders setup api client default headers
func (c *APIClient) WithHeaders(headers map[string]string) *APIClient {
	c.headers = headers
	return c
}

// WithSuccessStatus setup success status detection func
func (c *APIClient) WithSuccessStatus(
	isSuccessStatus func(statusCode int) bool) *APIClient {
	c.isSuccessStatus = isSuccessStatus
	return c
}

//newHTTPClient creates new http client with timeout
func newHTTPClient(timeoutMs int) *http.Client {
	transport := &http.Transport{
		MaxIdleConnsPerHost: 1024,
	}
	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(timeoutMs) * time.Millisecond,
	}
}

// GetJSON send get http request to url with given req.
// Stores the result  in the value pointed to by res
func (c *APIClient) GetJSON(url string, resp interface{}) error {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	for name, val := range c.headers {
		req.Header.Set(name, val)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if !c.isSuccessStatus(res.StatusCode) {
		io.Copy(ioutil.Discard, res.Body)
		return fmt.Errorf("response status code %d, url %s", res.StatusCode, url)
	}

	return json.NewDecoder(res.Body).Decode(resp)
}

// PostJSON send post http request to url with given req.
// Stores the result  in the value pointed to by resp
func (c *APIClient) PostJSON(url string, reqBody, resp interface{}) error {

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(reqBody); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return err
	}

	for name, val := range c.headers {
		req.Header.Set(name, val)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if !c.isSuccessStatus(res.StatusCode) {
		//for reuse http client connection
		io.Copy(ioutil.Discard, res.Body)

		return fmt.Errorf("response status code %d, url %s", res.StatusCode, url)
	}

	return json.NewDecoder(res.Body).Decode(resp)
}
