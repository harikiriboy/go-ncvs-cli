package ncvs

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"

	"github.com/harikiriboy/go-ncvs-cli/version"
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
	CreateScanTemplate(params CreateScanTemplateParams) (res string, err error)
	DeleteScanTemplate(params DeleteScanTemplateParams) (res string, err error)
	DescribeRulePackages(params DescribeRulePackagesParams) (res string, err error)
	DescribeRulePackageAttributes(params DescribeRulePackageAttributesParams) (res string, err error)
	DescribeScanHistories(params DescribeScanHistoriesParams) (res string, err error)
	DescribeScanResults(params DescribeScanResultsParams) (res string, err error)
	DescribeScanTemplates(params DescribeScanTemplatesParams) (res string, err error)
	DownloadScanResults(params DownloadScanResultsParams) (res string, err error)
	ExecuteScan(params ExecuteScanParams) (res string, err error)
	UpdateScanTemplate(params UpdateScanTemplateParams) (res string, err error)
}

// client struct
type client struct {
	endpoint string
}

// NewClient returns a new API client
func NewClient(endpoint string, ignoreSSLCertsErrors bool) Client {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: ignoreSSLCertsErrors,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig, Proxy: http.ProxyFromEnvironment}
	http.DefaultClient.Transport = transport

	return &client{
		endpoint: endpoint,
	}
}

// CreateScanTemplate call CreateScanTemplate API and return Response
func (c *client) CreateScanTemplate(params CreateScanTemplateParams) (res string, err error) {
	const action = "CreateScanTemplate"
	res, err = c.doRequest(c.makeRequest(action, params))
	return
}

// DeleteScanTemplate call DeleteScanTemplate API and return Response
func (c *client) DeleteScanTemplate(params DeleteScanTemplateParams) (res string, err error) {
	const action = "DeleteScanTemplate"
	res, err = c.doRequest(c.makeRequest(action, params))
	return
}

// DescribeRulePackageAttributes call DescribeRulePackageAttributes API and return Response
func (c *client) DescribeRulePackageAttributes(params DescribeRulePackageAttributesParams) (res string, err error) {
	const action = "DescribeRulePackageAttributes"
	res, err = c.doRequest(c.makeRequest(action, params))
	return
}

// DescribeRulePackages call DescribeRulePackages API and return Response
func (c *client) DescribeRulePackages(params DescribeRulePackagesParams) (res string, err error) {
	const action = "DescribeRulePackages"
	res, err = c.doRequest(c.makeRequest(action, params))
	return
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

// UpdateScanTemplate call UpdateScanTemplate API and return Response
func (c *client) UpdateScanTemplate(params UpdateScanTemplateParams) (res string, err error) {
	const action = "UpdateScanTemplate"
	res, err = c.doRequest(c.makeRequest(action, params))
	return
}

// makeRequest return API request
func (c *client) makeRequest(action string, params interface{}) (req *http.Request) {
	endpoint, _ := url.Parse(c.endpoint)
	req, _ = http.NewRequest("POST", endpoint.String(), c.encodeJSON(params))

	req.Header.Set("User-Agent", "go-ncvs-cli "+version.Version+" "+runtime.GOOS+" "+runtime.GOARCH)
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
