// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0-rc6

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type UserHTTPServer interface {
	DeleteUser(context.Context, *DeleteUserRequest) (*BaseResponse, error)
	GetUser(context.Context, *GetUserRequest) (*UserInfo, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	UpdateUser(context.Context, *UserInfo) (*BaseResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	UserRegister(context.Context, *UserRegisterRequest) (*BaseResponse, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/user/register", _User_UserRegister0_HTTP_Handler(srv))
	r.POST("/user/login", _User_UserLogin0_HTTP_Handler(srv))
	r.PUT("/user", _User_UpdateUser0_HTTP_Handler(srv))
	r.DELETE("/user/{id}", _User_DeleteUser0_HTTP_Handler(srv))
	r.GET("/user/{id}", _User_GetUser0_HTTP_Handler(srv))
	r.GET("/user", _User_ListUser0_HTTP_Handler(srv))
}

func _User_UserRegister0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.User/UserRegister")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserRegister(ctx, req.(*UserRegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BaseResponse)
		return ctx.Result(200, reply)
	}
}

func _User_UserLogin0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserLoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.User/UserLogin")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserLogin(ctx, req.(*UserLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserLoginResponse)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfo
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.User/UpdateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UserInfo))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BaseResponse)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.User/DeleteUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BaseResponse)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.User/GetUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfo)
		return ctx.Result(200, reply)
	}
}

func _User_ListUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.User/ListUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserResponse)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *BaseResponse, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *UserInfo, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserResponse, err error)
	UpdateUser(ctx context.Context, req *UserInfo, opts ...http.CallOption) (rsp *BaseResponse, err error)
	UserLogin(ctx context.Context, req *UserLoginRequest, opts ...http.CallOption) (rsp *UserLoginResponse, err error)
	UserRegister(ctx context.Context, req *UserRegisterRequest, opts ...http.CallOption) (rsp *BaseResponse, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*BaseResponse, error) {
	var out BaseResponse
	pattern := "/user/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.v1.User/DeleteUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*UserInfo, error) {
	var out UserInfo
	pattern := "/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.user.v1.User/GetUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserResponse, error) {
	var out ListUserResponse
	pattern := "/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.user.v1.User/ListUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UserInfo, opts ...http.CallOption) (*BaseResponse, error) {
	var out BaseResponse
	pattern := "/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.v1.User/UpdateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...http.CallOption) (*UserLoginResponse, error) {
	var out UserLoginResponse
	pattern := "/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.v1.User/UserLogin"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...http.CallOption) (*BaseResponse, error) {
	var out BaseResponse
	pattern := "/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.v1.User/UserRegister"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
