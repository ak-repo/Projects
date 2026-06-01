package tasks

import (
	"context"
	"log/slog"

	"cron-order-worker-api/internal/services"
)

type RetryFailedOrdersJob struct {
	orderService *services.OrderService
	logger       *slog.Logger
}

func NewRetryFailedOrdersJob(orderService *services.OrderService, logger *slog.Logger) *RetryFailedOrdersJob {
	return &RetryFailedOrdersJob{orderService: orderService, logger: logger}
}

func (j *RetryFailedOrdersJob) Name() string {
	return "retry_failed_orders"
}

func (j *RetryFailedOrdersJob) Description() string {
	return "Retries failed or pending order sync records"
}

func (j *RetryFailedOrdersJob) Run(ctx context.Context) error {
	processed, err := j.orderService.RetryFailedOrders(ctx, 20)
	if err != nil {
		return err
	}
	j.logger.Info("retry failed orders job completed", "processed", processed)
	return nil
}
