package commands

import (
	"bytes"
	"context"
	"flag"
	"testing"

	"github.com/harikiriboy/go-ncvs-cli/writer"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDescribeScanHistories(t *testing.T) {
	Convey("Tests DescribeScanHistories", t, func() {
		client := newMockClient()
		out := new(bytes.Buffer)

		cmd := DescribeScanHistories{Writer: writer.New(out), Client: client, startTime: "2012-03-04 05:06:07", endTime: "2012-03-04 05:06:07"}
		ctx := context.Background()
		flg := flag.NewFlagSet("", flag.ContinueOnError)

		Convey("should success name", func() {
			So(cmd.Name(), ShouldEqual, "describe-scan-histories")
		})
		Convey("should success synopsis", func() {
			So(cmd.Synopsis(), ShouldEqual, "Describe ScanHistories")
		})
		Convey("should success usage", func() {
			So(cmd.Usage(), ShouldEqual, "describe-scan-histories:\n	describe-scan-histories\n	  [-scan-template-name]\n	  [-start-time]\n	  [-end-time]\n")
		})
		Convey("should success set flags", func() {
			cmd.SetFlags(flg)

			So(flg, ShouldNotBeNil)
		})
		Convey("should success execute", func() {
			cmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "this is dummy response.")
		})
		Convey("should success return error message if api error", func() {
			errClient := newMockErrorClient()
			errCmd := DescribeScanHistories{Writer: writer.New(out), Client: errClient}
			errCmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "API request Error this is dummy error.")
		})
		Convey("should success return error message if start-time format error", func() {
			errClient := newMockErrorClient()
			errCmd := DescribeScanHistories{Writer: writer.New(out), Client: errClient, startTime: "invalidate"}
			errCmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "start-time format error")
		})
		Convey("should success return error message if end-time format error", func() {
			errClient := newMockErrorClient()
			errCmd := DescribeScanHistories{Writer: writer.New(out), Client: errClient, endTime: "invalidate"}
			errCmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "end-time format error")
		})
	})
}
