CREATE TABLE IF NOT EXISTS orders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    order_number VARCHAR(100) NOT NULL UNIQUE,
    customer_name VARCHAR(150) NOT NULL,
    amount DECIMAL(12,2) NOT NULL,
    sync_status ENUM('pending', 'failed', 'processing', 'synced') NOT NULL DEFAULT 'pending',
    attempts INT NOT NULL DEFAULT 0,
    last_error TEXT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_orders_sync_status_attempts (sync_status, attempts),
    INDEX idx_orders_updated_at (updated_at)
);

CREATE TABLE IF NOT EXISTS job_history (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    job_name VARCHAR(100) NOT NULL,
    status ENUM('running', 'success', 'failed', 'skipped') NOT NULL,
    started_at DATETIME NOT NULL,
    finished_at DATETIME NULL,
    duration_ms BIGINT NULL,
    error_message TEXT NULL,
    triggered_by ENUM('scheduler', 'manual') NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_job_history_job_name (job_name),
    INDEX idx_job_history_created_at (created_at)
);

INSERT INTO orders (order_number, customer_name, amount, sync_status, attempts, last_error)
VALUES
('ORD-1001', 'Customer One', 2500.00, 'failed', 0, 'ERP timeout'),
('ORD-1002', 'Customer Two', 1800.50, 'failed', 1, 'Connection reset'),
('ORD-1003', 'Customer Three', 950.00, 'pending', 0, NULL)
ON DUPLICATE KEY UPDATE order_number = order_number;
