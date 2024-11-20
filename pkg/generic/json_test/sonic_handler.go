package test

import (
	"context"

	"github.com/cloudwego/kitex/pkg/generic/json_test/kitex_gen/kitex/test/server"
)

type KitexHandler struct{}

func (*KitexHandler) TestGeneric(ctx context.Context, req *server.GenericRequest) (r *server.GenericResponse, err error) {
	resp := &server.GenericResponse{
		Token:    req.Token,
		Text:     req.Text,
		HttpCode: 200,
	}
	return resp, nil
}
