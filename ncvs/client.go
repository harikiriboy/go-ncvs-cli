package ncvs

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/smartystreets/go-aws-auth"
)

const (
	// DefaultAPIEndPoint is https://vss.api.cloud.nifty.com
	DefaultAPIEndPoint = "https://vss.api.cloud.nifty.com"
	// DefaultAPIVersion is 2017-02-23
	DefaultAPIVersion = "2017-02-23"
)

// Client interface
type Client interface {
	DescribeScanTemplates(params DescribeScanTemplatesParams) (res string, err error)
}

// client struct
type client struct {
	endpoint string
}

// NewClient returns a new API client
func NewClient(endpoint string) Client {
	return &client{
		endpoint: endpoint,
	}
}

// DescribeScanTemplates call DescribeScanTemplates API and return Response
func (c *client) DescribeScanTemplates(params DescribeScanTemplatesParams) (res string, err error) {
	const action = "DescribeScanTemplates"
	res, err = c.doRequest(c.makeRequest(action, params))
	return
}

// makeRequest return API request
func (c *client) makeRequest(action string, params interface{}) (req *http.Request) {
	endpoint, _ := url.Parse(c.endpoint)
	req, _ = http.NewRequest("POST", endpoint.String(), c.encodeJSON(params))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Amz-Target", DefaultAPIVersion+"."+action)
	awsauth.Sign4(req)
	return
}

// doRequest sends an API request and returns the API response decode json.
func (c *client) doRequest(req *http.Request) (res string, err error) {
	result, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer result.Body.Close()
	bodyBytes, err := ioutil.ReadAll(result.Body)
	res = string(bodyBytes)
	return
}

// encodeJSON return encoding json
func (c *client) encodeJSON(object interface{}) (reader io.Reader) {
	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(object)

	if err != nil {
		return
	}

	reader = buffer
	return
}
