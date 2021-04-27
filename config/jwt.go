package config

type Jwt struct {
	Secret string `json:"secret"`
	Expire int    `json:"expire"`
}
