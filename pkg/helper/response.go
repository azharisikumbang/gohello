package helper

import (
	"encoding/json"
	"net/http"

	"github.com/azharisikumbang/gohello/pkg/dto"
)

func NewStdResponse(data any, error []error) *dto.StdResponse {
	strErrors := make([]string, len(error))
	for i, e := range error {
		strErrors[i] = e.Error()
	}

	return &dto.StdResponse{
		Data:       data,
		Pagination: dto.Pagination{},
		Errors:     strErrors,
	}
}

func NewBadRequestJsonResponse(w http.ResponseWriter, errs []dto.FormError) {
	resp := dto.NewErrorResponse()

	for _, e := range errs {
		resp.AddErrValue(&e)
	}

	WriteJson(w, http.StatusBadRequest, resp)
}

func NewServerErrorJsonResponse(w http.ResponseWriter, e error) {
	resp := dto.NewErrorResponse()
	erv := NewServerError("server_error", e, "Internal server error")
	resp.AddErrValue(&erv)

	WriteJson(w, http.StatusInternalServerError, resp)
}

func NewOKJSONReponse(w http.ResponseWriter, data any, code int) {
	resp := &dto.StdResponse{
		Data:       data,
		Pagination: dto.Pagination{},
		Errors:     nil,
	}

	WriteJson(w, code, resp)
}

func WriteJson(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewErrValue(name string, value string) dto.ErrValue {
	return dto.ErrValue{
		Name:  name,
		Value: value,
	}
}
