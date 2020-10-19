package main

import (
	"github.com/BurntSushi/xgb"
	"github.com/safinsingh/stat"
)

func main() {
	_, err := xgb.NewConn()
	if err != nil {
		stat.Fail("Failed to initialize X connection: " + err.Error())
	}
}
