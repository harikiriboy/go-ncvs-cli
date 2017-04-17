package ncvs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestDescribeScanTemplatesResponse is dummy response of describeScanTemplates
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
		Convey("should success describe scan templates", func() {
			_, client := newTestClient(testDescribeScanTemplatesResponse)
			res, err := client.DescribeScanTemplates(DescribeScanTemplatesParams{})

			var actual, expected interface{}
			json.Unmarshal([]byte(res), &actual)
			json.Unmarshal([]byte(testDescribeScanTemplatesResponse), &expected)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
	})
}
