package ncvs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// testDescribeScanTemplatesResponse is dummy response of describeScanTemplates
const testDescribeScanTemplatesResponse = `{
  "RequestID":"test_request_id",
  "ScanTemplates": [
    {
      "ScanTemplateName": "test_scan_template_name",
      "SSHPort": "test_ssh_port",
      "RulePackageNames": ["test_rule_package_name"],
      "ScanTargets": [
        {
          "Region": "test_region",
          "InstanceUniqueId": "test_instance_unique_id"
        }
      ],
      "Description": "test_description",
      "CreatedTime": "test_created_time"
    }
  ]
}`

// testExecuteScanResponse is dummy response of executeScan
const testExecuteScanResponse = `{
  "RequestID":"test_request_id",
  "ScanHistory": {
    "ScanHistoryUUID": "test_scan_history_uuid",
    "Status": "test_status",
    "StartTime": "test_start_time",
    "EndTime": "test_end_time",
    "ScanTemplateName": "test_scan_template_name"
  }
}`

// testDescribeScanHistoriesResponse is dummy response of describeScanHistories
const testDescribeScanHistoriesResponse = `{
  "RequestID":"test_request_id",
  "ScanHistries": [
    {
      "ScanHistoryUUID": "test_scan_history_uuid",
      "Status": "test_status",
      "StartTime": "test_start_time",
      "EndTime": "test_end_time",
      "ScanTemplateName": "test_scan_template_name"
    }
  ]
}`

// testDescribeScanResultsResponse is dummy response of describeScanResults
const testDescribeScanResultsResponse = `{
  "RequestID":"test_request_id",
  "ScanResults": [
    {
      "Rule": {
        "RuleName": "test_rule_name",
        "Synopsis": "test_synopsis",
        "Description": "test_description",
        "Solution": "test_solution",
        "RiskFactor": "test_risk_factor",
        "CVSSBaseScore": "test_cvss_base_score"
      },
      "Severity": 0,
      "Count": 0,
      "ScanResultTargets": [
        {
          "IPAddress": "test_ip_address",
          "Port": "test_port"
        }
      ],
      "ScanHistoryUUID": "test_scan_history_uuid"
    }
  ]
}`

// testDownloadScanResultsResponse is dummy response of downloadScanResults
const testDownloadScanResultsResponse = `{
  "RequestID":"test_request_id",
  "DownloadURL": "test_download_url"

}`

// testDeleteScanTemplateResponse is dummy response of deleteScanTemplate
const testDeleteScanTemplateResponse = `{
  "RequestID":"test_request_id"
}`

// testCreateScanTemplateResponse is dummy response of createScanTemplate
const testCreateScanTemplateResponse = `{
  "RequestID":"test_request_id",
  "ScanTemplate": {
    "ScanTemplateName": "test_scan_template_name",
    "SSHPort": "test_ssh_port",
    "RulePackageNames": ["test_rule_packages"],
    "ScanTargets": [
      {
        "Region": "test_scan_targets",
        "InstanceUniqueId": "test_instance_unique_id"
      }
    "],
    "Description": "test_description",
    "CreatedTime": "test_created_time"
  }
}`

// testUpdateScanTemplateResponse is dummy response of updateScanTemplate
const testUpdateScanTemplateResponse = `{
  "RequestID":"test_request_id",
  "ScanTemplate": {
    "ScanTemplateName": "test_scan_template_name",
    "SSHPort": "test_ssh_port",
    "RulePackageNames": ["test_rule_packages"],
    "ScanTargets": [
      {
        "Region": "test_scan_targets",
        "InstanceUniqueId": "test_instance_unique_id"
      }
    "],
    "Description": "test_description",
    "CreatedTime": "test_created_time"
  }
}`

// testDescribeRulePackagesResponse is dummy response of describeRulePackages
const testDescribeRulePackagesResponse = `{
  "RequestID":"test_request_id",
}`

// testDescribeRulePackageAttributesResponse is dummy response of describeRulePackageAttributes
const testDescribeRulePackageAttributesResponse = `{
  "RequestID":"test_request_id",
  "RulePackageName": "test_rule_package_name",
  "Marker": "stest_marker",
  "Rules": [
    {
      "RuleName": "test_rule_name",
      "Synopsis": "test_synopsis",
      "Description": "test_description",
      "Solution": "test_solution",
      "RiskFactor": "test_risk_factor",
      "CVSSBaseScore": "test_cvss_base_score"
    }
  ]
}`

// NewTestClient creates dummy http client of ncvs
func newTestClient(body string) (*httptest.Server, Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintln(w, body)
	}))

	client := NewClient(server.URL)

	return server, client
}

func TestNCVSClient(t *testing.T) {
	Convey("Tests NCVS Client", t, func() {
		Convey("should success create scan template", func() {
			_, client := newTestClient(testCreateScanTemplateResponse)
			res, err := client.CreateScanTemplate(CreateScanTemplateParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testCreateScanTemplateResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should success delete scan template", func() {
			_, client := newTestClient(testDeleteScanTemplateResponse)
			res, err := client.DeleteScanTemplate(DeleteScanTemplateParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testDeleteScanTemplateResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should success describe scan templates", func() {
			_, client := newTestClient(testDescribeScanTemplatesResponse)
			res, err := client.DescribeScanTemplates(DescribeScanTemplatesParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testDescribeScanTemplatesResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should success execute scan", func() {
			_, client := newTestClient(testExecuteScanResponse)
			res, err := client.ExecuteScan(ExecuteScanParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testExecuteScanResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should success describe scan histories", func() {
			_, client := newTestClient(testDescribeScanHistoriesResponse)
			res, err := client.DescribeScanHistories(DescribeScanHistoriesParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testDescribeScanHistoriesResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should success describe scan results", func() {
			_, client := newTestClient(testDescribeScanResultsResponse)
			res, err := client.DescribeScanResults(DescribeScanResultsParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testDescribeScanResultsResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should success download scan results", func() {
			_, client := newTestClient(testDownloadScanResultsResponse)
			res, err := client.DownloadScanResults(DownloadScanResultsParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testDownloadScanResultsResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should success update scan template", func() {
			_, client := newTestClient(testUpdateScanTemplateResponse)
			res, err := client.UpdateScanTemplate(UpdateScanTemplateParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testUpdateScanTemplateResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should success describe rule packages", func() {
			_, client := newTestClient(testDescribeRulePackagesResponse)
			res, err := client.DescribeRulePackages(DescribeRulePackagesParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testDescribeRulePackagesResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should success describe rule package attributes", func() {
			_, client := newTestClient(testDescribeRulePackageAttributesResponse)
			res, err := client.DescribeRulePackageAttributes(DescribeRulePackageAttributesParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testDescribeRulePackageAttributesResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
	})
}
