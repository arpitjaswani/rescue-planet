package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response :
type Response struct {
	Status  int         `json:"status"`
	Content interface{} `json:"content"`
}

// WebResponse :
func WebResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	var response Response
	response.Status = statusCode
	response.Content = data
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("err encoding final data: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error, please come after sometime."))
	}
	return
}
