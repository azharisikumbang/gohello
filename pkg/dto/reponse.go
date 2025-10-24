package dto

type Pagination struct {
	Next string `json:"next"`
	Self string `json:"self"`
	Prev string `json:"prev"`
}

type StdResponse struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"links"`
	Errors     []string   `json:"errors"`
}

// data / ok response
type DataResponse struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"links"`
}

// error repsonse
type ErrValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ErrResponse struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"links"`
	Errors     []ErrValue `json:"errors"`
	Message    string     `json:"message"`
}

func NewErrorResponse() *ErrResponse {
	return &ErrResponse{
		Data:   nil,
		Errors: make([]ErrValue, 0),
	}
}

func (e *ErrResponse) AddErrValue(i ErrValue) *ErrResponse {
	e.Errors = append(e.Errors, i)

	return e
}

func (e *ErrResponse) AddUserMessage(m string) {
	e.Message = m
}
