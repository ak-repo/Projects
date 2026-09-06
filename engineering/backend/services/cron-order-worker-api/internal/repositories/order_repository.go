package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"cron-order-worker-api/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) List(ctx context.Context) ([]domain.Order, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, order_number, customer_name, amount, sync_status, attempts, last_error, created_at, updated_at
		FROM orders
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]domain.Order, 0)
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(
			&order.ID,
			&order.OrderNumber,
			&order.CustomerName,
			&order.Amount,
			&order.SyncStatus,
			&order.Attempts,
			&order.LastError,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, rows.Err()
}

func (r *OrderRepository) SeedFailedOrders(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO orders (order_number, customer_name, amount, sync_status, attempts, last_error)
		VALUES
		(CONCAT('ORD-DEMO-', UUID()), 'Demo Customer A', 1200.00, 'failed', 0, 'ERP timeout'),
		(CONCAT('ORD-DEMO-', UUID()), 'Demo Customer B', 2450.00, 'failed', 1, 'Temporary network error'),
		(CONCAT('ORD-DEMO-', UUID()), 'Demo Customer C', 800.00, 'pending', 0, NULL)
	`)
	return err
}

func (r *OrderRepository) FetchRetryableOrders(ctx context.Context, limit int) ([]domain.Order, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, order_number, customer_name, amount, sync_status, attempts, last_error, created_at, updated_at
		FROM orders
		WHERE sync_status IN ('failed', 'pending')
		  AND attempts < 3
		ORDER BY updated_at ASC
		LIMIT ?
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]domain.Order, 0)
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(
			&order.ID,
			&order.OrderNumber,
			&order.CustomerName,
			&order.Amount,
			&order.SyncStatus,
			&order.Attempts,
			&order.LastError,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, rows.Err()
}

func (r *OrderRepository) MarkProcessing(ctx context.Context, orderID int64) error {
	result, err := r.db.ExecContext(ctx, `
		UPDATE orders
		SET sync_status = 'processing', updated_at = NOW()
		WHERE id = ?
		  AND sync_status IN ('failed', 'pending')
	`, orderID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("order %d is not available for processing", orderID)
	}
	return nil
}

func (r *OrderRepository) MarkSynced(ctx context.Context, orderID int64) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE orders
		SET sync_status = 'synced', last_error = NULL, updated_at = NOW()
		WHERE id = ?
	`, orderID)
	return err
}

func (r *OrderRepository) MarkFailed(ctx context.Context, orderID int64, reason string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE orders
		SET sync_status = 'failed', attempts = attempts + 1, last_error = ?, updated_at = NOW()
		WHERE id = ?
	`, reason, orderID)
	return err
}
