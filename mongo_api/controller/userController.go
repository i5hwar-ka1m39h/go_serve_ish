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

type requestBodyGetUser struct {
	Id string `json:"id"`
}

type GetUserResp struct {
	Message string      `json:"message"`
	User    *model.User `json:"user"`
}

func (usrcntr *userController) GetUserById(w http.ResponseWriter, r *http.Request) {

	var requestBody requestBodyGetUser
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&requestBody); err != nil {
		utils.ErrorSend(err, "error reading body", w, http.StatusBadRequest)
		return
	}

	user, err := usrcntr.userUC.GetUserById(r.Context(), requestBody.Id)
	if err != nil {
		utils.ErrorSend(err, "error occured while getting the user", w, http.StatusInternalServerError)
		return
	}

	utils.SendJsonResponse(GetUserResp{
		Message: "user found with given id",
		User:    user,
	}, w, http.StatusOK)

}

type UserUpdateBody struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type UpdateResp struct {
	Message string `json:"message"`
}

func (usrcntr *userController) UpdateUserbyId(w http.ResponseWriter, r *http.Request) {
	var requestBody UserUpdateBody

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&requestBody); err != nil {
		utils.ErrorSend(err, "error occured while reading body", w, http.StatusBadRequest)

		return
	}

	update := make(map[string]any)

	if requestBody.Email != nil {
		update["email"] = *requestBody.Email
	}

	if requestBody.Name != nil {
		update["name"] = *requestBody.Name
	}

	if requestBody.Password != nil {
		update["password"] = *requestBody.Password
	}

	id := r.PathValue("id")
	err := usrcntr.userUC.UpdateUser(r.Context(), id, update)

	if err != nil {
		utils.ErrorSend(err, "error updating user", w, http.StatusInternalServerError)

		return
	}

	utils.SendJsonResponse(UpdateResp{Message: "Updated the user successfully"}, w, http.StatusOK)
}

type SearchReqBody struct {
	Email string `json:"email"`
}

type UserResp struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type SearchRsp struct {
	Users   []UserResp `json:"users"`
	Message string     `json:"message"`
}

func (usrCntrl *userController) SearchUser(w http.ResponseWriter, r *http.Request) {
	var searchRequestBody SearchReqBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&searchRequestBody); err != nil {
		utils.ErrorSend(err, "error occured while reading request", w, http.StatusBadRequest)
		return
	}

	users, err := usrCntrl.userUC.SearchUser(r.Context(), searchRequestBody.Email)
	if err != nil {
		utils.ErrorSend(err, "error getting the user", w, http.StatusNotFound)
		return
	}

	userRes := make([]UserResp, len(users))

	for _, usr := range users {
		var user = UserResp{
			Email: usr.Email,
			Name:  usr.Name,
		}
		userRes = append(userRes, user)
	}

	utils.SendJsonResponse(userRes, w, http.StatusOK)

}
