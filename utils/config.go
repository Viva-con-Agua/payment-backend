package utils

import (
	"github.com/jinzhu/configor"
)

var (
	Config = struct {
		Key string `required:"true"`
	}{}
)

func LoadConfig() {
	configor.Load(&Config, "config.yml")
}
