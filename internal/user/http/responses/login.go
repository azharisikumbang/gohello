package response

type ValidLoginReponse struct {
	Token string `json:"access_token"`
}

func NewValidLoginReponse(t string) *ValidLoginReponse {
	return &ValidLoginReponse{
		Token: t,
	}
}
