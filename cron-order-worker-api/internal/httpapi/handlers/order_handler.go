package handlers

import (
	"net/http"

	"cron-order-worker-api/internal/httpapi/response"
	"cron-order-worker-api/internal/services"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.ListOrders(r.Context())
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
		"data":   orders,
	})
}

func (h *OrderHandler) SeedOrders(w http.ResponseWriter, r *http.Request) {
	if err := h.orderService.SeedOrders(r.Context()); err != nil {
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusCreated, map[string]any{
		"status":  "ok",
		"message": "demo orders seeded",
	})
}
