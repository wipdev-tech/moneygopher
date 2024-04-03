// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: transactions/transactions.proto

package transactions

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

// TransactionsClient is the client API for Transactions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionsClient interface {
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error)
}

type transactionsClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionsClient(cc grpc.ClientConnInterface) TransactionsClient {
	return &transactionsClient{cc}
}

func (c *transactionsClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error) {
	out := new(DepositResponse)
	err := c.cc.Invoke(ctx, "/Transactions/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionsServer is the server API for Transactions service.
// All implementations must embed UnimplementedTransactionsServer
// for forward compatibility
type TransactionsServer interface {
	Deposit(context.Context, *DepositRequest) (*DepositResponse, error)
	mustEmbedUnimplementedTransactionsServer()
}

// UnimplementedTransactionsServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionsServer struct {
}

func (UnimplementedTransactionsServer) Deposit(context.Context, *DepositRequest) (*DepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedTransactionsServer) mustEmbedUnimplementedTransactionsServer() {}

// UnsafeTransactionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionsServer will
// result in compilation errors.
type UnsafeTransactionsServer interface {
	mustEmbedUnimplementedTransactionsServer()
}

func RegisterTransactionsServer(s grpc.ServiceRegistrar, srv TransactionsServer) {
	s.RegisterService(&Transactions_ServiceDesc, srv)
}

func _Transactions_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Transactions/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transactions_ServiceDesc is the grpc.ServiceDesc for Transactions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transactions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Transactions",
	HandlerType: (*TransactionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _Transactions_Deposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transactions/transactions.proto",
}