package utilsApi

import (
	"authentication-service/models/errorsModel"
	"authentication-service/models/messageModel"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func OkResponse(w http.ResponseWriter, r *http.Request, obj interface{}) {
	jsonResponse, err := json.Marshal(obj)
	if err != nil {
		ResponseByError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func NoContentResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func CreatedResponse(w http.ResponseWriter, r *http.Request, obj interface{}) {
	jsonResponse, err := json.Marshal(obj)
	if err != nil {
		ResponseByError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func ResponseByError(w http.ResponseWriter, r *http.Request, err error) {

	if err, ok := err.(*errorsModel.ValidationErrorModel); ok {
		responseMessage(w, r, err, http.StatusBadRequest)
		return
	}

	if err, ok := err.(*errorsModel.ConflictErrorModel); ok {
		responseMessage(w, r, err, http.StatusConflict)
		return
	}

	if err, ok := err.(*errorsModel.UnprocessableErrorModel); ok {
		responseMessage(w, r, err, http.StatusUnprocessableEntity)
		return
	}

	if err, ok := err.(*errorsModel.UnauthorizedErrorModel); ok {
		responseMessage(w, r, err, http.StatusUnauthorized)
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		responseMessage(w, r, err, http.StatusNotFound)
		return
	}

	responseMessage(w, r, err, http.StatusInternalServerError)

}

func responseMessage(w http.ResponseWriter, _ *http.Request, err error, status int) {
	message := messageModel.MessageModel{Message: err.Error()}
	jsonResponse, err := json.Marshal(message)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"Message": "%s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
}
