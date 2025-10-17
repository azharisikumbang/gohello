package helper

import (
	"encoding/json"
	"net/http"

	"github.com/azharisikumbang/gohello/pkg/dto"
)

func NewStdReponse(data any, error []error) *dto.StdReponse {
	strErrors := make([]string, len(error))
	for i, e := range error {
		strErrors[i] = e.Error()
	}

	return &dto.StdReponse{
		Data:       data,
		Pagination: dto.Pagination{},
		Errors:     strErrors,
	}
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
