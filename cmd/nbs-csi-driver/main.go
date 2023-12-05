package main

import (
	"flag"

	"github.com/t33m/nbs-csi-driver/internal/driver"
)

const driverName = "nbs.csi.nebius.ai"

var version = ""

func main() {
	if version == "" {
		version = "dev"
	}

	cfg := driver.Config{
		DriverName:    driverName,
		VendorVersion: version,
	}

	flag.StringVar(&cfg.Endpoint, "endpoint", "unix://csi/csi.sock", "CSI endpoint")
	flag.StringVar(&cfg.NodeID, "node-id", "", "Node id")

	flag.Parse()

	srv, err := driver.NewDriver(cfg)
	if err != nil {
		panic(err)
	}
	if err := srv.Run("/csi/csi.sock"); err != nil {
		panic(err)
	}
}
