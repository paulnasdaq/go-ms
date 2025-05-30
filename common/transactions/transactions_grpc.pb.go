// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/transactions.proto

package transactions

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TransactionsService_AddTransaction_FullMethodName = "/transactions.TransactionsService/AddTransaction"
)

// TransactionsServiceClient is the client API for TransactionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionsServiceClient interface {
	AddTransaction(ctx context.Context, in *AddTransactionRequest, opts ...grpc.CallOption) (*AddTransactionResponse, error)
}

type transactionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionsServiceClient(cc grpc.ClientConnInterface) TransactionsServiceClient {
	return &transactionsServiceClient{cc}
}

func (c *transactionsServiceClient) AddTransaction(ctx context.Context, in *AddTransactionRequest, opts ...grpc.CallOption) (*AddTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionsService_AddTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionsServiceServer is the server API for TransactionsService service.
// All implementations must embed UnimplementedTransactionsServiceServer
// for forward compatibility.
type TransactionsServiceServer interface {
	AddTransaction(context.Context, *AddTransactionRequest) (*AddTransactionResponse, error)
	mustEmbedUnimplementedTransactionsServiceServer()
}

// UnimplementedTransactionsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransactionsServiceServer struct{}

func (UnimplementedTransactionsServiceServer) AddTransaction(context.Context, *AddTransactionRequest) (*AddTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransaction not implemented")
}
func (UnimplementedTransactionsServiceServer) mustEmbedUnimplementedTransactionsServiceServer() {}
func (UnimplementedTransactionsServiceServer) testEmbeddedByValue()                             {}

// UnsafeTransactionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionsServiceServer will
// result in compilation errors.
type UnsafeTransactionsServiceServer interface {
	mustEmbedUnimplementedTransactionsServiceServer()
}

func RegisterTransactionsServiceServer(s grpc.ServiceRegistrar, srv TransactionsServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransactionsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransactionsService_ServiceDesc, srv)
}

func _TransactionsService_AddTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).AddTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionsService_AddTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).AddTransaction(ctx, req.(*AddTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionsService_ServiceDesc is the grpc.ServiceDesc for TransactionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transactions.TransactionsService",
	HandlerType: (*TransactionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTransaction",
			Handler:    _TransactionsService_AddTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transactions.proto",
}
