package driver

import (
	"context"
	"log"
	"path/filepath"

	"github.com/container-storage-interface/spec/lib/go/csi"
	nbsblockstorepublicapi "github.com/ydb-platform/nbs/cloud/blockstore/public/api/protos"
	nbsclient "github.com/ydb-platform/nbs/cloud/blockstore/public/sdk/go/client"
)

const unixSocketPath = "/tmp/nbs.sock" // TODO: calculate from request

const topologyKeyNode = "topology.nbs.csi/node"

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

	nodeID    string
	clientID  string
	nbsClient nbsclient.ClientIface
}

func newNodeService(nodeID, clientID string, nbsClient nbsclient.ClientIface) csi.NodeServer {
	return &nodeService{nodeID: nodeID, clientID: clientID, nbsClient: nbsClient}
}

func (s *nodeService) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	log.Printf("csi.NodeStageVolumeRequest: %+v", req)

	// TODO: check if endpoint exists
	//s.nbsClient.DescribeEndpoint()

	hostType := nbsblockstorepublicapi.EHostType_HOST_TYPE_DEFAULT

	startEndpointResp, err := s.nbsClient.StartEndpoint(ctx, &nbsblockstorepublicapi.TStartEndpointRequest{
		UnixSocketPath:            filepath.Join(req.StagingTargetPath, "nbs.sock"),
		DiskId:                    req.VolumeId,
		ClientId:                  s.clientID,
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

func (s *nodeService) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	log.Printf("csi.NodeUnstageVolumeRequest: %+v", req)

	// TODO: check if endpoint exists
	//s.nbsClient.DescribeEndpoint()

	stopEndpointResp, err := s.nbsClient.StopEndpoint(ctx, &nbsblockstorepublicapi.TStopEndpointRequest{
		UnixSocketPath: unixSocketPath,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("stopEndpointResp: %+v", stopEndpointResp)

	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (s *nodeService) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	log.Printf("csi.NodePublishVolumeRequest: %+v", req)

	return &csi.NodePublishVolumeResponse{}, nil
}

func (s *nodeService) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Printf("csi.NodeUnpublishVolumeRequest: %+v", req)

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (s *nodeService) NodeGetCapabilities(context.Context, *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: capabilities,
	}, nil
}

func (s *nodeService) NodeGetInfo(context.Context, *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{
		NodeId: s.nodeID,
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{topologyKeyNode: s.nodeID},
		},
	}, nil
}
