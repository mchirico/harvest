package configure

type SecretStruct struct {
	Id      string `json:"clientID"`
	Secret  string `json:"clientSecret"`
	Url     string `json:"url"`
	Code    string `json:"code"`
	Seconds string `json:"seconds"`
	Expire  string `json:"expire"`
}
