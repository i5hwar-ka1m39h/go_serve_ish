package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/utils"
)

type userController struct {
	userUC model.UserUsecase
}

type requestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Message  string `json:"message"`
	JWTtoken string `json:"token"`
}

func (userCntrl *userController) SignUpUser(w http.ResponseWriter, r *http.Request) {

	var reqBody requestBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&reqBody); err != nil {

		utils.ErrorSend(err, "error occured while reading body", w, http.StatusBadRequest)
		return
	}

	hashedPassowrd := utils.HashPass(reqBody.Password)

	rfToken, err := utils.CreateAccessToken(reqBody.Email)

	if err != nil {
		log.Println("failed to create refreshtoken", err)
	}
	user := model.User{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: hashedPassowrd,
		CreateAt: time.Now(),

		RefreshToken: rfToken,
		UpdatedAt:    time.Now(),
	}

	err = userCntrl.userUC.CreateUser(r.Context(), &user)

	if err != nil {
		utils.ErrorSend(err, "failed to create user", w, http.StatusInternalServerError)
		return
	}

	jwtToken, err := utils.CreateAccessToken(reqBody.Email)
	if err != nil {
		utils.ErrorSend(err, "error while getting jwt", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(SignUpResponse{
		Message:  "user signed up successfully",
		JWTtoken: jwtToken,
	}, w, http.StatusCreated)
}
