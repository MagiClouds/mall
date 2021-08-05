// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.3

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

type ProductHTTPServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductReply, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductReply, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductReply, error)
	ListProduct(context.Context, *ListProductRequest) (*ListProductReply, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductReply, error)
}

func RegisterProductHTTPServer(s *http.Server, srv ProductHTTPServer) {
	r := s.Route("/")
	r.POST("/product", _Product_CreateProduct0_HTTP_Handler(srv))
	r.PUT("/product", _Product_UpdateProduct0_HTTP_Handler(srv))
	r.DELETE("/product/{id}", _Product_DeleteProduct0_HTTP_Handler(srv))
	r.GET("/product/{id}", _Product_GetProduct0_HTTP_Handler(srv))
	r.GET("/product", _Product_ListProduct0_HTTP_Handler(srv))
}

func _Product_CreateProduct0_HTTP_Handler(srv ProductHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateProductRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.product.service.v1.Product/CreateProduct")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateProduct(ctx, req.(*CreateProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateProductReply)
		return ctx.Result(200, reply)
	}
}

func _Product_UpdateProduct0_HTTP_Handler(srv ProductHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateProductRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.product.service.v1.Product/UpdateProduct")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProduct(ctx, req.(*UpdateProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateProductReply)
		return ctx.Result(200, reply)
	}
}

func _Product_DeleteProduct0_HTTP_Handler(srv ProductHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteProductRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.product.service.v1.Product/DeleteProduct")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProduct(ctx, req.(*DeleteProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteProductReply)
		return ctx.Result(200, reply)
	}
}

func _Product_GetProduct0_HTTP_Handler(srv ProductHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProductRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.product.service.v1.Product/GetProduct")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProduct(ctx, req.(*GetProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProductReply)
		return ctx.Result(200, reply)
	}
}

func _Product_ListProduct0_HTTP_Handler(srv ProductHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListProductRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.product.service.v1.Product/ListProduct")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListProduct(ctx, req.(*ListProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListProductReply)
		return ctx.Result(200, reply)
	}
}

type ProductHTTPClient interface {
	CreateProduct(ctx context.Context, req *CreateProductRequest, opts ...http.CallOption) (rsp *CreateProductReply, err error)
	DeleteProduct(ctx context.Context, req *DeleteProductRequest, opts ...http.CallOption) (rsp *DeleteProductReply, err error)
	GetProduct(ctx context.Context, req *GetProductRequest, opts ...http.CallOption) (rsp *GetProductReply, err error)
	ListProduct(ctx context.Context, req *ListProductRequest, opts ...http.CallOption) (rsp *ListProductReply, err error)
	UpdateProduct(ctx context.Context, req *UpdateProductRequest, opts ...http.CallOption) (rsp *UpdateProductReply, err error)
}

type ProductHTTPClientImpl struct {
	cc *http.Client
}

func NewProductHTTPClient(client *http.Client) ProductHTTPClient {
	return &ProductHTTPClientImpl{client}
}

func (c *ProductHTTPClientImpl) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...http.CallOption) (*CreateProductReply, error) {
	var out CreateProductReply
	pattern := "/product"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.product.service.v1.Product/CreateProduct"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProductHTTPClientImpl) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...http.CallOption) (*DeleteProductReply, error) {
	var out DeleteProductReply
	pattern := "/product/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.product.service.v1.Product/DeleteProduct"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProductHTTPClientImpl) GetProduct(ctx context.Context, in *GetProductRequest, opts ...http.CallOption) (*GetProductReply, error) {
	var out GetProductReply
	pattern := "/product/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.product.service.v1.Product/GetProduct"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProductHTTPClientImpl) ListProduct(ctx context.Context, in *ListProductRequest, opts ...http.CallOption) (*ListProductReply, error) {
	var out ListProductReply
	pattern := "/product"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.product.service.v1.Product/ListProduct"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProductHTTPClientImpl) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...http.CallOption) (*UpdateProductReply, error) {
	var out UpdateProductReply
	pattern := "/product"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.product.service.v1.Product/UpdateProduct"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
