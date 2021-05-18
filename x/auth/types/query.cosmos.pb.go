// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

import (
	context "context"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Accounts returns all the existing accounts
	Accounts(ctx context.Context, in *QueryAccountsRequest, opts ...grpc.CallOption) (*QueryAccountsResponse, error)
	// Account returns account details based on address.
	Account(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*QueryAccountResponse, error)
	// Params queries all parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc        grpc.ClientConnInterface
	_Accounts types.Invoker
	_Account  types.Invoker
	_Params   types.Invoker
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc: cc}
}

func (c *queryClient) Accounts(ctx context.Context, in *QueryAccountsRequest, opts ...grpc.CallOption) (*QueryAccountsResponse, error) {
	if invoker := c._Accounts; invoker != nil {
		var out QueryAccountsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Accounts, err = invokerConn.Invoker("/cosmos.auth.v1beta1.Query/Accounts")
		if err != nil {
			var out QueryAccountsResponse
			err = c._Accounts(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryAccountsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.auth.v1beta1.Query/Accounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Account(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*QueryAccountResponse, error) {
	if invoker := c._Account; invoker != nil {
		var out QueryAccountResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Account, err = invokerConn.Invoker("/cosmos.auth.v1beta1.Query/Account")
		if err != nil {
			var out QueryAccountResponse
			err = c._Account(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryAccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.auth.v1beta1.Query/Account", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	if invoker := c._Params; invoker != nil {
		var out QueryParamsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Params, err = invokerConn.Invoker("/cosmos.auth.v1beta1.Query/Params")
		if err != nil {
			var out QueryParamsResponse
			err = c._Params(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.auth.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Accounts returns all the existing accounts
	Accounts(context.Context, *QueryAccountsRequest) (*QueryAccountsResponse, error)
	// Account returns account details based on address.
	Account(context.Context, *QueryAccountRequest) (*QueryAccountResponse, error)
	// Params queries all parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Accounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Accounts(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.auth.v1beta1.QueryAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Accounts(types.UnwrapSDKContext(ctx), req.(*QueryAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Account(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.auth.v1beta1.QueryAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Account(types.UnwrapSDKContext(ctx), req.(*QueryAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.auth.v1beta1.QueryParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(types.UnwrapSDKContext(ctx), req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.auth.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Accounts",
			Handler:    _Query_Accounts_Handler,
		},
		{
			MethodName: "Account",
			Handler:    _Query_Account_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Metadata: "cosmos/auth/v1beta1/query.proto",
}

const (
	QueryAccountsMethod = "/cosmos.auth.v1beta1.Query/Accounts"
	QueryAccountMethod  = "/cosmos.auth.v1beta1.Query/Account"
	QueryParamsMethod   = "/cosmos.auth.v1beta1.Query/Params"
)