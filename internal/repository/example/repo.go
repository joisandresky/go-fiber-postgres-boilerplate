package example

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type repo struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewRepository(db *gorm.DB, logger *zap.SugaredLogger) Repository {
	return &repo{
		db,
		logger,
	}
}

// you can write in this file or make seperate file like get_hello.go and same also for usecase or api handler
func (r *repo) Hello2(ctx context.Context, name string) error {
	r.logger.Infof("Hello Example Repository: %s", name)

	return nil
}
