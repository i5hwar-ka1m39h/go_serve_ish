package usecase

import (
	"context"
	"time"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
)

type productUsecase struct {
	productRepository model.ProductRepository
	contextTime       time.Duration
}

func NewProductUsecase(productRepo model.ProductRepository, time time.Duration) model.ProductUsecase {
	return &productUsecase{
		productRepository: productRepo,
		contextTime:       time,
	}
}

func (prdUC *productUsecase) CreateSingleProduct(c context.Context, product *model.Product) error {
	ctx, cancel := context.WithTimeout(c, prdUC.contextTime)

	defer cancel()

	return prdUC.productRepository.CreateSingle(ctx, product)
}
func (prdUC *productUsecase) CreateMultipeProduct(c context.Context, product map[string]any) error {
	ctx, cancel := context.WithTimeout(c, prdUC.contextTime)
	defer cancel()

	return prdUC.CreateMultipeProduct(ctx, product)
}
func (prdUC *productUsecase) GetSingleProduct(c context.Context, productId string) (model.Product, error) {
	ctx, cancel := context.WithTimeout(c, prdUC.contextTime)
	defer cancel()

	return prdUC.productRepository.GetSingleById(ctx, productId)

}
func (prdUC *productUsecase) GetAllProducts(c context.Context) ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(c, prdUC.contextTime)
	defer cancel()

	return prdUC.productRepository.GetAll(ctx)

}
func (prdUC *productUsecase) SearchProduct(c context.Context, productName string) ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(c, prdUC.contextTime)
	defer cancel()

	return prdUC.productRepository.GetByStr(ctx, productName)

}
func (prdUC *productUsecase) UpdateSingleProduct(c context.Context, productId string, updateData map[string]any) error {
	ctx, cancel := context.WithTimeout(c, prdUC.contextTime)
	defer cancel()

	return prdUC.UpdateSingleProduct(ctx, productId, updateData)

}
