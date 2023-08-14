// Copyright 2023, NVIDIA CORPORATION & AFFILIATES
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: nvidia/ipam/node/v1/node.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IPAMService_Allocate_FullMethodName    = "/nvidia.ipam.node.v1.IPAMService/Allocate"
	IPAMService_IsAllocated_FullMethodName = "/nvidia.ipam.node.v1.IPAMService/IsAllocated"
	IPAMService_Deallocate_FullMethodName  = "/nvidia.ipam.node.v1.IPAMService/Deallocate"
)

// IPAMServiceClient is the client API for IPAMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IPAMServiceClient interface {
	// Allocate is called as a part of CMD_ADD flow.
	// Returns response with allocated IPs if allocation succeeds or an error in case of failure.
	// If multiple pools are provided in the request, then allocation succeeds if it succeeds for all pools.
	// errors:
	//
	//	Canceled - request was canceled by the caller
	//	Internal - internal failure of the service, this error can't be fixed by the caller
	//	InvalidArgument - missing required argument or argument has wrong format, check message for details
	//	NotFound - allocation is requested for unknown IP pool
	//	AlreadyExists - container identified by IPAMParameters already has allocated IP in the pool
	//	ResourceExhausted - no free IP addresses available in the IP pool
	Allocate(ctx context.Context, in *AllocateRequest, opts ...grpc.CallOption) (*AllocateResponse, error)
	// IsAllocated is called as a part of CMD_CHECK flow
	// Returns empty response if a valid allocation already exists or an error otherwise.
	// If multiple pools are provided in the request, then check
	// succeed only if it is succeed for all pools.
	// errors:
	//
	//	Canceled - request was canceled by the caller
	//	Internal - internal failure of the service, this error can't be fixed by the caller
	//	InvalidArgument - missing required argument or argument has wrong format, check message for details
	//	NotFound - allocation not found
	IsAllocated(ctx context.Context, in *IsAllocatedRequest, opts ...grpc.CallOption) (*IsAllocatedResponse, error)
	// Deallocate is called as a part of CMD_DEL flow.
	// Returns empty response if allocation for all pools released successfully or error otherwise.
	// If multiple pools are provided in the request, then deallocation
	// succeeds if it is succeeds for all pools else no deallocation is performed.
	// errors:
	//
	//	Canceled - request was canceled by the caller
	//	Internal - internal failure of the service, this error can't be fixed by the caller
	//	InvalidArgument - missing required argument or argument has wrong format, check message for details
	Deallocate(ctx context.Context, in *DeallocateRequest, opts ...grpc.CallOption) (*DeallocateResponse, error)
}

type iPAMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIPAMServiceClient(cc grpc.ClientConnInterface) IPAMServiceClient {
	return &iPAMServiceClient{cc}
}

func (c *iPAMServiceClient) Allocate(ctx context.Context, in *AllocateRequest, opts ...grpc.CallOption) (*AllocateResponse, error) {
	out := new(AllocateResponse)
	err := c.cc.Invoke(ctx, IPAMService_Allocate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPAMServiceClient) IsAllocated(ctx context.Context, in *IsAllocatedRequest, opts ...grpc.CallOption) (*IsAllocatedResponse, error) {
	out := new(IsAllocatedResponse)
	err := c.cc.Invoke(ctx, IPAMService_IsAllocated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPAMServiceClient) Deallocate(ctx context.Context, in *DeallocateRequest, opts ...grpc.CallOption) (*DeallocateResponse, error) {
	out := new(DeallocateResponse)
	err := c.cc.Invoke(ctx, IPAMService_Deallocate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPAMServiceServer is the server API for IPAMService service.
// All implementations must embed UnimplementedIPAMServiceServer
// for forward compatibility
type IPAMServiceServer interface {
	// Allocate is called as a part of CMD_ADD flow.
	// Returns response with allocated IPs if allocation succeeds or an error in case of failure.
	// If multiple pools are provided in the request, then allocation succeeds if it succeeds for all pools.
	// errors:
	//
	//	Canceled - request was canceled by the caller
	//	Internal - internal failure of the service, this error can't be fixed by the caller
	//	InvalidArgument - missing required argument or argument has wrong format, check message for details
	//	NotFound - allocation is requested for unknown IP pool
	//	AlreadyExists - container identified by IPAMParameters already has allocated IP in the pool
	//	ResourceExhausted - no free IP addresses available in the IP pool
	Allocate(context.Context, *AllocateRequest) (*AllocateResponse, error)
	// IsAllocated is called as a part of CMD_CHECK flow
	// Returns empty response if a valid allocation already exists or an error otherwise.
	// If multiple pools are provided in the request, then check
	// succeed only if it is succeed for all pools.
	// errors:
	//
	//	Canceled - request was canceled by the caller
	//	Internal - internal failure of the service, this error can't be fixed by the caller
	//	InvalidArgument - missing required argument or argument has wrong format, check message for details
	//	NotFound - allocation not found
	IsAllocated(context.Context, *IsAllocatedRequest) (*IsAllocatedResponse, error)
	// Deallocate is called as a part of CMD_DEL flow.
	// Returns empty response if allocation for all pools released successfully or error otherwise.
	// If multiple pools are provided in the request, then deallocation
	// succeeds if it is succeeds for all pools else no deallocation is performed.
	// errors:
	//
	//	Canceled - request was canceled by the caller
	//	Internal - internal failure of the service, this error can't be fixed by the caller
	//	InvalidArgument - missing required argument or argument has wrong format, check message for details
	Deallocate(context.Context, *DeallocateRequest) (*DeallocateResponse, error)
	mustEmbedUnimplementedIPAMServiceServer()
}

// UnimplementedIPAMServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIPAMServiceServer struct {
}

func (UnimplementedIPAMServiceServer) Allocate(context.Context, *AllocateRequest) (*AllocateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allocate not implemented")
}
func (UnimplementedIPAMServiceServer) IsAllocated(context.Context, *IsAllocatedRequest) (*IsAllocatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAllocated not implemented")
}
func (UnimplementedIPAMServiceServer) Deallocate(context.Context, *DeallocateRequest) (*DeallocateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deallocate not implemented")
}
func (UnimplementedIPAMServiceServer) mustEmbedUnimplementedIPAMServiceServer() {}

// UnsafeIPAMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IPAMServiceServer will
// result in compilation errors.
type UnsafeIPAMServiceServer interface {
	mustEmbedUnimplementedIPAMServiceServer()
}

func RegisterIPAMServiceServer(s grpc.ServiceRegistrar, srv IPAMServiceServer) {
	s.RegisterService(&IPAMService_ServiceDesc, srv)
}

func _IPAMService_Allocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPAMServiceServer).Allocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPAMService_Allocate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPAMServiceServer).Allocate(ctx, req.(*AllocateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPAMService_IsAllocated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAllocatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPAMServiceServer).IsAllocated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPAMService_IsAllocated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPAMServiceServer).IsAllocated(ctx, req.(*IsAllocatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPAMService_Deallocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeallocateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPAMServiceServer).Deallocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPAMService_Deallocate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPAMServiceServer).Deallocate(ctx, req.(*DeallocateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IPAMService_ServiceDesc is the grpc.ServiceDesc for IPAMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IPAMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nvidia.ipam.node.v1.IPAMService",
	HandlerType: (*IPAMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allocate",
			Handler:    _IPAMService_Allocate_Handler,
		},
		{
			MethodName: "IsAllocated",
			Handler:    _IPAMService_IsAllocated_Handler,
		},
		{
			MethodName: "Deallocate",
			Handler:    _IPAMService_Deallocate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nvidia/ipam/node/v1/node.proto",
}