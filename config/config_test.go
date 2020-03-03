package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Given an environment with no environment variables set", t, func() {
		cfg, err := Get()

		Convey("When the config values are retrieved", func() {

			Convey("Then there should be no error returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("The values should be set to the expected defaults", func() {
				So(cfg.BindAddr, ShouldEqual, ":22300")
				So(cfg.MongoConfig.BindAddr, ShouldEqual, "http://localhost:27017")
				So(cfg.MongoConfig.Collection, ShouldEqual, "recipes")
				So(cfg.MongoConfig.Database, ShouldEqual, "recipe-db")
			})
		})
	})
}
