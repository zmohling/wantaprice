package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/zmohling/wantaprice/pkg/model"
)

type createUserRequest struct {
	User     model.User
	Password string
}

type createUserResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

func (r createUserResponse) error() error { return r.Err }

func makeCreateUserEndpoint(s IdentityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		id, err := s.CreateUser(req.User, req.Password)
		return createUserResponse{ID: id, Err: err}, nil
	}
}
