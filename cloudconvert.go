package cloudconvert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	BASE_URL = "https://api.cloudconvert.com"
)

type Client struct {
	client  *http.Client // HTTP client used to communicate with the API
	Token   string       // API Key obtained via box
	BaseURL *url.URL     // Base URL for API requests.

	Process    *ProcessService    // Service to fetch process related data
	Conversion *ConversionService // Service to fetch conversion related data
}

func NewClient(token string) *Client {
	baseURL, _ := url.Parse(BASE_URL)
	c := &Client{client: http.DefaultClient, Token: token, BaseURL: baseURL}
	c.Process = &ProcessService{client: c}
	c.Conversion = &ConversionService{client: c}
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
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

	req.Header.Add("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
				return resp, err
			}
		}
	}

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &err)
	return resp, err
}
