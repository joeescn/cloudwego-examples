package main

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	api "github.com/joeescn/hello-svc/kitex_gen/api"
)

// HelloImpl implements the last service interface defined in the IDL.
type HelloImpl struct{}

// Echo implements the HelloImpl interface.
func (s *HelloImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	klog.Info("[hello-svc] called Hello Echo service", req.String())
	if req.GetMessage() == "" {
		return nil, kerrors.ErrBiz.WithCause(errors.New("message cannot be empty"))
	}
	return &api.Response{
		Message: "input message:" + req.GetMessage(),
	}, nil
}

// Add implements the HelloImpl interface.
func (s *HelloImpl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
	// TODO: Your code here...
	klog.Info("[hello-svc] called Hello Add service", req.String())
	if req.First == 0 || req.Second == 0 {
		return nil, kerrors.ErrBiz.WithCause(errors.New("input cannot be empty"))
	}
	return &api.AddResponse{
		Sum: req.First + req.Second,
	}, nil
}
