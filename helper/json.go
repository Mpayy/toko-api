package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(request *http.Request, result interface{}) {
	err := json.NewDecoder(request.Body).Decode(result)
	PanicIfError(err)
}

func WriteResponseBody(writer http.ResponseWriter, result interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(result)
	PanicIfError(err)
}
