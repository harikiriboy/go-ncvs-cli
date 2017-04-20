package commands

import (
	"bytes"
	"context"
	"flag"
	"testing"

	"github.com/harikiriboy/go-ncvs-cli/writer"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdateScanTemplate(t *testing.T) {
	Convey("Tests UpdateScanTemplate", t, func() {
		client := newMockClient()
		out := new(bytes.Buffer)

		cmd := UpdateScanTemplate{Writer: writer.New(out), Client: client}
		ctx := context.Background()
		flg := flag.NewFlagSet("", flag.ContinueOnError)

		Convey("should success name", func() {
			So(cmd.Name(), ShouldEqual, "update-scan-template")
		})
		Convey("should success synopsis", func() {
			So(cmd.Synopsis(), ShouldEqual, "Update ScanTemplates")
		})
		Convey("should success usage", func() {
			So(cmd.Usage(), ShouldEqual, "update-scan-template:\n	update-scan-template\n	  [-scan-template-name]\n	  [-update-scan-template-name]\n	  [-ssh-port]\n	  [-description]\n	  [--rule-package-names]\n	  [--scan-targets <{'Region':'region','InstanceUniqueId': 'instanceUniqueId'}>]\n")
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
			errCmd := UpdateScanTemplate{Writer: writer.New(out), Client: errClient}
			errCmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "API request Error this is dummy error.")
		})
	})
}
