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
	DescribeScanHistories(params DescribeScanHistoriesParams) (res string, err error)
	DescribeScanResults(params DescribeScanResultsParams) (res string, err error)
	DescribeScanTemplates(params DescribeScanTemplatesParams) (res string, err error)
	DownloadScanResults(params DownloadScanResultsParams) (res string, err error)
	ExecuteScan(params ExecuteScanParams) (res string, err error)
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

// DescribeScanHistories call DescribeScanHistories API and return Response
func (c *client) DescribeScanHistories(params DescribeScanHistoriesParams) (res string, err error) {
	const action = "DescribeScanHistories"
	res, err = c.doRequest(c.makeRequest(action, params))
	return
}

// DescribeScanResults call DescribeScanResults API and return Response
func (c *client) DescribeScanResults(params DescribeScanResultsParams) (res string, err error) {
	const action = "DescribeScanResults"
	res, err = c.doRequest(c.makeRequest(action, params))
	return
}

// DescribeScanTemplates call DescribeScanTemplates API and return Response
func (c *client) DescribeScanTemplates(params DescribeScanTemplatesParams) (res string, err error) {
	const action = "DescribeScanTemplates"
	res, err = c.doRequest(c.makeRequest(action, params))
	return
}

// DownloadScanResults call DownloadScanResults API and return Response
func (c *client) DownloadScanResults(params DownloadScanResultsParams) (res string, err error) {
	const action = "DownloadScanResults"
	res, err = c.doRequest(c.makeRequest(action, params))
	return
}

// ExecuteScan call ExecuteScan API and return Response
func (c *client) ExecuteScan(params ExecuteScanParams) (res string, err error) {
	const action = "ExecuteScan"
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
