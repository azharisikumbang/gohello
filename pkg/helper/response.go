package helper

import (
	"encoding/json"
	"net/http"

	"github.com/azharisikumbang/gohello/pkg/dto"
)

func NewStdReponse(data any, error []error) *dto.StdResponse {
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

func NewErrorJsonReponse(w http.ResponseWriter, ers []dto.ErrValue, code int) {
	resp := dto.NewErrorResponse()

	for _, e := range ers {
		resp.AddErrValue(e)
	}

	ToJson(resp, w, code)
}

func NewOkJsonReponse(w http.ResponseWriter, data any, code int) {
	resp := &dto.StdResponse{
		Data:       data,
		Pagination: dto.Pagination{},
		Errors:     nil,
	}

	ToJson(resp, w, code)
}

func ToJson(data any, w http.ResponseWriter, code int) {
	newData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(newData)
}

func NewErrValue(key string, val string) dto.ErrValue {
	return dto.ErrValue{
		Key:   key,
		Value: val,
	}
}
