package utils

import (
	"github.com/lucsky/cuid"
)

func GenerateCUID() string {
	return cuid.New()
}
