// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: wordle/wordle/query.proto

package wordle

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
	Query_Params_FullMethodName    = "/wordle.wordle.Query/Params"
	Query_Wordle_FullMethodName    = "/wordle.wordle.Query/Wordle"
	Query_WordleAll_FullMethodName = "/wordle.wordle.Query/WordleAll"
	Query_Guess_FullMethodName     = "/wordle.wordle.Query/Guess"
	Query_GuessAll_FullMethodName  = "/wordle.wordle.Query/GuessAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Wordle items.
	Wordle(ctx context.Context, in *QueryGetWordleRequest, opts ...grpc.CallOption) (*QueryGetWordleResponse, error)
	WordleAll(ctx context.Context, in *QueryAllWordleRequest, opts ...grpc.CallOption) (*QueryAllWordleResponse, error)
	// Queries a list of Guess items.
	Guess(ctx context.Context, in *QueryGetGuessRequest, opts ...grpc.CallOption) (*QueryGetGuessResponse, error)
	GuessAll(ctx context.Context, in *QueryAllGuessRequest, opts ...grpc.CallOption) (*QueryAllGuessResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Wordle(ctx context.Context, in *QueryGetWordleRequest, opts ...grpc.CallOption) (*QueryGetWordleResponse, error) {
	out := new(QueryGetWordleResponse)
	err := c.cc.Invoke(ctx, Query_Wordle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) WordleAll(ctx context.Context, in *QueryAllWordleRequest, opts ...grpc.CallOption) (*QueryAllWordleResponse, error) {
	out := new(QueryAllWordleResponse)
	err := c.cc.Invoke(ctx, Query_WordleAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Guess(ctx context.Context, in *QueryGetGuessRequest, opts ...grpc.CallOption) (*QueryGetGuessResponse, error) {
	out := new(QueryGetGuessResponse)
	err := c.cc.Invoke(ctx, Query_Guess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GuessAll(ctx context.Context, in *QueryAllGuessRequest, opts ...grpc.CallOption) (*QueryAllGuessResponse, error) {
	out := new(QueryAllGuessResponse)
	err := c.cc.Invoke(ctx, Query_GuessAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Wordle items.
	Wordle(context.Context, *QueryGetWordleRequest) (*QueryGetWordleResponse, error)
	WordleAll(context.Context, *QueryAllWordleRequest) (*QueryAllWordleResponse, error)
	// Queries a list of Guess items.
	Guess(context.Context, *QueryGetGuessRequest) (*QueryGetGuessResponse, error)
	GuessAll(context.Context, *QueryAllGuessRequest) (*QueryAllGuessResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Wordle(context.Context, *QueryGetWordleRequest) (*QueryGetWordleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Wordle not implemented")
}
func (UnimplementedQueryServer) WordleAll(context.Context, *QueryAllWordleRequest) (*QueryAllWordleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WordleAll not implemented")
}
func (UnimplementedQueryServer) Guess(context.Context, *QueryGetGuessRequest) (*QueryGetGuessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Guess not implemented")
}
func (UnimplementedQueryServer) GuessAll(context.Context, *QueryAllGuessRequest) (*QueryAllGuessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuessAll not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Wordle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetWordleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Wordle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Wordle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Wordle(ctx, req.(*QueryGetWordleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_WordleAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllWordleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WordleAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_WordleAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WordleAll(ctx, req.(*QueryAllWordleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Guess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetGuessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Guess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Guess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Guess(ctx, req.(*QueryGetGuessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GuessAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllGuessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GuessAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GuessAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GuessAll(ctx, req.(*QueryAllGuessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wordle.wordle.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Wordle",
			Handler:    _Query_Wordle_Handler,
		},
		{
			MethodName: "WordleAll",
			Handler:    _Query_WordleAll_Handler,
		},
		{
			MethodName: "Guess",
			Handler:    _Query_Guess_Handler,
		},
		{
			MethodName: "GuessAll",
			Handler:    _Query_GuessAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wordle/wordle/query.proto",
}
