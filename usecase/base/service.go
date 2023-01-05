package base

import (
	"github.com/rs/zerolog"
)

type Service struct {
	Repository Repository
	Logger     *zerolog.Logger
}

func LoadService(r Repository, l *zerolog.Logger) *Service {
	return &Service{Repository: r, Logger: l}
}
