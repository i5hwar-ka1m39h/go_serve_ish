package utils

import (
	"log"
	"net/http"
)


type ErrorResp struct{
	Error string `json:"error"`
}

func ErrorSend(err error, customMessage string, w http.ResponseWriter, errCode int){
	log.Println("error: ",customMessage, err)

	SendJsonResponse(ErrorResp{
		Error: customMessage+" :"+err.Error(),
	},w, errCode)
}