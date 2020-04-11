package main

import (
	"github.com/apex/log"
	"github.com/vbstreetz/blzgo/examples/util"
	"os"
)

func main() {
	util.SetupLogging()
	util.LoadEnv()

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("key is required")
	}

	ctx, err := util.NewClient()
	if err != nil {
		log.Fatalf("%s", err)
	}

	key := args[0]

	log.Infof("checking if key(%s) exists...", key)

	if v, err := ctx.Has(key); err != nil {
		log.Fatalf("%s", err)
	} else {
		log.Infof("key(%s) exist status: %t", key, v)
	}
}
