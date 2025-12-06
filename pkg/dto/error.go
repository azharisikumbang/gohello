package dto

import (
	"log"
	"time"
)

type ErrValueInterface interface {
	ErrorName() string
	ErrorValue() string
}

type FormError struct {
	Name  string
	Value string
}

func (e *FormError) ErrorValue() string {
	return e.Value
}

func (e *FormError) ErrorName() string {
	return e.Name
}

type ServerError struct {
	Name      string
	Value     error
	DisplayAs string
	Time      time.Time
}

func (e *ServerError) ErrorName() string {
	return e.Name
}

func (e *ServerError) ErrorValue() string {
	return e.DisplayAs
}

func (e *ServerError) Error() string {
	return e.ErrorValue()
}

func (e *ServerError) Log() {
	log.Fatal(e.Value.Error())
}
