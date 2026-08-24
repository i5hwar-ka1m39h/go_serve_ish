package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type cartController struct {
	cartUC model.CartUsercase
}

type cartProductBody struct {
	ProductId string  `json:"productId"`
	Price     float64 `json:"price"`
	Quantity  uint64  `json:"quantity"`
}

type createCartBody struct {
	UserId string            `json:"userId"`
	Items  []cartProductBody `json:"items"`
}

type CreateCartResp struct {
	Message string      `json:"message"`
	Cart    *model.Cart `json:"cart"`
}

func (cartCntrl *cartController) CreateCart(w http.ResponseWriter, r *http.Request) {
	var reqBody createCartBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error occured while reading body", w, http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(reqBody.UserId)
	if err != nil {
		utils.ErrorSend(err, "invalid user id", w, http.StatusBadRequest)
		return
	}

	cartItems := make([]model.CartProduct, 0, len(reqBody.Items))
	totalPrice := 0.0

	for _, item := range reqBody.Items {
		productId, err := primitive.ObjectIDFromHex(item.ProductId)
		if err != nil {
			utils.ErrorSend(err, "invalid product id in cart items", w, http.StatusBadRequest)
			return
		}

		cartItems = append(cartItems, model.CartProduct{
			ProductID: productId,
			Price:     item.Price,
			Quantity:  item.Quantity,
		})

		totalPrice += item.Price * float64(item.Quantity)
	}

	cart := model.Cart{
		ID:         primitive.NewObjectID(),
		USerId:     userId,
		TotalPrice: totalPrice,
		Items:      cartItems,
		CreateAt:   time.Now(),
		UpdatedAt:  time.Now(),
	}

	err = cartCntrl.cartUC.CreateCart(r.Context(), &cart)
	if err != nil {
		utils.ErrorSend(err, "failed to create cart", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(CreateCartResp{
		Message: "cart created successfully",
		Cart:    &cart,
	}, w, http.StatusCreated)
}

type addToCartBody struct {
	CartId    string  `json:"cartId"`
	ProductId string  `json:"productId"`
	Price     float64 `json:"price"`
	Quantity  uint64  `json:"quantity"`
}

type CartUpdateResp struct {
	Message string `json:"message"`
}

func (cartCntrl *cartController) AddToCart(w http.ResponseWriter, r *http.Request) {
	var reqBody addToCartBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error occured while reading body", w, http.StatusBadRequest)
		return
	}

	productId, err := primitive.ObjectIDFromHex(reqBody.ProductId)
	if err != nil {
		utils.ErrorSend(err, "invalid product id", w, http.StatusBadRequest)
		return
	}

	cartData := map[string]any{
		"productId": productId,
		"price":     reqBody.Price,
		"quantity":  reqBody.Quantity,
	}

	err = cartCntrl.cartUC.AddToCart(r.Context(), reqBody.CartId, cartData)
	if err != nil {
		utils.ErrorSend(err, "failed to add product to cart", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(CartUpdateResp{Message: "product added to cart successfully"}, w, http.StatusOK)
}

type removeFromCartBody struct {
	CartId    string `json:"cartId"`
	ProductId string `json:"productId"`
}

func (cartCntrl *cartController) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	var reqBody removeFromCartBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error occured while reading body", w, http.StatusBadRequest)
		return
	}

	productId, err := primitive.ObjectIDFromHex(reqBody.ProductId)
	if err != nil {
		utils.ErrorSend(err, "invalid product id", w, http.StatusBadRequest)
		return
	}

	cartData := map[string]any{
		"productId": productId,
	}

	err = cartCntrl.cartUC.RemoveFromCart(r.Context(), reqBody.CartId, cartData)
	if err != nil {
		utils.ErrorSend(err, "failed to remove product from cart", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(CartUpdateResp{Message: "product removed from cart successfully"}, w, http.StatusOK)
}

type getCartBody struct {
	Id string `json:"id"`
}

type GetCartResp struct {
	Message string      `json:"message"`
	Cart    *model.Cart `json:"cart"`
}

func (cartCntrl *cartController) GetCartDetails(w http.ResponseWriter, r *http.Request) {
	var reqBody getCartBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error reading body", w, http.StatusBadRequest)
		return
	}

	cart, err := cartCntrl.cartUC.GetCartDetails(r.Context(), reqBody.Id)
	if err != nil {
		utils.ErrorSend(err, "error occured while getting the cart", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(GetCartResp{
		Message: "cart found with given id",
		Cart:    cart,
	}, w, http.StatusOK)
}
