// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: authentication/authentication_service.proto

package authentication_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	EcoService_Register_FullMethodName            = "/auth_service.EcoService/Register"
	EcoService_Login_FullMethodName               = "/auth_service.EcoService/Login"
	EcoService_GetProfile_FullMethodName          = "/auth_service.EcoService/GetProfile"
	EcoService_EditProfile_FullMethodName         = "/auth_service.EcoService/EditProfile"
	EcoService_ListUsers_FullMethodName           = "/auth_service.EcoService/ListUsers"
	EcoService_DeleteUser_FullMethodName          = "/auth_service.EcoService/DeleteUser"
	EcoService_ResetPassword_FullMethodName       = "/auth_service.EcoService/ResetPassword"
	EcoService_RefreshToken_FullMethodName        = "/auth_service.EcoService/RefreshToken"
	EcoService_Logout_FullMethodName              = "/auth_service.EcoService/Logout"
	EcoService_GetEcoPoints_FullMethodName        = "/auth_service.EcoService/GetEcoPoints"
	EcoService_AddEcoPoints_FullMethodName        = "/auth_service.EcoService/AddEcoPoints"
	EcoService_GetEcoPointsHistory_FullMethodName = "/auth_service.EcoService/GetEcoPointsHistory"
)

// EcoServiceClient is the client API for EcoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Xizmatlar (services)
type EcoServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetProfile(ctx context.Context, in *ProfilRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	EditProfile(ctx context.Context, in *EditProfileRequest, opts ...grpc.CallOption) (*EditProfileResponse, error)
	ListUsers(ctx context.Context, in *UsersListRequest, opts ...grpc.CallOption) (*UsersListResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	GetEcoPoints(ctx context.Context, in *EcoPointsRequest, opts ...grpc.CallOption) (*EcoPointsResponse, error)
	AddEcoPoints(ctx context.Context, in *AddEcoPointsRequest, opts ...grpc.CallOption) (*AddEcoPointsResponse, error)
	GetEcoPointsHistory(ctx context.Context, in *GetEcoPointsHistoryRequest, opts ...grpc.CallOption) (*EcoPointsHistoryResponse, error)
}

type ecoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEcoServiceClient(cc grpc.ClientConnInterface) EcoServiceClient {
	return &ecoServiceClient{cc}
}

func (c *ecoServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, EcoService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, EcoService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) GetProfile(ctx context.Context, in *ProfilRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, EcoService_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) EditProfile(ctx context.Context, in *EditProfileRequest, opts ...grpc.CallOption) (*EditProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditProfileResponse)
	err := c.cc.Invoke(ctx, EcoService_EditProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) ListUsers(ctx context.Context, in *UsersListRequest, opts ...grpc.CallOption) (*UsersListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UsersListResponse)
	err := c.cc.Invoke(ctx, EcoService_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, EcoService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, EcoService_ResetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, EcoService_RefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, EcoService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) GetEcoPoints(ctx context.Context, in *EcoPointsRequest, opts ...grpc.CallOption) (*EcoPointsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EcoPointsResponse)
	err := c.cc.Invoke(ctx, EcoService_GetEcoPoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) AddEcoPoints(ctx context.Context, in *AddEcoPointsRequest, opts ...grpc.CallOption) (*AddEcoPointsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddEcoPointsResponse)
	err := c.cc.Invoke(ctx, EcoService_AddEcoPoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecoServiceClient) GetEcoPointsHistory(ctx context.Context, in *GetEcoPointsHistoryRequest, opts ...grpc.CallOption) (*EcoPointsHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EcoPointsHistoryResponse)
	err := c.cc.Invoke(ctx, EcoService_GetEcoPointsHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EcoServiceServer is the server API for EcoService service.
// All implementations must embed UnimplementedEcoServiceServer
// for forward compatibility
//
// Xizmatlar (services)
type EcoServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetProfile(context.Context, *ProfilRequest) (*ProfileResponse, error)
	EditProfile(context.Context, *EditProfileRequest) (*EditProfileResponse, error)
	ListUsers(context.Context, *UsersListRequest) (*UsersListResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	GetEcoPoints(context.Context, *EcoPointsRequest) (*EcoPointsResponse, error)
	AddEcoPoints(context.Context, *AddEcoPointsRequest) (*AddEcoPointsResponse, error)
	GetEcoPointsHistory(context.Context, *GetEcoPointsHistoryRequest) (*EcoPointsHistoryResponse, error)
	mustEmbedUnimplementedEcoServiceServer()
}

// UnimplementedEcoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEcoServiceServer struct {
}

func (UnimplementedEcoServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedEcoServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedEcoServiceServer) GetProfile(context.Context, *ProfilRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedEcoServiceServer) EditProfile(context.Context, *EditProfileRequest) (*EditProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProfile not implemented")
}
func (UnimplementedEcoServiceServer) ListUsers(context.Context, *UsersListRequest) (*UsersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedEcoServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedEcoServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedEcoServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedEcoServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedEcoServiceServer) GetEcoPoints(context.Context, *EcoPointsRequest) (*EcoPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEcoPoints not implemented")
}
func (UnimplementedEcoServiceServer) AddEcoPoints(context.Context, *AddEcoPointsRequest) (*AddEcoPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEcoPoints not implemented")
}
func (UnimplementedEcoServiceServer) GetEcoPointsHistory(context.Context, *GetEcoPointsHistoryRequest) (*EcoPointsHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEcoPointsHistory not implemented")
}
func (UnimplementedEcoServiceServer) mustEmbedUnimplementedEcoServiceServer() {}

// UnsafeEcoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EcoServiceServer will
// result in compilation errors.
type UnsafeEcoServiceServer interface {
	mustEmbedUnimplementedEcoServiceServer()
}

func RegisterEcoServiceServer(s grpc.ServiceRegistrar, srv EcoServiceServer) {
	s.RegisterService(&EcoService_ServiceDesc, srv)
}

func _EcoService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfilRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).GetProfile(ctx, req.(*ProfilRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_EditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).EditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_EditProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).EditProfile(ctx, req.(*EditProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).ListUsers(ctx, req.(*UsersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_GetEcoPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EcoPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).GetEcoPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_GetEcoPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).GetEcoPoints(ctx, req.(*EcoPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_AddEcoPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEcoPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).AddEcoPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_AddEcoPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).AddEcoPoints(ctx, req.(*AddEcoPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcoService_GetEcoPointsHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEcoPointsHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcoServiceServer).GetEcoPointsHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcoService_GetEcoPointsHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcoServiceServer).GetEcoPointsHistory(ctx, req.(*GetEcoPointsHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EcoService_ServiceDesc is the grpc.ServiceDesc for EcoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EcoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.EcoService",
	HandlerType: (*EcoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _EcoService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _EcoService_Login_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _EcoService_GetProfile_Handler,
		},
		{
			MethodName: "EditProfile",
			Handler:    _EcoService_EditProfile_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _EcoService_ListUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _EcoService_DeleteUser_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _EcoService_ResetPassword_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _EcoService_RefreshToken_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _EcoService_Logout_Handler,
		},
		{
			MethodName: "GetEcoPoints",
			Handler:    _EcoService_GetEcoPoints_Handler,
		},
		{
			MethodName: "AddEcoPoints",
			Handler:    _EcoService_AddEcoPoints_Handler,
		},
		{
			MethodName: "GetEcoPointsHistory",
			Handler:    _EcoService_GetEcoPointsHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication/authentication_service.proto",
}
