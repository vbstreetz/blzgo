package main

import (
	"github.com/apex/log"
	"github.com/vbstreetz/blzgo/examples/util"
)

func main() {
	util.SetupLogging()
	util.LoadEnv()

	ctx, err := util.NewClient()
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Infof("getting keyvalues...")

	if keyValues, err := ctx.TxKeyValues(util.GasInfo()); err != nil {
		log.Fatalf("%s", err)
	} else {
		log.Infof("values:")
		for _, keyValue := range keyValues {
			log.Infof("key(%s) value(%+v)", keyValue.Key, keyValue.Value)
		}
	}
}
