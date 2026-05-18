package middleware

import (
	"net/http"

	"github.com/Mpayy/toko-api/helper"
	"github.com/Mpayy/toko-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "Unauthorized",
	}
	helper.WriteResponseBody(writer, &webResponse)

}
