// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thetaradix.proto

package thetaradix

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterRequest struct {
	Mail                 string   `protobuf:"bytes,1,opt,name=mail,proto3" json:"mail,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_108ea66c53471098, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *RegisterRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	PleaseVerifyPassword bool     `protobuf:"varint,2,opt,name=pleaseVerifyPassword,proto3" json:"pleaseVerifyPassword,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,4,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReply) Reset()         { *m = RegisterReply{} }
func (m *RegisterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()    {}
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_108ea66c53471098, []int{1}
}

func (m *RegisterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReply.Unmarshal(m, b)
}
func (m *RegisterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReply.Marshal(b, m, deterministic)
}
func (m *RegisterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReply.Merge(m, src)
}
func (m *RegisterReply) XXX_Size() int {
	return xxx_messageInfo_RegisterReply.Size(m)
}
func (m *RegisterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReply proto.InternalMessageInfo

func (m *RegisterReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RegisterReply) GetPleaseVerifyPassword() bool {
	if m != nil {
		return m.PleaseVerifyPassword
	}
	return false
}

func (m *RegisterReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type LoginRequest struct {
	Mail                 string   `protobuf:"bytes,1,opt,name=mail,proto3" json:"mail,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_108ea66c53471098, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Jwt                  string   `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_108ea66c53471098, []int{3}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LoginReply) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func (m *LoginReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "thetaradix.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "thetaradix.RegisterReply")
	proto.RegisterType((*LoginRequest)(nil), "thetaradix.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "thetaradix.LoginReply")
}

func init() { proto.RegisterFile("thetaradix.proto", fileDescriptor_108ea66c53471098) }

var fileDescriptor_108ea66c53471098 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x55, 0x48, 0x80, 0x70, 0x2a, 0x22, 0xb2, 0x10, 0x32, 0x61, 0x41, 0x99, 0x98, 0x3a, 0x94,
	0x89, 0x85, 0x81, 0x11, 0x81, 0x84, 0x2c, 0xc4, 0xc6, 0x60, 0xd2, 0x23, 0x18, 0xf2, 0x61, 0x7c,
	0xae, 0x4a, 0x7e, 0x01, 0x03, 0x7f, 0x1a, 0xd9, 0x4d, 0x68, 0x40, 0x01, 0x36, 0xbf, 0x7b, 0xf7,
	0xf5, 0xee, 0x19, 0x12, 0xfb, 0x84, 0x56, 0x1a, 0x39, 0x57, 0x6f, 0x53, 0x6d, 0x1a, 0xdb, 0x30,
	0x58, 0x47, 0xb2, 0x7b, 0xd8, 0x13, 0x58, 0x28, 0xb2, 0x68, 0x04, 0xbe, 0x2e, 0x90, 0x2c, 0x63,
	0x10, 0x55, 0x52, 0x95, 0x3c, 0x38, 0x0e, 0x4e, 0x76, 0x84, 0x7f, 0xb3, 0x14, 0xe2, 0x5a, 0xe5,
	0x2f, 0xb5, 0xac, 0x90, 0x87, 0x3e, 0xfe, 0x85, 0x1d, 0xa7, 0x25, 0xd1, 0xb2, 0x31, 0x73, 0xbe,
	0xb1, 0xe2, 0x7a, 0x9c, 0xbd, 0x07, 0xb0, 0xbb, 0xee, 0xaf, 0xcb, 0x96, 0x71, 0xd8, 0xa6, 0x45,
	0x9e, 0x23, 0x91, 0x1f, 0x10, 0x8b, 0x1e, 0xb2, 0x19, 0xec, 0xeb, 0x12, 0x25, 0xe1, 0x1d, 0x1a,
	0xf5, 0xd8, 0xde, 0x0c, 0x7b, 0xc6, 0x62, 0x94, 0x63, 0x19, 0x4c, 0xd0, 0x98, 0xc6, 0x5c, 0x23,
	0x91, 0x2c, 0x90, 0x47, 0x7e, 0xfe, 0xb7, 0xd8, 0x65, 0x14, 0x87, 0x49, 0x94, 0x9d, 0xc3, 0xe4,
	0xaa, 0x29, 0x54, 0xfd, 0x8f, 0xca, 0x5f, 0x95, 0x08, 0x80, 0xae, 0xfe, 0x6f, 0x15, 0x09, 0x84,
	0xcf, 0x4b, 0xdb, 0x95, 0xbb, 0xa7, 0xcb, 0xad, 0xba, 0xf5, 0x56, 0xa7, 0xeb, 0xe1, 0xec, 0x23,
	0x00, 0xb8, 0x75, 0x5e, 0x08, 0xe7, 0x05, 0xbb, 0x80, 0xb8, 0xbf, 0x15, 0x3b, 0x9a, 0x0e, 0x6c,
	0xfb, 0xe1, 0x50, 0x7a, 0x38, 0x4e, 0xba, 0xc5, 0xce, 0x60, 0xd3, 0xaf, 0xc9, 0xf8, 0x30, 0x67,
	0xa8, 0x3c, 0x3d, 0x18, 0x61, 0x74, 0xd9, 0x3e, 0x6c, 0xf9, 0xdf, 0x71, 0xfa, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x91, 0x0b, 0x08, 0x97, 0x31, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ThetaRadixClient is the client API for ThetaRadix service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ThetaRadixClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type thetaRadixClient struct {
	cc *grpc.ClientConn
}

func NewThetaRadixClient(cc *grpc.ClientConn) ThetaRadixClient {
	return &thetaRadixClient{cc}
}

func (c *thetaRadixClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/thetaradix.ThetaRadix/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetaRadixClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/thetaradix.ThetaRadix/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThetaRadixServer is the server API for ThetaRadix service.
type ThetaRadixServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterThetaRadixServer(s *grpc.Server, srv ThetaRadixServer) {
	s.RegisterService(&_ThetaRadix_serviceDesc, srv)
}

func _ThetaRadix_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetaRadixServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thetaradix.ThetaRadix/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetaRadixServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetaRadix_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetaRadixServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thetaradix.ThetaRadix/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetaRadixServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ThetaRadix_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thetaradix.ThetaRadix",
	HandlerType: (*ThetaRadixServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ThetaRadix_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ThetaRadix_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetaradix.proto",
}
