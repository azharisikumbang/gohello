package dto

type Pagination struct {
	Next string `json:"next"`
	Self string `json:"self"`
	Prev string `json:"prev"`
}

type StdReponse struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"links"`
	Errors     []error    `json:"errors"`
}
