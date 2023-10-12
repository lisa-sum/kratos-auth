// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.24.3
// source: admin/v1/admin.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAdminServiceGetAdminBoard = "/admin.v1.AdminService/GetAdminBoard"
const OperationAdminServiceGetModeratorBoard = "/admin.v1.AdminService/GetModeratorBoard"
const OperationAdminServiceGetPublicContent = "/admin.v1.AdminService/GetPublicContent"
const OperationAdminServiceGetUserBoard = "/admin.v1.AdminService/GetUserBoard"
const OperationAdminServiceListUser = "/admin.v1.AdminService/ListUser"
const OperationAdminServiceLogin = "/admin.v1.AdminService/Login"
const OperationAdminServiceLogout = "/admin.v1.AdminService/Logout"
const OperationAdminServiceRegister = "/admin.v1.AdminService/Register"

type AdminServiceHTTPServer interface {
	GetAdminBoard(context.Context, *emptypb.Empty) (*Content, error)
	GetModeratorBoard(context.Context, *emptypb.Empty) (*Content, error)
	GetPublicContent(context.Context, *emptypb.Empty) (*Content, error)
	GetUserBoard(context.Context, *emptypb.Empty) (*Content, error)
	// ListUser 用户列表
	ListUser(context.Context, *emptypb.Empty) (*ListUserReply, error)
	// Login 登陆
	Login(context.Context, *LoginReq) (*User, error)
	// Logout 登出
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	// Register 注册
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
}

func RegisterAdminServiceHTTPServer(s *http.Server, srv AdminServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/login", _AdminService_Login0_HTTP_Handler(srv))
	r.POST("/api/v1/logout", _AdminService_Logout0_HTTP_Handler(srv))
	r.POST("/api/v1/register", _AdminService_Register0_HTTP_Handler(srv))
	r.GET("/api/v1/users", _AdminService_ListUser0_HTTP_Handler(srv))
	r.GET("/api/v1/all", _AdminService_GetPublicContent0_HTTP_Handler(srv))
	r.GET("/api/v1/user", _AdminService_GetUserBoard0_HTTP_Handler(srv))
	r.GET("/api/v1/mod", _AdminService_GetModeratorBoard0_HTTP_Handler(srv))
	r.GET("/api/v1/admin", _AdminService_GetAdminBoard0_HTTP_Handler(srv))
}

func _AdminService_Login0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminServiceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*User)
		return ctx.Result(200, reply)
	}
}

func _AdminService_Logout0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminServiceLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _AdminService_Register0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminServiceRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _AdminService_ListUser0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminServiceListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _AdminService_GetPublicContent0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminServiceGetPublicContent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPublicContent(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Content)
		return ctx.Result(200, reply)
	}
}

func _AdminService_GetUserBoard0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminServiceGetUserBoard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserBoard(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Content)
		return ctx.Result(200, reply)
	}
}

func _AdminService_GetModeratorBoard0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminServiceGetModeratorBoard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetModeratorBoard(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Content)
		return ctx.Result(200, reply)
	}
}

func _AdminService_GetAdminBoard0_HTTP_Handler(srv AdminServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminServiceGetAdminBoard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAdminBoard(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Content)
		return ctx.Result(200, reply)
	}
}

type AdminServiceHTTPClient interface {
	GetAdminBoard(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Content, err error)
	GetModeratorBoard(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Content, err error)
	GetPublicContent(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Content, err error)
	GetUserBoard(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Content, err error)
	ListUser(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *User, err error)
	Logout(ctx context.Context, req *LogoutReq, opts ...http.CallOption) (rsp *LogoutReply, err error)
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterReply, err error)
}

type AdminServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewAdminServiceHTTPClient(client *http.Client) AdminServiceHTTPClient {
	return &AdminServiceHTTPClientImpl{client}
}

func (c *AdminServiceHTTPClientImpl) GetAdminBoard(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Content, error) {
	var out Content
	pattern := "/api/v1/admin"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminServiceGetAdminBoard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminServiceHTTPClientImpl) GetModeratorBoard(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Content, error) {
	var out Content
	pattern := "/api/v1/mod"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminServiceGetModeratorBoard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminServiceHTTPClientImpl) GetPublicContent(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Content, error) {
	var out Content
	pattern := "/api/v1/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminServiceGetPublicContent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminServiceHTTPClientImpl) GetUserBoard(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Content, error) {
	var out Content
	pattern := "/api/v1/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminServiceGetUserBoard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminServiceHTTPClientImpl) ListUser(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/api/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminServiceListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminServiceHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*User, error) {
	var out User
	pattern := "/api/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdminServiceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminServiceHTTPClientImpl) Logout(ctx context.Context, in *LogoutReq, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/api/v1/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdminServiceLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminServiceHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/api/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdminServiceRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
