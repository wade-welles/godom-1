package godom

import (
	"encoding/json"
	"io"
	"net/http"
)

// HTTPRequest .
type HTTPRequest struct {
	method    string
	url       string
	component *Component
	body      io.Reader
	header    http.Header
}

// FromJSON .
func (r *HTTPRequest) FromJSON(target interface{}, cb func(*http.Response)) error {
	req, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		return err
	}
	req.Header = r.header
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return err
	}
	if r.component.Quit != nil {
		cb(res)
	}
	return nil
}

// Header .
func (r *HTTPRequest) Header(key string, val string) {
	r.header.Set(key, val)
}
