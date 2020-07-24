package utils

import (
	"github.com/jinzhu/configor"
)

var (
	Config = struct {
		Key          string `required:"true"`
		Urlpath      string `required:"true"`
		Payday       int    `required:"true"`
		Alloworigins []string
	}{}
)

func LoadConfig() {
	configor.Load(&Config, "config/config.yml")
}
