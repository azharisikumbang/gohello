package helper

import (
	"time"

	"github.com/azharisikumbang/gohello/pkg/dto"
)

func NewFormError(n string, m string) dto.FormError {
	return dto.FormError{
		Name:  n,
		Value: m,
	}
}

func NewServerError(n string, m error, as string) dto.ServerError {
	return dto.ServerError{
		Name:      n,
		Value:     m,
		DisplayAs: as,
		Time:      time.Now(),
	}
}
