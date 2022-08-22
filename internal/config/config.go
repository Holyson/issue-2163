package config

import "github.com/Holyson/test-go-zero-cors/rest"

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
