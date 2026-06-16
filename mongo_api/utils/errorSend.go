package utils

import (
	"log"
	"net/http"
)




func ErrorSend(err error, customMessage string, w http.ResponseWriter, errCode int){
	log.Println("error: ",customMessage, err)
	http.Error(w, customMessage+err.Error(), errCode)

}