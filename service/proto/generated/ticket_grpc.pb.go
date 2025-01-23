// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.3
// source: service/proto/ticket.proto

package generated

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

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceClient interface {
	SubmitRequest(ctx context.Context, in *SubmitTicketRequest, opts ...grpc.CallOption) (*TicketReceipt, error)
	GetDetails(ctx context.Context, in *GetDetailsRequest, opts ...grpc.CallOption) (*TicketReceipt, error)
	GetUsersInSection(ctx context.Context, in *GetUsersInSectionRequest, opts ...grpc.CallOption) (*UsersInSection, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	ModifyUserSeat(ctx context.Context, in *ModifyUserSeatRequest, opts ...grpc.CallOption) (*ModifyUserSeatResponse, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) SubmitRequest(ctx context.Context, in *SubmitTicketRequest, opts ...grpc.CallOption) (*TicketReceipt, error) {
	out := new(TicketReceipt)
	err := c.cc.Invoke(ctx, "/ticket.TicketService/SubmitRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetDetails(ctx context.Context, in *GetDetailsRequest, opts ...grpc.CallOption) (*TicketReceipt, error) {
	out := new(TicketReceipt)
	err := c.cc.Invoke(ctx, "/ticket.TicketService/GetDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetUsersInSection(ctx context.Context, in *GetUsersInSectionRequest, opts ...grpc.CallOption) (*UsersInSection, error) {
	out := new(UsersInSection)
	err := c.cc.Invoke(ctx, "/ticket.TicketService/GetUsersInSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, "/ticket.TicketService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ModifyUserSeat(ctx context.Context, in *ModifyUserSeatRequest, opts ...grpc.CallOption) (*ModifyUserSeatResponse, error) {
	out := new(ModifyUserSeatResponse)
	err := c.cc.Invoke(ctx, "/ticket.TicketService/ModifyUserSeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations must embed UnimplementedTicketServiceServer
// for forward compatibility
type TicketServiceServer interface {
	SubmitRequest(context.Context, *SubmitTicketRequest) (*TicketReceipt, error)
	GetDetails(context.Context, *GetDetailsRequest) (*TicketReceipt, error)
	GetUsersInSection(context.Context, *GetUsersInSectionRequest) (*UsersInSection, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	ModifyUserSeat(context.Context, *ModifyUserSeatRequest) (*ModifyUserSeatResponse, error)
	mustEmbedUnimplementedTicketServiceServer()
}

// UnimplementedTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTicketServiceServer struct {
}

func (UnimplementedTicketServiceServer) SubmitRequest(context.Context, *SubmitTicketRequest) (*TicketReceipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitRequest not implemented")
}
func (UnimplementedTicketServiceServer) GetDetails(context.Context, *GetDetailsRequest) (*TicketReceipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetails not implemented")
}
func (UnimplementedTicketServiceServer) GetUsersInSection(context.Context, *GetUsersInSectionRequest) (*UsersInSection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersInSection not implemented")
}
func (UnimplementedTicketServiceServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTicketServiceServer) ModifyUserSeat(context.Context, *ModifyUserSeatRequest) (*ModifyUserSeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserSeat not implemented")
}
func (UnimplementedTicketServiceServer) mustEmbedUnimplementedTicketServiceServer() {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_SubmitRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).SubmitRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket.TicketService/SubmitRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).SubmitRequest(ctx, req.(*SubmitTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket.TicketService/GetDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetDetails(ctx, req.(*GetDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetUsersInSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersInSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetUsersInSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket.TicketService/GetUsersInSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetUsersInSection(ctx, req.(*GetUsersInSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket.TicketService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ModifyUserSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUserSeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ModifyUserSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket.TicketService/ModifyUserSeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ModifyUserSeat(ctx, req.(*ModifyUserSeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticket.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitRequest",
			Handler:    _TicketService_SubmitRequest_Handler,
		},
		{
			MethodName: "GetDetails",
			Handler:    _TicketService_GetDetails_Handler,
		},
		{
			MethodName: "GetUsersInSection",
			Handler:    _TicketService_GetUsersInSection_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TicketService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifyUserSeat",
			Handler:    _TicketService_ModifyUserSeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/proto/ticket.proto",
}
