package jobs

import "context"

type Job interface {
	Name() string
	Description() string
	Run(ctx context.Context) error
}
