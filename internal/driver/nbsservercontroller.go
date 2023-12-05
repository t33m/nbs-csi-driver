package driver

import (
	"context"
	"log"

	"github.com/container-storage-interface/spec/lib/go/csi"
	nbsblockstorepublicapi "github.com/ydb-platform/nbs/cloud/blockstore/public/api/protos"
	nbsclient "github.com/ydb-platform/nbs/cloud/blockstore/public/sdk/go/client"
	nbsstoragecoreapi "github.com/ydb-platform/nbs/cloud/storage/core/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const diskBlockSize uint32 = 4 * 1024

var nbsServerControllerServiceCapabilities = []*csi.ControllerServiceCapability{
	{
		Type: &csi.ControllerServiceCapability_Rpc{
			Rpc: &csi.ControllerServiceCapability_RPC{
				Type: csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
			},
		},
	},
	{
		Type: &csi.ControllerServiceCapability_Rpc{
			Rpc: &csi.ControllerServiceCapability_RPC{
				Type: csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
			},
		},
	},
}

type nbsServerControllerService struct {
	csi.ControllerServer

	nbsClient nbsclient.ClientIface
}

func newNBSServerControllerService() (csi.ControllerServer, error) {
	nbsClient, err := nbsclient.NewGrpcClient(
		&nbsclient.GrpcClientOpts{
			Endpoint: "localhost:9766",
			ClientId: driverClientID,
		}, nbsclient.NewStderrLog(nbsclient.LOG_DEBUG),
	)
	if err != nil {
		return nil, err
	}

	return &nbsServerControllerService{nbsClient: nbsClient}, nil
}

func (c *nbsServerControllerService) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	diskID, err := generateDiskID(req.Name)
	if err != nil {
		return nil, err
	}

	limitBytes := req.CapacityRange.LimitBytes
	if uint64(limitBytes)%uint64(diskBlockSize) != 0 {
		return nil, status.Errorf(codes.InvalidArgument,
			"incorrect value: limit_bytes %d, block size: %d", limitBytes, diskBlockSize,
		)
	}

	// TODO: check if volume exists
	//c.nbsClient.DescribeVolume()

	createVolumeResp, err := c.nbsClient.CreateVolume(ctx, &nbsblockstorepublicapi.TCreateVolumeRequest{
		DiskId:           diskID,
		BlockSize:        diskBlockSize,
		BlocksCount:      uint64(limitBytes) / uint64(diskBlockSize),
		StorageMediaKind: nbsstoragecoreapi.EStorageMediaKind_STORAGE_MEDIA_SSD,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("createVolumeResp: %+v", createVolumeResp)

	return &csi.CreateVolumeResponse{Volume: &csi.Volume{
		CapacityBytes: limitBytes,
		VolumeId:      diskID,
	}}, nil
}

func (c *nbsServerControllerService) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	// TODO: check if volume exists
	//c.nbsClient.DescribeVolume()

	deleteVolumeResp, err := c.nbsClient.DestroyVolume(ctx, &nbsblockstorepublicapi.TDestroyVolumeRequest{
		DiskId: req.VolumeId,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("deleteVolumeResp: %+v", deleteVolumeResp)

	return &csi.DeleteVolumeResponse{}, nil
}

func (c *nbsServerControllerService) ControllerGetCapabilities(context.Context, *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: nbsServerControllerServiceCapabilities,
	}, nil
}
