package c01

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/**
 * If the goconvey command does not exist, try the following steps:
 *
 * -> export GOPATH=/path/to/go
 * -> cd ./src
 * -> go build -o ../bin/goconvey github.com/smartystreets/goconvey
 */
func TestSpec(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 1
		Convey("When the integer is incremented", func() {
			x++
			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
