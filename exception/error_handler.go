package exception

import (
	"net/http"

	"github.com/Mpayy/toko-api/helper"
	"github.com/Mpayy/toko-api/model/web"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

//func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
//	if notFoundError(writer, request, err) {
//		return
//	}
//
//	if validationError(writer, request, err) {
//		return
//	}
//
//	internalServerError(writer, request, err)
//}

func NewErrorHandler(logger *logrus.Logger) func(http.ResponseWriter, *http.Request, interface{}) {
	return func(writer http.ResponseWriter, request *http.Request, err any) {
		logger.WithField("error", err).Error("Terjadi Error")
		if notFoundError(writer, request, err) {
			return
		}
		if validationError(writer, request, err) {
			return
		}
		internalServerError(writer, request, err)
	}
}

func validationError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   helper.TranslateValidationError(exception),
		}
		helper.WriteResponseBody(writer, webResponse)
		return true
	}
	return false
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteResponseBody(writer, webResponse)
		return true
	}
	return false
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteResponseBody(writer, webResponse)
}
