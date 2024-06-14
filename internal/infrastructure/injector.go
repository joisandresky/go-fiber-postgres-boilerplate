package infrastructure

import (
	exampleApi "go-fiber-postgres-boilerplate/internal/api/example/v1"
	exampleRepo "go-fiber-postgres-boilerplate/internal/repository/example"
	exampleUsecase "go-fiber-postgres-boilerplate/internal/usecase/example"

	"github.com/gofiber/fiber/v2"
)

func (srv *Server) InjectDependencies(app fiber.Router) {
	exampleRepo := exampleRepo.NewRepository(srv.psqlDB, srv.logger)
	exampleUC := exampleUsecase.NewUsecase(srv.logger, exampleRepo)

	exampleApi.NewHttpApi(exampleUC).RegisterRoutes(app)
}
