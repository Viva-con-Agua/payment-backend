package utils

import (
	"github.com/jinzhu/configor"
)

var (
	Config = struct {
		Key          string `required:"true"`
		Urlpath      string
		Alloworigins []string
	}{}
)

func LoadConfig() {
	configor.Load(&Config, "config/config.yml")
}
