package example

import "context"

func (r *repo) Hello(ctx context.Context, name string) error {
	r.logger.Infof("Hello Example Repository: %s", name)

	return nil
}
