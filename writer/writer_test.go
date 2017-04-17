package writer

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWriter(t *testing.T) {
	Convey("Tests Writer", t, func() {
		out := new(bytes.Buffer)
		w := New(out)
		Convey("should success print", func() {
			expected := "test"
			w.Print("test")
			So(out.String(), ShouldEqual, expected)
		})
		Convey("should success printf", func() {
			str := "printf"
			expected := fmt.Sprintf("test %s", str)
			w.Printf("test %s", str)
			So(out.String(), ShouldEqual, expected)
		})
	})
}
