package commands

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"testing"

	"github.com/harikiriboy/go-ncvs-cli/version"
	"github.com/harikiriboy/go-ncvs-cli/writer"
	. "github.com/smartystreets/goconvey/convey"
)

func TestVersion(t *testing.T) {
	Convey("Tests Version", t, func() {
		out := new(bytes.Buffer)
		cmd := Version{Writer: writer.New(out)}
		ctx := context.Background()
		flg := flag.NewFlagSet("", flag.ContinueOnError)

		Convey("should success name", func() {
			So(cmd.Name(), ShouldEqual, "version")
		})
		Convey("should success synopsis", func() {
			So(cmd.Synopsis(), ShouldEqual, "Show version")
		})
		Convey("should success usage", func() {
			So(cmd.Usage(), ShouldEqual, "version:\n	version\n")
		})
		Convey("should success set flags", func() {
			cmd.SetFlags(flg)

			So(flg, ShouldNotBeNil)
		})
		Convey("should success execute", func() {
			cmd.Execute(ctx, flg)

			expected := fmt.Sprintf("ncvs-cli %s\n", version.Version)
			So(out.String(), ShouldEqual, expected)
		})

	})
}
