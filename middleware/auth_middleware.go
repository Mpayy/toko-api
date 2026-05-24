package middleware

import (
	"net/http"

	"github.com/Mpayy/toko-api/helper"
	"github.com/Mpayy/toko-api/model/web"
	"github.com/sirupsen/logrus"
)

type AuthMiddleware struct {
	Handler http.Handler
	Logger  *logrus.Logger
}

func NewMiddleware(handler http.Handler, logger *logrus.Logger) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler, Logger: logger}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	middleware.Logger.WithFields(logrus.Fields{
		"method": request.Method,
		"path":   request.URL.Path,
	}).Info("Request received")

	if "RAHASIA" == request.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	middleware.Logger.WithFields(logrus.Fields{
		"method": request.Method,
		"path":   request.URL.Path,
	}).Warn("Unauthorized request")

	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "Unauthorized",
	}
	helper.WriteResponseBody(writer, &webResponse)

}
