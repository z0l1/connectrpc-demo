package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	userv1 "github.com/z0l1/connectrpc-demo/gen/services/user/v1"
	"github.com/z0l1/connectrpc-demo/gen/services/user/v1/userv1connect"
)

type UserServer struct {
}

func (s *UserServer) GetUserById(context.Context, *userv1.GetUserByIdRequest) (*userv1.GetUserByIdResponse, error) {
	resp := userv1.GetUserByIdResponse{
		Id:       1,
		Username: "anyad",
		Email:    "any@ad",
		Password: "redacted",
	}

	return &resp, nil
}

func (s *UserServer) GetUserByUsername(context.Context, *userv1.GetUserByUsernameRequest) (*userv1.GetUserByUsernameResponse, error) {
	resp := userv1.GetUserByUsernameResponse{
		Id:       1,
		Username: "anyad",
		Email:    "any@ad",
		Password: "redacted",
	}

	return &resp, nil

}

func main() {
	userSrv := &UserServer{}
	mux := http.NewServeMux()

	path, handler := userv1connect.NewUserServiceHandler(
		userSrv,
		// Validation via Protovalidate is almost always recommended
		connect.WithInterceptors(validate.NewInterceptor()),
	)

	fmt.Println("path for user service", path)

	mux.Handle(path, handler)

	p := new(http.Protocols)
	p.SetHTTP1(true)
	// Use h2c so we can serve HTTP/2 without TLS.
	p.SetUnencryptedHTTP2(true)
	s := http.Server{
		Addr:      "localhost:8080",
		Handler:   mux,
		Protocols: p,
	}
	s.ListenAndServe()

}
