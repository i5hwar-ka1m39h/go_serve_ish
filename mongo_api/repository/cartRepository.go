package repository

import (
	"context"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type cartRepository struct {
	database mongo.Database
	coll     string
}

func NewCartRepository(db mongo.Database, collection string) model.CartRepository {
	return &cartRepository{
		database: db,
		coll:     collection,
	}
}

func (cartRep *cartRepository) CreateCart(c context.Context, cart *model.Cart) error {
	coll := cartRep.database.Collection(cartRep.coll)

}
func (cartRep *cartRepository) UpdateCart(c context.Context, cartData map[string]any) error {}
func (cartRep *cartRepository) GetCart(c context.Context, cartId string) (*Cart, error)     {}
