package driver

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

type Driver struct {
	grpcServer *grpc.Server
}

func NewDriver() (*Driver, error) {
	grpcServer := grpc.NewServer()
	csi.RegisterIdentityServer(grpcServer, &identity{})

	controllerSvc, err := newNBSServerControllerService()
	if err != nil {
		return nil, err
	}
	csi.RegisterControllerServer(grpcServer, controllerSvc)

	nodeSvc, err := newNodeService()
	if err != nil {
		return nil, err
	}
	csi.RegisterNodeServer(grpcServer, nodeSvc)

	return &Driver{grpcServer: grpcServer}, nil
}

func (s *Driver) Run(network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGPIPE)

	var wg sync.WaitGroup
	defer wg.Wait()

	defer s.grpcServer.GracefulStop()

	done := make(chan error, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Print("Server finished")

		log.Print("Server started")
		if err := s.grpcServer.Serve(l); err != nil {
			done <- err
		}
	}()

	select {
	case err := <-done:
		if err != nil {
			log.Printf("Server failed: %+v", err)
		}
	case sig := <-signalChannel:
		log.Printf("Got unix signal %q", sig)
	}

	return nil
}
