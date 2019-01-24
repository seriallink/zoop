package zoop

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseUrl = "https://api.zoop.ws/v1/"

// Map data por post in the request
type Params map[string]interface{}

// Map extra request headers
type Headers map[string]string

// Make request and return the response
func (c *Client) execute(method string, path string, params interface{}, headers Headers, model interface{}) error {

	// init vars
	var url = baseUrl + path

	// init an empty payload
	payload := strings.NewReader("")

	// check for params
	if params != nil {

		// marshal params
		b, err := json.Marshal(params)
		if err != nil {
			return err
		}

		// set payload with params
		payload = strings.NewReader(string(b))

	}

	// set request
	request, _ := http.NewRequest(method, url, payload)
	request.Header.Add("Authorization", c.BasicAuth)
	request.Header.Add("accept", "application/json")
	request.Header.Add("content-type", "application/json")

	// add extra headers
	if headers != nil {
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// read response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// init zoop error response
	er := &ErrResponse{}

	// check for error message
	if err = json.Unmarshal(data, er); err == nil && er.ErrObject != nil {
		return er.ErrObject
	}

	// parse data
	return json.Unmarshal(data, model)

}

// Execute GET requests
func (c *Client) Get(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute("GET", path, params, headers, model)
}

// Execute POST requests
func (c *Client) Post(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute("POST", path, params, headers, model)
}

// Execute PUT requests
func (c *Client) Put(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute("PUT", path, params, headers, model)
}

// Execute DELETE requests
func (c *Client) Delete(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute("DELETE", path, params, headers, model)
}
