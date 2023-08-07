// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/v1/dictionary.proto

package v1

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for DictionaryService service

func NewDictionaryServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DictionaryService service

type DictionaryService interface {
	Paginate(ctx context.Context, in *DictionaryPaginateRequest, opts ...client.CallOption) (*DictionaryPaginateResponse, error)
	Create(ctx context.Context, in *DictionaryCreateRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *DictionaryUpdateRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	Share(ctx context.Context, in *IDRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *IDRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	Show(ctx context.Context, in *IDRequest, opts ...client.CallOption) (*DictionaryItem, error)
	ShowBySlug(ctx context.Context, in *ShowBySlugRequest, opts ...client.CallOption) (*DictionaryItem, error)
}

type dictionaryService struct {
	c    client.Client
	name string
}

func NewDictionaryService(name string, c client.Client) DictionaryService {
	return &dictionaryService{
		c:    c,
		name: name,
	}
}

func (c *dictionaryService) Paginate(ctx context.Context, in *DictionaryPaginateRequest, opts ...client.CallOption) (*DictionaryPaginateResponse, error) {
	req := c.c.NewRequest(c.name, "DictionaryService.Paginate", in)
	out := new(DictionaryPaginateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryService) Create(ctx context.Context, in *DictionaryCreateRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "DictionaryService.Create", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryService) Update(ctx context.Context, in *DictionaryUpdateRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "DictionaryService.Update", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryService) Share(ctx context.Context, in *IDRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "DictionaryService.Share", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryService) Delete(ctx context.Context, in *IDRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "DictionaryService.Delete", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryService) Show(ctx context.Context, in *IDRequest, opts ...client.CallOption) (*DictionaryItem, error) {
	req := c.c.NewRequest(c.name, "DictionaryService.Show", in)
	out := new(DictionaryItem)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryService) ShowBySlug(ctx context.Context, in *ShowBySlugRequest, opts ...client.CallOption) (*DictionaryItem, error) {
	req := c.c.NewRequest(c.name, "DictionaryService.ShowBySlug", in)
	out := new(DictionaryItem)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DictionaryService service

type DictionaryServiceHandler interface {
	Paginate(context.Context, *DictionaryPaginateRequest, *DictionaryPaginateResponse) error
	Create(context.Context, *DictionaryCreateRequest, *emptypb.Empty) error
	Update(context.Context, *DictionaryUpdateRequest, *emptypb.Empty) error
	Share(context.Context, *IDRequest, *emptypb.Empty) error
	Delete(context.Context, *IDRequest, *emptypb.Empty) error
	Show(context.Context, *IDRequest, *DictionaryItem) error
	ShowBySlug(context.Context, *ShowBySlugRequest, *DictionaryItem) error
}

func RegisterDictionaryServiceHandler(s server.Server, hdlr DictionaryServiceHandler, opts ...server.HandlerOption) error {
	type dictionaryService interface {
		Paginate(ctx context.Context, in *DictionaryPaginateRequest, out *DictionaryPaginateResponse) error
		Create(ctx context.Context, in *DictionaryCreateRequest, out *emptypb.Empty) error
		Update(ctx context.Context, in *DictionaryUpdateRequest, out *emptypb.Empty) error
		Share(ctx context.Context, in *IDRequest, out *emptypb.Empty) error
		Delete(ctx context.Context, in *IDRequest, out *emptypb.Empty) error
		Show(ctx context.Context, in *IDRequest, out *DictionaryItem) error
		ShowBySlug(ctx context.Context, in *ShowBySlugRequest, out *DictionaryItem) error
	}
	type DictionaryService struct {
		dictionaryService
	}
	h := &dictionaryServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DictionaryService{h}, opts...))
}

type dictionaryServiceHandler struct {
	DictionaryServiceHandler
}

func (h *dictionaryServiceHandler) Paginate(ctx context.Context, in *DictionaryPaginateRequest, out *DictionaryPaginateResponse) error {
	return h.DictionaryServiceHandler.Paginate(ctx, in, out)
}

func (h *dictionaryServiceHandler) Create(ctx context.Context, in *DictionaryCreateRequest, out *emptypb.Empty) error {
	return h.DictionaryServiceHandler.Create(ctx, in, out)
}

func (h *dictionaryServiceHandler) Update(ctx context.Context, in *DictionaryUpdateRequest, out *emptypb.Empty) error {
	return h.DictionaryServiceHandler.Update(ctx, in, out)
}

func (h *dictionaryServiceHandler) Share(ctx context.Context, in *IDRequest, out *emptypb.Empty) error {
	return h.DictionaryServiceHandler.Share(ctx, in, out)
}

func (h *dictionaryServiceHandler) Delete(ctx context.Context, in *IDRequest, out *emptypb.Empty) error {
	return h.DictionaryServiceHandler.Delete(ctx, in, out)
}

func (h *dictionaryServiceHandler) Show(ctx context.Context, in *IDRequest, out *DictionaryItem) error {
	return h.DictionaryServiceHandler.Show(ctx, in, out)
}

func (h *dictionaryServiceHandler) ShowBySlug(ctx context.Context, in *ShowBySlugRequest, out *DictionaryItem) error {
	return h.DictionaryServiceHandler.ShowBySlug(ctx, in, out)
}
