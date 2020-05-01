package main

import (
	"github.com/apex/log"
	"github.com/vbstreetz/blzgo"
)

func main() {
	bluzelle.SetupLogging()
	bluzelle.LoadEnv()

	ctx, err := bluzelle.NewTestClient()
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Infof("getting keys...")

	if v, err := ctx.Keys(); err != nil {
		log.Fatalf("%s", err)
	} else {
		log.Infof("keys(%+v)", v)
	}
}
