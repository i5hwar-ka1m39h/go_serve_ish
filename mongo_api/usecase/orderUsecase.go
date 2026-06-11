package usecase

import (
	"context"
	"time"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
)

type orderUsecase struct {
	orderRepository model.OrderRepository
	contextTime     time.Duration
}

func NewOrderUsecase(orderRep model.OrderRepository, time time.Duration) model.OrderUsecases {
	return &orderUsecase{
		orderRepository: orderRep,
		contextTime:     time,
	}
}

func (ordUC *orderUsecase) CreateOrder(c context.Context, order *model.Order) error {
	ctx, cancel := context.WithTimeout(c, ordUC.contextTime)
	defer cancel()

	return ordUC.orderRepository.CreateSingle(ctx, order)
}

func (ordUC *orderUsecase) UpdateOrder(c context.Context, orderId string, updateData map[string]any) error {
	ctx, cancel := context.WithTimeout(c, ordUC.contextTime)
	defer cancel()

	return ordUC.orderRepository.UpdateSingle(ctx, orderId, updateData)
}

func (ordUC *orderUsecase) GetOrderById(c context.Context, orderId string) (*model.Order, error) {
	ctx, cancel := context.WithTimeout(c, ordUC.contextTime)
	defer cancel()

	return ordUC.orderRepository.GetSingleId(ctx, orderId)
}

func (ordUc *orderUsecase) GetAllOrderForUser(c context.Context, userId string) ([]model.Order, error) {
	ctx, cancel := context.WithTimeout(c, ordUc.contextTime)
	defer cancel()

	return ordUc.orderRepository.GetAllForUser(ctx, userId)
}
