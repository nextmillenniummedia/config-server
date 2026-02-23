package utils

import (
	"config-server/app/errors"
	"encoding/json"
	lib_errors "errors"
	"net/http"

	loggergo "github.com/nextmillenniummedia/logger-go"
)

func SendHttpResponse(res http.ResponseWriter, data any, logger loggergo.ILogger) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(data); err != nil {
		SendHttpError(res, err, logger)
	}
}

func SendHttpError(res http.ResponseWriter, err error, logger loggergo.ILogger) {
	code := http.StatusInternalServerError

	if lib_errors.Is(err, errors.CommandNotFoundError) {
		code = http.StatusNotFound
	}
	if lib_errors.Is(err, errors.BadRequestError) {
		code = http.StatusBadRequest
	}

	if code == http.StatusInternalServerError {
		logger.Error("response error", "err", err.Error())
	}

	message := err.Error()
	payload := errors.HttpError{
		Code:    code,
		Message: message,
	}

	res.WriteHeader(code)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(&payload)
}

func SendHttpForbidden(res http.ResponseWriter) {
	res.WriteHeader(http.StatusForbidden)
	res.Write([]byte("{}"))
}
