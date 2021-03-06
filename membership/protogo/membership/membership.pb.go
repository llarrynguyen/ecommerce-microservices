// Code generated by protoc-gen-go. DO NOT EDIT.
// source: membership.proto

/*
Package membership is a generated protocol buffer package.

It is generated from these files:
	membership.proto

It has these top-level messages:
	QueryRequest
	MembershipRequest
	MembershipResponse
*/
package membership

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QueryRequest struct {
	ID    string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=Email" json:"Email,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *QueryRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type MembershipRequest struct {
	ID           string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	FirstName    string `protobuf:"bytes,2,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName     string `protobuf:"bytes,3,opt,name=LastName" json:"LastName,omitempty"`
	Email        string `protobuf:"bytes,4,opt,name=Email" json:"Email,omitempty"`
	Password     string `protobuf:"bytes,5,opt,name=Password" json:"Password,omitempty"`
	PasswordSalt string `protobuf:"bytes,6,opt,name=PasswordSalt" json:"PasswordSalt,omitempty"`
	BirthDate    string `protobuf:"bytes,7,opt,name=BirthDate" json:"BirthDate,omitempty"`
}

func (m *MembershipRequest) Reset()                    { *m = MembershipRequest{} }
func (m *MembershipRequest) String() string            { return proto.CompactTextString(m) }
func (*MembershipRequest) ProtoMessage()               {}
func (*MembershipRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MembershipRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *MembershipRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *MembershipRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *MembershipRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *MembershipRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *MembershipRequest) GetPasswordSalt() string {
	if m != nil {
		return m.PasswordSalt
	}
	return ""
}

func (m *MembershipRequest) GetBirthDate() string {
	if m != nil {
		return m.BirthDate
	}
	return ""
}

type MembershipResponse struct {
	ID           string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	FirstName    string `protobuf:"bytes,2,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName     string `protobuf:"bytes,3,opt,name=LastName" json:"LastName,omitempty"`
	Email        string `protobuf:"bytes,4,opt,name=Email" json:"Email,omitempty"`
	Password     string `protobuf:"bytes,5,opt,name=Password" json:"Password,omitempty"`
	PasswordSalt string `protobuf:"bytes,6,opt,name=PasswordSalt" json:"PasswordSalt,omitempty"`
	BirthDate    string `protobuf:"bytes,7,opt,name=BirthDate" json:"BirthDate,omitempty"`
	Version      int32  `protobuf:"varint,8,opt,name=Version" json:"Version,omitempty"`
	CreatedAt    string `protobuf:"bytes,9,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
	UpdatedAt    string `protobuf:"bytes,10,opt,name=UpdatedAt" json:"UpdatedAt,omitempty"`
}

func (m *MembershipResponse) Reset()                    { *m = MembershipResponse{} }
func (m *MembershipResponse) String() string            { return proto.CompactTextString(m) }
func (*MembershipResponse) ProtoMessage()               {}
func (*MembershipResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MembershipResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *MembershipResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *MembershipResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *MembershipResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *MembershipResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *MembershipResponse) GetPasswordSalt() string {
	if m != nil {
		return m.PasswordSalt
	}
	return ""
}

func (m *MembershipResponse) GetBirthDate() string {
	if m != nil {
		return m.BirthDate
	}
	return ""
}

func (m *MembershipResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *MembershipResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *MembershipResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "membership.QueryRequest")
	proto.RegisterType((*MembershipRequest)(nil), "membership.MembershipRequest")
	proto.RegisterType((*MembershipResponse)(nil), "membership.MembershipResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MembershipService service

type MembershipServiceClient interface {
	Save(ctx context.Context, in *MembershipRequest, opts ...grpc.CallOption) (*MembershipResponse, error)
	FindByID(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (MembershipService_FindByIDClient, error)
	FindByEmail(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (MembershipService_FindByEmailClient, error)
}

type membershipServiceClient struct {
	cc *grpc.ClientConn
}

func NewMembershipServiceClient(cc *grpc.ClientConn) MembershipServiceClient {
	return &membershipServiceClient{cc}
}

func (c *membershipServiceClient) Save(ctx context.Context, in *MembershipRequest, opts ...grpc.CallOption) (*MembershipResponse, error) {
	out := new(MembershipResponse)
	err := grpc.Invoke(ctx, "/membership.MembershipService/Save", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipServiceClient) FindByID(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (MembershipService_FindByIDClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MembershipService_serviceDesc.Streams[0], c.cc, "/membership.MembershipService/FindByID", opts...)
	if err != nil {
		return nil, err
	}
	x := &membershipServiceFindByIDClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MembershipService_FindByIDClient interface {
	Recv() (*MembershipResponse, error)
	grpc.ClientStream
}

type membershipServiceFindByIDClient struct {
	grpc.ClientStream
}

func (x *membershipServiceFindByIDClient) Recv() (*MembershipResponse, error) {
	m := new(MembershipResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *membershipServiceClient) FindByEmail(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (MembershipService_FindByEmailClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MembershipService_serviceDesc.Streams[1], c.cc, "/membership.MembershipService/FindByEmail", opts...)
	if err != nil {
		return nil, err
	}
	x := &membershipServiceFindByEmailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MembershipService_FindByEmailClient interface {
	Recv() (*MembershipResponse, error)
	grpc.ClientStream
}

type membershipServiceFindByEmailClient struct {
	grpc.ClientStream
}

func (x *membershipServiceFindByEmailClient) Recv() (*MembershipResponse, error) {
	m := new(MembershipResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MembershipService service

type MembershipServiceServer interface {
	Save(context.Context, *MembershipRequest) (*MembershipResponse, error)
	FindByID(*QueryRequest, MembershipService_FindByIDServer) error
	FindByEmail(*QueryRequest, MembershipService_FindByEmailServer) error
}

func RegisterMembershipServiceServer(s *grpc.Server, srv MembershipServiceServer) {
	s.RegisterService(&_MembershipService_serviceDesc, srv)
}

func _MembershipService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.MembershipService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServiceServer).Save(ctx, req.(*MembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MembershipService_FindByID_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MembershipServiceServer).FindByID(m, &membershipServiceFindByIDServer{stream})
}

type MembershipService_FindByIDServer interface {
	Send(*MembershipResponse) error
	grpc.ServerStream
}

type membershipServiceFindByIDServer struct {
	grpc.ServerStream
}

func (x *membershipServiceFindByIDServer) Send(m *MembershipResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MembershipService_FindByEmail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MembershipServiceServer).FindByEmail(m, &membershipServiceFindByEmailServer{stream})
}

type MembershipService_FindByEmailServer interface {
	Send(*MembershipResponse) error
	grpc.ServerStream
}

type membershipServiceFindByEmailServer struct {
	grpc.ServerStream
}

func (x *membershipServiceFindByEmailServer) Send(m *MembershipResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MembershipService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "membership.MembershipService",
	HandlerType: (*MembershipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _MembershipService_Save_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindByID",
			Handler:       _MembershipService_FindByID_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindByEmail",
			Handler:       _MembershipService_FindByEmail_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "membership.proto",
}

func init() { proto.RegisterFile("membership.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x25, 0xf9, 0x9a, 0x36, 0xb9, 0x5f, 0x11, 0x1d, 0x5c, 0x0c, 0x45, 0xa5, 0x64, 0xd5, 0x55,
	0x11, 0xf5, 0x05, 0xac, 0x69, 0x20, 0xa0, 0xa2, 0x09, 0xba, 0x9f, 0x9a, 0x0b, 0x1d, 0x68, 0x7e,
	0x9c, 0x99, 0x56, 0xfa, 0x46, 0xbe, 0x8e, 0x0f, 0xe2, 0x3b, 0x48, 0x32, 0xf9, 0x2b, 0x52, 0x5c,
	0xb8, 0x72, 0x97, 0x73, 0x4e, 0xce, 0xb9, 0x3f, 0xb9, 0x81, 0xc3, 0x04, 0x93, 0x05, 0x0a, 0xb9,
	0xe4, 0xf9, 0x34, 0x17, 0x99, 0xca, 0x08, 0xb4, 0x8c, 0x7b, 0x05, 0xc3, 0xc7, 0x35, 0x8a, 0x6d,
	0x88, 0xaf, 0x6b, 0x94, 0x8a, 0x1c, 0x80, 0x19, 0x78, 0xd4, 0x18, 0x1b, 0x13, 0x27, 0x34, 0x03,
	0x8f, 0x1c, 0x83, 0x35, 0x4f, 0x18, 0x5f, 0x51, 0xb3, 0xa4, 0x34, 0x70, 0x3f, 0x0c, 0x38, 0xba,
	0x6b, 0x42, 0xf6, 0x79, 0x4f, 0xc0, 0xf1, 0xb9, 0x90, 0xea, 0x9e, 0x25, 0x58, 0xf9, 0x5b, 0x82,
	0x8c, 0xc0, 0xbe, 0x65, 0x95, 0xf8, 0xaf, 0x14, 0x1b, 0xdc, 0x56, 0xed, 0x75, 0xaa, 0x16, 0x8e,
	0x07, 0x26, 0xe5, 0x5b, 0x26, 0x62, 0x6a, 0x69, 0x47, 0x8d, 0x89, 0x0b, 0xc3, 0xfa, 0x39, 0x62,
	0x2b, 0x45, 0xfb, 0xa5, 0xbe, 0xc3, 0x15, 0xfd, 0xcc, 0xb8, 0x50, 0x4b, 0x8f, 0x29, 0xa4, 0x03,
	0xdd, 0x4f, 0x43, 0xb8, 0xef, 0x26, 0x90, 0xee, 0x4c, 0x32, 0xcf, 0x52, 0x89, 0x7f, 0x7b, 0x28,
	0x42, 0x61, 0xf0, 0x8c, 0x42, 0xf2, 0x2c, 0xa5, 0xf6, 0xd8, 0x98, 0x58, 0x61, 0x0d, 0x0b, 0xdf,
	0x8d, 0x40, 0xa6, 0x30, 0xbe, 0x56, 0xd4, 0xd1, 0xbe, 0x86, 0x28, 0xd4, 0xa7, 0x3c, 0xae, 0x54,
	0xd0, 0x6a, 0x43, 0x5c, 0x7c, 0xee, 0x7c, 0xfe, 0x08, 0xc5, 0x86, 0xbf, 0x20, 0x99, 0x43, 0x2f,
	0x62, 0x1b, 0x24, 0xa7, 0xd3, 0xce, 0xc5, 0x7d, 0xbb, 0x92, 0xd1, 0xd9, 0x3e, 0xb9, 0x5a, 0xb8,
	0x0f, 0xb6, 0xcf, 0xd3, 0x78, 0xb6, 0x0d, 0x3c, 0x42, 0xbb, 0xef, 0x76, 0xef, 0xf4, 0xa7, 0x94,
	0x73, 0x83, 0x04, 0xf0, 0x5f, 0xe7, 0xe8, 0x3d, 0xff, 0x22, 0x6a, 0xd1, 0x2f, 0xff, 0x9b, 0xcb,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x04, 0x92, 0x0c, 0x4b, 0x03, 0x00, 0x00,
}
