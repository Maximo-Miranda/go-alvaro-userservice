// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go-alvaro-userservice.proto

package goalvaro

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// User
type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lastname             string   `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Phonenumber          string   `protobuf:"bytes,6,opt,name=phonenumber,proto3" json:"phonenumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_0eeaa148f71ae881, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetPhonenumber() string {
	if m != nil {
		return m.Phonenumber
	}
	return ""
}

type UserId struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_0eeaa148f71ae881, []int{1}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Response
type Response struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0eeaa148f71ae881, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// UserResponse
type UserResponse struct {
	Meta                 *Response `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data                 *User     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0eeaa148f71ae881, []int{3}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetMeta() *Response {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UserResponse) GetData() *User {
	if m != nil {
		return m.Data
	}
	return nil
}

// UserBoolResponse
type UserBoolResponse struct {
	Meta                 *Response `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data                 bool      `protobuf:"varint,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserBoolResponse) Reset()         { *m = UserBoolResponse{} }
func (m *UserBoolResponse) String() string { return proto.CompactTextString(m) }
func (*UserBoolResponse) ProtoMessage()    {}
func (*UserBoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0eeaa148f71ae881, []int{4}
}

func (m *UserBoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBoolResponse.Unmarshal(m, b)
}
func (m *UserBoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBoolResponse.Marshal(b, m, deterministic)
}
func (m *UserBoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBoolResponse.Merge(m, src)
}
func (m *UserBoolResponse) XXX_Size() int {
	return xxx_messageInfo_UserBoolResponse.Size(m)
}
func (m *UserBoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserBoolResponse proto.InternalMessageInfo

func (m *UserBoolResponse) GetMeta() *Response {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UserBoolResponse) GetData() bool {
	if m != nil {
		return m.Data
	}
	return false
}

// UsersResponse
type UsersResponse struct {
	Meta                 *Response `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data                 []*User   `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UsersResponse) Reset()         { *m = UsersResponse{} }
func (m *UsersResponse) String() string { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()    {}
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0eeaa148f71ae881, []int{5}
}

func (m *UsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersResponse.Unmarshal(m, b)
}
func (m *UsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersResponse.Marshal(b, m, deterministic)
}
func (m *UsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersResponse.Merge(m, src)
}
func (m *UsersResponse) XXX_Size() int {
	return xxx_messageInfo_UsersResponse.Size(m)
}
func (m *UsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersResponse proto.InternalMessageInfo

func (m *UsersResponse) GetMeta() *Response {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UsersResponse) GetData() []*User {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "goalvaro.User")
	proto.RegisterType((*UserId)(nil), "goalvaro.UserId")
	proto.RegisterType((*Response)(nil), "goalvaro.Response")
	proto.RegisterType((*UserResponse)(nil), "goalvaro.UserResponse")
	proto.RegisterType((*UserBoolResponse)(nil), "goalvaro.UserBoolResponse")
	proto.RegisterType((*UsersResponse)(nil), "goalvaro.UsersResponse")
}

func init() { proto.RegisterFile("go-alvaro-userservice.proto", fileDescriptor_0eeaa148f71ae881) }

var fileDescriptor_0eeaa148f71ae881 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x6b, 0xe2, 0x40,
	0x18, 0x25, 0x31, 0xba, 0xd9, 0x2f, 0xbb, 0x22, 0xc3, 0xe2, 0x86, 0x78, 0x91, 0x1c, 0x16, 0x2f,
	0x46, 0x70, 0x85, 0x5d, 0x76, 0x6f, 0xd6, 0x52, 0x84, 0xd2, 0x43, 0x8a, 0x97, 0xf6, 0x34, 0x3a,
	0x5f, 0xd3, 0x40, 0x92, 0x09, 0x33, 0xd1, 0xe2, 0x5f, 0xe9, 0x5f, 0xe8, 0x9f, 0x2c, 0x33, 0x31,
	0x6a, 0x2c, 0x1e, 0xa4, 0xb7, 0x79, 0xdf, 0xfb, 0xde, 0x9b, 0x97, 0xc9, 0x83, 0x5e, 0xc4, 0x87,
	0x34, 0xd9, 0x50, 0xc1, 0x87, 0x6b, 0x89, 0x42, 0xa2, 0xd8, 0xc4, 0x2b, 0x0c, 0x72, 0xc1, 0x0b,
	0x4e, 0xec, 0x88, 0x97, 0x9c, 0xd7, 0x8b, 0x38, 0x8f, 0x12, 0x1c, 0xe9, 0xf9, 0x72, 0xfd, 0x34,
	0xc2, 0x34, 0x2f, 0xb6, 0xe5, 0x9a, 0xff, 0x6a, 0x80, 0xb5, 0x90, 0x28, 0x48, 0x1b, 0xcc, 0x98,
	0xb9, 0x46, 0xdf, 0x18, 0x34, 0x42, 0x33, 0x66, 0x84, 0x80, 0x95, 0xd1, 0x14, 0x5d, 0xb3, 0x6f,
	0x0c, 0xbe, 0x86, 0xfa, 0x4c, 0x3c, 0xb0, 0x13, 0x2a, 0x0b, 0x3d, 0x6f, 0xe8, 0xf9, 0x1e, 0x93,
	0x1f, 0xd0, 0xc4, 0x94, 0xc6, 0x89, 0x6b, 0x69, 0xa2, 0x04, 0x4a, 0x91, 0x53, 0x29, 0x5f, 0xb8,
	0x60, 0x6e, 0xb3, 0x54, 0x54, 0x98, 0xf4, 0xc1, 0xc9, 0x9f, 0x79, 0x86, 0xd9, 0x3a, 0x5d, 0xa2,
	0x70, 0x5b, 0x9a, 0x3e, 0x1e, 0xf9, 0x2e, 0xb4, 0x54, 0xb6, 0x39, 0x3b, 0x4d, 0xe7, 0xff, 0x05,
	0x3b, 0x44, 0x99, 0xf3, 0x4c, 0xa2, 0x4a, 0xba, 0xe2, 0x0c, 0x77, 0xac, 0x3e, 0x13, 0x17, 0xbe,
	0xa4, 0x28, 0x25, 0x8d, 0xaa, 0x0f, 0xa8, 0xa0, 0xff, 0x00, 0xdf, 0x94, 0xe7, 0x5e, 0xfd, 0x0b,
	0xac, 0x14, 0x0b, 0xaa, 0xd5, 0xce, 0x98, 0x04, 0xd5, 0xb3, 0x05, 0xd5, 0x46, 0xa8, 0x79, 0xe2,
	0x83, 0xc5, 0x68, 0x41, 0xb5, 0x9d, 0x33, 0x6e, 0x1f, 0xf6, 0xb4, 0x9b, 0xe6, 0xfc, 0x3b, 0xe8,
	0x28, 0x34, 0xe5, 0x3c, 0xb9, 0xd8, 0x9f, 0x1c, 0xf9, 0xdb, 0x3b, 0xbf, 0x47, 0xf8, 0xae, 0xfc,
	0xe4, 0x27, 0xc2, 0x36, 0xce, 0x85, 0x1d, 0xbf, 0x99, 0xe0, 0x28, 0x78, 0x5f, 0xd6, 0x86, 0xfc,
	0x07, 0xfb, 0x36, 0x96, 0x85, 0x2e, 0x43, 0x37, 0x28, 0x3b, 0x13, 0x54, 0x9d, 0x09, 0xae, 0x55,
	0x67, 0xbc, 0x9f, 0x75, 0xa7, 0x43, 0xb0, 0x09, 0xc0, 0x95, 0x40, 0x5a, 0x60, 0xd9, 0xa5, 0xfa,
	0x9a, 0xd7, 0x3d, 0x09, 0x50, 0xa9, 0xfe, 0x80, 0x73, 0x83, 0xfa, 0xc6, 0xe9, 0x76, 0x3e, 0x23,
	0x9d, 0xfa, 0xda, 0x9c, 0x9d, 0x15, 0x4e, 0x00, 0x16, 0x39, 0xbb, 0xf4, 0xba, 0x7f, 0x00, 0x33,
	0x4c, 0x70, 0xa7, 0xfa, 0x78, 0x9b, 0x57, 0x9f, 0x1c, 0xff, 0xc6, 0x65, 0x4b, 0xbf, 0xc4, 0xef,
	0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x43, 0x20, 0x53, 0x74, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// ListUsers
	ListUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UsersResponse, error)
	// CreateUser
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	// GetUserById
	GetUserByID(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserResponse, error)
	// UpdateUser
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	// DeleteUser
	DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserBoolResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ListUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/goalvaro.UserService/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/goalvaro.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByID(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/goalvaro.UserService/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/goalvaro.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserBoolResponse, error) {
	out := new(UserBoolResponse)
	err := c.cc.Invoke(ctx, "/goalvaro.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// ListUsers
	ListUser(context.Context, *empty.Empty) (*UsersResponse, error)
	// CreateUser
	CreateUser(context.Context, *User) (*UserResponse, error)
	// GetUserById
	GetUserByID(context.Context, *UserId) (*UserResponse, error)
	// UpdateUser
	UpdateUser(context.Context, *User) (*UserResponse, error)
	// DeleteUser
	DeleteUser(context.Context, *UserId) (*UserBoolResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) ListUser(ctx context.Context, req *empty.Empty) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) GetUserByID(ctx context.Context, req *UserId) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(ctx context.Context, req *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServiceServer) DeleteUser(ctx context.Context, req *UserId) (*UserBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalvaro.UserService/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUser(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalvaro.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalvaro.UserService/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByID(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalvaro.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalvaro.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goalvaro.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _UserService_ListUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _UserService_GetUserByID_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go-alvaro-userservice.proto",
}