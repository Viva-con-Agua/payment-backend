package utils

import (
	"github.com/jinzhu/configor"
)

var (
	Config = struct {
		Key          string `required:"true"`
		Alloworigins []string
	}{}
)

func LoadConfig() {
	configor.Load(&Config, "config/config.yml")
}
