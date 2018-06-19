// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: master_service.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MasterService service

type MasterServiceClient interface {
	// Creates a session.
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	// Extends a session.
	ExtendSession(ctx context.Context, in *ExtendSessionRequest, opts ...grpc.CallOption) (*ExtendSessionResponse, error)
	// Prepares future partial run calls.
	PartialRunSetup(ctx context.Context, in *PartialRunSetupRequest, opts ...grpc.CallOption) (*PartialRunSetupResponse, error)
	// Drives the graph computation.
	RunStep(ctx context.Context, in *RunStepRequest, opts ...grpc.CallOption) (*RunStepResponse, error)
	// Closes a session.
	CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*CloseSessionResponse, error)
	// List the devices usable by the master.
	ListDevices(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error)
	// Close and abandon all existing sessions.  Ongoing computations
	// will no longer affect fresh ones via the resources in containers listed in
	// the ResetRequest.  See ResetRequest for more details.
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error)
}

type masterServiceClient struct {
	cc *grpc.ClientConn
}

func NewMasterServiceClient(cc *grpc.ClientConn) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	err := grpc.Invoke(ctx, "/tensorflow.MasterService/CreateSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) ExtendSession(ctx context.Context, in *ExtendSessionRequest, opts ...grpc.CallOption) (*ExtendSessionResponse, error) {
	out := new(ExtendSessionResponse)
	err := grpc.Invoke(ctx, "/tensorflow.MasterService/ExtendSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) PartialRunSetup(ctx context.Context, in *PartialRunSetupRequest, opts ...grpc.CallOption) (*PartialRunSetupResponse, error) {
	out := new(PartialRunSetupResponse)
	err := grpc.Invoke(ctx, "/tensorflow.MasterService/PartialRunSetup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) RunStep(ctx context.Context, in *RunStepRequest, opts ...grpc.CallOption) (*RunStepResponse, error) {
	out := new(RunStepResponse)
	err := grpc.Invoke(ctx, "/tensorflow.MasterService/RunStep", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*CloseSessionResponse, error) {
	out := new(CloseSessionResponse)
	err := grpc.Invoke(ctx, "/tensorflow.MasterService/CloseSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) ListDevices(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error) {
	out := new(ListDevicesResponse)
	err := grpc.Invoke(ctx, "/tensorflow.MasterService/ListDevices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := grpc.Invoke(ctx, "/tensorflow.MasterService/Reset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MasterService service

type MasterServiceServer interface {
	// Creates a session.
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	// Extends a session.
	ExtendSession(context.Context, *ExtendSessionRequest) (*ExtendSessionResponse, error)
	// Prepares future partial run calls.
	PartialRunSetup(context.Context, *PartialRunSetupRequest) (*PartialRunSetupResponse, error)
	// Drives the graph computation.
	RunStep(context.Context, *RunStepRequest) (*RunStepResponse, error)
	// Closes a session.
	CloseSession(context.Context, *CloseSessionRequest) (*CloseSessionResponse, error)
	// List the devices usable by the master.
	ListDevices(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error)
	// Close and abandon all existing sessions.  Ongoing computations
	// will no longer affect fresh ones via the resources in containers listed in
	// the ResetRequest.  See ResetRequest for more details.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)
}

func RegisterMasterServiceServer(s *grpc.Server, srv MasterServiceServer) {
	s.RegisterService(&_MasterService_serviceDesc, srv)
}

func _MasterService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.MasterService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_ExtendSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtendSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).ExtendSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.MasterService/ExtendSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).ExtendSession(ctx, req.(*ExtendSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_PartialRunSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartialRunSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).PartialRunSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.MasterService/PartialRunSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).PartialRunSetup(ctx, req.(*PartialRunSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_RunStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).RunStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.MasterService/RunStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).RunStep(ctx, req.(*RunStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_CloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).CloseSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.MasterService/CloseSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).CloseSession(ctx, req.(*CloseSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.MasterService/ListDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).ListDevices(ctx, req.(*ListDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.MasterService/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MasterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _MasterService_CreateSession_Handler,
		},
		{
			MethodName: "ExtendSession",
			Handler:    _MasterService_ExtendSession_Handler,
		},
		{
			MethodName: "PartialRunSetup",
			Handler:    _MasterService_PartialRunSetup_Handler,
		},
		{
			MethodName: "RunStep",
			Handler:    _MasterService_RunStep_Handler,
		},
		{
			MethodName: "CloseSession",
			Handler:    _MasterService_CloseSession_Handler,
		},
		{
			MethodName: "ListDevices",
			Handler:    _MasterService_ListDevices_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _MasterService_Reset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master_service.proto",
}

func init() { proto.RegisterFile("master_service.proto", fileDescriptorMasterService) }

var fileDescriptorMasterService = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xbd, 0x4e, 0x84, 0x40,
	0x14, 0x85, 0x25, 0xf1, 0x27, 0x19, 0x77, 0x63, 0x32, 0x5a, 0x28, 0x26, 0xec, 0xaa, 0xfd, 0x16,
	0xda, 0x5a, 0x81, 0x76, 0x6a, 0x10, 0xac, 0x6c, 0x0c, 0xba, 0x57, 0x33, 0x09, 0x3b, 0x83, 0x73,
	0x2f, 0xea, 0x63, 0xf8, 0x40, 0x3e, 0x80, 0xa5, 0x8f, 0x60, 0xf0, 0x45, 0x0c, 0x30, 0xe8, 0xcc,
	0x86, 0xdd, 0xf6, 0x7c, 0xe7, 0x7c, 0x09, 0xc3, 0x65, 0x3b, 0xb3, 0x0c, 0x09, 0xf4, 0x1d, 0x82,
	0x7e, 0x11, 0x0f, 0x30, 0x29, 0xb4, 0x22, 0xc5, 0x19, 0x81, 0x44, 0xa5, 0x1f, 0x73, 0xf5, 0xea,
	0x0f, 0xda, 0x46, 0x4b, 0x8e, 0x3f, 0x56, 0xd9, 0xf0, 0xb2, 0x09, 0xd2, 0x76, 0xc1, 0x6f, 0xd8,
	0x30, 0xd2, 0x90, 0x11, 0xa4, 0x80, 0x28, 0x94, 0xe4, 0xe3, 0xc9, 0xff, 0x7a, 0xe2, 0xa0, 0x04,
	0x9e, 0x4b, 0x40, 0xf2, 0x0f, 0x96, 0x34, 0xb0, 0x50, 0x12, 0x1b, 0xeb, 0xf9, 0x1b, 0x81, 0x9c,
	0xf6, 0x5a, 0x1d, 0xd4, 0x6b, 0x9d, 0x6b, 0x18, 0xeb, 0x2d, 0xdb, 0x8a, 0x33, 0x4d, 0x22, 0xcb,
	0x93, 0x52, 0xa6, 0x40, 0x65, 0xc1, 0x0f, 0xed, 0xd5, 0x1c, 0xec, 0xcc, 0x47, 0x4b, 0x3b, 0xc6,
	0x1d, 0xb2, 0x8d, 0x3a, 0x23, 0x28, 0xb8, 0x6f, 0xf7, 0x4d, 0xd8, 0xb9, 0xf6, 0x7b, 0x99, 0x71,
	0x5c, 0xb3, 0x41, 0x94, 0x2b, 0xfc, 0x7b, 0xca, 0x91, 0xf3, 0x50, 0x16, 0xe9, 0x6c, 0xe3, 0xc5,
	0x05, 0xa3, 0xbc, 0x62, 0x9b, 0x17, 0x02, 0xe9, 0x0c, 0xea, 0x9f, 0x85, 0x3c, 0xb0, 0x07, 0x16,
	0xe8, 0x84, 0xa3, 0x85, 0xdc, 0xf8, 0x4e, 0xd9, 0x5a, 0x02, 0x08, 0xc4, 0x77, 0x9d, 0x0f, 0xa9,
	0xa3, 0xce, 0xb1, 0xd7, 0x43, 0xda, 0x75, 0x18, 0x7d, 0x56, 0x81, 0xf7, 0x55, 0x05, 0xde, 0x77,
	0x15, 0x78, 0xef, 0x3f, 0xc1, 0x0a, 0xf3, 0x95, 0x7e, 0xb2, 0xfb, 0x53, 0x81, 0xa4, 0x4b, 0x49,
	0x62, 0x06, 0xe1, 0xb6, 0x73, 0x69, 0x71, 0x7d, 0x80, 0x18, 0x7b, 0xf7, 0xeb, 0xcd, 0x29, 0x9e,
	0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x6b, 0x9a, 0xba, 0xbc, 0x02, 0x00, 0x00,
}
