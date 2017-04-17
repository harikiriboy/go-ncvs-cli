package commands

import (
	"bytes"
	"context"
	"flag"
	"testing"

	"github.com/harikiriboy/go-ncvs-cli/writer"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDescribeScanTemplates(t *testing.T) {
	Convey("Tests DescribeTemplates", t, func() {
		client := newMockClient()
		out := new(bytes.Buffer)

		cmd := DescribeScanTemplates{Writer: writer.New(out), Client: client}
		ctx := context.Background()
		flg := flag.NewFlagSet("", flag.ContinueOnError)

		Convey("should success name", func() {
			So(cmd.Name(), ShouldEqual, "describe-scan-templates")
		})
		Convey("should success synopsis", func() {
			So(cmd.Synopsis(), ShouldEqual, "Describe ScanTemplates")
		})
		Convey("should success usage", func() {
			So(cmd.Usage(), ShouldEqual, "describe-scan-templates:\n	describe-scan-templates\n	  [-scan-template-name]\n")
		})
		Convey("should success set flags", func() {
			cmd.SetFlags(flg)

			So(flg, ShouldNotBeNil)
		})
		Convey("should success execute", func() {
			cmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "this is dummy response.")
		})
		Convey("should success return error message", func() {
			errClient := newMockErrorClient()
			errCmd := DescribeScanTemplates{Writer: writer.New(out), Client: errClient}
			errCmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "API request Error this is dummy error.")
		})
	})
}
