// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0--rc1
// source: article.proto

package __

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
	GetMyArticles_GetMyArticles_FullMethodName = "/GetMyArticles/GetMyArticles"
)

// GetMyArticlesClient is the client API for GetMyArticles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetMyArticlesClient interface {
	GetMyArticles(ctx context.Context, in *GetMyArticleRequest, opts ...grpc.CallOption) (*ArticleResponse, error)
}

type getMyArticlesClient struct {
	cc grpc.ClientConnInterface
}

func NewGetMyArticlesClient(cc grpc.ClientConnInterface) GetMyArticlesClient {
	return &getMyArticlesClient{cc}
}

func (c *getMyArticlesClient) GetMyArticles(ctx context.Context, in *GetMyArticleRequest, opts ...grpc.CallOption) (*ArticleResponse, error) {
	out := new(ArticleResponse)
	err := c.cc.Invoke(ctx, GetMyArticles_GetMyArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetMyArticlesServer is the server API for GetMyArticles service.
// All implementations must embed UnimplementedGetMyArticlesServer
// for forward compatibility
type GetMyArticlesServer interface {
	GetMyArticles(context.Context, *GetMyArticleRequest) (*ArticleResponse, error)
	mustEmbedUnimplementedGetMyArticlesServer()
}

// UnimplementedGetMyArticlesServer must be embedded to have forward compatible implementations.
type UnimplementedGetMyArticlesServer struct {
}

func (UnimplementedGetMyArticlesServer) GetMyArticles(context.Context, *GetMyArticleRequest) (*ArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyArticles not implemented")
}
func (UnimplementedGetMyArticlesServer) mustEmbedUnimplementedGetMyArticlesServer() {}

// UnsafeGetMyArticlesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetMyArticlesServer will
// result in compilation errors.
type UnsafeGetMyArticlesServer interface {
	mustEmbedUnimplementedGetMyArticlesServer()
}

func RegisterGetMyArticlesServer(s grpc.ServiceRegistrar, srv GetMyArticlesServer) {
	s.RegisterService(&GetMyArticles_ServiceDesc, srv)
}

func _GetMyArticles_GetMyArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetMyArticlesServer).GetMyArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetMyArticles_GetMyArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetMyArticlesServer).GetMyArticles(ctx, req.(*GetMyArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetMyArticles_ServiceDesc is the grpc.ServiceDesc for GetMyArticles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetMyArticles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GetMyArticles",
	HandlerType: (*GetMyArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMyArticles",
			Handler:    _GetMyArticles_GetMyArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}
