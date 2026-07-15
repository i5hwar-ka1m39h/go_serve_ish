package utils

import (
	"encoding/json"
	"log"
	"net/http"
)




func SendJsonResponse( resPayload any, w http.ResponseWriter, code int){

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	if err :=json.NewEncoder(w).Encode(resPayload); err != nil{
		log.Println("error in encoding json", err)
	}


}