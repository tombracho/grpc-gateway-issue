// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: userservice/userservice.proto

package volunteering

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
	VolunteeringHttp_Login_FullMethodName = "/volunteering_http.VolunteeringHttp/Login"
)

// VolunteeringHttpClient is the client API for VolunteeringHttp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VolunteeringHttpClient interface {
	// Login
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// ## {{.RequestType.Name}}
	// | Field ID    | Name      | Type                                                       | Description                  |
	// | ----------- | --------- | ---------------------------------------------------------  | ---------------------------- | {{range .RequestType.Fields}}
	// | {{.Number}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	// ## {{.ResponseType.Name}}
	// | Field ID    | Name      | Type                                                       | Description                  |
	// | ----------- | --------- | ---------------------------------------------------------- | ---------------------------- | {{range .ResponseType.Fields}}
	// | {{.Number}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	Login(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error)
}

type volunteeringHttpClient struct {
	cc grpc.ClientConnInterface
}

func NewVolunteeringHttpClient(cc grpc.ClientConnInterface) VolunteeringHttpClient {
	return &volunteeringHttpClient{cc}
}

func (c *volunteeringHttpClient) Login(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, VolunteeringHttp_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VolunteeringHttpServer is the server API for VolunteeringHttp service.
// All implementations must embed UnimplementedVolunteeringHttpServer
// for forward compatibility.
type VolunteeringHttpServer interface {
	// Login
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// ## {{.RequestType.Name}}
	// | Field ID    | Name      | Type                                                       | Description                  |
	// | ----------- | --------- | ---------------------------------------------------------  | ---------------------------- | {{range .RequestType.Fields}}
	// | {{.Number}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	// ## {{.ResponseType.Name}}
	// | Field ID    | Name      | Type                                                       | Description                  |
	// | ----------- | --------- | ---------------------------------------------------------- | ---------------------------- | {{range .ResponseType.Fields}}
	// | {{.Number}} | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	Login(context.Context, *UserId) (*User, error)
	mustEmbedUnimplementedVolunteeringHttpServer()
}

// UnimplementedVolunteeringHttpServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVolunteeringHttpServer struct{}

func (UnimplementedVolunteeringHttpServer) Login(context.Context, *UserId) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedVolunteeringHttpServer) mustEmbedUnimplementedVolunteeringHttpServer() {}
func (UnimplementedVolunteeringHttpServer) testEmbeddedByValue()                          {}

// UnsafeVolunteeringHttpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VolunteeringHttpServer will
// result in compilation errors.
type UnsafeVolunteeringHttpServer interface {
	mustEmbedUnimplementedVolunteeringHttpServer()
}

func RegisterVolunteeringHttpServer(s grpc.ServiceRegistrar, srv VolunteeringHttpServer) {
	// If the following call pancis, it indicates UnimplementedVolunteeringHttpServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VolunteeringHttp_ServiceDesc, srv)
}

func _VolunteeringHttp_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolunteeringHttpServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolunteeringHttp_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolunteeringHttpServer).Login(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

// VolunteeringHttp_ServiceDesc is the grpc.ServiceDesc for VolunteeringHttp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VolunteeringHttp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "volunteering_http.VolunteeringHttp",
	HandlerType: (*VolunteeringHttpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _VolunteeringHttp_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userservice/userservice.proto",
}
