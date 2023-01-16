package main

import (
	"github.com/mingrammer/flog/pkg"
	"math/rand"
	"time"

	"github.com/mingrammer/cfmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	opts := pkg.ParseOptions()
	if err := pkg.Run(opts); err != nil {
		cfmt.Warningln(err.Error())
	}
}
