package runtime

import (
	"context"
	"ha/pkg/log"
	"os"
)

var HostName string

func init() {
	var err error
	HostName, err = os.Hostname()
	if err != nil {
		log.G(context.TODO()).Fatalf("can't resolve hostname. %+v", err)
		os.Exit(1)
	}
}
