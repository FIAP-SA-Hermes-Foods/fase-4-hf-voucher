// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: voucher.proto

package voucher_proto

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

// VoucherClient is the client API for Voucher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VoucherClient interface {
	CreateVoucher(ctx context.Context, in *CreateVoucherRequest, opts ...grpc.CallOption) (*CreateVoucherResponse, error)
	GetVoucherByID(ctx context.Context, in *GetVoucherByIDRequest, opts ...grpc.CallOption) (*GetVoucherByIDResponse, error)
	UpdateVoucherByID(ctx context.Context, in *UpdateVoucherByIDRequest, opts ...grpc.CallOption) (*UpdateVoucherByIDResponse, error)
}

type voucherClient struct {
	cc grpc.ClientConnInterface
}

func NewVoucherClient(cc grpc.ClientConnInterface) VoucherClient {
	return &voucherClient{cc}
}

func (c *voucherClient) CreateVoucher(ctx context.Context, in *CreateVoucherRequest, opts ...grpc.CallOption) (*CreateVoucherResponse, error) {
	out := new(CreateVoucherResponse)
	err := c.cc.Invoke(ctx, "/Voucher/CreateVoucher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voucherClient) GetVoucherByID(ctx context.Context, in *GetVoucherByIDRequest, opts ...grpc.CallOption) (*GetVoucherByIDResponse, error) {
	out := new(GetVoucherByIDResponse)
	err := c.cc.Invoke(ctx, "/Voucher/GetVoucherByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voucherClient) UpdateVoucherByID(ctx context.Context, in *UpdateVoucherByIDRequest, opts ...grpc.CallOption) (*UpdateVoucherByIDResponse, error) {
	out := new(UpdateVoucherByIDResponse)
	err := c.cc.Invoke(ctx, "/Voucher/UpdateVoucherByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoucherServer is the server API for Voucher service.
// All implementations must embed UnimplementedVoucherServer
// for forward compatibility
type VoucherServer interface {
	CreateVoucher(context.Context, *CreateVoucherRequest) (*CreateVoucherResponse, error)
	GetVoucherByID(context.Context, *GetVoucherByIDRequest) (*GetVoucherByIDResponse, error)
	UpdateVoucherByID(context.Context, *UpdateVoucherByIDRequest) (*UpdateVoucherByIDResponse, error)
	mustEmbedUnimplementedVoucherServer()
}

// UnimplementedVoucherServer must be embedded to have forward compatible implementations.
type UnimplementedVoucherServer struct {
}

func (UnimplementedVoucherServer) CreateVoucher(context.Context, *CreateVoucherRequest) (*CreateVoucherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVoucher not implemented")
}
func (UnimplementedVoucherServer) GetVoucherByID(context.Context, *GetVoucherByIDRequest) (*GetVoucherByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVoucherByID not implemented")
}
func (UnimplementedVoucherServer) UpdateVoucherByID(context.Context, *UpdateVoucherByIDRequest) (*UpdateVoucherByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVoucherByID not implemented")
}
func (UnimplementedVoucherServer) mustEmbedUnimplementedVoucherServer() {}

// UnsafeVoucherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VoucherServer will
// result in compilation errors.
type UnsafeVoucherServer interface {
	mustEmbedUnimplementedVoucherServer()
}

func RegisterVoucherServer(s grpc.ServiceRegistrar, srv VoucherServer) {
	s.RegisterService(&Voucher_ServiceDesc, srv)
}

func _Voucher_CreateVoucher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVoucherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoucherServer).CreateVoucher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Voucher/CreateVoucher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoucherServer).CreateVoucher(ctx, req.(*CreateVoucherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Voucher_GetVoucherByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVoucherByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoucherServer).GetVoucherByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Voucher/GetVoucherByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoucherServer).GetVoucherByID(ctx, req.(*GetVoucherByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Voucher_UpdateVoucherByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVoucherByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoucherServer).UpdateVoucherByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Voucher/UpdateVoucherByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoucherServer).UpdateVoucherByID(ctx, req.(*UpdateVoucherByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Voucher_ServiceDesc is the grpc.ServiceDesc for Voucher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Voucher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Voucher",
	HandlerType: (*VoucherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVoucher",
			Handler:    _Voucher_CreateVoucher_Handler,
		},
		{
			MethodName: "GetVoucherByID",
			Handler:    _Voucher_GetVoucherByID_Handler,
		},
		{
			MethodName: "UpdateVoucherByID",
			Handler:    _Voucher_UpdateVoucherByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voucher.proto",
}
