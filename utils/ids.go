package utils

import (
	"github.com/go-basic/uuid"
	"strings"
)

func NewUuid() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}


