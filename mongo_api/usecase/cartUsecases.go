package usecase

import (
	"time"

	"context"
	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
)

type cartUsercase struct {
	cartRepository model.CartRepository
	contextTime    time.Duration
}

func NewCartUsecase(cartRepo model.CartRepository, time time.Duration) model.CartUsercase {
	return &cartUsercase{
		cartRepository: cartRepo,
		contextTime:    time,
	}
}

func (cartUC *cartUsercase) CreateCart(c context.Context, cart *model.Cart) error {
	ctx, cancel := context.WithTimeout(c, cartUC.contextTime)
	defer cancel()

	return cartUC.cartRepository.CreateCart(ctx, cart)
}
func (cartUC *cartUsercase) AddToCart(c context.Context, cartId string, cartData map[string]any) error {
	ctx, cancel := context.WithTimeout(c, cartUC.contextTime)
	defer cancel()

	return  cartUC.cartRepository.UpdateCart(ctx, cartId, cartData )

}
func (cartUC *cartUsercase) RemoveFromCart(c context.Context,cartId string, cartData map[string]any) error {
	ctx, cancel := context.WithTimeout(c, cartUC.contextTime)
	defer cancel()

	return  cartUC.cartRepository.UpdateCart(ctx, cartId, cartData )

}
func (cartUC *cartUsercase) GetCartDetails(c context.Context, cartId string) (*model.Cart, error) {
	ctx, cancel := context.WithTimeout(c, cartUC.contextTime)
	defer cancel()

	return  cartUC.cartRepository.GetCart(ctx, cartId)
}
