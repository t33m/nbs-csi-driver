package main

import "github.com/t33m/nbs-csi-driver/internal/driver"

func main() {
	srv, err := driver.NewDriver()
	if err != nil {
		panic(err)
	}
	if err := srv.Run("unix", "/csi/csi.sock"); err != nil {
		panic(err)
	}
}
