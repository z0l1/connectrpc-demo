package main

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	userv1 "github.com/z0l1/connectrpc-demo/gen/grpc/user/v1"
	"github.com/z0l1/connectrpc-demo/gen/grpc/user/v1/userv1connect"
)

type UserServer struct {
}

func (s *UserServer) GetUserById(c context.Context, req *userv1.GetUserByIdRequest) (*userv1.GetUserByIdResponse, error) {
	slog.Info("GetUserById", slog.Int64("id", req.Id))

	resp := userv1.GetUserByIdResponse{
		Id:       req.Id,
		Username: "anyad",
		Email:    "any@ad",
		Password: "redacted",
	}

	return &resp, nil
}

func (s *UserServer) GetUserByUsername(c context.Context, req *userv1.GetUserByUsernameRequest) (*userv1.GetUserByUsernameResponse, error) {
	slog.Info("GetUserByUsername", slog.String("username", req.Username))

	resp := userv1.GetUserByUsernameResponse{
		Id:       1,
		Username: req.Username,
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

	slog.Info("grpc user service initiated", slog.String("path", path))

	mux.Handle(path, handler)

	p := new(http.Protocols)
	p.SetHTTP1(true)
	// Use h2c so we can serve HTTP/2 without TLS.
	p.SetUnencryptedHTTP2(true)

	addr := "localhost:8080"
	s := http.Server{
		Addr:      addr,
		Handler:   mux,
		Protocols: p,
	}
	slog.Info("server listening", slog.String("addr", addr))

	err := s.ListenAndServe()

	slog.Info("server shutdown")
	if err != nil {
		slog.Info("server error", slog.String("err", err.Error()))
	}
}
