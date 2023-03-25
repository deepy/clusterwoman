package cloudprovider

import (
	"context"
	"errors"
	"github.com/bufbuild/connect-go"
	"github.com/deepy/externalgrpc-connect/gen/clusterautoscaler/cloudprovider/v1/externalgrpc"
	metav1 "github.com/deepy/externalgrpc-connect/gen/k8s.io/apimachinery/pkg/apis/meta/v1"
)

var nodeGroup = externalgrpc.NodeGroup{Id: "metal", MinSize: 1, MaxSize: 3}

type CloudProviderServer struct {
	Nodes map[string]string
}

func (c CloudProviderServer) NodeGroups(ctx context.Context, c2 *connect.Request[externalgrpc.NodeGroupsRequest]) (*connect.Response[externalgrpc.NodeGroupsResponse], error) {
	//TODO implement me
	return connect.NewResponse(&externalgrpc.NodeGroupsResponse{NodeGroups: []*externalgrpc.NodeGroup{&nodeGroup}}), nil
}

func (c CloudProviderServer) NodeGroupForNode(ctx context.Context, c2 *connect.Request[externalgrpc.NodeGroupForNodeRequest]) (*connect.Response[externalgrpc.NodeGroupForNodeResponse], error) {
	//TODO implement me
	return connect.NewResponse(&externalgrpc.NodeGroupForNodeResponse{NodeGroup: &nodeGroup}), nil
}

func (c CloudProviderServer) PricingNodePrice(ctx context.Context, c2 *connect.Request[externalgrpc.PricingNodePriceRequest]) (*connect.Response[externalgrpc.PricingNodePriceResponse], error) {
	//Implementation optional.
	return nil, connect.NewError(connect.CodeUnknown, errors.New("not implemented"))
}

func (c CloudProviderServer) PricingPodPrice(ctx context.Context, c2 *connect.Request[externalgrpc.PricingPodPriceRequest]) (*connect.Response[externalgrpc.PricingPodPriceResponse], error) {
	//Implementation optional.
	return nil, connect.NewError(connect.CodeUnknown, errors.New("not implemented"))
}

func (c CloudProviderServer) GPULabel(ctx context.Context, c2 *connect.Request[externalgrpc.GPULabelRequest]) (*connect.Response[externalgrpc.GPULabelResponse], error) {
	//LONGTERM implement me
	return connect.NewResponse(&externalgrpc.GPULabelResponse{}), nil
}

func (c CloudProviderServer) GetAvailableGPUTypes(ctx context.Context, c2 *connect.Request[externalgrpc.GetAvailableGPUTypesRequest]) (*connect.Response[externalgrpc.GetAvailableGPUTypesResponse], error) {
	//TODO implement me
	return connect.NewResponse(&externalgrpc.GetAvailableGPUTypesResponse{GpuTypes: nil}), nil
}

func (c CloudProviderServer) Cleanup(ctx context.Context, c2 *connect.Request[externalgrpc.CleanupRequest]) (*connect.Response[externalgrpc.CleanupResponse], error) {
	//TODO implement me
	return connect.NewResponse(&externalgrpc.CleanupResponse{}), nil
}

func (c CloudProviderServer) Refresh(ctx context.Context, c2 *connect.Request[externalgrpc.RefreshRequest]) (*connect.Response[externalgrpc.RefreshResponse], error) {
	//TODO implement me
	return connect.NewResponse(&externalgrpc.RefreshResponse{}), nil
}

func (c CloudProviderServer) NodeGroupTargetSize(ctx context.Context, c2 *connect.Request[externalgrpc.NodeGroupTargetSizeRequest]) (*connect.Response[externalgrpc.NodeGroupTargetSizeResponse], error) {
	//TODO implement me
	//c2.Msg.Id - nodegroup ID
	return connect.NewResponse(&externalgrpc.NodeGroupTargetSizeResponse{TargetSize: int32(len(c.Nodes))}), nil
	panic("implement me NodeGroupTargetSize")
}

func (c CloudProviderServer) NodeGroupIncreaseSize(ctx context.Context, c2 *connect.Request[externalgrpc.NodeGroupIncreaseSizeRequest]) (*connect.Response[externalgrpc.NodeGroupIncreaseSizeResponse], error) {
	//TODO implement me
	panic("implement me NodeGroupIncreaseSize")
}

func (c CloudProviderServer) NodeGroupDeleteNodes(ctx context.Context, c2 *connect.Request[externalgrpc.NodeGroupDeleteNodesRequest]) (*connect.Response[externalgrpc.NodeGroupDeleteNodesResponse], error) {
	//TODO implement me
	panic("implement me NodeGroupDeleteNodes")
}

func (c CloudProviderServer) NodeGroupDecreaseTargetSize(ctx context.Context, c2 *connect.Request[externalgrpc.NodeGroupDecreaseTargetSizeRequest]) (*connect.Response[externalgrpc.NodeGroupDecreaseTargetSizeResponse], error) {
	//TODO implement me
	panic("implement me NodeGroupDecreaseTargetSize")
}

func (c CloudProviderServer) NodeGroupNodes(ctx context.Context, c2 *connect.Request[externalgrpc.NodeGroupNodesRequest]) (*connect.Response[externalgrpc.NodeGroupNodesResponse], error) {
	//TODO implement me
	return connect.NewResponse(&externalgrpc.NodeGroupNodesResponse{Instances: nil}), nil
}

func (c CloudProviderServer) NodeGroupTemplateNodeInfo(ctx context.Context, c2 *connect.Request[externalgrpc.NodeGroupTemplateNodeInfoRequest]) (*connect.Response[externalgrpc.NodeGroupTemplateNodeInfoResponse], error) {
	//Implementation optional.
	// Can I just error here?
	return nil, connect.NewError(connect.CodeUnknown, errors.New("not implemented"))
}

func (c CloudProviderServer) NodeGroupGetOptions(ctx context.Context, c2 *connect.Request[externalgrpc.NodeGroupAutoscalingOptionsRequest]) (*connect.Response[externalgrpc.NodeGroupAutoscalingOptionsResponse], error) {
	//Implementation optional.
	// NodeGroup. Returning a grpc error will result in using default options.
	minute := int64(60)
	duration := metav1.Duration{Duration: &minute}
	//TODO implement me properly
	return connect.NewResponse(&externalgrpc.NodeGroupAutoscalingOptionsResponse{NodeGroupAutoscalingOptions: &externalgrpc.NodeGroupAutoscalingOptions{
		ScaleDownUtilizationThreshold:    0,
		ScaleDownGpuUtilizationThreshold: 0,
		ScaleDownUnneededTime:            &duration,
		ScaleDownUnreadyTime:             &duration,
	}}), nil
	//return nil, connect.NewError(connect.CodeUnknown, errors.New("not implemented"))
}
