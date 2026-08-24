package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type orderController struct {
	orderUC model.OrderUsecases
}

type orderItemBody struct {
	ProductId string  `json:"productId"`
	Qunatity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
}

type createOrderBody struct {
	UserId            string              `json:"userId"`
	Subtotal          uint64              `json:"subtotal"`
	Tax               uint64              `json:"tax"`
	ShippingCost      uint64              `json:"shippingCost"`
	Total             uint64              `json:"total"`
	Status            model.Status        `json:"status"`
	PaymentStatus     model.PaymentStatus `json:"paymentStatus"`
	ShippingAddressId string              `json:"shippingAddressId"`
	Item              []orderItemBody     `json:"orderItems"`
}

type CreateOrderResp struct {
	Message string       `json:"message"`
	Order   *model.Order `json:"order"`
}

func (ordCntrl *orderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var reqBody createOrderBody
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

	shipAddrId, err := primitive.ObjectIDFromHex(reqBody.ShippingAddressId)
	if err != nil {
		utils.ErrorSend(err, "invalid shipping address id", w, http.StatusBadRequest)
		return
	}

	orderItems := make([]model.OrderItem, 0, len(reqBody.Item))

	for _, item := range reqBody.Item {
		productId, err := primitive.ObjectIDFromHex(item.ProductId)
		if err != nil {
			utils.ErrorSend(err, "invalid product id in order items", w, http.StatusBadRequest)
			return
		}

		orderItems = append(orderItems, model.OrderItem{
			ProductId: productId,
			Qunatity:  item.Qunatity,
			Price:     item.Price,
		})
	}

	order := model.Order{
		ID:                primitive.NewObjectID(),
		UserId:            userId,
		Subtotal:          reqBody.Subtotal,
		Tax:               reqBody.Tax,
		ShippingCost:      reqBody.ShippingCost,
		Total:             reqBody.Total,
		Status:            reqBody.Status,
		PaymentStatus:     reqBody.PaymentStatus,
		ShippingAddressId: shipAddrId,
		Item:              orderItems,
		CreateAt:          time.Now(),
		UpdatedAt:         time.Now(),
	}

	err = ordCntrl.orderUC.CreateOrder(r.Context(), &order)
	if err != nil {
		utils.ErrorSend(err, "failed to create order", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(CreateOrderResp{
		Message: "order created successfully",
		Order:   &order,
	}, w, http.StatusCreated)
}

type getOrderBody struct {
	Id string `json:"id"`
}

type GetOrderResp struct {
	Message string       `json:"message"`
	Order   *model.Order `json:"order"`
}

func (ordCntrl *orderController) GetOrderById(w http.ResponseWriter, r *http.Request) {
	var reqBody getOrderBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error reading body", w, http.StatusBadRequest)
		return
	}

	order, err := ordCntrl.orderUC.GetOrderById(r.Context(), reqBody.Id)
	if err != nil {
		utils.ErrorSend(err, "error occured while getting the order", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(GetOrderResp{
		Message: "order found with given id",
		Order:   order,
	}, w, http.StatusOK)
}

type getUserOrdersBody struct {
	UserId string `json:"userId"`
}

type GetUserOrdersResp struct {
	Message string        `json:"message"`
	Orders  []model.Order `json:"orders"`
}

func (ordCntrl *orderController) GetAllOrdersForUser(w http.ResponseWriter, r *http.Request) {
	var reqBody getUserOrdersBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error occured while reading request", w, http.StatusBadRequest)
		return
	}

	orders, err := ordCntrl.orderUC.GetAllOrderForUser(r.Context(), reqBody.UserId)
	if err != nil {
		utils.ErrorSend(err, "error getting orders for user", w, http.StatusNotFound)
		return
	}

	utils.SendJsonResponse(GetUserOrdersResp{
		Message: "orders fetched successfully",
		Orders:  orders,
	}, w, http.StatusOK)
}

type OrderUpdateBody struct {
	Status        *model.Status        `json:"status"`
	PaymentStatus *model.PaymentStatus `json:"paymentStatus"`
	ShippingCost  *uint64              `json:"shippingCost"`
}

type UpdateOrderResp struct {
	Message string `json:"message"`
}

func (ordCntrl *orderController) UpdateOrderById(w http.ResponseWriter, r *http.Request) {
	var reqBody OrderUpdateBody

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error occured while reading body", w, http.StatusBadRequest)
		return
	}

	update := make(map[string]any)

	if reqBody.Status != nil {
		update["status"] = *reqBody.Status
	}

	if reqBody.PaymentStatus != nil {
		update["paymentStatus"] = *reqBody.PaymentStatus
	}

	if reqBody.ShippingCost != nil {
		update["shippingCost"] = *reqBody.ShippingCost
	}

	id := r.PathValue("id")
	err := ordCntrl.orderUC.UpdateOrder(r.Context(), id, update)
	if err != nil {
		utils.ErrorSend(err, "error updating order", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(UpdateOrderResp{Message: "Updated the order successfully"}, w, http.StatusOK)
}
