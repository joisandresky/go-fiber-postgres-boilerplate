package example

import (
	"context"

	"go-fiber-postgres-boilerplate/internal/dto"
	"go-fiber-postgres-boilerplate/pkg/response"
)

type Usecase interface {
	SayHello(ctx context.Context, req dto.SayHelloRequest) (response.API, error)
}
