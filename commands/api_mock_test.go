package commands

import (
	"fmt"

	"github.com/harikiriboy/go-ncvs-cli/ncvs"
)

type mockClient struct{}
type mockErrorClient struct{}

func newMockClient() ncvs.Client {
	return &mockClient{}
}

func newMockErrorClient() ncvs.Client {
	return &mockErrorClient{}
}

func (m *mockClient) CreateScanTemplate(params ncvs.CreateScanTemplateParams) (res string, err error) {
	return "this is dummy response.", nil
}

func (m *mockErrorClient) CreateScanTemplate(params ncvs.CreateScanTemplateParams) (res string, err error) {
	err = fmt.Errorf("this is dummy error.")
	return
}

func (m *mockClient) DescribeScanTemplates(params ncvs.DescribeScanTemplatesParams) (res string, err error) {
	return "this is dummy response.", nil
}

func (m *mockErrorClient) DescribeScanTemplates(params ncvs.DescribeScanTemplatesParams) (res string, err error) {
	err = fmt.Errorf("this is dummy error.")
	return
}

func (m *mockClient) ExecuteScan(params ncvs.ExecuteScanParams) (res string, err error) {
	return "this is dummy response.", nil
}

func (m *mockErrorClient) ExecuteScan(params ncvs.ExecuteScanParams) (res string, err error) {
	err = fmt.Errorf("this is dummy error.")
	return
}

func (m *mockClient) DescribeScanResults(params ncvs.DescribeScanResultsParams) (res string, err error) {
	return "this is dummy response.", nil
}

func (m *mockErrorClient) DescribeScanResults(params ncvs.DescribeScanResultsParams) (res string, err error) {
	err = fmt.Errorf("this is dummy error.")
	return
}

func (m *mockClient) DescribeScanHistories(params ncvs.DescribeScanHistoriesParams) (res string, err error) {
	return "this is dummy response.", nil
}

func (m *mockErrorClient) DescribeScanHistories(params ncvs.DescribeScanHistoriesParams) (res string, err error) {
	err = fmt.Errorf("this is dummy error.")
	return
}

func (m *mockClient) DownloadScanResults(params ncvs.DownloadScanResultsParams) (res string, err error) {
	return "this is dummy response.", nil
}

func (m *mockErrorClient) DownloadScanResults(params ncvs.DownloadScanResultsParams) (res string, err error) {
	err = fmt.Errorf("this is dummy error.")
	return
}

func (m *mockClient) DeleteScanTemplate(params ncvs.DeleteScanTemplateParams) (res string, err error) {
	return "this is dummy response.", nil
}

func (m *mockErrorClient) DeleteScanTemplate(params ncvs.DeleteScanTemplateParams) (res string, err error) {
	err = fmt.Errorf("this is dummy error.")
	return
}

func (m *mockClient) UpdateScanTemplate(params ncvs.UpdateScanTemplateParams) (res string, err error) {
	return "this is dummy response.", nil
}

func (m *mockErrorClient) UpdateScanTemplate(params ncvs.UpdateScanTemplateParams) (res string, err error) {
	err = fmt.Errorf("this is dummy error.")
	return
}

func (m *mockClient) DescribeRulePackages(params ncvs.DescribeRulePackagesParams) (res string, err error) {
	return "this is dummy response.", nil
}

func (m *mockErrorClient) DescribeRulePackages(params ncvs.DescribeRulePackagesParams) (res string, err error) {
	err = fmt.Errorf("this is dummy error.")
	return
}

func (m *mockClient) DescribeRulePackageAttributes(params ncvs.DescribeRulePackageAttributesParams) (res string, err error) {
	return "this is dummy response.", nil
}

func (m *mockErrorClient) DescribeRulePackageAttributes(params ncvs.DescribeRulePackageAttributesParams) (res string, err error) {
	err = fmt.Errorf("this is dummy error.")
	return
}
