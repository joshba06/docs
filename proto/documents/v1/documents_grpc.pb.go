// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: documents/v1/documents.proto

package documentsv1

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
	DocumentsService_AddDocument_FullMethodName   = "/documents.v1.DocumentsService/AddDocument"
	DocumentsService_GetDocument_FullMethodName   = "/documents.v1.DocumentsService/GetDocument"
	DocumentsService_ListDocuments_FullMethodName = "/documents.v1.DocumentsService/ListDocuments"
	DocumentsService_GetFile_FullMethodName       = "/documents.v1.DocumentsService/GetFile"
)

// DocumentsServiceClient is the client API for DocumentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentsServiceClient interface {
	AddDocument(ctx context.Context, in *AddDocumentRequest, opts ...grpc.CallOption) (*Document, error)
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*Document, error)
	ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (DocumentsService_ListDocumentsClient, error)
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*File, error)
}

type documentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentsServiceClient(cc grpc.ClientConnInterface) DocumentsServiceClient {
	return &documentsServiceClient{cc}
}

func (c *documentsServiceClient) AddDocument(ctx context.Context, in *AddDocumentRequest, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, DocumentsService_AddDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentsServiceClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, DocumentsService_GetDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentsServiceClient) ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (DocumentsService_ListDocumentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &DocumentsService_ServiceDesc.Streams[0], DocumentsService_ListDocuments_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &documentsServiceListDocumentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DocumentsService_ListDocumentsClient interface {
	Recv() (*Document, error)
	grpc.ClientStream
}

type documentsServiceListDocumentsClient struct {
	grpc.ClientStream
}

func (x *documentsServiceListDocumentsClient) Recv() (*Document, error) {
	m := new(Document)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *documentsServiceClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, DocumentsService_GetFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentsServiceServer is the server API for DocumentsService service.
// All implementations should embed UnimplementedDocumentsServiceServer
// for forward compatibility
type DocumentsServiceServer interface {
	AddDocument(context.Context, *AddDocumentRequest) (*Document, error)
	GetDocument(context.Context, *GetDocumentRequest) (*Document, error)
	ListDocuments(*ListDocumentsRequest, DocumentsService_ListDocumentsServer) error
	GetFile(context.Context, *GetFileRequest) (*File, error)
}

// UnimplementedDocumentsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDocumentsServiceServer struct {
}

func (UnimplementedDocumentsServiceServer) AddDocument(context.Context, *AddDocumentRequest) (*Document, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDocument not implemented")
}
func (UnimplementedDocumentsServiceServer) GetDocument(context.Context, *GetDocumentRequest) (*Document, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocument not implemented")
}
func (UnimplementedDocumentsServiceServer) ListDocuments(*ListDocumentsRequest, DocumentsService_ListDocumentsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListDocuments not implemented")
}
func (UnimplementedDocumentsServiceServer) GetFile(context.Context, *GetFileRequest) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}

// UnsafeDocumentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentsServiceServer will
// result in compilation errors.
type UnsafeDocumentsServiceServer interface {
	mustEmbedUnimplementedDocumentsServiceServer()
}

func RegisterDocumentsServiceServer(s grpc.ServiceRegistrar, srv DocumentsServiceServer) {
	s.RegisterService(&DocumentsService_ServiceDesc, srv)
}

func _DocumentsService_AddDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServiceServer).AddDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentsService_AddDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServiceServer).AddDocument(ctx, req.(*AddDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentsService_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServiceServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentsService_GetDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServiceServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentsService_ListDocuments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListDocumentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DocumentsServiceServer).ListDocuments(m, &documentsServiceListDocumentsServer{stream})
}

type DocumentsService_ListDocumentsServer interface {
	Send(*Document) error
	grpc.ServerStream
}

type documentsServiceListDocumentsServer struct {
	grpc.ServerStream
}

func (x *documentsServiceListDocumentsServer) Send(m *Document) error {
	return x.ServerStream.SendMsg(m)
}

func _DocumentsService_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentsServiceServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentsService_GetFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentsServiceServer).GetFile(ctx, req.(*GetFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentsService_ServiceDesc is the grpc.ServiceDesc for DocumentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "documents.v1.DocumentsService",
	HandlerType: (*DocumentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDocument",
			Handler:    _DocumentsService_AddDocument_Handler,
		},
		{
			MethodName: "GetDocument",
			Handler:    _DocumentsService_GetDocument_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _DocumentsService_GetFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListDocuments",
			Handler:       _DocumentsService_ListDocuments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "documents/v1/documents.proto",
}
