package services

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"cron-order-worker-api/internal/domain"
	"cron-order-worker-api/internal/repositories"
)

type OrderService struct {
	orderRepo *repositories.OrderRepository
	logger    *slog.Logger
}

func NewOrderService(orderRepo *repositories.OrderRepository, logger *slog.Logger) *OrderService {
	return &OrderService{orderRepo: orderRepo, logger: logger}
}

func (s *OrderService) ListOrders(ctx context.Context) ([]domain.Order, error) {
	return s.orderRepo.List(ctx)
}

func (s *OrderService) SeedOrders(ctx context.Context) error {
	return s.orderRepo.SeedFailedOrders(ctx)
}

func (s *OrderService) RetryFailedOrders(ctx context.Context, limit int) (int, error) {
	orders, err := s.orderRepo.FetchRetryableOrders(ctx, limit)
	if err != nil {
		return 0, err
	}

	processed := 0
	for _, order := range orders {
		if err := s.retrySingleOrder(ctx, order); err != nil {
			s.logger.Error("failed to retry order", "order_id", order.ID, "order_number", order.OrderNumber, "error", err)
			continue
		}
		processed++
	}

	return processed, nil
}

func (s *OrderService) retrySingleOrder(ctx context.Context, order domain.Order) error {
	if err := s.orderRepo.MarkProcessing(ctx, order.ID); err != nil {
		return err
	}

	// Simulated external ERP call.
	// In real production code, move this to an external client package/interface.
	select {
	case <-time.After(500 * time.Millisecond):
	case <-ctx.Done():
		return ctx.Err()
	}

	// Demo rule: first attempt fails, later attempts sync successfully.
	if order.Attempts >= 1 {
		return s.orderRepo.MarkSynced(ctx, order.ID)
	}

	return s.orderRepo.MarkFailed(ctx, order.ID, fmt.Sprintf("temporary ERP failure for order %s", order.OrderNumber))
}
