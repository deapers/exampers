package trade

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/wendal/go-oci8"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadFundInfo(t *testing.T) {
	Convey("加载FundInfo", t, func() {
		os.Setenv("NLS_LANG", "")
		dbcon, err := sql.Open("oci8", "cta1/oracle@127.0.0.1/orcl.js.local")
		So(err, ShouldBeNil)
		So(dbcon, ShouldNotBeNil)
		Convey("全局Map为nil时的加载", func() {
			So(LoadFundInfo(dbcon), ShouldBeNil)
		})
		Convey("全局Map非nil时的加载", func() {
			So(LoadFundInfo(dbcon), ShouldBeNil)
		})
	})
}
