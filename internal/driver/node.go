package driver

import (
	"context"
	"log"

	"github.com/container-storage-interface/spec/lib/go/csi"
	nbsblockstorepublicapi "github.com/ydb-platform/nbs/cloud/blockstore/public/api/protos"
	nbsclient "github.com/ydb-platform/nbs/cloud/blockstore/public/sdk/go/client"
)

const unixSocketPath = "/tmp/nbs.sock"

var capabilities = []*csi.NodeServiceCapability{
	{
		Type: &csi.NodeServiceCapability_Rpc{
			Rpc: &csi.NodeServiceCapability_RPC{
				Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
			},
		},
	},
}

type nodeService struct {
	csi.NodeServer

	nbsClient nbsclient.ClientIface
}

func newNodeService() (csi.NodeServer, error) {
	nbsClient, err := nbsclient.NewGrpcClient(
		&nbsclient.GrpcClientOpts{
			Endpoint: "localhost:9766",
			ClientId: driverClientID,
		}, nbsclient.NewStderrLog(nbsclient.LOG_DEBUG),
	)
	if err != nil {
		return nil, err
	}

	return &nodeService{nbsClient: nbsClient}, nil
}

func (c *nodeService) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	// TODO: check if endpoint exists
	//c.nbsClient.DescribeEndpoint()

	hostType := nbsblockstorepublicapi.EHostType_HOST_TYPE_DEFAULT

	startEndpointResp, err := c.nbsClient.StartEndpoint(ctx, &nbsblockstorepublicapi.TStartEndpointRequest{
		UnixSocketPath:            unixSocketPath,
		DiskId:                    req.VolumeId,
		ClientId:                  driverClientID,
		DeviceName:                req.VolumeId,
		MountSeqNumber:            0, // TODO: ?
		IpcType:                   nbsblockstorepublicapi.EClientIpcType_IPC_VHOST,
		VhostQueuesCount:          8,
		VolumeAccessMode:          nbsblockstorepublicapi.EVolumeAccessMode_VOLUME_ACCESS_READ_WRITE,
		VolumeMountMode:           nbsblockstorepublicapi.EVolumeMountMode_VOLUME_MOUNT_REMOTE,
		UnalignedRequestsDisabled: false,
		ClientProfile: &nbsblockstorepublicapi.TClientProfile{
			//CpuUnitCount: nil, // TODO: ?
			HostType: &hostType,
		},
		//ClientCGroups:             clientCgroups, // TODO: ?
	})
	if err != nil {
		return nil, err
	}

	log.Printf("startEndpointResp: %+v", startEndpointResp)

	return &csi.NodeStageVolumeResponse{}, err
}

func (c *nodeService) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	// TODO: check if endpoint exists
	//c.nbsClient.DescribeEndpoint()

	stopEndpointResp, err := c.nbsClient.StopEndpoint(ctx, &nbsblockstorepublicapi.TStopEndpointRequest{
		UnixSocketPath: unixSocketPath,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("stopEndpointResp: %+v", stopEndpointResp)

	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (c *nodeService) NodeGetCapabilities(context.Context, *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: capabilities,
	}, nil
}
