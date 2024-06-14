package example

import (
	"context"

	"go-fiber-postgres-boilerplate/internal/dto"
	"go-fiber-postgres-boilerplate/internal/repository/example"
	"go-fiber-postgres-boilerplate/pkg/response"

	"go.uber.org/zap"
)

type uc struct {
	logger *zap.SugaredLogger
	repo   example.Repository
}

func NewUsecase(
	logger *zap.SugaredLogger,
	repo example.Repository,
) Usecase {
	return &uc{
		logger,
		repo,
	}
}

func (u *uc) SayHello(ctx context.Context, req dto.SayHelloRequest) (response.API, error) {
	if req.Name == "" {
		return response.BadRequest(nil, "Please provide valid name!")
	}

	err := u.repo.Hello(ctx, req.Name)
	if err != nil {
		return response.ResponseDBError(err, "failed to process saying hello :(")
	}

	return response.API{
		Status:  200,
		Success: true,
		Message: "Hey it's working you know!",
	}, nil
}
