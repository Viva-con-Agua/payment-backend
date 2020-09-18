package utils

import (
	"github.com/jinzhu/configor"
)

var (
  Config = struct {
		Alloworigins []string
	}{}
)

func LoadConfig() {
	configor.Load(&Config, "config/config.yml")
}
