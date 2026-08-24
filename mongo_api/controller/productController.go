package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/utils"
)

type productController struct {
	productUC model.ProductUsecase
}

type createProductBody struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    uint64  `json:"quantity"`
}

type CreateProductResp struct {
	Message string         `json:"message"`
	Product *model.Product `json:"product"`
}

func (prdCntrl *productController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var reqBody createProductBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error occured while reading body", w, http.StatusBadRequest)
		return
	}

	product := model.Product{
		Title:       reqBody.Title,
		Description: reqBody.Description,
		Price:       reqBody.Price,
		Quantity:    reqBody.Quantity,
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := prdCntrl.productUC.CreateSingleProduct(r.Context(), &product)
	if err != nil {
		utils.ErrorSend(err, "failed to create product", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(CreateProductResp{
		Message: "product created successfully",
		Product: &product,
	}, w, http.StatusCreated)
}

type getProductBody struct {
	Id string `json:"id"`
}

type GetProductResp struct {
	Message string        `json:"message"`
	Product model.Product `json:"product"`
}

func (prdCntrl *productController) GetProductById(w http.ResponseWriter, r *http.Request) {
	var reqBody getProductBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error reading body", w, http.StatusBadRequest)
		return
	}

	product, err := prdCntrl.productUC.GetSingleProduct(r.Context(), reqBody.Id)
	if err != nil {
		utils.ErrorSend(err, "error occured while getting the product", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(GetProductResp{
		Message: "product found with given id",
		Product: product,
	}, w, http.StatusOK)
}

type GetAllProductsResp struct {
	Message  string          `json:"message"`
	Products []model.Product `json:"products"`
}

func (prdCntrl *productController) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	products, err := prdCntrl.productUC.GetAllProducts(r.Context())
	if err != nil {
		utils.ErrorSend(err, "error occured while getting products", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(GetAllProductsResp{
		Message:  "products fetched successfully",
		Products: products,
	}, w, http.StatusOK)
}

type searchProductBody struct {
	Name string `json:"name"`
}

type SearchProductResp struct {
	Message  string          `json:"message"`
	Products []model.Product `json:"products"`
}

func (prdCntrl *productController) SearchProduct(w http.ResponseWriter, r *http.Request) {
	var reqBody searchProductBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error occured while reading request", w, http.StatusBadRequest)
		return
	}

	products, err := prdCntrl.productUC.SearchProduct(r.Context(), reqBody.Name)
	if err != nil {
		utils.ErrorSend(err, "error getting the products", w, http.StatusNotFound)
		return
	}

	utils.SendJsonResponse(SearchProductResp{
		Message:  "products found",
		Products: products,
	}, w, http.StatusOK)
}

type ProductUpdateBody struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Quantity    *uint64  `json:"quantity"`
}

type UpdateProductResp struct {
	Message string `json:"message"`
}

func (prdCntrl *productController) UpdateProductById(w http.ResponseWriter, r *http.Request) {
	var reqBody ProductUpdateBody

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		utils.ErrorSend(err, "error occured while reading body", w, http.StatusBadRequest)
		return
	}

	update := make(map[string]any)

	if reqBody.Title != nil {
		update["title"] = *reqBody.Title
	}

	if reqBody.Description != nil {
		update["description"] = *reqBody.Description
	}

	if reqBody.Price != nil {
		update["price"] = *reqBody.Price
	}

	if reqBody.Quantity != nil {
		update["quantity"] = *reqBody.Quantity
	}

	id := r.PathValue("id")
	err := prdCntrl.productUC.UpdateSingleProduct(r.Context(), id, update)
	if err != nil {
		utils.ErrorSend(err, "error updating product", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(UpdateProductResp{Message: "Updated the product successfully"}, w, http.StatusOK)
}
