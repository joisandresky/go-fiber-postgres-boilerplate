package example

import "context"

type Repository interface {
	Hello(ctx context.Context, name string) error
	Hello2(ctx context.Context, name string) error
}
