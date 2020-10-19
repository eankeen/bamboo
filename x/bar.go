package x

import (
	"github.com/eankeen/bamboo/types"

	"github.com/BurntSushi/xgb"
	"github.com/safinsingh/stat"
)

func Draw(conf types.Config) {
	_, err := xgb.NewConn()
	if err != nil {
		stat.Fail("Failed to initialize X connection: " + err.Error())
	}
}
