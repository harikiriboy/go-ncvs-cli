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
