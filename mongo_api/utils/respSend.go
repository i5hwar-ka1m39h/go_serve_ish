package utils

import (
	"encoding/json"
	"log"
	"net/http"
)




func SendJsonResponse( resPayload any, w http.ResponseWriter, code int){

	data, err := json.Marshal(resPayload)

	if err != nil{
		log.Println("error marshaling response object ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)

	w.Write(data)





}